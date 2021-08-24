package pars

import "strings"

type LinearDOM struct {
	Content []LinearDOMElement
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
