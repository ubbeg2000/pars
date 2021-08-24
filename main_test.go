package pars

import (
	"fmt"
	"io"
	"os"
	"testing"
)

var (
	reader io.Reader
	dp     DOMParser
	lp     LinearParser
	dom    DOM
	ldom   LinearDOM
)

func TestMain(m *testing.M) {
	dp = NewDOMParser()
	lp = NewLinearParser()

	reader, _ = os.Open("./test/index.html")
	dom = dp.Parse(reader)

	reader, _ = os.Open("./test/index.html")
	ldom = lp.Parse(reader)

	ldom.Traverse(func(el Element) {
		fmt.Println(el)
	})

	m.Run()
}
