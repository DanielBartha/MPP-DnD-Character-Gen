package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"path/filepath"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/characterClasses"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/repository"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/service"
)

func usage() {
	fmt.Printf(`Usage:
  %s create -name CHARACTER_NAME -race RACE -class CLASS -level N -str N -dex N -con N -intel N -wis N -cha N
  %s view -name CHARACTER_NAME
  %s list
  %s delete -name CHARACTER_NAME
  %s equip -name CHARACTER_NAME -weapon WEAPON_NAME -slot SLOT
  %s equip -name CHARACTER_NAME -armor ARMOR_NAME
  %s equip -name CHARACTER_NAME -shield SHIELD_NAME
  %s learn-spell -name CHARACTER_NAME -spell SPELL_NAME
  %s prepare-spell -name CHARACTER_NAME -spell SPELL_NAME 
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	cmd := os.Args[1]

	err := service.LoadSpellsCSV("5e-SRD-Spells.csv")
	if err != nil {
		fmt.Println("Error loading spell list:", err)
		os.Exit(1)
	}

	allWeps, err := service.LoadEnrichedWeapons("5e-SRD-Equipment.csv")
	if err != nil {
		fmt.Println("Error loading weapons:", err)
		os.Exit(1)
	}
	simple := service.SimpleWeapons(allWeps)
	martial := service.MartialWeapons(allWeps)

	characterClasses.InitWeapons(allWeps, simple, martial)

	switch cmd {
	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)

		name := createCmd.String("name", "", "character name (required)")
		race := createCmd.String("race", "", "character race (required)")
		// "acolyte" as default value
		background := createCmd.String("background", "acolyte", "character background (required)")
		class := createCmd.String("class", "", "character class (required)")
		level := createCmd.Int("level", 1, "character level (required)")

		str := createCmd.Int("str", 10, "strength is required")
		dex := createCmd.Int("dex", 10, "dexterity is required")
		con := createCmd.Int("con", 10, "constitution is required")
		intel := createCmd.Int("int", 10, "intelligence is required")
		wis := createCmd.Int("wis", 10, "wisdom is required")
		cha := createCmd.Int("cha", 10, "charisma is required")

		err := createCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("error parsing flags")
			createCmd.Usage()
			os.Exit(2)
		}

		if *name == "" {
			fmt.Println("name is required")
			os.Exit(2)
		}
		if *race == "" {
			fmt.Println("race is required")
			os.Exit(2)
		}
		if *class == "" {
			fmt.Println("class is required")
			os.Exit(2)
		}
		if *level <= 0 {
			fmt.Println("level is required")
			os.Exit(2)
		}

		characterCreate := domain.Character{
			Name:       *name,
			Race:       *race,
			Background: *background,
			Class:      *class,
			Level:      *level,
			Stats: domain.Stats{
				Str:   *str,
				Dex:   *dex,
				Con:   *con,
				Intel: *intel,
				Wis:   *wis,
				Cha:   *cha,
			},
		}

		svc := service.NewCharacterService()

		characterCreate.Skills = svc.GetClassSkills(&characterCreate)
		svc.ApplyRacialBonuses(&characterCreate)
		svc.UpdateProficiency(&characterCreate)
		svc.InitSpellcasting(&characterCreate)

		characterCreate.Equipment = domain.Equipment{
			Weapon: map[string]string{
				"main hand": "",
				"off hand":  "",
			},
			Armor:  "",
			Shield: "",
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "settings.json"))
		if err := repo.Save(&characterCreate); err != nil {
			fmt.Println("error saving character:", err)
			os.Exit(2)
		}

		fmt.Printf("saved character %+v\n", characterCreate.Name)

	case "view":
		viewCmd := flag.NewFlagSet("view", flag.ExitOnError)
		name := viewCmd.String("name", "", "character name (required)")
		_ = viewCmd.Parse(os.Args[2:])

		if *name == "" {
			fmt.Println("name is required")
			os.Exit(2)
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "settings.json"))
		character, err := repo.Load(*name)
		if err != nil {
			fmt.Printf("character %q not found\n", *name)
			return
		}

		fmt.Printf(
			"Name: %s\n"+
				"Class: %s\n"+
				"Race: %s\n"+
				"Background: %s\n"+
				"Level: %d\n"+
				"Ability scores:\n"+
				"  STR: %d (%+d)\n"+
				"  DEX: %d (%+d)\n"+
				"  CON: %d (%+d)\n"+
				"  INT: %d (%+d)\n"+
				"  WIS: %d (%+d)\n"+
				"  CHA: %d (%+d)\n"+
				"Proficiency bonus: +%d\n"+
				"Skill proficiencies: %s\n",
			character.Name,
			strings.ToLower(character.Class),
			strings.ToLower(character.Race),
			character.Background,
			character.Level,
			character.Stats.Str, character.Stats.StrMod,
			character.Stats.Dex, character.Stats.DexMod,
			character.Stats.Con, character.Stats.ConMod,
			character.Stats.Intel, character.Stats.IntelMod,
			character.Stats.Wis, character.Stats.WisMod,
			character.Stats.Cha, character.Stats.ChaMod,
			character.Proficiency,
			strings.Join(character.Skills.Skills, ", "),
		)

		if character.Spellcasting != nil && character.Spellcasting.CanCast {
			fmt.Println("Spell slots:")

			if character.Spellcasting.CantripsKnown > 0 {
				fmt.Printf("  Level 0: %d\n", character.Spellcasting.CantripsKnown)
			}

			for lvl := 1; lvl <= 9; lvl++ {
				if count, ok := character.Spellcasting.MaxSlots[lvl]; ok && count > 0 {
					fmt.Printf("  Level %d: %d\n", lvl, count)
				}
			}
		}

		if weapon, ok := character.Equipment.Weapon["main hand"]; ok && weapon != "" {
			fmt.Printf("Main hand: %s\n", weapon)
		}

		if weapon, ok := character.Equipment.Weapon["off hand"]; ok && weapon != "" {
			fmt.Printf("Off hand: %s\n", weapon)
		}
		if character.Equipment.Armor != "" {
			fmt.Printf("Armor: %s\n", character.Equipment.Armor)
		}

		if character.Equipment.Shield != "" {
			fmt.Printf("Shield: %s\n", character.Equipment.Shield)
		}

	case "list":
		repo := repository.NewJsonRepository(filepath.Join("data", "settings.json"))
		characters, err := repo.List()
		if err != nil {
			fmt.Println("error listing characters:", err)
			os.Exit(2)
		}
		if len(characters) == 0 {
			fmt.Println("no characters found")
			return
		}
		for _, c := range characters {
			fmt.Printf("- %s (%s %s)\n", c.Name, c.Race, c.Class)
		}

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		name := deleteCmd.String("name", "", "character name (required)")
		_ = deleteCmd.Parse(os.Args[2:])

		if *name == "" {
			fmt.Println("name is required")
			os.Exit(2)
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "settings.json"))
		if err := repo.Delete(*name); err != nil {
			fmt.Println("error deleting character:", err)
			os.Exit(2)
		}
		fmt.Printf("deleted %s\n", *name)

	case "equip":
		equipCmd := flag.NewFlagSet("equip", flag.ExitOnError)
		name := equipCmd.String("name", "", "character name (required)")
		weapon := equipCmd.String("weapon", "", "weapon to equip")
		slot := equipCmd.String("slot", "", "slot to equip weapon to")
		armor := equipCmd.String("armor", "", "armor to equip")
		shield := equipCmd.String("shield", "", "shield to equip")
		_ = equipCmd.Parse(os.Args[2:])

		if *name == "" {
			fmt.Println("name is required")
			os.Exit(2)
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "settings.json"))
		character, err := repo.Load(*name)
		if err != nil {
			fmt.Printf("character %q not found\n", *name)
			return
		}

		if *weapon != "" && *slot != "" {
			if character.Equipment.Weapon == nil {
				character.Equipment.Weapon = make(map[string]string)
			}

			if existing, ok := character.Equipment.Weapon[*slot]; ok && existing != "" {
				fmt.Printf("%s already occupied\n", *slot)
				return
			}

			character.Equipment.Weapon[*slot] = *weapon
			if err := repo.Save(character); err != nil {
				fmt.Printf("error saving character: %v\n", err)
				os.Exit(2)
			}
			fmt.Printf("Equipped weapon %s to %s\n", *weapon, *slot)
			return
		}

		if *armor != "" {
			character.Equipment.Armor = *armor
			if err := repo.Save(character); err != nil {
				fmt.Println("error saving character:", *armor)
				os.Exit(2)
			}
			fmt.Printf("Equipped armor %s\n", *armor)
			return
		}

		if *shield != "" {
			character.Equipment.Shield = *shield
			if err := repo.Save(character); err != nil {
				fmt.Println("error saving character:", err)
				os.Exit(2)
			}
			fmt.Printf("Equipped shield %s\n", *shield)
			return
		}

		fmt.Println("no equipment provided")

	case "learn-spell":
		learnCmd := flag.NewFlagSet("learn-spell", flag.ExitOnError)
		name := learnCmd.String("name", "", "character name (required)")
		spell := learnCmd.String("spell", "", "spell name (required)")
		_ = learnCmd.Parse(os.Args[2:])

		if *name == "" || *spell == "" {
			fmt.Println("usage: learn-spell -name <name> -spell <spell>")
			os.Exit(2)
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "settings.json"))
		character, err := repo.Load(*name)
		if err != nil {
			fmt.Printf("character %q not found\n", *name)
			return
		}

		if character.Spellcasting == nil || !character.Spellcasting.CanCast {
			fmt.Printf("this class can't cast spells\n")
			return
		}

		for _, s := range character.Spellcasting.LearnedSpells {
			if strings.EqualFold(s, *spell) {
				fmt.Printf("%s already learned\n", *spell)
				return
			}
		}

		if character.Spellcasting.PreparedMode {
			fmt.Printf("this class prepares spells and can't learn them\n")
			return
		}

		// checks for non-existing spells (csv)
		level, err := service.GetSpellLevel(*spell)
		if err != nil {
			fmt.Println(err)
			return
		}

		if !service.IsSpellForClass(*spell, character.Class) {
			fmt.Printf("%s cannot learn %s\n", character.Class, *spell)
			return
		}

		if level > 0 {
			if slots, ok := character.Spellcasting.Slots[level]; !ok || slots == 0 {
				fmt.Printf("the spell has higher level than the available spell slots\n")
				return
			}
		}

		character.Spellcasting.LearnedMode = true
		character.Spellcasting.LearnedSpells = append(character.Spellcasting.LearnedSpells, *spell)

		if err := repo.Save(character); err != nil {
			fmt.Println("error saving character:", err)
			os.Exit(2)
		}

		fmt.Printf("Learned spell %s\n", *spell)

	case "prepare-spell":
		prepareCmd := flag.NewFlagSet("prepare-spell", flag.ExitOnError)
		name := prepareCmd.String("name", "", "character name (required)")
		spell := prepareCmd.String("spell", "", "spell name (required)")
		_ = prepareCmd.Parse(os.Args[2:])

		if *name == "" || *spell == "" {
			fmt.Println("usage: prepare-spell -name <name> -spell <spell>")
			os.Exit(2)
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "settings.json"))
		character, err := repo.Load(*name)
		if err != nil {
			fmt.Printf("character %q not found\n", *name)
			return
		}

		if character.Spellcasting == nil || !character.Spellcasting.CanCast {
			fmt.Printf("this class can't cast spells\n")
			return
		}

		for _, s := range character.Spellcasting.PreparedSpells {
			if strings.EqualFold(s, *spell) {
				fmt.Printf("%s already prepared\n", *spell)
				return
			}
		}

		if character.Spellcasting.LearnedMode {
			fmt.Printf("this class learns spells and can't prepare them\n")
			return
		}

		// check for non-existing spells (csv)
		level, err := service.GetSpellLevel(*spell)
		if err != nil {
			fmt.Println(err)
			return
		}

		if !service.IsSpellForClass(*spell, character.Class) {
			fmt.Printf("%s cannot prepare %s\n", character.Class, *spell)
			return
		}

		if level > 0 {
			if slots, ok := character.Spellcasting.Slots[level]; !ok || slots == 0 {
				fmt.Printf("the spell has higher level than the available spell slots\n")
				return
			}
		}

		character.Spellcasting.PreparedMode = true
		character.Spellcasting.PreparedSpells = append(character.Spellcasting.PreparedSpells, *spell)

		if err := repo.Save(character); err != nil {
			fmt.Println("error saving character:", err)
			os.Exit(2)
		}

		fmt.Printf("Prepared spell %s\n", *spell)

	case "enrich-spells":
		input := "5e-SRD-Spells.csv"
		output := "data/enriched/5e-SRD-Spells-enriched.csv"

		if err := service.EnrichSpells(input, output); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(2)
		}

	case "enrich-weapons":
		input := "5e-SRD-Equipment.csv"
		output := "data/enriched/5e-SRD-Weapons-Enriched.csv"

		if err := service.EnrichWeapons(input, output); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(2)
		}

	case "enrich-armor":
		input := "5e-SRD-Equipment.csv"
		output := "data/enriched/5e-SRD-Armor-Enriched.csv"

		if err := service.EnrichArmor(input, output); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(2)
		}

	default:
		usage()
		os.Exit(2)
	}
}
