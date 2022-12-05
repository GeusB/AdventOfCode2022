package tools

import (
	"bufio"
	"os"
)

func ReadFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
