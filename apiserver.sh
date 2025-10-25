VERSION_PACKAGE="intelligent-investor/pkg/version"
output_dir="./_output"
if [ -z "$GIT_VERSION" ]; then
	GIT_VERSION=$(git describe --tags --always --match="v*")
fi

GIT_TREE_STATE=$(git diff --no-ext-diff --quiet)
if [ -z "$GIT_TREE_STATE" ]; then
	GIT_TREE_STATE="dirty"
fi
GIT_COMMIT=$(git rev-parse HEAD)

GO_FLAGS="-X ${VERSION_PACKAGE}.gitVersion=${GIT_VERSION} -X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT} -X ${VERSION_PACKAGE}.gitTreeState=${GIT_TREE_STATE} -X ${VERSION_PACKAGE}.buildTime=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
echo "GO_FLAGS: ${GO_FLAGS}"
go build -v -ldflags "${GO_FLAGS}" -o $output_dir/apiserver.exe ./cmd/api-server/main.go
