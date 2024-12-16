// Package internal contains data used throughout the project.
package internal

const (
	ExperimentEnableStdFrontendFetch = "enable-std-frontend-fetch"
)

// Experiments represents all of the active experiments in the codebase and
// a description of each experiment.
var Experiments = map[string]string{
	ExperimentEnableStdFrontendFetch: "Enable frontend fetching for module std.",
}

// Experiment holds data associated with an experimental feature for frontend
// or worker.
type Experiment struct {
	// This struct is used to decode dynamic config (see
	// internal/config/dynconfig). Make sure that changes to this struct are
	// coordinated with the deployment of config files.

	// Name is the name of the feature.
	Name string

	// Rollout is the percentage of requests enrolled in the experiment.
	Rollout uint

	// Description provides a description of the experiment.
	Description string
}
