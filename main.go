package main

import (
	"md2"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("md2", func(w *unison.Window) {
		w.Content().AddChild(md2.New().Layout())
	})
}
