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

func ChunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func RemoveDuplicates(intSlice []string) []string {
	keys := make(map[string]bool)
	result := []string{}

	for _, x := range intSlice {
		if _, val := keys[x]; !val {
			keys[x] = true
			result = append(result, x)
		}
	}
	return result
}
