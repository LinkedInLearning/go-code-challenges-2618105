package wordwrapper

import "strings"

func Wrap(text string, width int) string {
	wrappedText := ""
	words := strings.Split(text, " ")

	currentLine := ""

	for _, word := range words {
		if len(currentLine) > 0 && len(currentLine)+len(word) > width {
			// Word wrap to next line
			wrappedText = wrappedText[:len(wrappedText)-1]
			wrappedText += "\n"
			currentLine = ""
		}
		wrappedText += word
		wrappedText += " "

		currentLine += word
		currentLine += " "
	}

	return strings.TrimSpace(wrappedText)
}
