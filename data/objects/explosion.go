package objects

import (
	t "github.com/nsf/termbox-go"
	// "galactixBash/lib/utils"
	"galactixBash/data/assets"
)

type Explosion struct {
    Base
    FrameIndex int
}

func NewExplosion(x,y int) Explosion {

	return Explosion {
		Base: Base {
			X: x,
			Y: y,
			ForeColor: 2,
			BgColor: 18,
			Body: 'Ϫ',
			Sleep: 0,
    		TimeToSleep: 2,
    		BodyMatrix: assets.Explosion1F1(),
		},
		FrameIndex: 1,
	}
}

func (e *Explosion) NextFrame() bool {

	switch {
		case e.FrameIndex == 1:
			e.Base.BodyMatrix =  assets.Explosion1F2()
		case e.FrameIndex == 2:
			e.Base.BodyMatrix =  assets.Explosion1F3()
		case e.FrameIndex == 3:
			e.Base.BodyMatrix =  assets.Explosion1F4()
		default:
			return false
	}

	e.FrameIndex++;
	return true
}

func (e *Explosion) Render() {
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