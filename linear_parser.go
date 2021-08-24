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
	var (
		domElements []LinearDOMElement = make([]LinearDOMElement, 0)
		newDom      *LinearDOMElement  = nil
		// prevDom     *LinearDOMElement  = nil
		done bool = false
	)

	z := html.NewTokenizer(res)

	for {
		tokenType := z.Next()

		if done {
			break
		}

		switch tokenType {
		case html.StartTagToken:
			t, _ := z.TagName()
			tn := string(t)

			newDom = new(LinearDOMElement)
			newDom.TagName = tn
			newDom.SelfEnclosed = tokenType == html.SelfClosingTagToken
			newDom.Attributes = ParseAttributes(z)
			newDom.Text = ""

			domElements = append(domElements, *newDom)

		case html.SelfClosingTagToken:
			t, _ := z.TagName()
			tn := string(t)

			newDom = new(LinearDOMElement)
			newDom.TagName = tn
			newDom.SelfEnclosed = tokenType == html.SelfClosingTagToken
			newDom.Attributes = ParseAttributes(z)
			newDom.Text = ""

			domElements = append(domElements, *newDom)

		case html.TextToken:
			txt := strings.TrimSpace(string(z.Text()))
			if len(txt) > 0 {
				newDom = new(LinearDOMElement)
				newDom.TagName = ""
				newDom.SelfEnclosed = true
				newDom.Text = txt

				domElements = append(domElements, *newDom)
			}

		case html.EndTagToken:
			newDom = nil

		case html.ErrorToken:
			done = true
		}
	}

	return LinearDOM{Content: domElements}
}
