package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
)

const DIR_PERM = 0755
const FILE_PERM = 0644

//go:embed assets/index.html
var htmlFile []byte

//go:embed assets/style.css
var cssFile []byte

//go:embed assets/script.js
var jsFile []byte

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Printf("Usage: sandman SANDBOX_NAME\n")
		os.Exit(0)
	}

	name := args[1]

	err := os.Mkdir(name, DIR_PERM)
	if err != nil {
		fmt.Printf("A directory with the same name already exists: %s\n", name)
		os.Exit(0)
	}

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

	fmt.Printf("Sandbox \"%s\" created with success!", name)
}
