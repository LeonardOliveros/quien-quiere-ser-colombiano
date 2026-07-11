// Stub module: keeps the Go toolchain (go build/vet/test ./...) from
// descending into infra/, which is a TypeScript CDK app whose node_modules
// contains non-buildable .go template files.
module infra-not-a-go-package

go 1.25
