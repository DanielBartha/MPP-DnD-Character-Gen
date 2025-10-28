package service

import (
	"fmt"
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

type SpellService struct{}

func NewSpellService() *SpellService {
	return &SpellService{}
}

func (s *SpellService) LearnSpell(c *domain.Character, spellName string) (string, error) {
	if c.Spellcasting == nil || !c.Spellcasting.CanCast {
		return "", fmt.Errorf("this class can't cast spells")
	}

	for _, s := range c.Spellcasting.LearnedSpells {
		if strings.EqualFold(s, spellName) {
			return "", fmt.Errorf("%s already learned", spellName)
		}
	}

	if c.Spellcasting.PreparedMode {
		return "", fmt.Errorf("this class prepares spells and can't learn them")
	}

	level, err := GetSpellLevel(spellName)
	if err != nil {
		return "", err
	}

	if !IsSpellForClass(spellName, c.Class) {
		return "", fmt.Errorf("%s cannot learn %s", c.Class, spellName)
	}

	if level > 0 {
		if slots, ok := c.Spellcasting.Slots[level]; !ok || slots == 0 {
			return "", fmt.Errorf("the spell has higher level than the available spell slots")
		}
	}

	c.Spellcasting.LearnedMode = true
	c.Spellcasting.LearnedSpells = append(c.Spellcasting.LearnedSpells, spellName)

	return fmt.Sprintf("Learned spell %s", spellName), nil
}
