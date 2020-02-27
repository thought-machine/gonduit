package responses

// FileContentQueryResponse represents a response of the
// diffusion.filecontentquery call.
type FileContentQueryResponse struct {
	TooSlow  bool   `json:"tooSlow"`
	TooHuge  bool   `json:"tooHuge"`
	FilePHID string `json:"filePHID"`
}
