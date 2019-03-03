package requests

import "github.com/thought-machine/gonduit/constants"

// HarbormasterCreateArtifactRequest represents a request to the harbormaster.createartifact call.
type HarbormasterCreateArtifactRequest struct {
	BuildTargetPHID string                             `json:"buildTargetPHID"`
	ArtifactKey     string                             `json:"artifactKey"`
	ArtifactType    constants.HarbormasterArtifactType `json:"artifactType"`
	ArtifactData    map[string]interface{}             `json:"artifactData"`
	Request
}
