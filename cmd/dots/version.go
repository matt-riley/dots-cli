package dots

import "fmt"

// SetVersionInfo gets the version number from the git hash
func SetVersionInfo(version string, commit string, date string) {
	rootCmd.Version = fmt.Sprintf(
		"%s (Built on %s from Git SHA %s)",
		version,
		date,
		commit,
	)
}
