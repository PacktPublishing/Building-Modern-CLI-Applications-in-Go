# audiofile
In Chapter 14, Publish your Go binary as a Homebrew Formula with GoReleaser, we do exactly that. 

## Steps to get setup with GoReleaser:

### Install GoReleaser
Visit https://goreleaser.com/install/

### Initialize GoReleaser
`goreleaser init`

### Run a "local-only" release to see if it works
`goreleaser release --snapshot --clean`

### Create a new tag
`git tag -a v0.1 -m "Initial deploy"`
`git push origin v0.1`