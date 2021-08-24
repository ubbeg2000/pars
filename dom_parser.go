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
	var (
		dom DOM = DOM{
			Tree:    nil,
			Ids:     make(map[string]*DOMNode),
			Classes: make(map[string][]*DOMNode),
			Tags:    make(map[string][]*DOMNode),
		}
		domTree   *DOMNode = nil
		parentDom *DOMNode = domTree
		newDom    *DOMNode = nil
		done      bool     = false
	)

	z := html.NewTokenizer(reader)

	for {
		tokenType := z.Next()

		if done {
			break
		}

		switch tokenType {
		case html.DoctypeToken:
			dom.Doctype = string(z.Raw())

		case html.StartTagToken, html.SelfClosingTagToken, html.TextToken:
			t, _ := z.TagName()
			tn := string(t)
			attrs := ParseAttributes(z)

			if tokenType == html.StartTagToken {
				newDom = new(DOMNode)
				newDom.TagName = tn
				newDom.Text = ""
				newDom.Attributes = attrs
				newDom.Children = []*DOMNode{}
				newDom.Parent = parentDom
				newDom.SelfEnclosed = false

				if domTree == nil {
					domTree = newDom
				} else {
					parentDom.AppendChild(newDom)
				}

				parentDom = newDom

				dom.RegisterToMaps(newDom)
			}

			if tokenType == html.SelfClosingTagToken {
				newDom = new(DOMNode)
				newDom.TagName = tn
				newDom.Text = ""
				newDom.Attributes = attrs
				newDom.Children = []*DOMNode{}
				newDom.Parent = parentDom
				newDom.SelfEnclosed = true

				parentDom.AppendChild(newDom)

				dom.RegisterToMaps(newDom)
			}

			if tokenType == html.TextToken {
				txt := strings.TrimSpace(string(z.Text()))

				if len(txt) > 0 {
					newDom = new(DOMNode)
					newDom.TagName = tn
					newDom.Text = txt
					newDom.Attributes = attrs
					newDom.Children = []*DOMNode{}
					newDom.Parent = parentDom
					newDom.SelfEnclosed = true

					parentDom.AppendChild(newDom)
				}
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
