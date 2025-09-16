package characterBase

func (c Character) ProficiencyBonus() int {
	switch {
	case c.Level >= 17:
		return 6
	case c.Level >= 13:
		return 5
	case c.Level >= 9:
		return 4
	case c.Level >= 5:
		return 3
	default:
		return 2
	}
}

type Character struct {
	Name        string
	Race        string
	Background  string
	Class       string
	Level       int
	Proficiency int
	Stats       Stats
	// Equipment Equipment
}

type Stats struct {
	Str   int
	Dex   int
	Con   int
	Intel int
	Wis   int
	Cha   int
}

// type Equipment struct {
// 	Armaments string
// 	Armor     string
// 	Gear      string
// 	Tools     string
// 	Mounts    string
// }
