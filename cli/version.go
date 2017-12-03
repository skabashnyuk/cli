package cli

// Default build-time variable.
// These values are overriding via ldflags
var (
	Version    = "unknown-version"
	GitCommit  = "unknown-commit"
	GitSummary = "unknown-summary"
	GitBranch  = "unknown-branch"
	BuildDate  = "unknown-buildtime"
)
