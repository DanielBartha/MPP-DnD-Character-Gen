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
	result := (score - 10) / 2
	if (score-10)%2 < 0 {
		result--
	}
	return result
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

func (s *CharacterService) GetClassSkills(c *domain.Character) domain.ClassLoadout {
	classKey := strings.ToLower(strings.TrimSpace(c.Class))
	cs, ok := characterClasses.Classes[classKey]

	if !ok {
		return domain.ClassLoadout{
			MaxAllowed: 0,
			Skills:     []string{"insight", "religion"},
			Armor:      []string{},
			Shields:    "",
			Weapons:    []domain.WeaponInfo{},
			MainHand:   "",
			OffHand:    "",
		}
	}

	src := cs.Skills
	localSkills := make([]string, len(src))
	copy(localSkills, src)

	selected := []string{}
	if cs.MaxAllowed > 0 && len(localSkills) > 0 {
		limit := cs.MaxAllowed
		if limit > len(localSkills) {
			limit = len(localSkills)
		}
		selected = localSkills[:limit]
	}

	selected = append(selected, "insight", "religion")

	return domain.ClassLoadout{
		MaxAllowed: cs.MaxAllowed,
		Skills:     selected,
		Armor:      cs.Armor,
		Shields:    cs.Shields,
		Weapons:    cs.Weapons,
		MainHand:   cs.MainHand,
		OffHand:    cs.OffHand,
	}
}

func (s *CharacterService) ApplyRacialBonuses(character *domain.Character) {
	switch strings.ToLower(character.Race) {
	case "dwarf":
		character.Stats.Con += 2

	case "elf":
		character.Stats.Dex += 2

	case "halfling", "stout halfling":
		character.Stats.Dex += 2

	case "lightfoot halfling":
		character.Stats.Dex += 2
		character.Stats.Cha++

	case "human":
		character.Stats.Str++
		character.Stats.Dex++
		character.Stats.Con++
		character.Stats.Intel++
		character.Stats.Wis++
		character.Stats.Cha++

	case "dragonborn":
		character.Stats.Str += 2
		character.Stats.Cha++

	case "gnome":
		character.Stats.Intel += 2

	// TODO: half-eelves get to choose which stats to increase besides the rizz, for now dex and wis as defaults
	case "half-elf":
		character.Stats.Cha += 2
		character.Stats.Dex++
		character.Stats.Wis++

	case "half-orc":
		character.Stats.Str += 2
		character.Stats.Con++

	case "tiefling":
		character.Stats.Cha += 2
		character.Stats.Intel++
	}

	character.Stats.StrMod = (character.Stats.Str - 10) / 2
	character.Stats.DexMod = (character.Stats.Dex - 10) / 2
	character.Stats.ConMod = (character.Stats.Con - 10) / 2
	character.Stats.IntelMod = (character.Stats.Intel - 10) / 2
	character.Stats.WisMod = (character.Stats.Wis - 10) / 2
	character.Stats.ChaMod = (character.Stats.Cha - 10) / 2
}
