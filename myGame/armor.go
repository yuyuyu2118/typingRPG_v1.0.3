package myGame

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

type ArmorState int

const (
	armorNil ArmorState = iota
	armor1
	armor2
	armor3
	armor4
	armor5
	armor6
	armor7
	armor8
	armor9
	armor10
)

var keyToArmor = map[pixelgl.Button]ArmorState{
	pixelgl.Key1: armor1,
	pixelgl.Key2: armor2,
	pixelgl.Key3: armor3,
	pixelgl.Key4: armor4,
	pixelgl.Key5: armor5,
	pixelgl.Key6: armor6,
	pixelgl.Key7: armor7,
	pixelgl.Key8: armor8,
	pixelgl.Key9: armor9,
	pixelgl.Key0: armor10,
}

var armorSlice = []string{"1. ???", "2. ???", "3. ???", "4. ???", "5. ???", "6. ???", "7. ???", "8. ???", "9. ???", "0. ???"}
var armorNum = []string{"armor0", "armor1", "armor2", "armor3", "armor4", "armor5", "armor6", "armor7", "armor8", "armor9"}
var armorName = []string{"草織りのローブ", "フルーツアーマー", "木の鎧", "ソウルバインドプレート", "スタンプレート", "鉄の鎧", "飛翔のマント", "勇者の鎧", "刃舞の衣", "冥界の鎧"}

var (
	armorPath = "assets/shop/armor.csv"
	descArmor = CsvToSlice(armorPath)
)
var currentarmorState ArmorState

func InitArmor(win *pixelgl.Window, Txt *text.Text, botText string) {
	xOffSet, yOffSet, txtPos := myUtil.ShopInitAndText(win, myUtil.ScreenTxt, botText)

	for i, v := range armorName {
		if event.ArmorPurchaseEventInstance.Armors[i] {
			armorSlice[i] = strconv.Itoa(i+1) + ". " + v
		}
	}
	if event.ArmorPurchaseEventInstance.Armors[9] {
		armorSlice[9] = "0. " + armorName[9]
	}

	buttonSliceArmor = myUtil.DisplayShopLineup(win, armorSlice, buttonSliceArmor, 30.0, colornames.White, myUtil.DescriptionTxt, xOffSet, yOffSet, txtPos)

	for i := 0; i < len(keyToArmor)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if win.Pressed(key) && event.ArmorPurchaseEventInstance.Armors[i] {
			currentarmorState = ArmorState(i + 1)
			break
		}
	}
	if win.Pressed(pixelgl.Key0) && event.ArmorPurchaseEventInstance.Armors[9] {
		currentarmorState = armor10
	}
	if currentarmorState >= armor1 && currentarmorState <= armor10 {
		DescriptionArmor(win, descArmor, int(currentarmorState)-1)
	}
}

func ArmorClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {
	var tempArmor = ""

	for i := 0; i < len(keyToArmor)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (buttonSliceArmor[i].Contains(mousePos) || win.Pressed(key)) && event.ArmorPurchaseEventInstance.Armors[i] && myState.CurrentGS == myState.ArmorShop {
			currentarmorState = ArmorState(i + 1)
			//CreateArmorEvent(descArmor, 0)
			log.Println("防具屋->防具", i+1)
			break
		}
	}

	if (buttonSliceArmor[9].Contains(mousePos) || win.JustPressed(pixelgl.Key0)) && event.ArmorPurchaseEventInstance.Armors[9] && myState.CurrentGS == myState.ArmorShop {
		currentarmorState = armor10
		log.Println("防具屋->防具10")
	} else if win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.ArmorShop {
		myState.CurrentGS = myState.TownScreen
		log.Println("防具屋->町")
	}

	if (buySellSliceArmor[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyB)) && player.Gold >= 100 {
		loadContent := SaveFileLoad(SaveFilePath)
		//TODO: お金が足りないときの処理を記述
		for i := 0; i < len(keyToArmor)-1; i++ {
			if currentarmorState == ArmorState(i+1) {
				requiredGold, _ := strconv.Atoi(descArmor[i+1][4])
				belongArmor, _ := strconv.Atoi(loadContent[4][i])
				//log.Println(loadContent)
				log.Println(belongArmor)
				if belongArmor == 0 {
					if player.Gold >= requiredGold {
						log.Println(descArmor[i+1][4], "買える", player.Gold)
						createOk := CreateArmorEvent(descArmor, i)
						if createOk {
							player.Gold -= requiredGold
							tempArmor = "armor" + strconv.Itoa(i+1)
						}
					} else {
						log.Println(descArmor[i+1][4], "お金が足りない", player.Gold)
					}
				} else {
					log.Println("すでに持っている")
					break
				}
			}
		}
		if currentarmorState == armor10 {
			requiredGold, _ := strconv.Atoi(descArmor[10][4])
			if player.Gold >= requiredGold {
				log.Println(descArmor[10][4], "買える", player.Gold)
			} else {
				log.Println(descArmor[10][4], "お金が足りない", player.Gold)
			}
			log.Println(descArmor[10][4])
			tempArmor = "armor" + strconv.Itoa(10)
		}

		if tempArmor != "" {
			SaveArmorPurchaseEvent(SaveFilePath, 4, tempArmor, player)
			SaveGame(SaveFilePath, 1, player)
		}
	}

	return myState.CurrentGS
}

