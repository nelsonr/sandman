package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
)

// Directory and file permissions
const DIR_PERM = 0755
const FILE_PERM = 0644

// The "embed" directive creates a virtual file system of the
// "assets" directory and its contents during compilation so that
// it can be read later on when creating the sandbox.

//go:embed assets/*
var assetsFS embed.FS

func main() {
	args := os.Args

	// Show help message and exit when the provided number of arguments is wrong
	if len(args) != 2 {
		fmt.Printf("Usage: sandman NAME\n")
		os.Exit(0)
	}

	// Name of the directory to be created
	name := args[1]

	// Attempts to create the named directory
	// Fails and exits if there's already one with the same name
	err := os.Mkdir(name, DIR_PERM)
	if err != nil {
		fmt.Printf("A directory with the same name already exists: %s\n", name)
		os.Exit(0)
	}

	// Creates the files inside the sandbox directory
	fileEntries, err := fs.ReadDir(assetsFS, "assets")
	if err != nil {
		panic(err)
	}

	for i := range fileEntries {
		fileContents, err := fs.ReadFile(assetsFS, path.Join("assets", fileEntries[i].Name()))
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(path.Join(name, fileEntries[i].Name()), fileContents, FILE_PERM)
		if err != nil {
			panic(err)
		}
	}

	// Displays a success message
	fmt.Printf("Sandbox \"%s\" created with success!\n", name)
}
