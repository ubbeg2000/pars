package pars

import (
	"fmt"
	"strings"
)

type DOMNode struct {
	TagName      string
	Text         string
	Parent       *DOMNode
	Children     []*DOMNode
	Attributes   map[string]string
	SelfEnclosed bool
}

func (d DOMNode) GetText() string {
	var retval string = ""

	for _, c := range d.Children {
		if c.TagName == "" && c.Text != "" {
			retval += fmt.Sprintf(" %s", c.Text)
		}
	}

	return strings.TrimSpace(retval)
}

func (d DOMNode) String() string {
	retval := fmt.Sprintf("<%s", d.TagName)

	for key, value := range d.Attributes {
		retval += fmt.Sprintf(" %s=\"%s\"", key, value)
	}

	if d.SelfEnclosed {
		retval += "/>"
	} else {
		retval += fmt.Sprintf(">%s</%s>", d.GetText(), d.TagName)
	}

	return retval
}

func (d DOMNode) Traverse(cb func(el DOMNode)) {
	var f func(node DOMNode)

	f = func(node DOMNode) {
		cb(node)
		for _, c := range node.Children {
			f(*c)
		}
	}

	f(d)
}

func (d DOMNode) GetElementByID(id string) *DOMNode {
	var retval *DOMNode = nil

	d.Traverse(func(el DOMNode) {
		if el.Attributes["id"] == id {
			retval = &el
		}
	})

	return retval
}

func (d DOMNode) GetElementsByTagName(tag string) []*DOMNode {
	var retval []*DOMNode = make([]*DOMNode, 0)

	d.Traverse(func(el DOMNode) {
		if el.TagName == tag {
			retval = append(retval, &el)
		}
	})

	return retval
}

func (d DOMNode) GetElementsByClassName(class string) []*DOMNode {
	var retval []*DOMNode = make([]*DOMNode, 0)

	d.Traverse(func(el DOMNode) {
		if strings.Contains(el.Attributes["class"], class) {
			retval = append(retval, &el)
		}
	})

	return retval
}

func (d DOMNode) QuerySelector(selector string) []*DOMNode {
	tagName, id, classList := ParseSelector(selector)
	retval := make([]*DOMNode, 0)
	classList = FormatClassList(classList)

	d.Traverse(func(el DOMNode) {
		var (
			tn string = tagName
			i  string = id
			cl string = classList
		)
		elClassList := FormatClassList(el.Attributes["class"])

		if tn == "" {
			tn = el.TagName
		}

		if i == "" {
			i = el.Attributes["id"]
		}

		if cl == "" {
			cl = elClassList
		}

		if tn == el.TagName {
			if i == el.Attributes["id"] {
				if cl == elClassList {
					retval = append(retval, &el)
				}
			}
		}
	})

	return retval
}

func (d *DOMNode) AppendChild(nd *DOMNode) {
	d.Children = append(d.Children, nd)
}

func (d DOMNode) Render() string {
	var f func(node DOMNode, gap string) string

	f = func(node DOMNode, gap string) string {
		var retval string = ""

		if node.TagName != "" {
			retval = fmt.Sprintf("%s<%s", gap, node.TagName)

			for key, value := range node.Attributes {
				retval += fmt.Sprintf(" %s=\"%s\"", key, value)
			}

			if node.SelfEnclosed {
				retval += "/>\n"
			} else {
				retval += ">\n"

				for _, c := range node.Children {
					retval += f(*c, gap+"  ")
				}

				retval += fmt.Sprintf("%s</%s>\n", gap, node.TagName)
			}

		} else {
			retval = fmt.Sprintf("%s%s\n", gap, node.Text)
		}

		return retval
	}

	return f(d, "")
}
