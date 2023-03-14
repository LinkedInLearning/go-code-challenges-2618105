package wordcountreader

import (
	"bufio"
	"os"
	"strings"
)

type FileContentInfo struct {
	Words, Characters, CharactersExcludingSpaces int
}

func ReadFileContent(filename string) FileContentInfo {
	words := 0
	characters := 0
	charactersExcludingSpaces := 0

	fileReader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fileReader.Close()

	fileScanner := bufio.NewScanner(fileReader)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		processLine(fileScanner.Text(), &words, &characters, &charactersExcludingSpaces)
	}

	return FileContentInfo{words, characters, charactersExcludingSpaces}
}

func processLine(line string, words, characters, charactersExcludingSpaces *int) {
	*characters += len(line)
	*charactersExcludingSpaces += len(strings.Replace(line, " ", "", -1))

	*words += len(strings.Fields(line))
}
