package wordcountreader

import "testing"

func TestReadFileContent(t *testing.T) {

	fi := ReadFileContent("simple_words.txt")

	if fi.Words != 18 {
		t.Errorf("Expected 18 words, got %d", fi.Words)
	}

	if fi.Characters != 86 {
		t.Errorf("Expected 86 characters, got %d", fi.Characters)
	}

	if fi.CharactersExcludingSpaces != 70 {
		t.Errorf("Expected 70 characters excluding spaces, got %d",
			fi.CharactersExcludingSpaces)
	}

}
