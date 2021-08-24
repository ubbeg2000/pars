package pars

import "strings"

type LinearDOM struct {
	Content []*LinearDOMElement
}

func (ld LinearDOM) Traverse(cb func(el LinearDOMElement)) {
	for _, el := range ld.Content {
		cb(*el)
	}
}

func (ld LinearDOM) GetElementByID(id string) *LinearDOMElement {
	for _, el := range ld.Content {
		if el.Attributes["id"] == id {
			return el
		}
	}

	return nil
}

func (ld LinearDOM) GetElementsByClassName(class string) []*LinearDOMElement {
	var retval []*LinearDOMElement
	for _, el := range ld.Content {
		if strings.Contains(el.Attributes["class"], class) {
			retval = append(retval, el)
		}
	}

	return retval
}

func (ld LinearDOM) GetElementsByTagName(tag string) []*LinearDOMElement {
	var retval []*LinearDOMElement
	for _, el := range ld.Content {
		if el.TagName == tag {
			retval = append(retval, el)
		}
	}

	return retval
}

func (d LinearDOM) QuerySelector(selector string) []*LinearDOMElement {
	tagName, id, classList := ParseSelector(selector)
	retval := make([]*LinearDOMElement, 0)
	classList = FormatClassList(classList)

	d.Traverse(func(el LinearDOMElement) {
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
