package manpage

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/yuin/goldmark"
)

func TestManpage(t *testing.T) {
	src, err := ioutil.ReadFile("manpage.md")
	if err != nil {
		t.Fatal(err)
	}

	md := goldmark.New(
		goldmark.WithRenderer(Renderer()),
	)
	if err := md.Convert(src, os.Stdout); err != nil {
		t.Fatal(err)
	}
}
