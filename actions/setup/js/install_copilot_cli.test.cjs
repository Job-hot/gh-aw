// @ts-check
import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { createRequire } from "module";
import * as fs from "fs";
import * as os from "os";
import * as path from "path";

const req = createRequire(import.meta.url);
const { parseSemver, cmpSemver, rowMatchesGhAw, copilotRows, pickRange, detectArch, findCachedCopilot } = req("./install_copilot_cli.cjs");

describe("parseSemver", () => {
  it("parses release versions", () => {
    expect(parseSemver("1.2.3")).toEqual([1, 2, 3, ""]);
  });

  it("parses pre-release versions", () => {
    expect(parseSemver("1.2.3-beta.1")).toEqual([1, 2, 3, "beta.1"]);
  });

  it("rejects non-string input", () => {
    expect(parseSemver(null)).toBeNull();
    expect(parseSemver(undefined)).toBeNull();
    expect(parseSemver(123)).toBeNull();
  });

  it("rejects malformed strings", () => {
    expect(parseSemver("dev")).toBeNull();
    expect(parseSemver("1.2")).toBeNull();
    expect(parseSemver("v1.2.3")).toBeNull();
    expect(parseSemver("1.2.3.4")).toBeNull();
  });
});

describe("cmpSemver", () => {
  it("compares major/minor/patch", () => {
    expect(cmpSemver([1, 0, 0, ""], [2, 0, 0, ""])).toBe(-1);
    expect(cmpSemver([1, 5, 0, ""], [1, 2, 0, ""])).toBe(1);
    expect(cmpSemver([1, 2, 3, ""], [1, 2, 3, ""])).toBe(0);
    expect(cmpSemver([1, 2, 3, ""], [1, 2, 4, ""])).toBe(-1);
  });

  it("treats pre-release as lower than release", () => {
    expect(cmpSemver([1, 0, 0, "beta"], [1, 0, 0, ""])).toBe(-1);
    expect(cmpSemver([1, 0, 0, ""], [1, 0, 0, "beta"])).toBe(1);
  });
});

describe("rowMatchesGhAw", () => {
  it("matches wildcard rows regardless of compiler version", () => {
    expect(rowMatchesGhAw({ "max-gh-aw": "*" }, null)).toBe(true);
    expect(rowMatchesGhAw({ "max-gh-aw": "*" }, [1, 2, 3, ""])).toBe(true);
  });

  it("matches when compiler version is <= max-gh-aw", () => {
    expect(rowMatchesGhAw({ "max-gh-aw": "2.0.0" }, [1, 5, 0, ""])).toBe(true);
    expect(rowMatchesGhAw({ "max-gh-aw": "2.0.0" }, [2, 0, 0, ""])).toBe(true);
  });

  it("rejects when compiler version exceeds max-gh-aw", () => {
    expect(rowMatchesGhAw({ "max-gh-aw": "1.0.0" }, [2, 0, 0, ""])).toBe(false);
  });

  it("rejects non-wildcard rows when compiler version is unparseable", () => {
    expect(rowMatchesGhAw({ "max-gh-aw": "2.0.0" }, null)).toBe(false);
  });
});

describe("copilotRows", () => {
  it("returns rows from a well-formed matrix", () => {
    const rows = copilotRows({
      "agent-compat-v1": { copilot: [{ "max-gh-aw": "*" }] },
    });
    expect(rows).toHaveLength(1);
  });

  it("returns [] for malformed inputs", () => {
    expect(copilotRows(null)).toEqual([]);
    expect(copilotRows({})).toEqual([]);
    expect(copilotRows({ "agent-compat-v1": null })).toEqual([]);
    expect(copilotRows({ "agent-compat-v1": { copilot: "not-an-array" } })).toEqual([]);
  });
});

