package main

import "testing"

func TestExtractURL(t *testing.T) {
	testSet := map[string]string{
		"https://github.com/foo/bar.git":  "https://github.com/foo/bar",
		"https://github.com/foo/bar":      "https://github.com/foo/bar",
		"git@github.com:foo/bar.git":      "https://github.com/foo/bar",
		"git@github.com:foo/bar":          "https://github.com/foo/bar",
		"https://git.example.com/a/b.git": "https://git.example.com/a/b",
		"https://git.example.com/a/b":     "https://git.example.com/a/b",
		"git@git.example.com:a/b.git":     "https://git.example.com/a/b",
		"git@git.example.com:a/b":         "https://git.example.com/a/b",
	}

	for in, out := range testSet {
		actual, err := extractURL(in)
		if err != nil {
			t.Error(err)
		} else if actual != out {
			t.Errorf("'%s':\nexpected '%s'\n     got '%s'", in, out, actual)
		}
	}
}
