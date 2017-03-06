package objects

import (
	t "github.com/nsf/termbox-go"
)

type Base struct {
    X, Y int
    Life int
    ForeColor t.Attribute
    BgColor t.Attribute
    Body rune
    Sleep int
    TimeToSleep int
    BodyMatrix [][]rune
}

func (e *Base) Render() {
	if(e.BodyMatrix != nil) {

		for y,valY := range e.BodyMatrix {
			for x,valX := range valY {
				t.SetCell(e.X+x, e.Y+y, valX, e.ForeColor, e.BgColor)
			}
		}
	} else {
		t.SetCell(e.X, e.Y, e.Body, e.ForeColor, e.BgColor)
	}
}

func (b *Base) GetWidth() int {
	return len(b.BodyMatrix[0])
}

func (b *Base) GetHeight() int {
	return len(b.BodyMatrix)
}

func (b *Base) CheckCollision(xx,yy int) bool {
	

	minX := b.X
	maxX := b.X + b.GetWidth() - 1

	minY := b.Y
	maxY := b.Y + b.GetHeight() - 1

	if( (xx >= minX && xx <= maxX) && ( yy >= minY && yy <=maxY )) {
		return true
	} else {
		return false
	}
}