func DescriptionArmor(win *pixelgl.Window, descArmor [][]string, num int) {
	//TODO: Tabを押している間は強化素材等の情報を表示する
	num++
	xOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).X + 300
	yOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).Y - 50
	txtPos := pixel.V(0, 0)

	myUtil.DescriptionTxt.Color = colornames.White

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descArmor[0][1]+": "+descArmor[num][1], "   カラー: "+descArmor[num][17])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descArmor[0][2]+": "+descArmor[num][2], descArmor[0][3]+": "+descArmor[num][3])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descArmor[0][4]+": "+descArmor[num][4]+"S ")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "素材: "+descArmor[num][5], descArmor[num][6]+"個, ", descArmor[num][7], descArmor[num][8]+"個, ", descArmor[num][9], descArmor[num][10]+"個")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "説明: "+descArmor[num][11])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, " "+descArmor[num][12])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "特殊能力: "+descArmor[num][14])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, " "+descArmor[num][15])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descArmor[num][16])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	myUtil.DescriptionTxt.Color = colornames.White
	fmt.Fprintln(myUtil.DescriptionTxt, "B. 作ってもらう")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)
	buySellSliceArmor = append(buySellSliceArmor, myUtil.DescriptionTxt.Bounds().Moved(txtPos))
}

func CreateArmorEvent(descArmor [][]string, num int) bool {
	//TODO: 素材が足りるかどうかの判定実装中
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false}

	for name, count := range tempSlice {
		if name == descArmor[num][5] {
			tempCount, _ := strconv.Atoi(descArmor[num][6])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[0] = true
			}
		}
		if name == descArmor[num][7] {
			tempCount, _ := strconv.Atoi(descArmor[num][8])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		}
		if name == descArmor[num][9] {
			tempCount, _ := strconv.Atoi(descArmor[num][10])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if name == descArmor[num][5] {
				tempCount, _ := strconv.Atoi(descArmor[num][6])
				tempSlice[name] -= tempCount
			}
			if name == descArmor[num][7] {
				tempCount, _ := strconv.Atoi(descArmor[num][8])
				tempSlice[name] -= tempCount
			}
			if name == descArmor[num][9] {
				tempCount, _ := strconv.Atoi(descArmor[num][10])
				tempSlice[name] -= tempCount
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費して防具を作成しました。")
		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}
}

func InitArmorBelongScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "持ち物/防具"
	InitArmorBelong(win, Txt, botText)
}

func InitArmorBelong(win *pixelgl.Window, Txt *text.Text, botText string) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)

	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White
	fmt.Fprintln(myUtil.ScreenTxt, botText, "Tabで切り替え", "BackSpace.戻る")
	tempPosition = myPos.BotCenPos(win, myUtil.ScreenTxt)
	myPos.DrawPos(win, myUtil.ScreenTxt, tempPosition)

	loadContent := SaveFileLoad(SaveFilePath)
	counts := make(map[string]int)
	elements := loadContent[4]

	for i, val := range elements {
		num, err := strconv.Atoi(val)
		if err == nil {
			armorKey := fmt.Sprintf("armor%d", i)
			counts[armorKey] = num
		}
	}

	for i, value := range armorName {
		if counts["armor"+strconv.Itoa(i)] != 0 {
			tempInt := counts["armor"+strconv.Itoa(i)]
			equipmentSlice = append(equipmentSlice, value+": "+strconv.Itoa(tempInt))
		}
	}

	for _, equipmentName := range equipmentSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, equipmentName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		equipmentButtonSlice = append(equipmentButtonSlice, Txt.Bounds().Moved(txtPos))
	}
	equipmentSlice = equipmentSlice[:0]
}

func ArmorBelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.StageSelect
		log.Println("GoToScreen->Dungeon")
	} else if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) {
		myState.CurrentGS = myState.TownScreen
		log.Println("GoToScreen->Town")
	} else if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) {
		myState.CurrentGS = myState.EquipmentScreen
		log.Println("GoToScreen->Equipment")
	} else if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.JobSelect
		log.Println("GoToScreen->JobSelect")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentBelong = myState.ArmorBelong
		myState.CurrentGS = myState.StartScreen
		log.Println("所持品/防具->GoTo")
	}
	return myState.CurrentGS
}