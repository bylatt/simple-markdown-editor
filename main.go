package main

import (
	"github.com/gomarkdown/markdown"
	"syscall/js"
)

func convert(src, dest string) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		document := js.Global().Get("document")
		input := document.Call("getElementById", src)
		mdStr := input.Get("value").String()
		htmlByte := markdown.ToHTML([]byte(mdStr), nil, nil)
		html := string(htmlByte)
		output := document.Call("getElementById", dest)
		output.Set("innerHTML", html)
		return nil
	})
	js.Global().Get("document").Call("getElementById", "input").Call("addEventListener", "keyup", cb)
}

func main() {
	var done chan struct{}
	convert("input", "output")
	<-done
}
