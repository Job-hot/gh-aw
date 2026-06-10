// @ts-check
import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import fs from "fs";
import path from "path";
import crypto from "crypto";

const ASSET_DIR = "/tmp/gh-aw/safeoutputs/assets";

/**
 * Writes a temp agent output file, sets GH_AW_AGENT_OUTPUT, and returns its path.
 * @param {object|string} data
 * @returns {string}
 */
function setAgentOutput(data) {
  const filePath = path.join("/tmp", `test_agent_output_${Date.now()}_${Math.random().toString(36).slice(2)}.json`);
  const content = typeof data === "string" ? data : JSON.stringify(data);
  fs.writeFileSync(filePath, content);
  process.env.GH_AW_AGENT_OUTPUT = filePath;
  return filePath;
}

/**
 * Creates a test asset file in the asset staging directory.
 * @param {string} fileName
 * @param {string} [content]
 * @returns {{ assetPath: string, sha: string, size: number }}
 */
function createAssetFile(fileName, content = "fake asset data") {
  fs.mkdirSync(ASSET_DIR, { recursive: true });
  const assetPath = path.join(ASSET_DIR, fileName);
  fs.writeFileSync(assetPath, content);
  const sha = crypto.createHash("sha256").update(fs.readFileSync(assetPath)).digest("hex");
  return { assetPath, sha, size: content.length };
}

