# TreeGen

A lightweight CLI tool built in Go to generate visual directory tree structures.

## Installation

Ensure you have Go installed, then run:

```bash
go install github.com/vassn/treegen@latest
```

## Usage

Generate a tree for any directory by providing the path:

```bash
treegen ./my-project
```

### Options

- `-c` , `--copy` - Copy the generated tree to the clipboard
- `-o` , `--output` - Save the result to a specific file path
- `-q` , `--quiet` - Suppress output to stdout

### Examples

**Copy to clipboard and suppress terminal output:**
```bash
treegen . -c -q
```

**Save output to a file:**
```bash
treegen ../another-dir -o output.txt
```

## Output Format

The tool generates a clean, Unicode-based visual representation:

```text
~/my-project/
├── cmd/
│   └── root.go
├── internal/
│   └── tree.go
├── go.mod
└── main.go
```

## Features

- Folders are listed before files, sorted alphabetically.
-  When using `--output`, missing parent directories are created automatically.
-  Fast copy to clipboard via the `-c` flag for easy pasting.
