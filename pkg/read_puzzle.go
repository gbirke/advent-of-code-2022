package pkg

import (
	"bufio"
	"os"
)

func ReadPuzzle() ([]string, error) {
	puzzleFile, err := os.Open("puzzle.txt")

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(puzzleFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	puzzleFile.Close()

	return lines, nil
}
