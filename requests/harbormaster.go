package requests

import (
	"github.com/thought-machine/gonduit/constants"
	"github.com/thought-machine/gonduit/entities"
)

// HarbormasterCreateArtifactRequest represents a request to the harbormaster.createartifact call.
type HarbormasterCreateArtifactRequest struct {
	BuildTargetPHID string                             `json:"buildTargetPHID"`
	ArtifactKey     string                             `json:"artifactKey"`
	ArtifactType    constants.HarbormasterArtifactType `json:"artifactType"`
	ArtifactData    map[string]interface{}             `json:"artifactData"`
	Request
}

// HarbormasterSendMessageRequest represents a request to the harbormaster.sendmessage call.
type HarbormasterSendMessageRequest struct {
	BuildTargetPHID string                            `json:"buildTargetPHID"`
	Type            constants.HarbormasterMessageType `json:"type"`
	Unit            []entities.UnitResult             `json:"unit"`
	Lint            []entities.LintResult             `json:"lint"`
	Request
}
