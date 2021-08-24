package pars

import (
	"testing"
)

func TestLinearDOMTraversal(t *testing.T) {
	i := 0
	ldom.Traverse(func(el LinearDOMElement) {
		i++
	})

	if i != 21 {
		t.Errorf("Traversed %d elements, should be %d\n", i, 21)
	}
}

func TestLinearDOMByTagSearch(t *testing.T) {
	res := ldom.GetElementsByTagName("div")
	if len(res) != 8 {
		t.Errorf("Found %d elements, should be %d\n", len(res), 8)
	}
}

func TestLinearDOMByClassSearch(t *testing.T) {
	res := ldom.GetElementsByClassName("a-class")
	if len(res) != 3 {
		t.Errorf("Found %d elements, should be %d\n", len(res), 3)
	}
}

func TestLinearDOMByIDSearch(t *testing.T) {
	res := ldom.GetElementByID("an-id")
	if res.TagName != "div" {
		t.Error("failed to search id")
	}
}
