//go:build buildChecks

package main

import (
	"fmt"
	"go/build"
)

func buildChecks() {
	ctx := build.Context{}
	p1, err := ctx.Import(".", ".", build.AllowBinary)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("Dir:", p1.Dir)
	fmt.Println("Package name: ", p1.Name)
	fmt.Println("AllTags: ", p1.AllTags)
	fmt.Println("GoFiles: ", p1.GoFiles)
	fmt.Println("Imports: ", p1.Imports)
	fmt.Println("isCommand/main package: ", p1.IsCommand())
	fmt.Println("IsLocalImport: ", build.IsLocalImport("."))
}
