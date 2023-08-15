package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type app struct {
}

func newApp() (*app, error) {
	a := &app{}

	ebiten.SetWindowTitle("絵文字サーチ")

	return a, nil
}

func (a *app) Update() error {
	return nil
}

func (a *app) Draw(screen *ebiten.Image) {

}

func (a *app) Layout(ww, wh int) (sw, sh int) {
	return ww, wh
}

func main() {
	app, err := newApp()
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(app); err != nil {
		panic(err)
	}
}
