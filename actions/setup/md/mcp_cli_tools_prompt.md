<mcp-clis>
CLI servers are available on `PATH`:
__GH_AW_MCP_CLI_SERVERS_LIST__
Use `<server> --help` for tool names, parameters, and examples before calling any command.
If a listed CLI wrapper is unavailable, continue using the corresponding MCP tool directly. Do not treat a missing CLI wrapper as a missing capability.
To pass many arguments safely, pipe a JSON object on stdin with `printf` and pass `.` as the payload sentinel: `printf '%s\n' '{"param":"value","count":1}' | <server> <tool> .`
</mcp-clis>
