package scraper

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GetCommandOutput ..
func GetCommandOutput(cmd string) []byte {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	output, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	return output
}

// RemoveLineFeed ..
func RemoveLineFeed(output []byte) []byte {
	output = output[:len(output)-1]

	return output
}

// SplitLines ..
func SplitLines(output []byte) [][]byte {
	separator := []byte("\n")
	lines := bytes.Split(output, separator)

	return lines
}
