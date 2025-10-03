package service

import (
	"fmt"
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

var fullCasterSlots = map[int]map[int]int{
	// need to extend
	1: {1: 4},
	2: {1: 3, 2: 2},
	3: {1: 4, 2: 2},
}

var halfCasterSlots = map[int]map[int]int{
	1: {},
	2: {1: 2},
}

var pactCasterSlots = map[int]map[int]int{
	1: {1: 1},
	2: {1: 2},
}

func GetSlotsForClassLevel(class string, level int) (map[int]int, error) {
	classKey := strings.ToLower(strings.TrimSpace(class))
	casterType, _ := domain.SpellcastingType[classKey]
	switch casterType {
	case "full":
		if m, ok := fullCasterSlots[level]; ok {
			return copyIntMap(m), nil
		}
		return map[int]int{}, fmt.Errorf("no slot table for full caster level %d", level)
	case "half":
		if m, ok := halfCasterSlots[level]; ok {
			return copyIntMap(m), nil
		}
		return map[int]int{}, fmt.Errorf("no slot table for half caster level %d", level)
	case "pact":
		if m, ok := pactCasterSlots[level]; ok {
			return copyIntMap(m), nil
		}
		return map[int]int{}, fmt.Errorf("no pact table for level %d", level)
	default:
		return map[int]int{}, nil // not a spellcaster
	}
}

func copyIntMap(in map[int]int) map[int]int {
	out := make(map[int]int, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

// to call during create
func (s *CharacterService) InitSpellcasting(c *domain.Character) {
	if c.Spellcasting == nil {
		c.Spellcasting = &domain.Spellcasting{
			LearnedSpells:  []string{},
			PreparedSpells: []string{},
			Slots:          map[int]int{},
			MaxSlots:       map[int]int{},
		}
	}

	slots, err := GetSlotsForClassLevel(c.Class, c.Level)
	if err != nil {
		return
	}

	for lvl, count := range slots {
		c.Spellcasting.Slots[lvl] = count
		c.Spellcasting.MaxSlots[lvl] = count
	}
}
