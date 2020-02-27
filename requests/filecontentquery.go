package requests

// FileContentQueryRequest represents a request to the
// diffusion.filecontentquery call.
type FileContentQueryRequest struct {
	Path       string `json:"path"`
	Commit     string `json:"string"`
	Timeout    int    `json:"timeout"`
	ByteLimit  int    `json:"byteLimit"`
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
}
