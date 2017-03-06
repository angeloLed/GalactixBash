package objects

import (
	t "github.com/nsf/termbox-go"
	"galactixBash/lib/utils"
	"galactixBash/data/assets"
)

type Enemy struct {
    Base
}

func NewEnemy(x,y int) Enemy {

	return Enemy {
		Base: Base {
			X: x,
			Y: y,
			ForeColor: t.ColorGreen,
			BgColor: 18,
			Body: 'Ϫ',
			Sleep: 0,
    		TimeToSleep: utils.GetRandomNum(1,10),
    		BodyMatrix: assets.Ship2(),
		},
	}
}