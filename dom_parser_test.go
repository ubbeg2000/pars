package pars

import (
	"testing"
)

func TestDOMTraversal(t *testing.T) {
	i := 0
	dom.Traverse(func(el DOMNode) {
		if el.TagName != "" {
			i++
		}
	})

	if i != 17 {
		t.Errorf("traversed %d nodes, should be %d", i, 10)
	}
}

func TestDOMNodeTraversal(t *testing.T) {
	i := 0
	dom.GetElementByID("dom-node").Traverse(func(el DOMNode) {
		if el.TagName != "" {
			i++
		}
	})

	if i != 10 {
		t.Errorf("traversed %d nodes, should be %d", i, 10)
	}
}

func TestDOMNodeTagSearch(t *testing.T) {
	res := dom.GetElementsByTagName("body")[0].GetElementsByTagName("div")
	if len(res) != 8 {
		t.Error("failed to search tags")
	}
}

func TestDOMNodeClassNameSearch(t *testing.T) {
	res := dom.GetElementsByTagName("body")[0].GetElementsByClassName("a-class")
	if len(res) != 3 {
		t.Error("failed to search classes")
	}
}

func TestDOMNodeIDSearch(t *testing.T) {
	res := dom.GetElementsByTagName("body")[0].GetElementByID("an-id")
	if res.TagName != "div" || res.GetText() != "div with id asdf" {
		t.Error(res.GetText())
	}
}

func TestDOMNodeQuerySelector(t *testing.T) {
	res := dom.Tree.QuerySelector("div .class1 .class2")
	if len(res) != 1 {
		t.Error(len(res))
	}
}
