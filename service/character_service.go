package service

import (
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/characterClasses"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

type CharacterService struct {
	// CharacterSkillProvider in the future for loading data from db
	// c'est possible
}

func NewCharacterService() *CharacterService {
	return &CharacterService{}
}

func abilityModifier(score int) int {
	return (score - 10) / 2
}

func (s *CharacterService) UpdateProficiency(character *domain.Character) {
	switch {
	case character.Level >= 17:
		character.Proficiency = 6
	case character.Level >= 13:
		character.Proficiency = 5
	case character.Level >= 9:
		character.Proficiency = 4
	case character.Level >= 5:
		character.Proficiency = 3
	default:
		character.Proficiency = 2
	}

	character.Stats.StrMod = abilityModifier(character.Stats.Str)
	character.Stats.DexMod = abilityModifier(character.Stats.Dex)
	character.Stats.ConMod = abilityModifier(character.Stats.Con)
	character.Stats.IntelMod = abilityModifier(character.Stats.Intel)
	character.Stats.WisMod = abilityModifier(character.Stats.Wis)
	character.Stats.ChaMod = abilityModifier(character.Stats.Cha)
}

func (s *CharacterService) AssignClassSkills(character *domain.Character) {
	classKey := strings.ToLower(strings.TrimSpace(character.Class))

	classSkills, ok := characterClasses.Classes[classKey]
	if !ok {
		character.Skills = characterClasses.ClassSkills{
			MaxAllowed: 0,
			Skills:     []string{"insight", "religion"},
		}
		return
	}

	// local copy of skills to not modify global map on accident
	src := classSkills.Skills
	localSlice := make([]string, len(src))
	copy(localSlice, src)

	selected := []string{}
	if classSkills.MaxAllowed > 0 && len(localSlice) > 0 {
		limit := classSkills.MaxAllowed
		if limit > len(localSlice) {
			limit = len(localSlice)
		}
		selected = localSlice[:limit]
	}

	selected = append(selected, "insight", "religion")

	character.Skills = characterClasses.ClassSkills{
		MaxAllowed: classSkills.MaxAllowed,
		Skills:     selected,
	}
}
