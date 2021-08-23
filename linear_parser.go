package pars

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type LinearParser struct{}

func NewLinearParser() LinearParser {
	return LinearParser{}
}

func (lp *LinearParser) Parse(res io.Reader) LinearDOM {
	var domElements []LinearDOMElement = make([]LinearDOMElement, 0)
	var newDom *LinearDOMElement
	var done bool = false

	z := html.NewTokenizer(res)

	for {
		tt := z.Next()

		if done {
			break
		}

		switch tt {
		case html.StartTagToken, html.SelfClosingTagToken:
			t, _ := z.TagName()
			tn := string(t)

			if newDom != nil {
				domElements = append(domElements, *newDom)
				newDom = nil
			}

			newDom = new(LinearDOMElement)
			newDom.TagName = tn
			newDom.SelfEnclosed = tt == html.SelfClosingTagToken
			newDom.Attributes = ParseAttributes(z)
			newDom.Text = ""

			if tt == html.SelfClosingTagToken {
				domElements = append(domElements, *newDom)
				newDom = nil
			}

		case html.TextToken:
			if newDom != nil {
				newDom.Text = strings.TrimSpace(string(z.Text()))
			}

		case html.EndTagToken:
			if newDom != nil {
				domElements = append(domElements, *newDom)
				newDom = nil
			}

		case html.ErrorToken:
			done = true
		}
	}

	return LinearDOM{Content: domElements}
}
