package recordjar

import (
	"bufio"
	"bytes"
	"strings"
)

// Parse will parse content into slice.
func Parse(content []byte) []map[string][]string {
	s := &scanner{
		s:    bufio.NewScanner(bytes.NewReader(content)),
		used: true,
	}
	p := &parser{
		data:       make([]map[string][]string, 0),
		currentKey: "",
		isNewItem:  true,
	}

	for s.Scan() {
		text := s.Text()
		if isComment(text) {
			p.isNewItem = true
			continue
		}

		if strings.HasPrefix(text, "  ") {
			p.parse([]string{text})
			continue
		}
		p.parse(strings.Split(text, ":"))
	}

	return p.data
}
