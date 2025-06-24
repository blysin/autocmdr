// Package version manages the application version information and provides version-related utilities.
package version

import (
	"fmt"
	"runtime"
)

var (
	// Version is the current version of the application
	Version = "v1.0.0"
	// GitCommit is the git commit hash
	GitCommit = "unknown"
	// BuildDate is the build date
	BuildDate = "unknown"
	// GoVersion is the Go version used to build the binary
	GoVersion = runtime.Version()
)

// Info contains version information
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

// Get returns version information
func Get() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		GoVersion: GoVersion,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String returns a formatted version string
func (i Info) String() string {
	return fmt.Sprintf("Version: %s\nGit Commit: %s\nBuild Date: %s\nGo Version: %s\nPlatform: %s",
		i.Version, i.GitCommit, i.BuildDate, i.GoVersion, i.Platform)
}
