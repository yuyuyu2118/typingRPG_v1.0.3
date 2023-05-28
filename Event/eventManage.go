package event

import (
	"strconv"
)

type WeaponPurchaseEvent struct {
	Weapons map[int]bool
}

var WeaponPurchaseEventInstance *WeaponPurchaseEvent

// TODO: 同じような変数のインスタンスを生成するものはすべてマップに置き換えることで、for文に変換できる
// weaponが購入されていればweaponPurchaseBoolに対応した武器がtrueになる
func CreateWeaponPurchaseEvent(value []string) {
	var weaponPurchaseBool = make(map[int]bool)
	for i, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt >= 1 {
			weaponPurchaseBool[i] = true
		} else {
			weaponPurchaseBool[i] = false
		}
	}
	WeaponPurchaseEventInstance = &WeaponPurchaseEvent{
		Weapons: weaponPurchaseBool,
	}
}

type ArmorPurchaseEvent struct {
	Armors map[int]bool
}

var ArmorPurchaseEventInstance *ArmorPurchaseEvent

func CreateArmorPurchaseEvent(value []string) {
	var armorPurchaseBool = make(map[int]bool)
	for i, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt >= 1 {
			armorPurchaseBool[i] = true
		} else {
			armorPurchaseBool[i] = false
		}
	}
	ArmorPurchaseEventInstance = &ArmorPurchaseEvent{
		Armors: armorPurchaseBool,
	}
}

type AccessoryPurchaseEvent struct {
	Accessorys map[int]bool
}

var AccessoryPurchaseEventInstance *AccessoryPurchaseEvent

func CreateAccessoryPurchaseEvent(value []string) {
	var accessoryPurchaseBool = make(map[int]bool)
	for i, b := range value {
		tempInt, _ := strconv.Atoi(b)
		if tempInt >= 1 {
			accessoryPurchaseBool[i] = true
		} else {
			accessoryPurchaseBool[i] = false
		}
	}
	AccessoryPurchaseEventInstance = &AccessoryPurchaseEvent{
		Accessorys: accessoryPurchaseBool,
	}
}

// type defeatedEnemyEvent struct {
// 	Slime    bool
// 	Bird     bool
// 	Plant    bool
// 	Goblin   bool
// 	Zombie   bool
// 	Fairy    bool
// 	Skull    bool
// 	Wizard   bool
// 	Solidier bool
// 	Dragon   bool
// }

// var DefeatedEnemyEventInstance defeatedEnemyEvent

// func CreateDefeatedEnemyEvent(value []string) {
// 	//temp := myIo.CsvToSliceAll(myGame.SaveFilePath)

// 	var defeatedEnemyBool []bool
// 	for _, b := range value {
// 		tempInt, _ := strconv.Atoi(b)
// 		if tempInt >= 1 {
// 			defeatedEnemyBool = append(defeatedEnemyBool, true)
// 		} else {
// 			defeatedEnemyBool = append(defeatedEnemyBool, false)
// 		}
// 	}
// 	DefeatedEnemyEventInstance = defeatedEnemyEvent{
// 		Slime:    defeatedEnemyBool[0],
// 		Bird:     defeatedEnemyBool[1],
// 		Plant:    defeatedEnemyBool[2],
// 		Goblin:   defeatedEnemyBool[3],
// 		Zombie:   defeatedEnemyBool[4],
// 		Fairy:    defeatedEnemyBool[5],
// 		Skull:    defeatedEnemyBool[6],
// 		Wizard:   defeatedEnemyBool[7],
// 		Solidier: defeatedEnemyBool[8],
// 		Dragon:   defeatedEnemyBool[9],
// 	}
// }
