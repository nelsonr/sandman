package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
)

// Directory and file permissions
const DIR_PERM = 0755
const FILE_PERM = 0644

// The "embed" directives set the contents of the files
// into the variables below during compile time.
// Paths must be relative to this file.

//go:embed assets/index.html
var htmlFile []byte

//go:embed assets/style.css
var cssFile []byte

//go:embed assets/script.js
var jsFile []byte

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
	err = os.WriteFile(path.Join(name, "index.html"), []byte(htmlFile), FILE_PERM)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path.Join(name, "style.css"), []byte(cssFile), FILE_PERM)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path.Join(name, "script.js"), []byte(jsFile), FILE_PERM)
	if err != nil {
		panic(err)
	}

	// Displays a success message
	fmt.Printf("Sandbox \"%s\" created with success!", name)
}
