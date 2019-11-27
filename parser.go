package recordjar

import (
	"fmt"
	"log"
	"strings"
)

type parser struct {
	data []map[string][]string

	currentKey string
	isNewItem  bool
}

func (p *parser) parse(data []string) {
	var m map[string][]string
	if p.isNewItem {
		m = make(map[string][]string)
		p.data = append(p.data, m)
		p.isNewItem = false
	} else {
		m = p.data[len(p.data)-1]
	}
	log.Printf("%v", data)

	if len(data) > 2 && len(data) < 1 {
		panic(fmt.Errorf("parse data failed: %s", data))
	}

	// Handle multiline attr, for example: Comments
	if len(data) != 2 {
		key := p.currentKey
		// Add a space before every line.
		m[key][len(m[key])-1] += " " + strings.TrimSpace(data[0])
		return
	}

	key := data[0]
	// Update current key every time get a new key
	p.currentKey = key

	m[key] = append(m[key], strings.TrimSpace(data[1]))
}
