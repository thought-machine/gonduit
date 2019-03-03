package constants

// HarbormasterArtifactType are the types of artifact that can be attached through Harbormaster.
type HarbormasterArtifactType string

const (
	// HarbormasterHost references a host lease from Drydock.
	HarbormasterHost HarbormasterArtifactType = "host"
	// HarbormasterWorkingCopy references a working copy lease from Drydock.
	HarbormasterWorkingCopy = "working-copy"
	// HarbormasterFile stores a reference to file data which has been uploaded to Phabricator.
	HarbormasterFile = "file"
	// HarbormasterURI stores a URI.
	HarbormasterURI = "uri"
)
