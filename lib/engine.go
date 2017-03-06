package lib

import (
	t "github.com/nsf/termbox-go"
	"galactixBash/data/objects"
	"galactixBash/lib/utils"
	"strconv"
)

func engine(g *game) {

	select {
		case ev := <-g.termboxQ:
			if ev.Type == t.EventKey {
				switch {
				case ev.Key == t.KeySpace:
					playerFire(g)
				case ev.Key == t.KeyArrowUp:
					g.player.Base.Y -= 1
				case ev.Key == t.KeyArrowDown:
					g.player.Base.Y += 1
				case ev.Key == t.KeyArrowLeft:
					g.player.Base.X -= 1
				case ev.Key == t.KeyArrowRight:
					g.player.Base.X += 1
				case ev.Key == t.KeyEsc:
					g.quit = true
				}
			}
		default:
	}

	fixOutOfBoard(g, &g.player.Base)
	firesEngine(g)
	enemiesEngine(g)
	explosionsEngine(g)

	if len(g.enemies) == 0 {
		nextLevel(g)
	}
}

//************************************************
//******************************** Generals

func nextLevel(g *game) {

	g.seed += utils.GetRandomNum(1,3);

	for i := 0; i < utils.GetRandomNum(g.seed * 2, g.seed * 5); i++ {
		enemy := objects.NewEnemy( utils.GetRandomNum(5,g.boardWidth-20), utils.GetRandomNum(1,5) )
		g.enemies = append(g.enemies, enemy )
	}
}

func checkOutOfBoard(g *game, b *objects.Base) bool {
	return ( b.X < 0 || b.X + b.GetWidth() > g.boardWidth -1 || b.Y < 0 || b.Y + b.GetHeight() > g.boardHeight -1 )
}

func fixOutOfBoard(g *game, b *objects.Base) {
	
	width := b.GetWidth()
	height := b.GetHeight()

	if b.X < 0 {
		b.X = 0
	}

	if b.X + width > g.boardWidth {
		b.X = g.boardWidth - width
	}

	if b.Y < 0 {
		b.Y = 0
	}

	if b.Y + height > g.boardHeight -1 {
		b.Y = g.boardHeight - height
	}
}

//************************************************
//******************************** Engine of player

func killPlayer(g *game) {
	addExplosion(g, g.player.Base.X, g.player.Base.Y)
}

func hurtPlayer(g *game, damage int) {
	g.player.Life -= damage
	if g.player.Life <= 0 {
		killPlayer(g)
	}
}

func playerFire(g *game) {
	if getPlayerFires(g) < g.maxPlayerFires {

		values := map[string]string{
			"x": strconv.Itoa(g.player.Base.X),
			"y": strconv.Itoa(g.player.Base.Y-1),
			"timeToSleep": strconv.Itoa(25),
			"color": strconv.Itoa(2),
			"friendly": strconv.FormatBool(true),
			"speed": strconv.Itoa(-1),
			"damage": strconv.Itoa(10),
		}

		fire(g, values)
	}
}

func getPlayerFires(g *game) int {
	count := 0
	for _,f := range g.fires {
		if f.Friendly {
			count++
		}
	}

	return count
}

//************************************************
//******************************** Engine of fires

func fire(g *game, values map[string]string) {
	g.fires = append(g.fires, objects.NewFire(values))
}

func firesEngine(g *game) {
	for i,_ := range g.fires {

		if(i < len(g.fires)) {

			if g.fires[i].Base.Sleep <= 0 {
				g.fires[i].Base.Y += g.fires[i].Speed
				g.fires[i].Base.Sleep += g.fires[i].TimeToSleep
			} else {
				g.fires[i].Base.Sleep --
			}

			if checkOutOfBoard(g, &g.fires[i].Base) { // far away
				killFire(g, i)

			} else if g.fires[i].Friendly { //is fire of player

				collision, enemyIndex := checkFireCollitionsWhitEnemies(g, &g.fires[i])
				if collision {
					g.player.Points++
					killEnemy(g, enemyIndex)
					killFire(g, i)
				}
			} else { //is fire of enemy
				if checkFireCollitionWhitPlayer(g, &g.fires[i]) {
					hurtPlayer(g, g.fires[i].Damage)
					killFire(g, i)
				}
			}
		}
	}
}

func checkFireCollitionsWhitEnemies(g *game, fire *objects.Fire) (bool, int) {
	
	for i,enemy := range g.enemies {
		if enemy.Base.CheckCollision(fire.Base.X, fire.Base.Y) {
			return true, i
		}
	}

	return false, -1
}

func checkFireCollitionWhitPlayer(g *game, fire *objects.Fire) (bool) {
	
	if g.player.Base.CheckCollision(fire.Base.X, fire.Base.Y) {
		return true
	}

	return false
}

func killFire(g *game, index int) {
	g.fires = append(g.fires[:index], g.fires[index+1:]...)
}


//**************************************************
//******************************** Engine of enemies

func enemiesEngine(g *game) {

	for i,_ := range g.enemies {

		if g.enemies[i].Base.Sleep > 0 {
			g.enemies[i].Base.Sleep --
		} else {
			opPercentage := utils.GetRandomNum(0,124)
			switch {
			case opPercentage >= 0 && opPercentage < 25:
				g.enemies[i].Base.Y -= 1
			case opPercentage >= 25 && opPercentage < 50:
				g.enemies[i].Base.Y += 1
			case opPercentage >= 50 && opPercentage < 75:
				g.enemies[i].Base.X -= 1
			case opPercentage >= 75 && opPercentage < 100:
				g.enemies[i].Base.X += 1
			case opPercentage >= 100 && opPercentage < 125:

				color := utils.GetRandomNum(170,175)
				damage := (((color/3) * 20)/100) + g.seed

				values := map[string]string{
					"x": strconv.Itoa(g.enemies[i].Base.X),
					"y": strconv.Itoa(g.enemies[i].Base.Y+1),
					"timeToSleep": strconv.Itoa(utils.GetRandomNum(5,25) - (g.seed)/3),
					"color": strconv.Itoa(color),
					"friendly": strconv.FormatBool(false),
					"speed": strconv.Itoa(1),
					"damage": strconv.Itoa(damage),
				}
				fire(g, values)
			}

			g.enemies[i].Base.Sleep += utils.GetRandomNum(5,150)

			fixOutOfBoard(g, &g.enemies[i].Base)
		}
	}
}

func killEnemy(g *game, index int) {
	addExplosion(g, g.enemies[index].Base.X, g.enemies[index].Base.Y)
	g.enemies = append(g.enemies[:index], g.enemies[index+1:]...)
}

//**************************************************
//******************************** Explosions
func explosionsEngine(g *game) {
	for i,_ := range g.explosions {

		if(i < len(g.explosions)) {
			
			if g.explosions[i].Base.Sleep > 0 {
				g.explosions[i].Base.Sleep --
			} else {
				g.explosions[i].Base.Sleep += 5
				
				if ! g.explosions[i].NextFrame() {
					g.explosions = append(g.explosions[:i], g.explosions[i+1:]...)
				}
			}
		}
	}
}

func addExplosion(g *game, x,y int) {
	g.explosions = append(g.explosions, objects.NewExplosion(x,y))
}