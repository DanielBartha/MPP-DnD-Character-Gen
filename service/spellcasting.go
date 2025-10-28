package service

import (
	"fmt"
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

var fullCasterSlots = map[int]map[int]int{
	// 1(character level): {1(spell level): 2(spell slots)}
	1:  {1: 2},
	2:  {1: 3, 2: 0},
	3:  {1: 4, 2: 2},
	4:  {1: 4, 2: 3},
	5:  {1: 4, 2: 3, 3: 2},
	6:  {1: 4, 2: 3, 3: 3},
	7:  {1: 4, 2: 3, 3: 3, 4: 1},
	8:  {1: 4, 2: 3, 3: 3, 4: 2},
	9:  {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
	10: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
	11: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
	12: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
	13: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
	14: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
	15: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
	16: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
	17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1, 9: 1},
	18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 1, 7: 1, 8: 1, 9: 1},
	19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 2, 7: 1, 8: 1, 9: 1},
	20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 2, 7: 2, 8: 1, 9: 1},
}

var halfCasterSlots = map[int]map[int]int{
	1:  {},
	2:  {1: 2},
	3:  {1: 3},
	4:  {1: 3},
	5:  {1: 4, 2: 2},
	6:  {1: 4, 2: 2},
	7:  {1: 4, 2: 3},
	8:  {1: 4, 2: 3},
	9:  {1: 4, 2: 3, 3: 2},
	10: {1: 4, 2: 3, 3: 2},
	11: {1: 4, 2: 3, 3: 3},
	12: {1: 4, 2: 3, 3: 3},
	13: {1: 4, 2: 3, 3: 3, 4: 1},
	14: {1: 4, 2: 3, 3: 3, 4: 1},
	15: {1: 4, 2: 3, 3: 3, 4: 2},
	16: {1: 4, 2: 3, 3: 3, 4: 2},
	17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
	18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
	19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
	20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
}

var pactCasterSlots = map[int]map[int]int{
	1:  {1: 1},
	2:  {1: 2},
	3:  {2: 2},
	4:  {2: 2},
	5:  {3: 2},
	6:  {3: 2},
	7:  {4: 2},
	8:  {4: 2},
	9:  {5: 2},
	10: {5: 2},
	11: {5: 3},
	12: {5: 3},
	13: {5: 3},
	14: {5: 3},
	15: {5: 3},
	16: {5: 3},
	17: {5: 4},
	18: {5: 4},
	19: {5: 4},
	20: {5: 4},
}

func GetSlotsForClassLevel(class string, level int) (map[int]int, string, error) {
	classKey := strings.ToLower(strings.TrimSpace(class))
	casterType := domain.GetSpellcastingType(classKey)
	if casterType == "" {
		return nil, "", fmt.Errorf("class %s is not a spellcaster", class)
	}

	switch casterType {
	case "full":
		if m, ok := fullCasterSlots[level]; ok {
			return copyIntMap(m), "full", nil
		}
	case "half":
		if m, ok := halfCasterSlots[level]; ok {
			return copyIntMap(m), "half", nil
		}
	case "pact":
		if m, ok := pactCasterSlots[level]; ok {
			return copyIntMap(m), "pact", nil
		}
	//default aka not spellcaster
	default:
		return nil, "", fmt.Errorf("no slot table for class %s", class)
	}

	return nil, casterType, fmt.Errorf("no slot table for %s level %d", class, level)
}

func GetSpellcastingAbility(class string) string {
	switch strings.ToLower(strings.TrimSpace(class)) {
	case "bard", "paladin", "sorcerer", "warlock":
		return "charisma"

	case "cleric", "druid", "ranger":
		return "wisdom"

	case "wizard", "artificer":
		return "intelligence"

	default:
		return ""
	}
}

func copyIntMap(in map[int]int) map[int]int {
	mapCopy := make(map[int]int, len(in))
	for k, v := range in {
		mapCopy[k] = v
	}
	return mapCopy
}

func (s *CharacterService) InitSpellcasting(c *domain.Character) {
	if c.Spellcasting == nil {
		c.Spellcasting = &domain.Spellcasting{
			LearnedSpells:  []string{},
			PreparedSpells: []string{},
			Slots:          map[int]int{},
			MaxSlots:       map[int]int{},
			CanCast:        false,
		}
	}

	slots, casterType, err := GetSlotsForClassLevel(c.Class, c.Level)
	if err != nil {
		c.Spellcasting.CanCast = false
		return
	}

	for lvl, count := range slots {
		c.Spellcasting.Slots[lvl] = count
		c.Spellcasting.MaxSlots[lvl] = count
	}

	c.Spellcasting.CasterType = casterType

	canCast := true
	switch strings.ToLower(c.Class) {
	case "paladin", "ranger":
		if c.Level < 2 {
			canCast = false
		}
	case "eldritch knight", "arcane trickster":
		if c.Level < 3 {
			canCast = false
		}
	case "fighter", "rogue":
		canCast = false
	}

	c.Spellcasting.CanCast = canCast
	if !canCast {
		return
	}

	ability := GetSpellcastingAbility(c.Class)
	c.Spellcasting.Ability = ability

	var abilityMod int
	switch ability {
	case "intelligence":
		abilityMod = c.Stats.IntelMod
	case "wisdom":
		abilityMod = c.Stats.WisMod
	case "charisma":
		abilityMod = c.Stats.ChaMod
	}

	if ability != "" {
		c.Spellcasting.SpellSaveDC = 8 + c.Proficiency + abilityMod
		c.Spellcasting.SpellAttackBonus = c.Proficiency + abilityMod
	}

	switch casterType {
	case "full":
		if c.Level >= 10 {
			c.Spellcasting.CantripsKnown = 5
		} else if c.Level >= 4 {
			c.Spellcasting.CantripsKnown = 4
		} else {
			c.Spellcasting.CantripsKnown = 3
		}
	case "half":
		c.Spellcasting.CantripsKnown = 0
	case "pact":
		if c.Level >= 10 {
			c.Spellcasting.CantripsKnown = 4
		} else if c.Level >= 4 {
			c.Spellcasting.CantripsKnown = 3
		} else {
			c.Spellcasting.CantripsKnown = 2
		}
	}
}
