package main

import "syscall/js"

func newBootState() js.Value {
	return js.ValueOf(map[string]interface{}{
		"preload": newFunc(`
game.load.image('progressBar', 'assets/progressBar.png');
		`),
		"create": js.Global().Get("Function").New(`
game.stage.backgroundColor = '#6AD7E5';

game.scale.scaleMode = Phaser.ScaleManager.SHOW_ALL;
document.body.style.backgroundColor = '#6AD7E5';

game.scale.minWidth = 250;
game.scale.minHeight = 170;
game.scale.maxWidth = 1000;
game.scale.maxHeight = 680;

game.scale.pageAlignHorizontally = true;
game.scale.pageAlignVertically = true;
game.scale.setScreenSize(true);

game.state.start('load');
		`),
	})
}
