package diamond

import (
	"strings"
)

const StartingLetter = 'A'

func Diamond(letter rune) string {

	diamondRows := rowsTillMiddle(letter)
	mirroredRows := mirrorRows(diamondRows)

	diamondRows = append(diamondRows, mirroredRows[1:]...)

	return strings.Join(diamondRows, "\n")
}

func rowsTillMiddle(letter rune) []string {
	var rows []string
	numberOfLetters := letter - StartingLetter

	for i := 0; i <= int(numberOfLetters); i++ {
		rows = append(rows, createRow(i, letter))
	}
	return rows
}

func createRow(index int, finalLetter rune) string {
	currentLetter := StartingLetter + rune(index)

	paddingSize := int(finalLetter - currentLetter)
	insidePaddingSize := int(currentLetter-StartingLetter)*2 - 1
	if insidePaddingSize < 0 {
		insidePaddingSize = 0
	}

	outerPadding := strings.Repeat(" ", paddingSize)
	innerPadding := strings.Repeat(" ", insidePaddingSize)

	row := outerPadding + string(currentLetter) + innerPadding
	if currentLetter != StartingLetter {
		row += string(currentLetter)
	}
	row += outerPadding

	return row
}

func mirrorRows(rows []string) []string {
	mirroredRows := make([]string, len(rows))
	for i, row := range rows {
		mirroredRows[len(rows)-i-1] = row
	}
	return mirroredRows
}