describe("upload_assets.cjs", () => {
  /** @type {ReturnType<typeof vi.fn>} */
  let mockCore;
  /** @type {{ exec: ReturnType<typeof vi.fn> }} */
  let mockExec;
  /** @type {string[]} */
  let tempFiles;
  /** @type {string} */
  let scriptText;

  /**
   * Evaluates the script with mocked globals.
   */
  const executeScript = async () => {
    global.core = mockCore;
    global.exec = mockExec;
    await eval(`(async () => { ${scriptText}; await main(); })()`);
  };

  beforeEach(() => {
    vi.clearAllMocks();
    tempFiles = [];

    delete process.env.GH_AW_ASSETS_BRANCH;
    delete process.env.GH_AW_AGENT_OUTPUT;
    delete process.env.GH_AW_SAFE_OUTPUTS_STAGED;

    mockCore = {
      info: vi.fn(),
      warning: vi.fn(),
      error: vi.fn(),
      setFailed: vi.fn(),
      setOutput: vi.fn(),
      summary: {
        addRaw: vi.fn().mockReturnThis(),
        write: vi.fn().mockResolvedValue(undefined),
      },
    };

    mockExec = { exec: vi.fn().mockResolvedValue(0) };

    scriptText = fs.readFileSync(path.join(__dirname, "upload_assets.cjs"), "utf8");
  });

  afterEach(() => {
    for (const f of tempFiles) {
      if (fs.existsSync(f)) fs.unlinkSync(f);
    }
  });

  describe("branch env var validation", () => {
    it("fails with ERR_CONFIG when GH_AW_ASSETS_BRANCH is not set", async () => {
      setAgentOutput({ items: [] });
      await executeScript();
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("ERR_CONFIG"));
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("GH_AW_ASSETS_BRANCH"));
    });
  });

  describe("agent output handling", () => {
    it("outputs upload_count=0 and branch_name when agent output fails to load", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test";
      process.env.GH_AW_AGENT_OUTPUT = "/nonexistent/path.json";
      await executeScript();
      expect(mockCore.setOutput).toHaveBeenCalledWith("upload_count", "0");
      expect(mockCore.setOutput).toHaveBeenCalledWith("branch_name", "assets/test");
      expect(mockCore.setFailed).not.toHaveBeenCalled();
    });

    it("outputs upload_count=0 when there are no upload-asset items", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test";
      setAgentOutput({ items: [{ type: "other_type" }] });
      await executeScript();
      expect(mockCore.setOutput).toHaveBeenCalledWith("upload_count", "0");
      expect(mockCore.setFailed).not.toHaveBeenCalled();
    });
  });

  describe("branch name normalization", () => {
    it("normalizes branch names with special characters", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/My Branch!@#$%";
      setAgentOutput({ items: [] });
      await executeScript();
      const call = mockCore.setOutput.mock.calls.find(/** @param {any[]} c */ c => c[0] === "branch_name");
      expect(call).toBeDefined();
      expect(call[1]).toBe("assets/My-Branch");
    });
  });

  describe("asset validation", () => {
    it("fails with ERR_VALIDATION when SHA does not match", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      createAssetFile("sha-mismatch.png", "real content");
      setAgentOutput({
        items: [
          {
            type: "upload_asset",
            fileName: "sha-mismatch.png",
            sha: "0000000000000000000000000000000000000000000000000000000000000000",
            size: 12,
            targetFileName: "out.png",
            url: "https://example.com/sha-mismatch.png",
          },
        ],
      });
      mockExec.exec.mockImplementation(async (_cmd, /** @type {string[]} */ args) => {
        if (args?.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("ERR_VALIDATION"));
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("SHA mismatch"));
    });

    it("fails with ERR_SYSTEM when the asset source file does not exist", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      setAgentOutput({
        items: [
          {
            type: "upload_asset",
            fileName: "ghost.png",
            sha: "abc123",
            size: 100,
            targetFileName: "out.png",
            url: "https://example.com/ghost.png",
          },
        ],
      });
      mockExec.exec.mockImplementation(async (_cmd, /** @type {string[]} */ args) => {
        if (args?.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("ERR_SYSTEM"));
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("ghost.png"));
    });

    it("fails with ERR_VALIDATION when a required asset field is missing", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      // sha and targetFileName are intentionally omitted
      setAgentOutput({ items: [{ type: "upload_asset", fileName: "test.png" }] });
      mockExec.exec.mockImplementation(async (_cmd, /** @type {string[]} */ args) => {
        if (args?.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("ERR_VALIDATION"));
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("Invalid asset entry"));
    });
  });

  describe("branch prefix validation", () => {
    it("creates an orphaned branch when the assets/ branch does not exist on origin", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test-workflow";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      const { sha, size } = createAssetFile("orphan-test.png");
      tempFiles.push("orphan-test.png");
      setAgentOutput({
        items: [{ type: "upload_asset", fileName: "orphan-test.png", sha, size, targetFileName: "orphan-test.png", url: "https://example.com/orphan-test.png" }],
      });
      let orphanCreated = false;
      mockExec.exec.mockImplementation(async (cmd, /** @type {string[]} */ args) => {
        const full = Array.isArray(args) ? `${cmd} ${args.join(" ")}` : cmd;
        if (full.includes("checkout --orphan")) orphanCreated = true;
        if (full.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      expect(orphanCreated).toBe(true);
      expect(mockCore.setFailed).not.toHaveBeenCalled();
    });

    it("fails when trying to create an orphaned branch without the assets/ prefix", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "custom/branch-name";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      const { sha, size } = createAssetFile("prefix-test.png");
      setAgentOutput({
        items: [{ type: "upload_asset", fileName: "prefix-test.png", sha, size, targetFileName: "prefix-test.png", url: "https://example.com/prefix-test.png" }],
      });
      mockExec.exec.mockImplementation(async (_cmd, /** @type {string[]} */ args) => {
        if (args?.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("assets/"));
      expect(mockCore.setFailed).toHaveBeenCalledWith(expect.stringContaining("custom/branch-name"));
    });

    it("checks out an existing branch regardless of its prefix", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "custom/existing-branch";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      const { sha, size } = createAssetFile("existing-branch-test.png");
      tempFiles.push("existing-branch-test.png");
      setAgentOutput({
        items: [{ type: "upload_asset", fileName: "existing-branch-test.png", sha, size, targetFileName: "existing-branch-test.png", url: "https://example.com/existing-branch-test.png" }],
      });
      let existingCheckedOut = false;
      let orphanCreated = false;
      mockExec.exec.mockImplementation(async (cmd, /** @type {string[]} */ args) => {
        const full = Array.isArray(args) ? `${cmd} ${args.join(" ")}` : cmd;
        if (full.includes("checkout -B")) existingCheckedOut = true;
        if (full.includes("checkout --orphan")) orphanCreated = true;
        return 0; // rev-parse succeeds: branch exists on origin
      });
      await executeScript();
      expect(existingCheckedOut).toBe(true);
      expect(orphanCreated).toBe(false);
      expect(mockCore.setFailed).not.toHaveBeenCalled();
    });
  });

  describe("file upload behaviour", () => {
    it("skips an asset whose target file already exists in the branch", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      const { sha, size } = createAssetFile("already-there.png");
      // Simulate target file already present in the working tree
      fs.writeFileSync("already-there.png", "existing content");
      tempFiles.push("already-there.png");
      setAgentOutput({
        items: [{ type: "upload_asset", fileName: "already-there.png", sha, size, targetFileName: "already-there.png", url: "https://example.com/already-there.png" }],
      });
      mockExec.exec.mockImplementation(async (_cmd, /** @type {string[]} */ args) => {
        if (args?.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      expect(mockCore.setFailed).not.toHaveBeenCalled();
      const uploadCountCall = mockCore.setOutput.mock.calls.find(/** @param {any[]} c */ c => c[0] === "upload_count");
      expect(uploadCountCall?.[1]).toBe("0");
    });

    it("skips git push in staged mode", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "true";
      const { sha, size } = createAssetFile("staged.png");
      tempFiles.push("staged.png");
      setAgentOutput({
        items: [{ type: "upload_asset", fileName: "staged.png", sha, size, targetFileName: "staged.png", url: "https://example.com/staged.png" }],
      });
      mockExec.exec.mockImplementation(async (_cmd, /** @type {string[]} */ args) => {
        if (args?.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      const pushCalls = mockExec.exec.mock.calls.filter(c => Array.isArray(c[1]) && c[1].includes("push"));
      expect(pushCalls).toHaveLength(0);
      expect(mockCore.setFailed).not.toHaveBeenCalled();
    });

    it("pushes to origin after upload in non-staged mode", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      const { sha, size } = createAssetFile("push-test.png");
      tempFiles.push("push-test.png");
      setAgentOutput({
        items: [{ type: "upload_asset", fileName: "push-test.png", sha, size, targetFileName: "push-test.png", url: "https://example.com/push-test.png" }],
      });
      mockExec.exec.mockImplementation(async (_cmd, /** @type {string[]} */ args) => {
        if (args?.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      const pushCalls = mockExec.exec.mock.calls.filter(c => Array.isArray(c[1]) && c[1].includes("push"));
      expect(pushCalls.length).toBeGreaterThan(0);
      expect(mockCore.setFailed).not.toHaveBeenCalled();
    });

    it("uses array args for git commit (no shell injection risk)", async () => {
      process.env.GH_AW_ASSETS_BRANCH = "assets/test-workflow";
      process.env.GH_AW_SAFE_OUTPUTS_STAGED = "false";
      const { sha, size } = createAssetFile("injection-test.png");
      tempFiles.push("injection-test.png");
      setAgentOutput({
        items: [{ type: "upload_asset", fileName: "injection-test.png", sha, size, targetFileName: "injection-test.png", url: "https://example.com/injection-test.png" }],
      });
      mockExec.exec.mockImplementation(async (_cmd, /** @type {string[]} */ args) => {
        if (args?.includes("rev-parse")) throw new Error("branch not found");
        return 0;
      });
      await executeScript();
      const commitCall = mockExec.exec.mock.calls.find(c => c[0] === "git" && Array.isArray(c[1]) && c[1].includes("commit"));
      expect(commitCall).toBeDefined();
      const mIdx = commitCall[1].indexOf("-m");
      const msg = commitCall[1][mIdx + 1];
      expect(msg).toBeDefined();
      expect(typeof msg).toBe("string");
      expect(msg).not.toMatch(/^"/);
      expect(msg).not.toMatch(/"$/);
      expect(msg).toContain("[skip-ci]");
      expect(msg).toContain("asset(s)");
    });
  });
});
