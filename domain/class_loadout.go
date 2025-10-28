package domain

type ClassLoadout struct {
	MaxAllowed int
	Skills     []string
	Armor      []string
	Shields    string
	Weapons    []WeaponInfo
	MainHand   string
	OffHand    string
}