describe("pickRange", () => {
  it("returns the first matching row's range", () => {
    const rows = [{ "max-gh-aw": "*", "min-agent": "1.0.50", "max-agent": "1.0.54" }];
    const r = pickRange(rows, null);
    expect(r).not.toBeNull();
    expect(r.min.slice(0, 3)).toEqual([1, 0, 50]);
    expect(r.max.slice(0, 3)).toEqual([1, 0, 54]);
  });

  it("skips rows whose max-gh-aw does not cover the compiler version", () => {
    const rows = [
      { "max-gh-aw": "1.0.0", "min-agent": "1.0.40", "max-agent": "1.0.45" },
      { "max-gh-aw": "*", "min-agent": "1.0.50", "max-agent": "1.0.54" },
    ];
    const r = pickRange(rows, [2, 0, 0, ""]);
    expect(r.min.slice(0, 3)).toEqual([1, 0, 50]);
  });

  it("skips rows with unparseable min/max-agent", () => {
    const rows = [
      { "max-gh-aw": "*", "min-agent": "bad", "max-agent": "1.0.54" },
      { "max-gh-aw": "*", "min-agent": "1.0.50", "max-agent": "1.0.54" },
    ];
    const r = pickRange(rows, null);
    expect(r.min.slice(0, 3)).toEqual([1, 0, 50]);
  });

  it("returns null when no row matches", () => {
    expect(pickRange([], null)).toBeNull();
  });
});

describe("detectArch", () => {
  it("returns the current process arch", () => {
    expect(typeof detectArch()).toBe("string");
    expect(detectArch().length).toBeGreaterThan(0);
  });
});

describe("findCachedCopilot", () => {
  /** @type {string} */
  let toolCacheRoot;
  const arch = "x64";

  beforeEach(() => {
    toolCacheRoot = fs.mkdtempSync(path.join(os.tmpdir(), "tc-test-"));
  });

  afterEach(() => {
    fs.rmSync(toolCacheRoot, { recursive: true, force: true });
  });

  /**
   * Materialize the `$AGENT_TOOLSDIRECTORY/copilot-cli/<version>/<arch>/bin/copilot`
   * layout plus the `<arch>.complete` sibling marker that matches what
   * `install-copilot-cli.sh` produces in runner-images.
   */
  function seedToolcache(version, { complete = true, withBin = true } = {}) {
    const versionDir = path.join(toolCacheRoot, "copilot-cli", version);
    const archDir = path.join(versionDir, arch);
    fs.mkdirSync(archDir, { recursive: true });
    if (withBin) {
      const binDir = path.join(archDir, "bin");
      fs.mkdirSync(binDir, { recursive: true });
      fs.writeFileSync(path.join(binDir, "copilot"), "#!/bin/sh\n", { mode: 0o755 });
    }
    if (complete) {
      fs.writeFileSync(`${archDir}.complete`, "");
    }
  }

  it("returns the highest version in range", () => {
    seedToolcache("1.0.50");
    seedToolcache("1.0.52");
    seedToolcache("1.0.54");
    const r = findCachedCopilot(toolCacheRoot, arch, {
      min: parseSemver("1.0.50"),
      max: parseSemver("1.0.54"),
    });
    expect(r).not.toBeNull();
    expect(r.version).toBe("1.0.54");
    expect(r.binDir).toBe(path.join(toolCacheRoot, "copilot-cli", "1.0.54", arch, "bin"));
  });

  it("ignores versions outside the range", () => {
    seedToolcache("1.0.49");
    seedToolcache("1.0.55");
    const r = findCachedCopilot(toolCacheRoot, arch, {
      min: parseSemver("1.0.50"),
      max: parseSemver("1.0.54"),
    });
    expect(r).toBeNull();
  });

  it("ignores entries without a .complete marker", () => {
    seedToolcache("1.0.52", { complete: false });
    const r = findCachedCopilot(toolCacheRoot, arch, {
      min: parseSemver("1.0.50"),
      max: parseSemver("1.0.54"),
    });
    expect(r).toBeNull();
  });

  it("ignores entries missing the copilot binary", () => {
    seedToolcache("1.0.52", { withBin: false });
    const r = findCachedCopilot(toolCacheRoot, arch, {
      min: parseSemver("1.0.50"),
      max: parseSemver("1.0.54"),
    });
    expect(r).toBeNull();
  });

  it("ignores non-semver directory names", () => {
    fs.mkdirSync(path.join(toolCacheRoot, "copilot-cli", "junk"), { recursive: true });
    seedToolcache("1.0.52");
    const r = findCachedCopilot(toolCacheRoot, arch, {
      min: parseSemver("1.0.50"),
      max: parseSemver("1.0.54"),
    });
    expect(r.version).toBe("1.0.52");
  });

  it("returns null when the copilot-cli directory does not exist", () => {
    const r = findCachedCopilot(toolCacheRoot, arch, {
      min: parseSemver("1.0.50"),
      max: parseSemver("1.0.54"),
    });
    expect(r).toBeNull();
  });
});
