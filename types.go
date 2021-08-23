package pars

import (
	"fmt"
	"strings"
)

type Parser interface {
	Parse() ParsedDocument
}

type ParsedDocument interface {
	GetElementByID(id string) Element
	GetElementsByTagName(tag string) Element
	GetElementsByClassName(class string) Element
	Traverse(callback func(el Element))
}

type Element interface {
	GetTagName() string
	GetAttributes() map[string]string
	GetText() string
	String() string
}

type DOM struct {
	Doctype string
	Tree    *DOMNode
	Ids     map[string]*DOMNode
	Tags    map[string][]*DOMNode
	Classes map[string][]*DOMNode
}

type LinearDOM struct {
	Content []LinearDOMElement
}

type DOMNode struct {
	TagName      string
	Text         string
	Parent       *DOMNode
	Children     []*DOMNode
	Attributes   map[string]string
	SelfEnclosed bool
}

type LinearDOMElement struct {
	TagName      string
	Text         string
	Attributes   map[string]string
	SelfEnclosed bool
}

func (d LinearDOMElement) GetTagName() string {
	return d.TagName
}

func (d LinearDOMElement) GetAttributes() map[string]string {
	return d.Attributes
}

func (d LinearDOMElement) GetText() string {
	return d.Text
}

func (d LinearDOMElement) String() string {
	retval := fmt.Sprintf("<%s", d.TagName)

	for key, value := range d.Attributes {
		retval += fmt.Sprintf(" %s=\"%s\"", key, value)
	}

	if d.SelfEnclosed {
		retval += "/>"
	} else {
		retval += fmt.Sprintf(">%s</%s>", d.Text, d.TagName)
	}

	return retval
}

func (ld LinearDOM) Traverse(cb func(el Element)) {
	for _, el := range ld.Content {
		cb(el)
	}
}

func (ld LinearDOM) GetElementByID(id string) LinearDOMElement {
	for _, el := range ld.Content {
		if el.GetAttributes()["id"] == id {
			return el
		}
	}

	return LinearDOMElement{}
}

func (ld LinearDOM) GetElementsByClassName(class string) []LinearDOMElement {
	var retval []LinearDOMElement
	for _, el := range ld.Content {
		if strings.Contains(el.GetAttributes()["class"], class) {
			retval = append(retval, el)
		}
	}

	return retval
}

func (ld LinearDOM) GetElementsByTagName(tag string) []LinearDOMElement {
	var retval []LinearDOMElement
	for _, el := range ld.Content {
		if el.GetTagName() == tag {
			retval = append(retval, el)
		}
	}

	return retval
}

func (d DOM) GetElementByID(id string) *DOMNode {
	return d.Ids[id]
}

func (d DOM) GetElementsByClassName(class string) []*DOMNode {
	return d.Classes[class]
}

func (d DOM) GetElementsByTagName(tag string) []*DOMNode {
	return d.Tags[tag]
}

func (d DOM) Traverse(cb func(el Element)) {
	var f func(node DOMNode)

	f = func(node DOMNode) {
		if len(node.Children) == 0 {
			cb(node)
		} else {
			for _, c := range node.Children {
				f(*c)
			}
		}
	}

	f(*d.Tree)
}

func (d *DOM) Render() string {
	return d.Doctype + "\n" + d.Tree.Render()
}

func (d *DOM) RegisterToMaps(nd *DOMNode) {
	d.Tags[nd.TagName] = append(d.Tags[nd.TagName], nd)

	if len(nd.Attributes["id"]) > 0 {
		d.Ids[nd.Attributes["id"]] = nd
	}

	if len(nd.Attributes["class"]) > 0 {
		for _, c := range strings.Split(nd.Attributes["class"], " ") {
			d.Classes[c] = append(d.Classes[c], nd)
		}
	}
}

func (d DOMNode) GetTagName() string {
	return d.TagName
}

func (d DOMNode) GetAttributes() map[string]string {
	return d.Attributes
}

func (d DOMNode) GetText() string {
	return d.Text
}

func (d DOMNode) String() string {
	retval := fmt.Sprintf("<%s", d.TagName)

	for key, value := range d.Attributes {
		retval += fmt.Sprintf(" %s=\"%s\"", key, value)
	}

	if d.SelfEnclosed {
		retval += "/>"
	} else {
		retval += fmt.Sprintf(">%s</%s>", d.Text, d.TagName)
	}

	return retval
}

func (d *DOMNode) AppendChild(nd *DOMNode) {
	d.Children = append(d.Children, nd)
}

func (d *DOMNode) Render() string {
	var f func(node DOMNode, gap string) string
	f = func(node DOMNode, gap string) string {
		retval := fmt.Sprintf("%s<%s", gap, d.TagName)

		for key, value := range d.Attributes {
			retval += fmt.Sprintf(" %s=\"%s\"", key, value)
		}

		if d.SelfEnclosed {
			retval += "/>\n"
		} else {
			retval += ">\n"

			for _, c := range d.Children {
				retval += f(*c, gap+"  ")
			}

			if len(d.Text) > 0 {
				retval += fmt.Sprintf("%s  %s\n%s</%s>\n", gap, d.Text, gap, d.TagName)
			} else {
				retval += fmt.Sprintf("%s</%s>\n", gap, d.TagName)
			}

		}

		return retval
	}

	return f(*d, "  ")
}