package entities

import "github.com/thought-machine/gonduit/constants"

// A UnitResult represents the result of a unit test, as sent to Harbormaster.
type UnitResult struct {
	Name      string                           `json:"name"`
	Result    constants.HarbormasterUnitResult `json:"result"`
	Namespace string                           `json:"namespace"`
	Engine    string                           `json:"engine"`
	Duration  float64                          `json:"duration"`
	Path      string                           `json:"path"`
	Coverage  map[string]string                `json:"coverage"`
	Details   string                           `json:"details"`
	Format    constants.HarbormasterUnitFormat `json:"format"`
}

// A LintResult represents the result of a lint warning, as sent to Harbormaster.
type LintResult struct {
	Name        string                             `json:"name"`
	Code        string                             `json:"code"`
	Severity    constants.HarbormasterLintSeverity `json:"severity"`
	Path        string                             `json:"path"`
	Line        int                                `json:"line"`
	Char        int                                `json:"char"`
	Description string                             `json:"description"`
}
