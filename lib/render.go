package lib

import (
	t "github.com/nsf/termbox-go"
	"time"
	"strconv"
)


func render(g *game) {
	t.Clear(t.Attribute(1), t.Attribute(1))

	renderBoard(g)
	renderPalette(g)

	//log
	renderString(1, 32, t.ColorBlack, t.ColorWhite, time.Now().Format(time.RFC1123))
	renderString(1, 33, t.ColorBlack, t.ColorWhite, "Enemies : " + strconv.Itoa(len(g.enemies)))
	renderString(1, 34, t.ColorBlack, t.ColorWhite, "Fires : " + strconv.Itoa(len(g.fires)))
	renderString(40, 32, t.ColorBlack, t.ColorWhite, "Player position : " + strconv.Itoa(g.player.Base.X) + ", " + strconv.Itoa(g.player.Base.Y))
	renderString(40, 33, t.ColorBlack, t.ColorWhite, "Player points : " + strconv.Itoa(g.player.Points))
	renderString(40, 34, t.ColorBlack, t.ColorWhite, "Seed: " + strconv.Itoa(g.seed))

	//print player
	renderPlayerLife(g.player.Life)
	if g.player.Life > 0 {
		g.player.Render()
	}

	//print enemies
	for _,enemy := range g.enemies {
		enemy.Render()
	}

	//print fires
	for _,fire := range g.fires {
		fire.Render()
	}

	//print explosions
	for _,explosion := range g.explosions {
		explosion.Render()
	}

	//render game over
	if g.player.Life <= 0 {
		renderGameOver(g)
	}

	t.Flush()
}

func renderGameOver(g *game) {

	String := "!! GAME OVER !!"
	renderString((g.boardWidth / 2 ) - (len(String)/2) , g.boardHeight/2, t.ColorBlack, t.ColorWhite, String)
}

func renderPlayerLife(life int) {
	renderString(1, 30, t.Attribute(19), t.Attribute(40), "Life : ")

	// life : 100 = x : lenghtBar
	prop := (life * 90)/100
	if prop <= 0 && life > 0 {
		prop = 1
	}

	//color
	color := 40
	switch {
	case life >= 50 && life < 85:
		color = 80
	case life >= 25 && life < 50:
		color = 30
	case life <= 25:
		color = 150
	}


	for i := 0; i < prop; i++ {
		t.SetCell(i+10, 30, ' ', t.Attribute(color), t.Attribute(color))
	}
}

func renderBoard(g *game) {
	for y := 0; y < g.boardHeight; y++ {
		for x := 0; x < g.boardWidth; x++ {
			t.SetCell(x, y, ' ', t.Attribute(bgColor), t.Attribute(bgColor))
		}
	}
}

func renderPalette(g *game) {
	i := 0
	for y := g.boardHeight+7; y < g.boardHeight*2; y++ {
		for x := 0; x < g.boardWidth; x++ {
			i++
			if(i>254) {
				return
			}
			t.SetCell(x, y, 'a', t.Attribute(16), t.Attribute(i))
			t.SetCell(x+1, y, ' ', t.Attribute(16), t.Attribute(i))
		}
	}
}

func renderString(x, y int, fg, bg t.Attribute, msg string) {
	for _, c := range msg {
		t.SetCell(x, y, c, fg, bg)
		x++
	}
}