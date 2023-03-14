package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {

	if !IsPalindrome("abba") {
		t.Errorf("Expected abba to be a palindrome")
	}

	if !IsPalindrome("racecar") {
		t.Errorf("Expected racecar to be a palindrome")
	}

	if IsPalindrome("abc") {
		t.Errorf("Expected abc to not be a palindrome")
	}

}
