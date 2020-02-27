package responses

// DiffusionFileContentQueryResponse represents a response of the
// diffusion.filecontentquery call.
type DiffusionFileContentQueryResponse struct {
	TooSlow  bool   `json:"tooSlow"`
	TooHuge  bool   `json:"tooHuge"`
	FilePHID string `json:"filePHID"`
}
