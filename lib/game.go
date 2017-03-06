package lib

import (
	t "github.com/nsf/termbox-go"
	"galactixBash/data/objects"
	"galactixBash/lib/utils"
	"time"
)

const (
	animationSpeed = 20 * time.Millisecond
	bgColor = 18
)

type game struct {
	seed int
	maxPlayerFires int
	termboxQ chan t.Event
	gameQ chan string
	quit bool
	boardHeight int
	boardWidth int
	fires []objects.Fire
	enemies []objects.Enemy
	explosions []objects.Explosion
	player objects.Player
	logString string
}

func NewGame() game {
	return game{}
}

func (g *game) Run() {

	//initialize game
	g.initialize()

	//initialize Termbox
	t.SetOutputMode(t.Output256)
	err := t.Init()
	if err != nil {
		panic(err)
	}
	defer t.Close()
	g.termboxQ = make(chan t.Event)
	go func() {
		for {
			g.termboxQ <- t.PollEvent()
		}
	}()

	// !! loop of hell !!
	for {
		engine(g)
		if g.quit {
			return
		}
		render(g)
		time.Sleep(animationSpeed)
	}
}

func (g *game) initialize() {

	g.boardHeight = 30
	g.boardWidth = 100
	g.maxPlayerFires = 5
	g.seed = 5

	//initialize player
	g.player = objects.NewPlayer( utils.GetRandomNum(5,g.boardWidth-5), utils.GetRandomNum(g.boardHeight-5, g.boardHeight-1))

	//initialize enemies
	for i := 0; i < utils.GetRandomNum(7, 40); i++ {
		enemy := objects.NewEnemy( utils.GetRandomNum(5,g.boardWidth-20), utils.GetRandomNum(1,5) )
		g.enemies = append(g.enemies, enemy )
	}
}