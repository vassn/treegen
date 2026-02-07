package internal

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type node struct {
	entry    os.DirEntry
	contents []node
}

func buildNodes(path string) ([]node, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var nodes []node
	for _, entry := range entries {
		child := node{entry: entry}

		if entry.IsDir() {
			childPath := filepath.Join(path, entry.Name())
			children, err := buildNodes(childPath)
			if err == nil {
				child.contents = children
			}
		}
		nodes = append(nodes, child)
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].entry.IsDir() != nodes[j].entry.IsDir() {
			return nodes[i].entry.IsDir()
		}
		return nodes[i].entry.Name() < nodes[j].entry.Name()
	})

	return nodes, nil
}

func buildTree(sb *strings.Builder, n node, prefix string) {
	for i, child := range n.contents {
		name := child.entry.Name()
		if child.entry.IsDir() {
			name += "/"
		}

		connector := "├── "
		if i == len(n.contents)-1 {
			connector = "└── "
		}

		sb.WriteString(prefix)
		sb.WriteString(connector)
		sb.WriteString(name)
		sb.WriteString("\n")

		newPrefix := prefix
		if i == len(n.contents)-1 {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}

		buildTree(sb, child, newPrefix)
	}
}

func RenderTree(path string) string {
	var sb strings.Builder

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err.Error()
	}
	rootName := filepath.Base(absPath)

	info, err := os.Lstat(path)
	if err != nil {
		return err.Error()
	}

	root := node{
		entry: fs.FileInfoToDirEntry(info),
	}

	children, err := buildNodes(path)
	if err != nil {
		return err.Error()
	}
	root.contents = children

	sb.WriteString("~/")
	sb.WriteString(rootName)
	sb.WriteString("/")
	sb.WriteString("\n")

	buildTree(&sb, root, "")

	return sb.String()
}