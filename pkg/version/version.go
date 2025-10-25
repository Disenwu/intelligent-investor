package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	gitVersion   = "v0.0.0-master+$Format:%h$"
	buildTime    = "1970-01-01T00:00:00Z"
	gitCommit    = "$Format:%H$"
	gitTreeState = "$Format:%d$"
)

type VersionInfo struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildTime    string `json:"buildTime"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (version *VersionInfo) Version() string {
	return version.GitVersion
}

func (version *VersionInfo) ToJSON() string {
	s, _ := json.Marshal(version)
	return string(s)
}

func (version VersionInfo) String() string {
	return fmt.Sprintf("GitVersion: %s\nGitCommit: %s\nGitTreeState: %s\nBuildTime: %s\nGoVersion: %s\nCompiler: %s\nPlatform: %s",
		version.GitVersion, version.GitCommit, version.GitTreeState, version.BuildTime, version.GoVersion, version.Compiler, version.Platform)
}

func Get() VersionInfo {
	return VersionInfo{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildTime:    buildTime,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
