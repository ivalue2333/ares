package cryptos

import (
	"testing"
)

func TestSha256Str(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},  // Example test case
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},       // Empty input
		{"OpenAI", "8b7d1a3187ab355dc31bc683aaa71ab5ed217940c12196a9cd5f4ca984babfa4"}, // Custom test case
	}

	for _, tc := range testCases {
		output := Sha256Str(tc.input)
		if output != tc.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, output)
		}
	}
}
