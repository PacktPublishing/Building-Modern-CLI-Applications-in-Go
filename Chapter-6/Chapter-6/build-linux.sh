#!/bin/bash

# Generate linux builds
    echo "installing for linux..."  
	go install cmd/uppercase/uppercase.go
	go install cmd/lettercount/lettercount.go
	go install cmd/pages/pages.go
	go install cmd/timeout/timeout.go
	go install cmd/panic/panic.go
	go install cmd/error/error.go
	go install cmd/api/api.go