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

	as := [][]rune{}
	var color int

	if utils.GetRandomNum(1,10) > 5 {
		as = assets.Ship1()
		color = 3
	} else {
		as = assets.Ship2()
		color = 4
	}

	return Enemy {
		Base: Base {
			X: x,
			Y: y,
			ForeColor: t.Attribute(color),
			BgColor: 18,
			Body: 'Ϫ',
			Sleep: 0,
    		TimeToSleep: utils.GetRandomNum(1,10),
    		BodyMatrix: as,
		},
	}
}