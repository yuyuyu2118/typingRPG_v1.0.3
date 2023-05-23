package myGame

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

type stageInf struct {
	stageNum int
}

func NewStageInf(stageNum int) *stageInf {
	return &stageInf{stageNum}
}

var (
	stage1Button = pixel.Rect{}
)

var (
	stageButtonSlice = []pixel.Rect{}
)

func InitStage(win *pixelgl.Window, Txt *text.Text) {
	xOffSet := myPos.TopLefPos(win, Txt).X + 300
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	//stageSlice := []string{"1. Slime", "2. Bird", "3. Plant", "4. Goblin", "5. Zombie", "6. Fairy", "7. Skull", "8. Wizard", "9. Solidier", "10. Dragon", "BackSpace. EXIT"}
	stageSlice := []string{"1. スライム", "2. バード", "3. プラント", "4. ゴブリン", "5. ゾンビ", "6. フェアリー", "7. スカル", "8. ウィザード", "9. ソルジャー", "BackSpace. 戻る"}

	for _, stageName := range stageSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, stageName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		stageButtonSlice = append(stageButtonSlice, Txt.Bounds().Moved(txtPos))
	}
}

func StageClickEvent(win *pixelgl.Window, mousePos pixel.Vec, stage *stageInf) GameState {

	if stageButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		CurrentGS = PlayingScreen
		log.Println("PlayStage is VS knight")
		stage.stageNum = 1
	} else if stageButtonSlice[9].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) {
		CurrentGS = GoToScreen
		log.Println("StageScreen -> GoToScreen")
	}
	log.Println("YourJob is", stage.stageNum)
	return CurrentGS
}
