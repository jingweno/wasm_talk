package main

import (
	"syscall/js"
)

var Phaser = js.Global().Get("Phaser")

// FIXME: `js.Callback` doesn't properly bind to "this".
// Hack to use `eval` and `Function` to evaluate code.
// See https://github.com/golang/go/issues/27441
func main() {
	game := Phaser.Get("Game").New(500, 340, Phaser.Get("AUTO"), "gameDiv")
	js.Global().Set("game", game)

	bootState := newBootState()
	js.Global().Set("bootState", bootState)

	loadState := newLoadState()
	js.Global().Set("loadState", loadState)

	menuState := newMenuState()
	js.Global().Set("menuState", menuState)

	playState := newPlayState()
	js.Global().Set("playState", playState)

	startGame()
}

func newFunc(code string) js.Value {
	return js.Global().Get("Function").New(code)
}

func eval(code string) {
	js.Global().Call("eval", code)
}
