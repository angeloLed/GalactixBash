package objects

import (
	t "github.com/nsf/termbox-go"
	// "galactixBash/lib/utils"
	"galactixBash/data/assets"
	"strconv"
	"log"
)

type Fire struct {
    Base
    Speed int
    Friendly bool
    Damage int
}

func NewFire(values map[string]string) Fire {

	x, err := strconv.Atoi(values["x"])
	y, err := strconv.Atoi(values["y"])
	color, err := strconv.Atoi(values["color"])
	speed, err := strconv.Atoi(values["speed"])
	timeToSleep, err := strconv.Atoi(values["timeToSleep"])
	friendly, err := strconv.ParseBool(values["friendly"])
	damage, err := strconv.Atoi(values["damage"])

    if err != nil {
    	 log.Fatal(err)
    }

	//create fire
	return Fire {
		Base: Base {
			X: x,
			Y: y,
			ForeColor: t.Attribute(color),
			BgColor: t.Attribute(18),
			Body: '|',
			Sleep: 0,
    		TimeToSleep: timeToSleep,
    		BodyMatrix: assets.Laser1(),
		},
		Speed: speed,
		Friendly: friendly,
		Damage: damage,
	}
}

func (f *Fire) Render() {
	if(f.Base.BodyMatrix != nil) {
		for y,valY := range f.Base.BodyMatrix {
			for x,valX := range valY {
				t.SetCell(f.Base.X+x, f.Base.Y+y, valX, f.Base.ForeColor, f.Base.BgColor)
			}
		}
	} else {
		t.SetCell(f.Base.X, f.Base.Y, f.Base.Body, f.Base.ForeColor, f.Base.BgColor)
	}
}