package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"path/filepath"

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
		svc.AssignClassSkills(&characterCreate)
		svc.ApplyRacialBonuses(&characterCreate)
		svc.UpdateProficiency(&characterCreate)

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
			character.Class,
			character.Race,
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

	case "learn-spell":

	case "prepare-spell":

	default:
		usage()
		os.Exit(2)
	}
}
