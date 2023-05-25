package battle

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

func BattleTypingSkill(win *pixelgl.Window, player *player.PlayerStatus, enemy *enemy.EnemyStatus) {
	if win.JustPressed(pixelgl.KeySpace) {
		log.Println("Skill!!!")
		if player.SP == 50 {
			index = 0
			player.SP = 0
			myGame.CurrentGS = myGame.SkillScreen
		} else {
			log.Println("skillポイントが足りない")
		}
	}
}

var (
	RookieSkillCount = 0
	RookieSkillWords = []string{"oreno", "kenngiwo", "kuraeee"}
)

func BattleTypingRookieSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myGame.GameState {
	question := RookieSkillWords[RookieSkillCount]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.OP // - elapsed.Seconds()

	if myGame.CurrentGS == myGame.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= 4
					if index == len(question) {
						index = 0
						RookieSkillCount++
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
						tempWordDamage = 0.0
						if RookieSkillCount == 3 {
							RookieSkillCount = 0
							myGame.CurrentGS = myGame.PlayingScreen
						}
					}
				} else {
					missType++
				}
			}
		} else {
			myGame.CurrentGS = myGame.SkillScreen
		}
	}

	myGame.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myGame.CurrentGS)
	if myGame.CurrentGS == myGame.EndScreen {
		RookieSkillCount = 0
	}
	return myGame.CurrentGS
}

func InitBattleTextRookieSkill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myGame.CurrentGS == myGame.SkillScreen {
		tempWords := RookieSkillWords[RookieSkillCount]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprint(Txt, "> ")
		Txt.Color = colornames.Darkslategray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.Orange
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))

		offset := Txt.Bounds().W()
		TxtOrigX := Txt.Dot.X
		spacing := 100.0
		if len(RookieSkillWords)-RookieSkillCount != 1 {
			Txt.Color = colornames.Orange
			offset := Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, RookieSkillWords[RookieSkillCount+1])
			myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
			Txt.Dot.X = TxtOrigX
		}
		if !(len(RookieSkillWords)-RookieSkillCount == 2 || len(RookieSkillWords)-RookieSkillCount == 1) {
			Txt.Color = colornames.Orange
			offset += Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, RookieSkillWords[RookieSkillCount+2])
			myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
		}
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

var bulletLoadingSkill = []bool{false, false, false, false, false}
var bulletDamageSkill = []int{0, 0, 0, 0, 0}

func BattleTypingHunterSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myGame.GameState {
	xOffSet := 50.0
	yOffSet := myPos.TopLefPos(win, myUtil.ScreenTxt).Y - 100
	txtPos := pixel.V(xOffSet, yOffSet)
	myUtil.ScreenTxt.Color = colornames.White
	myUtil.HunterBulletTxt.Clear()
	myUtil.HunterBulletTxt.Color = colornames.White
	fmt.Fprintln(myUtil.HunterBulletTxt, "*拡張装填*")
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	myUtil.HunterBulletTxt.Draw(win, tempPosition)

	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.OP // - elapsed.Seconds()

	if myGame.CurrentGS == myGame.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= 3
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					if index == len(question) {
						index = 0
						score++
						//enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						//PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
						//tempWordDamage = 0.0
						if bulletLoadingSkill[3] {
							bulletLoadingSkill[4] = true
						} else if bulletLoadingSkill[2] {
							bulletLoadingSkill[3] = true
						} else if bulletLoadingSkill[1] {
							bulletLoadingSkill[2] = true
						} else if bulletLoadingSkill[0] {
							bulletLoadingSkill[1] = true
						}
						bulletLoadingSkill[0] = true

						if bulletLoadingSkill[0] {
							bulletDamageSkill[0] = int(tempWordDamage)
						}
						if bulletLoadingSkill[1] {
							bulletDamageSkill[1] = int(tempWordDamage)
						}
						if bulletLoadingSkill[2] {
							bulletDamageSkill[2] = int(tempWordDamage)
						}
						if bulletLoadingSkill[3] {
							bulletDamageSkill[3] = int(tempWordDamage)
						}
						if bulletLoadingSkill[4] {
							bulletDamageSkill[4] = int(tempWordDamage)
						}
						tempWordDamage = 0.0
						log.Println(bulletLoadingSkill)
					}
				} else {
					missType++
				}
			}
		} else {
			myGame.CurrentGS = myGame.SkillScreen
		}
	}

	if bulletLoadingSkill[0] && !bulletLoadingSkill[1] && !bulletLoadingSkill[2] && !bulletLoadingSkill[3] && !bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoadingSkill[0] && bulletLoadingSkill[1] && !bulletLoadingSkill[2] && !bulletLoadingSkill[3] && !bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoadingSkill[0] && bulletLoadingSkill[1] && bulletLoadingSkill[2] && !bulletLoadingSkill[3] && !bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamageSkill[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-3], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoadingSkill[0] && bulletLoadingSkill[1] && bulletLoadingSkill[2] && bulletLoadingSkill[3] && !bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[3])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamageSkill[2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-3], bulletDamageSkill[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-4], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoadingSkill[0] && bulletLoadingSkill[1] && bulletLoadingSkill[2] && bulletLoadingSkill[3] && bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[4])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamageSkill[3])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-3], bulletDamageSkill[2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-4], bulletDamageSkill[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-5], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	}

	if win.JustPressed(pixelgl.KeyEnter) {
		bulletDamageSkills := bulletDamageSkill[0] + bulletDamageSkill[1] + bulletDamageSkill[2] + bulletDamageSkill[3] + bulletDamageSkill[4]
		//enemy.EnemySettings[myGame.StageNum].HP += float64(bulletDamageSkills) //TODO: debug用
		PlayerAttack(win, bulletDamageSkill[0], win.Bounds().Center().Sub(pixel.V(50, -200)))
		PlayerAttack(win, bulletDamageSkill[1], win.Bounds().Center().Sub(pixel.V(-100, -200)))
		PlayerAttack(win, bulletDamageSkill[2], win.Bounds().Center().Sub(pixel.V(200, -200)))
		PlayerAttack(win, bulletDamageSkill[3], win.Bounds().Center().Sub(pixel.V(-200, -200)))
		PlayerAttack(win, bulletDamageSkill[4], win.Bounds().Center().Sub(pixel.V(300, -200)))
		enemy.EnemySettings[myGame.StageNum].HP += float64(bulletDamageSkills)
		for i := 0; i < 5; i++ {
			bulletDamageSkill[i] = 0
			bulletLoadingSkill[i] = false
		}
		log.Println("Bang!!")
		myGame.CurrentGS = myGame.PlayingScreen
	}

	myGame.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myGame.CurrentGS)
	return myGame.CurrentGS
}

func InitBattleTextHunterSkill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myGame.CurrentGS == myGame.SkillScreen {
		tempWords := words[score]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprint(Txt, "> ")
		Txt.Color = colornames.Darkslategray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.White
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))

		offset := Txt.Bounds().W()
		TxtOrigX := Txt.Dot.X
		spacing := 100.0
		if len(words)-score != 1 {
			Txt.Color = colornames.Darkgray
			offset := Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, words[score+1])
			myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
			Txt.Dot.X = TxtOrigX
		}
		if !(len(words)-score == 2 || len(words)-score == 1) {
			Txt.Color = colornames.Gray
			offset += Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, words[score+2])
			myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
		}
	} else if myGame.CurrentGS == myGame.BattleEnemyScreen {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, "EnemyAttack!!")
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

var (
	MonkSkillWords = []string{
		"dadadadadadadadadadadadaddada!!!!!",
		"mudamudamudamudamudamudamuda!!!!!!!",
		"oraoraoraoraoraoraoraoraoraora!!!!!",
	}
	MonkSkillWord = ""
)

func BattleTypingMonkSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myGame.GameState {
	if MonkSkillWord == "" {
		MonkSkillWord = MonkSkillWords[rand.Intn(3)]
	}
	question := MonkSkillWord
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.OP // - elapsed.Seconds()

	if myGame.CurrentGS == myGame.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= float64(rand.Intn(3))
					enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
					PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
					tempWordDamage = 0.0
					if index == len(question) {
						index = 0
						MonkSkillWord = ""
						myGame.CurrentGS = myGame.PlayingScreen
					}
				} else {
					missType++
				}
			}
		} else {
			myGame.CurrentGS = myGame.SkillScreen
		}
	}

	myGame.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myGame.CurrentGS)
	if myGame.CurrentGS == myGame.EndScreen {
		//index?
		RookieSkillCount = 0
	}
	return myGame.CurrentGS
}

func InitBattleTextMonkSkill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myGame.CurrentGS == myGame.SkillScreen {
		tempWords := MonkSkillWord
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprint(Txt, "> ")
		Txt.Color = colornames.Gray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.Red
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}