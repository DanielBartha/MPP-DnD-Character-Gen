package service

import "github.com/DanielBartha/MPP-DnD-Character-Gen/domain"

type ArmorClass struct {
	DefaultValue      int
	ArmorType         string
	ArmorTypeModifier int
	ShieldEquipped    bool
}

func (d *ArmorClass) CalculateArmorClass(s *domain.Stats, e *domain.Equipment) {
	d.DefaultValue = 10 + s.DexMod

	d.ArmorType = e.Armor
}

// TODO: function to check armor type +shield and return value that modifies default ac

func (d *ArmorClass) GetArmorType(e *domain.Equipment) {
	d.ArmorType = e.Armor

	// if d.ArmorType != "" && d.ShieldEquipped == true {
	// 	switch
	// }
}

func (d *ArmorClass) CheckArmorModifier(e *domain.Equipment) {
	// TODO: get values from ArmorTypes map in classes.go, then assign ac value based on armor
}
