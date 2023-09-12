// Description: Entrypoint for the dots-cli application
package main

import (
	"fmt"

	"github.com/matt-riley/dots-cli/cmd/dots"
)

var (
	version = "0.0.1"
	commit  = "none"
	date    = "unknown"
)

func main() {
	dots.SetVersionInfo(version, commit, date)
	err := dots.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
