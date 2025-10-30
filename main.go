package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"path/filepath"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain/class"
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

	switch cmd {
	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)

		name := createCmd.String("name", "", "character name (required)")
		race := createCmd.String("race", "", "character race (required)")
		// "acolyte" default
		background := createCmd.String("background", "acolyte", "character background (required)")
		className := createCmd.String("class", "", "character class (required)")
		level := createCmd.Int("level", 1, "character level (required)")

		str := createCmd.Int("str", 10, "strength is required")
		dex := createCmd.Int("dex", 10, "dexterity is required")
		con := createCmd.Int("con", 10, "constitution is required")
		intel := createCmd.Int("int", 10, "intelligence is required")
		wis := createCmd.Int("wis", 10, "wisdom is required")
		cha := createCmd.Int("cha", 10, "charisma is required")

		if err := createCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println("error parsing flags:", err)
			os.Exit(2)
		}

		stats := domain.Stats{
			Str: *str, Dex: *dex, Con: *con, Intel: *intel, Wis: *wis, Cha: *cha,
		}

		char, err := domain.NewCharacter(*name, *race, *background, *className, *level, stats)
		if err != nil {
			fmt.Println("Error creating character:", err)
			os.Exit(2)
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "characters.json"))
		classRepo := class.NewClassRepository(allWeps, simple, martial)
		facade := service.NewCharacterFacade(repo, classRepo)

		if err := facade.CreateCharacter(char); err != nil {
			fmt.Println("Error creating character: ", err)
			os.Exit(2)
		}

		fmt.Printf("saved character %s\n", char.Name)

	case "view":
		viewCmd := flag.NewFlagSet("view", flag.ExitOnError)
		name := viewCmd.String("name", "", "character name (required)")
		_ = viewCmd.Parse(os.Args[2:])

		if *name == "" {
			fmt.Println("name is required")
			os.Exit(2)
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "characters.json"))
		facade := service.NewCharacterFacade(repo, nil)

		char, err := facade.ViewCharacter(*name)
		if err != nil {
			fmt.Printf("character %q not found\n", *name)
			return
		}

		fmt.Print(service.FormatCharacterView(char))

	case "list":
		repo := repository.NewJsonRepository(filepath.Join("data", "characters.json"))
		facade := service.NewCharacterFacade(repo, nil)

		chars, err := facade.ListCharacters()
		if err != nil {
			fmt.Println("error listing characters: ", err)
			os.Exit(2)
		}

		if len(chars) == 0 {
			fmt.Println("no characters found")
			return
		}

		for _, c := range chars {
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

		repo := repository.NewJsonRepository(filepath.Join("data", "characters.json"))
		facade := service.NewCharacterFacade(repo, nil)

		if err := facade.DeleteCharacter(*name); err != nil {
			fmt.Println("error deleting character: ", err)
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

		repo := repository.NewJsonRepository(filepath.Join("data", "characters.json"))
		classRepo := class.NewClassRepository(allWeps, simple, martial)
		facade := service.NewCharacterFacade(repo, classRepo)

		err := facade.EquipItem(*name, *weapon, *slot, *armor, *shield)
		if err != nil {
			switch err {
			case domain.ErrSlotOccupied:
				fmt.Printf("%s already occupied\n", *slot)
				return

			default:
				fmt.Println("error equipping:", err)
				os.Exit(2)
			}
		}

		if *weapon != "" && *slot != "" {
			fmt.Printf("Equipped weapon %s to %s\n", *weapon, *slot)
			return
		}
		if *armor != "" {
			fmt.Printf("Equipped armor %s\n", *armor)
			return
		}
		if *shield != "" {
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

		repo := repository.NewJsonRepository(filepath.Join("data", "characters.json"))
		facade := service.NewCharacterFacade(repo, nil)

		message, err := facade.LearnSpell(*name, *spell)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(message)

	case "prepare-spell":
		prepareCmd := flag.NewFlagSet("prepare-spell", flag.ExitOnError)
		name := prepareCmd.String("name", "", "character name (required)")
		spell := prepareCmd.String("spell", "", "spell name (required)")
		_ = prepareCmd.Parse(os.Args[2:])

		if *name == "" || *spell == "" {
			fmt.Println("usage: prepare-spell -name <name> -spell <spell>")
			os.Exit(2)
		}

		repo := repository.NewJsonRepository(filepath.Join("data", "characters.json"))
		facade := service.NewCharacterFacade(repo, nil)

		message, err := facade.PrepareSpell(*name, *spell)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Println(message)

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

	case "serve":
		fs := http.FileServer(http.Dir("."))
		fmt.Printf("Serving on http://localhost:8080\n")
		fmt.Printf("Open: http://localhost:8080/charactersheet.html\n")
		log.Fatal(http.ListenAndServe(":8080", fs))

	default:
		usage()
		os.Exit(2)
	}
}
