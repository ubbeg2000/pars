package pars

import (
	"sort"
	"strings"

	"golang.org/x/net/html"
)

func ParseAttributes(z *html.Tokenizer) map[string]string {
	retval := map[string]string{}

	for {
		key, value, hasMore := z.TagAttr()

		if string(key) != "" {
			retval[string(key)] = string(value)
		}

		if !hasMore {
			break
		}
	}

	return retval
}

func ParseSelector(selector string) (string, string, string) {
	var (
		el      []string = strings.Split(selector, " ")
		tag     string   = ""
		id      string   = ""
		classes []string = make([]string, 0)
	)

	for _, e := range el {
		if e[0] == '.' {
			classes = append(classes, e[1:])
		} else if e[0] == '#' {
			id = e[1:]
		} else {
			tag = e
		}
	}

	return tag, id, strings.Join(classes, " ")
}

func FormatClassList(classList string) string {
	temp := strings.Split(classList, " ")
	sort.Strings(temp)
	return strings.Join(temp, " ")
}
