package vending

var (
	Version   string
	Commit    string
	BuildTime string
)

func init() {
	// If version, commit, or build time are not set, make that clear.
	if Version == "" {
		Version = "unknown"
	}
	if Commit == "" {
		Commit = "unknown"
	}
	if BuildTime == "" {
		BuildTime = "unknown"
	}
}
