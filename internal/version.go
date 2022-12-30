package labs

// https://stackoverflow.com/questions/11354518/application-auto-build-versioning
// https://icinga.com/blog/2022/05/25/embedding-git-commit-information-in-go-binaries

// go tool nm build/labs | grep Version
// "github.com/niqdev/gopher-labs/internal.Version"
var Version string
