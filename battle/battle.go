package battle

import (
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

var (
	score    = 0
	words    = InitializeQuestion()
	index    = 0
	yourTime = 0.0

	gainGold = 0
	lostGold = 0

	AttackCount = 3.0
	tempCount   = 0.0
	lock        = false
	lock2       = false
	pressEnter  = false
)

func BattleTypingV1(win *pixelgl.Window, player *player.PlayerStatus, enemy *enemy.EnemyStatus, elapsed time.Duration) myGame.GameState {
	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	if typed != "" {
		if typed[0] == temp[index] && index < len(question) {
			index++
			collectType++

			enemy.HP -= player.OP
			player.SP += player.BaseSP

			if index == len(question) {
				index = 0
				score++

				if score == len(words) {
					myGame.CurrentGS = myGame.EndScreen
					yourTime = float64(elapsed.Seconds())
				}
			}
		} else {
			missType++
		}
	}

	BattleTypingSkill(win, player, enemy)
	myGame.CurrentGS = DeathFlug(player, enemy, elapsed, myGame.CurrentGS)
	return myGame.CurrentGS
}

func BattleTypingSkill(win *pixelgl.Window, player *player.PlayerStatus, enemy *enemy.EnemyStatus) {
	if win.JustPressed(pixelgl.KeySpace) {
		log.Println("Skill!!!")
		if player.SP == 50 {
			player.SP = 0
			if player.Job == "Warrior" {
				enemy.HP -= 15
			} else if player.Job == "Priest" {
				//TODO 僧侶の回復スキル
			} else if player.Job == "Wizard" {
				//TODO 魔法使いの時止めスキル
			}
		} else {
			log.Println("skillポイントが足りない")
		}
	}
}

func DeathFlug(player *player.PlayerStatus, enemy *enemy.EnemyStatus, elapsed time.Duration, currentGameState myGame.GameState) myGame.GameState {
	if player.HP <= 0 {
		yourTime = float64(elapsed.Seconds())
		min := int(float64(enemy.Gold) * 0.7)
		max := int(float64(enemy.Gold) * 1.3)
		lostGold = rand.Intn(max-min+1) + min
		player.Gold -= lostGold
		log.Println("GameOver!!")
		currentGameState = myGame.EndScreen
	}
	if enemy.HP <= 0 {
		//GoldRandom
		min := int(float64(enemy.Gold) * 0.7)
		max := int(float64(enemy.Gold) * 1.3)
		gainGold = rand.Intn(max-min+1) + min
		player.Gold += gainGold
		index = 0
		score++
		currentGameState = myGame.EndScreen
		yourTime = float64(elapsed.Seconds())
	}
	return currentGameState
}

func BattleTypingV2(win *pixelgl.Window, player *player.PlayerStatus, enemy *enemy.EnemyStatus, elapsed time.Duration) myGame.GameState {
	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = enemy.AttackTick - elapsed.Seconds()
	//log.Println(tempCount)
	//log.Println(elapsed.Seconds())

	if myGame.CurrentGS == myGame.PlayingScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					enemy.HP -= player.OP
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					PlayerAttack(win, int(player.OP), win.Bounds().Center().Sub(pixel.V(50, 150)))
					player.SP += player.BaseSP
					if index == len(question) {
						index = 0
						score++
					}
				} else {
					missType++
				}
			}
		} else {
			myGame.CurrentGS = myGame.BattleEnemyScreen
		}
	} else if myGame.CurrentGS == myGame.BattleEnemyScreen {
		//攻撃判定
		//PressEnter
		if win.JustPressed(pixelgl.KeyEnter) {
			pressEnter = true
		}
		if pressEnter == true {
			if enemy.EnemySize >= 4.0 && enemy.EnemySize < 4.5 && lock == false && lock2 == false {
				enemy.EnemySize = enemy.EnemySize * 1.05
				if enemy.EnemySize > 4.5 {
					lock = true
					win.Canvas().Clear(colornames.Red)
				}
			} else if enemy.EnemySize >= 4 && lock == true && lock2 == false {
				enemy.EnemySize = enemy.EnemySize * 0.95
				if enemy.EnemySize < 4 {
					lock2 = true
				}
			} else if lock == true && lock2 == true {
				enemy.EnemySize = 4.0
				lock = false
				lock2 = false
				player.HP -= enemy.OP
				myGame.CurrentGS = myGame.PlayingScreen
				pressEnter = false
				index = 0
			}
		}
	}

	BattleTypingSkill(win, player, enemy)
	myGame.CurrentGS = DeathFlug(player, enemy, elapsed, myGame.CurrentGS)
	return myGame.CurrentGS
}