// Package version provides build version information.
package version

import (
	"fmt"
	"runtime"
)

var (
	// Version is the main version number that is being run at the moment.
	Version = "dev"

	// GitCommit is the git commit that was compiled.
	GitCommit = "unknown"

	// BuildDate is the date the binary was built.
	BuildDate = "unknown"
)

// BuildInfo represents the build information.
type BuildInfo struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

// Get returns the complete build information.
func Get() BuildInfo {
	return BuildInfo{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String returns a formatted version string.
func String() string {
	info := Get()
	return fmt.Sprintf("Version: %s\nGit Commit: %s\nBuild Date: %s\nGo Version: %s\nPlatform: %s",
		info.Version,
		info.GitCommit,
		info.BuildDate,
		info.GoVersion,
		info.Platform,
	)
}
