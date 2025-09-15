package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/characterBase"
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
		// You could use the Flag package like this
		// But feel free to do it differently!
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)

		name := createCmd.String("name", "", "character name (required)")
		race := createCmd.String("race", "", "character race (required)")
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

		if *name == "" || *race == "" || *class == "" {
			fmt.Println("name, race, class and level are required")
			createCmd.Usage()
			os.Exit(2)
		}

		characterCreate := characterBase.Character{
			Name:  *name,
			Race:  *race,
			Class: *class,
			Level: *level,
			Skills: characterBase.Stats{
				Str:   *str,
				Dex:   *dex,
				Con:   *con,
				Intel: *intel,
				Wis:   *wis,
				Cha:   *cha,
			},
		}

		fmt.Printf("Created character: %+v\n", characterCreate.Name)

	case "view":

	case "list":

	case "delete":

	case "equip":

	case "learn-spell":

	case "prepare-spell":

	default:
		usage()
		os.Exit(2)
	}
}
