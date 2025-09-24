package stringx

// node represents a node in the Aho-Corasick automaton.
// It's used for efficient string matching and replacement operations.
type node struct {
	children map[rune]*node // Child nodes indexed by rune
	fail     *node          // Failure link for Aho-Corasick algorithm
	depth    int            // Depth of the node in the trie
	end      bool           // Whether this node represents the end of a word
}

// scope represents a matched substring with start and stop positions.
type scope struct {
	start int // Start position of the match
	stop  int // End position of the match
}

// add adds a word to the trie structure.
// Creates nodes as needed and marks the final node as an end node.
func (n *node) add(word string) {
	chars := []rune(word)
	if len(chars) == 0 {
		return
	}

	nd := n
	for i, char := range chars {
		if nd.children == nil {
			child := new(node)
			child.depth = i + 1
			nd.children = map[rune]*node{char: child}
			nd = child
		} else if child, ok := nd.children[char]; ok {
			nd = child
		} else {
			child := new(node)
			child.depth = i + 1
			nd.children[char] = child
			nd = child
		}
	}

	nd.end = true
}

// build constructs the failure links for the Aho-Corasick automaton.
// This enables efficient pattern matching by precomputing failure transitions.
func (n *node) build() {
	var nodes []*node
	for _, child := range n.children {
		child.fail = n
		nodes = append(nodes, child)
	}
	for len(nodes) > 0 {
		nd := nodes[0]
		nodes = nodes[1:]
		for key, child := range nd.children {
			nodes = append(nodes, child)
			cur := nd
			for cur != nil {
				if cur.fail == nil {
					child.fail = n
					break
				}
				if fail, ok := cur.fail.children[key]; ok {
					child.fail = fail
					break
				}
				cur = cur.fail
			}
		}
	}
}

// find searches for all patterns in the input text using the Aho-Corasick algorithm.
// Returns a slice of scopes representing all matched substrings.
func (n *node) find(chars []rune) []scope {
	var scopes []scope
	size := len(chars)
	cur := n

	for i := 0; i < size; i++ {
		child, ok := cur.children[chars[i]]
		if ok {
			cur = child
		} else {
			for cur != n {
				cur = cur.fail
				if child, ok = cur.children[chars[i]]; ok {
					cur = child
					break
				}
			}

			if child == nil {
				continue
			}
		}

		for child != n {
			if child.end {
				scopes = append(scopes, scope{
					start: i + 1 - child.depth,
					stop:  i + 1,
				})
			}
			child = child.fail
		}
	}

	return scopes
}
