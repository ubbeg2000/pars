package pars

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type DOMParser struct{}

func NewDOMParser() DOMParser {
	return DOMParser{}
}

func (dp *DOMParser) Parse(reader io.Reader) DOM {
	var dom DOM = DOM{
		Tree:    nil,
		Ids:     make(map[string]*DOMNode),
		Classes: make(map[string][]*DOMNode),
		Tags:    make(map[string][]*DOMNode),
	}

	var domTree *DOMNode = nil
	var parentDom *DOMNode = domTree
	var newDom *DOMNode
	var done bool = false

	z := html.NewTokenizer(reader)

	for {
		tokenType := z.Next()

		if done {
			break
		}

		switch tokenType {
		case html.DoctypeToken:
			dom.Doctype = string(z.Raw())

		case html.StartTagToken, html.SelfClosingTagToken:
			t, _ := z.TagName()
			tn := string(t)
			attrs := ParseAttributes(z)

			newDom = new(DOMNode)
			newDom.TagName = tn
			newDom.Text = ""
			newDom.Attributes = attrs
			newDom.Children = []*DOMNode{}
			newDom.Parent = parentDom
			newDom.SelfEnclosed = tokenType == html.SelfClosingTagToken

			dom.RegisterToMaps(newDom)

			if domTree == nil {
				domTree = newDom
			} else {
				parentDom.AppendChild(newDom)
			}

			if tokenType == html.StartTagToken {
				parentDom = newDom
			}

		case html.TextToken:
			if parentDom != nil {
				parentDom.Text = strings.TrimSpace(string(z.Text()))
			}

		case html.EndTagToken:
			parentDom = parentDom.Parent

		case html.ErrorToken:
			done = true
		}
	}

	dom.Tree = domTree

	return dom
}
