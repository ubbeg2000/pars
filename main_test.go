package pars

import (
	"io"
	"os"
	"testing"
)

var (
	reader io.Reader
	dom    DOM
	ldom   LinearDOM
)

func TestMain(m *testing.M) {
	reader, _ = os.Open("./test/index.html")
	dom = ParseToDOM(reader)

	reader, _ = os.Open("./test/index.html")
	ldom = ParseToLinear(reader)

	m.Run()
}
