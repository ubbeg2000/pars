package pars

import "strings"

type DOM struct {
	Doctype string
	Tree    *DOMNode
	Ids     map[string]*DOMNode
	Tags    map[string][]*DOMNode
	Classes map[string][]*DOMNode
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

func (d DOM) Traverse(cb func(el DOMNode)) {
	var f func(node DOMNode)

	f = func(node DOMNode) {
		cb(node)
		for _, c := range node.Children {
			f(*c)
		}
	}

	f(*d.Tree)
}

func (d DOM) Render() string {
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
