package essai

var buildInfo BuildInfo

type BuildInfo struct {
    Version         string
    GoVersion       string
    Commit          string
    Date            string
}

func SetBuildInfo(version, goVersion, commit, date string) {
    buildInfo = BuildInfo{
        Version: version,
        GoVersion: goVersion,
        Commit: commit,
        Date: date,
    }
}

func GetBuildInfo() BuildInfo {
    return buildInfo
}
