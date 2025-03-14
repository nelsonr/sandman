# Sandman

Command line utility to create a quick web sandbox. Made with [Go](https://go.dev).

It creates a sandbox directory with 3 files: `index.html`, `style.css` and `script.js`.

## How to install

Clone the repo, then run `go build` to compile the project into a single executable, then move it to a directory of your preference.

Update your PATH environment variable to contain the path to the executable.

Alternatively, you can install it by running the `install.sh` shell script, which will place it on `~/.local/bin/sandman`.

## How to use

Run `sandman my-sandbox` on your terminal to create a sandbox named `my-sandbox`. Change the name accordingly to your liking.

## Development

This project uses the [Go](https://go.dev) programming language.

Relevant project files:

```
├── main.go
├── install.sh
└── assets
    ├── index.html
    ├── style.css
    └── script.js
```

The `main.go` file contains the main source code.

The `install.sh` shell script will compile and install the executable into `~/.local/bin/` directory (it assumes it already exists).

The files inside the `assets` directory are the files generated when the sandbox is created. Change the contens and compile again to create your own custom sandbox generator.
