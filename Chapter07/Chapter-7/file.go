package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func file() {
	originalWorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("getting working directory: ", err)
	}
	fmt.Println("working directory: ", originalWorkingDir)
	examplesDir, err := createExamplesDir()
	if err != nil {
		fmt.Println("creating examples directory: ", err)
	}
	err = os.Chdir(examplesDir)
	if err != nil {
		fmt.Println("changing directory error:", err)
	}
	fmt.Println("changed working directory: ", examplesDir)
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("getting working directory: ", err)
	}
	fmt.Println("working directory: ", workingDir)
	createFiles()
	err = printFiles(workingDir)
	if err != nil {
		fmt.Printf("Error printing files in %s\n", workingDir)
	}
	err = os.Chdir(originalWorkingDir)
	if err != nil {
		fmt.Println("changing directory error: ", err)
	}
	fmt.Println("working directory: ", workingDir)
	symlink := filepath.Join(originalWorkingDir, "examplesLink")
	err = os.Symlink(examplesDir, symlink)
	if err != nil {
		fmt.Println("error creating symlink: ", err)
	}
	fmt.Printf("created symlink, %s, to %s\n", symlink, examplesDir)
	err = printFiles(symlink)
	if err != nil {
		fmt.Printf("Error printing files in %s\n", workingDir)
	}
	file := filepath.Join(examplesDir, "file1")
	linkedFile := filepath.Join(symlink, "file1")
	err = sameFileCheck(file, linkedFile)
	if err != nil {
		fmt.Println("unable to do same file check: ", err)
	}

	// cleanup
	err = os.Remove(symlink)
	if err != nil {
		fmt.Println("removing symlink error: ", err)
	}
	err = os.RemoveAll(examplesDir)
	if err != nil {
		fmt.Println("removing directory error: ", err)
	}
}

func createFiles() error {
	filename1 := "file1"
	filename2 := "file2"
	filename3 := "file3"
	f1, err := os.Create(filename1)
	if err != nil {
		return fmt.Errorf("error creating %s: %v\n", filename1, err)
	}
	defer f1.Close()
	f1.WriteString("abc")
	f2, err := os.Create(filename2)
	if err != nil {
		return fmt.Errorf("error creating %s: %v\n", filename2, err)
	}
	defer f2.Close()
	f2.WriteString("123")
	f3, err := os.Create(filename3)
	if err != nil {
		return fmt.Errorf("error creating %s: %v", filename3, err)
	}
	defer f3.Close()
	f3.WriteString("xyz")
	return nil
}

func createExamplesDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting user's home directory: %v\n", err)
	}
	fmt.Println("home directory: ", homeDir)
	examplesDir := filepath.Join(homeDir, "examples")
	err = os.Mkdir(examplesDir, os.FileMode(int(0777)))
	if err != nil {
		return "", fmt.Errorf("making directory error: %v\n", err)
	}
	fmt.Println("created: ", examplesDir)
	return examplesDir, nil
}

func printFiles(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read directory error: %s\n", err)
	}
	fmt.Printf("files in %s:\n", dir)

	for i, file := range files {
		fmt.Printf("  %v %v\n", i, file.Name())
	}
	return nil
}

func sameFileCheck(f1, f2 string) error {
	fileInfo0, err := os.Lstat(f1)
	if err != nil {
		return fmt.Errorf("getting fileinfo: %v", err)
	}
	fileInfo0Linked, err := os.Lstat(f2)
	if err != nil {
		return fmt.Errorf("getting fileinfo: %v", err)
	}
	isSameFile := os.SameFile(fileInfo0, fileInfo0Linked)
	if isSameFile {
		fmt.Printf("%s and %s are the same file.\n", fileInfo0.Name(), fileInfo0Linked.Name())
	} else {
		fmt.Printf("%s and %s are NOT the same file.\n", fileInfo0.Name(), fileInfo0Linked.Name())
	}
	return nil
}
