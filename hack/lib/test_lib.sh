function listPkgs() {
	go list ./cmd/... ./pkg/... | grep -v generated
}

function listFiles() {
	# pipeline is much faster than for loop
	listPkgs | xargs -I {} find "${GOPATH}/src/{}" -name '*.go' | grep -v generated
}
