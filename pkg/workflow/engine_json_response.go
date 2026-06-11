package workflow

// EngineJSONResponse represents the shared JSON structure returned by engines
// that emit a final response body plus stats metadata.
type EngineJSONResponse struct {
	Response string         `json:"response"`
	Stats    map[string]any `json:"stats"`
}
