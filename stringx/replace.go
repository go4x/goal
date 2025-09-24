package stringx

import (
	"sort"
	"strings"
)

// replaceTimes is the maximum number of replacement iterations.
// Replace more than once to avoid overlapped keywords after replace.
// Only try 2 times to avoid too many or infinite loops.
const replaceTimes = 2

type (
	// Replacer interface wraps the Replace method for string replacement operations.
	Replacer interface {
		Replace(text string) string
	}

	// replacer implements the Replacer interface using Aho-Corasick algorithm.
	// It provides efficient multi-pattern string replacement.
	replacer struct {
		*node
		mapping map[string]string // Maps patterns to their replacements
	}
)

// NewReplacer creates a new Replacer with the given pattern-to-replacement mapping.
// The replacer uses the Aho-Corasick algorithm for efficient multi-pattern matching.
func NewReplacer(mapping map[string]string) Replacer {
	rep := &replacer{
		node:    new(node),
		mapping: mapping,
	}
	for k := range mapping {
		rep.add(k)
	}
	rep.build()

	return rep
}

// Replace performs string replacement using the configured patterns.
// It iterates up to replaceTimes to handle overlapping patterns.
func (r *replacer) Replace(text string) string {
	for i := 0; i < replaceTimes; i++ {
		var replaced bool
		if text, replaced = r.doReplace(text); !replaced {
			return text
		}
	}

	return text
}

// doReplace performs a single replacement pass on the text.
// Returns the modified text and a boolean indicating if any replacements were made.
func (r *replacer) doReplace(text string) (string, bool) {
	chars := []rune(text)
	scopes := r.find(chars)
	if len(scopes) == 0 {
		return text, false
	}

	// Sort scopes by start position, with longer matches taking precedence for same start
	sort.Slice(scopes, func(i, j int) bool {
		if scopes[i].start < scopes[j].start {
			return true
		}
		if scopes[i].start == scopes[j].start {
			return scopes[i].stop > scopes[j].stop
		}
		return false
	})

	var buf strings.Builder
	var index int
	for i := 0; i < len(scopes); i++ {
		scp := &scopes[i]
		if scp.start < index {
			continue
		}

		buf.WriteString(string(chars[index:scp.start]))
		buf.WriteString(r.mapping[string(chars[scp.start:scp.stop])])
		index = scp.stop
	}
	if index < len(chars) {
		buf.WriteString(string(chars[index:]))
	}

	return buf.String(), true
}
