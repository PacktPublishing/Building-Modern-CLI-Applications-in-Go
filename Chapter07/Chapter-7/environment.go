package main

import (
	"fmt"
	"os"
	"strings"
)

func environment() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working directory:", err)
	}
	fmt.Println("retrieved working directory: ", dir)

	fmt.Println("setting WORKING_DIR to", dir)
	err = os.Setenv("WORKING_DIR", dir)
	if err != nil {
		fmt.Println("error setting working directory:", err)
	}
	fmt.Println(os.ExpandEnv("WORKING_DIR=${WORKING_DIR}"))

	environmentMap := map[string]string{}
	fmt.Printf("There are %d environment variables.\n", len(os.Environ()))
	for _, envar := range os.Environ() {
		keyValue := strings.Split(envar, "=")
		key := keyValue[0]
		value := keyValue[1]
		environmentMap[key] = value
	}

	fmt.Printf("key=%s, value=%s\n", "WORKING_DIR", environmentMap["WORKING_DIR"])

	fmt.Println("unsetting WORKING_DIR")
	err = os.Unsetenv("WORKING_DIR")
	if err != nil {
		fmt.Println("error unsetting working directory:", err)
	}
	fmt.Println(os.ExpandEnv("WORKING_DIR=${WORKING_DIR}"))
}
