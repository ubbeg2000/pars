package pars

import "fmt"

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
	var retval string = ""

	if d.TagName != "" {
		retval = fmt.Sprintf("<%s", d.TagName)

		for key, value := range d.Attributes {
			retval += fmt.Sprintf(" %s=\"%s\"", key, value)
		}

		if d.SelfEnclosed {
			retval += "/>"
		} else {
			retval += fmt.Sprintf(">%s</%s>", d.Text, d.TagName)
		}

		return retval
	} else {
		return d.Text
	}
}
