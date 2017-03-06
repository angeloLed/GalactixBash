package objects

import (
	t "github.com/nsf/termbox-go"
	// "galactixBash/lib/utils"
	"galactixBash/data/assets"
)

type Player struct {
    Base
    Points int
    Life int
}

func NewPlayer(x,y int) Player {

	return Player{
		Life: 100,
		Base: Base{
			X: x,
			Y: y,
			ForeColor: t.ColorWhite,
			BgColor: 18,
			Body: '本',
			BodyMatrix: assets.Ship3(),
		},
	}
}

func (e *Player) Render() {
	if(e.Base.BodyMatrix != nil) {

		for y,valY := range e.Base.BodyMatrix {
			for x,valX := range valY {

				t.SetCell(e.Base.X+x, e.Base.Y+y, valX, e.Base.ForeColor, e.Base.BgColor)
			}
		}
	} else {
		t.SetCell(e.Base.X, e.Base.Y, e.Base.Body, e.Base.ForeColor, e.Base.BgColor)
	}
}