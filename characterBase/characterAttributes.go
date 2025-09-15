package characterBase

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
