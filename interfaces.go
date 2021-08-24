package pars

type Element interface {
	GetTagName() string
	GetAttributes() map[string]string
	GetText() string
	String() string
}

type Parser interface {
	Parse() ParsedDocument
}

type ParsedDocument interface {
	GetElementByID(id string) Element
	GetElementsByTagName(tag string) Element
	GetElementsByClassName(class string) Element
	Traverse(callback func(el Element))
}
