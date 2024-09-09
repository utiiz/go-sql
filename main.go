package main

import (
	"os"

	"github.com/fumiama/go-docx"
)

func main() {
	w := docx.New().WithDefaultTheme()
	// add new paragraph
	para1 := w.AddParagraph()
	// add text
	para1.AddText("test").AddTab()
	para1.AddText("size").Size("44").AddTab()
	f, err := os.Create("generated.docx")
	// save to file
	if err != nil {
		panic(err)
	}
	_, err = w.WriteTo(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
