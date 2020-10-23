package rassemble

import "testing"

func TestJoin(t *testing.T) {
	testCases := []struct {
		name     string
		patterns []string
		expected string
	}{
		{
			name:     "empty",
			patterns: []string{},
			expected: "",
		},
		{
			name:     "empty literal",
			patterns: []string{""},
			expected: "(?:)",
		},
		{
			name:     "single literal",
			patterns: []string{"abc"},
			expected: "abc",
		},
		{
			name:     "multiple literals",
			patterns: []string{"abc", "def", "ghi"},
			expected: "abc|def|ghi",
		},
		{
			name:     "same literals",
			patterns: []string{"abc", "def", "abc", "def"},
			expected: "abc|def",
		},
		{
			name:     "same prefixes with different length",
			patterns: []string{"abcd", "abcf", "abc", "abce", "abcgh", "abdc"},
			expected: "ab(?:c(?:[d-f]|gh)?|dc)",
		},
		{
			name:     "same prefixes with same length",
			patterns: []string{"abcde", "abcfg", "abcgh"},
			expected: "abc(?:de|fg|gh)",
		},
		{
			name:     "same prefixes in increasing order",
			patterns: []string{"a", "ab", "abc", "abcd"},
			expected: "a(?:b(?:cd?)?)?",
		},
		{
			name:     "same prefixes in decreasing order",
			patterns: []string{"abcd", "abc", "ab", "a"},
			expected: "a(?:b(?:cd?)?)?",
		},
		{
			name:     "multiple prefix groups",
			patterns: []string{"abc", "ab", "abcd", "a", "bcd", "bcdef", "cdef", "cdeh"},
			expected: "a(?:b(?:cd?)?)?|bcd(?:ef)?|cde[fh]",
		},
		{
			name:     "merge literal to quest",
			patterns: []string{"abc(?:def)?", "abc"},
			expected: "abc(?:def)?",
		},
		{
			name:     "merge literal to star",
			patterns: []string{"abc(?:def)*", "abc"},
			expected: "abc(?:def)*",
		},
		{
			name:     "merge literal to plus",
			patterns: []string{"abc(?:def)+", "abc"},
			expected: "abc(?:def)*",
		},
		{
			name:     "merge literal to alternate",
			patterns: []string{"abc(?:de|f)", "abc"},
			expected: "abc(?:de|f)?",
		},
		{
			name:     "merge literal to concat",
			patterns: []string{"abca*b*", "abc"},
			expected: "abc(?:a*b*)?",
		},
		{
			name:     "merge literal to concat",
			patterns: []string{"abca*b*", "abcde"},
			expected: "abc(?:a*b*|de)",
		},
		{
			name:     "merge literal to quest with suffix",
			patterns: []string{"abc(?:def)?ghi", "abcd"},
			expected: "abc(?:(?:def)?ghi|d)",
		},
		{
			name:     "merge literal to alternate with same prefix",
			patterns: []string{"abcfd|def", "abcdef"},
			expected: "abc(?:fd|def)|def",
		},
		{
			name:     "merge literal to alternate with different prefix",
			patterns: []string{"abc|def", "ghi"},
			expected: "abc|def|ghi",
		},
		{
			name:     "character class",
			patterns: []string{"a", "1", "z", "2"},
			expected: "[12az]",
		},
		{
			name:     "character class with prefix",
			patterns: []string{"aa", "ab"},
			expected: "a[ab]",
		},
		{
			name:     "unmerge character class",
			patterns: []string{"a", "c", "e", "ab", "cd", "ef"},
			expected: "ab?|cd?|ef?",
		},
		{
			name:     "successive character class",
			patterns: []string{"aa", "ab", "ac"},
			expected: "a[a-c]",
		},
		{
			name:     "successive character class in random order",
			patterns: []string{"ac", "aa", "ae", "ab", "ad"},
			expected: "a[a-e]",
		},
		{
			name:     "numbers",
			patterns: []string{"1", "9", "2", "6", "3"},
			expected: "[1-369]",
		},
		{
			name:     "numbers 0 to 10",
			patterns: []string{"1", "9", "2", "6", "3", "7", "10", "8", "0", "5", "4"},
			expected: "[0-9]|10",
		},
		{
			name:     "numbers with prefix",
			patterns: []string{"a2", "a1", "a0", "a8", "a3", "a5", "a6", "a4", "a7", "a2", "a9", "a0", "a10"},
			expected: "a(?:[0-9]|10)",
		},
		{
			name:     "regexps",
			patterns: []string{"ab*c", "abc+", "bc+"},
			expected: "ab*c|abc+|bc+",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Join(tc.patterns)
			if err != nil {
				t.Fatalf("got an error: %s", err)
			}
			if got != tc.expected {
				t.Errorf("expected: %s, got: %s", tc.expected, got)
			}
		})
	}
}
