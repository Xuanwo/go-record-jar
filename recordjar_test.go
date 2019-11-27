package recordjar

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		expected []map[string][]string
	}{
		{"simple item", "case1", []map[string][]string{
			{
				"File-Date": []string{"2019-09-16"},
			},
		}},
		{"simple item with comments", "case2", []map[string][]string{
			{
				"File-Date": []string{"2019-09-16"},
			},
			{
				"Type":        []string{"language"},
				"Subtag":      []string{"aa"},
				"Description": []string{"Afar"},
				"Added":       []string{"2005-10-16"},
			},
		}},
		{"multiline comments", "case3", []map[string][]string{
			{
				"Type":        []string{"region"},
				"Subtag":      []string{"GB"},
				"Description": []string{"United Kingdom"},
				"Added":       []string{"2005-10-16"},
				"Comments":    []string{"as of 2006-03-29 GB no longer includes the Channel Islands and Isle of Man; see GG, JE, IM"},
			},
		}},
		{"multi exist attr", "case4", []map[string][]string{
			{
				"Type":        []string{"region"},
				"Subtag":      []string{"SZ"},
				"Description": []string{"Eswatini", "eSwatini", "Swaziland"},
				"Added":       []string{"2005-10-16"},
			},
		}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			content, err := ioutil.ReadFile(path.Join("testdata", tt.filename))
			if err != nil {
				t.Fatal(err)
			}

			assert.EqualValues(t, tt.expected, Parse(content))
		})
	}
}
