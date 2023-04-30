package dots

import "fmt"

func SetVersionInfo(version string, commit string, date string) {
	rootCmd.Version = fmt.Sprintf(
		"%s (Built on %s from Git SHA %s)",
		version,
		date,
		commit,
	)
}
