// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

type predResp struct {
	predicate func(string) bool
	response  string
}

var whatever = "Whatever."
var bobSays = []predResp{
	{
		predicate: silencep,
		response:  "Fine. Be that way!",
	},
	{
		predicate: yellQuestionp,
		response:  "Calm down, I know what I'm doing!",
	},
	{
		predicate: yellp,
		response:  "Whoa, chill out!",
	},
	{
		predicate: questionp,
		response:  "Sure.",
	},
	{
		predicate: anythingp,
		response:  whatever,
	},
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.Trim(remark, "\n\r\t ")
	for _, presp := range bobSays {
		if presp.predicate(remark) {
			return presp.response
		}
	}
	// This is, technically, redundant, but here just for safety.
	return whatever
}

func yellp(s string) bool {
	hasLetter := false
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		hasLetter = true
		if unicode.IsLower(r) {
			return false
		}
	}
	if !hasLetter {
		return false
	}
	return true
}

func questionp(s string) bool {
	return s[len(s)-1] == '?'
}

func yellQuestionp(s string) bool {
	return yellp(s) && questionp(s)
}

func silencep(s string) bool {
	return s == ""
}

func anythingp(s string) bool {
	return true
}
