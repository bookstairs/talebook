package config

import (
	"fmt"
	"runtime"
)

// This file could be replaced by debug/buildinfo after switching to go 1.18.
var (
	gitMajor = "1" // the major part of the version.
	gitMinor = "0" // the minor part of the version.
	gitPatch = "0" // the patch part of the version.

	gitCommit    = "" // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState = "" // state of git tree, either "clean" or "dirty"

	buildDate = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

const versionTmpl = "Version: %s.%s.%s, Build Date: %s\nGit Commit: %s"

// Version contains versioning information.
// how we'll want to distribute that information.
type Version struct {
	Major        string `json:"major"`
	Minor        string `json:"minor"`
	Patch        string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String use to display the version information in command line.
func (v *Version) String() string {
	return fmt.Sprintf(versionTmpl, v.Major, v.Minor, v.Patch, v.BuildDate, v.GitCommit)
}

// SemanticString use to display in API response.
func (v *Version) SemanticString() string {
	return v.Major + "." + v.Minor + "." + v.Patch
}

// TalebookVersion will return and print the current project version.
// These properties are provided by goreleaser.
func TalebookVersion() *Version {
	return &Version{
		Major:        gitMajor,
		Minor:        gitMinor,
		Patch:        gitPatch,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
