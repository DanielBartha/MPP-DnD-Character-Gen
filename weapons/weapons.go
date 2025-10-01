package weapons

type Weapon struct {
	Name      string
	TwoHanded bool
}

// defaulting all special weapons to be one hand, for now ig
var SimpleMelee = []Weapon{
	{Name: "club", TwoHanded: false},
	{Name: "dagger", TwoHanded: false},
	{Name: "greatclub", TwoHanded: true},
	{Name: "handaxe", TwoHanded: false},
	{Name: "javelin", TwoHanded: false},
	{Name: "light hammer", TwoHanded: false},
	{Name: "mace", TwoHanded: false},
	{Name: "quarterstaff", TwoHanded: false},
	{Name: "sickle", TwoHanded: false},
	{Name: "spear", TwoHanded: false},
}

var SimpleRanged = []Weapon{
	{Name: "light crossbow", TwoHanded: true},
	{Name: "dart", TwoHanded: false},
	{Name: "shortbow", TwoHanded: true},
	{Name: "sling", TwoHanded: true},
}

var MartialMelee = []Weapon{
	{Name: "battleaxe", TwoHanded: false},
	{Name: "flail", TwoHanded: false},
	{Name: "glaive", TwoHanded: true},
	{Name: "greataxe", TwoHanded: true},
	{Name: "greatsword", TwoHanded: true},
	{Name: "halberd", TwoHanded: true},
	{Name: "lance", TwoHanded: true},
	{Name: "longsword", TwoHanded: false},
	{Name: "maul", TwoHanded: true},
	{Name: "morningstar", TwoHanded: false},
	{Name: "pike", TwoHanded: true},
	{Name: "rapier", TwoHanded: false},
	{Name: "scimitar", TwoHanded: false},
	{Name: "shortsword", TwoHanded: false},
	{Name: "trident", TwoHanded: false},
	{Name: "war pick", TwoHanded: false},
	{Name: "warhammer", TwoHanded: false},
	{Name: "whip", TwoHanded: false},
}

var MartialRanged = []Weapon{
	{Name: "blowgun", TwoHanded: true},
	{Name: "hand crossbow", TwoHanded: false},
	{Name: "heavy crossbow", TwoHanded: true},
	{Name: "longbow", TwoHanded: true},
	{Name: "net", TwoHanded: false},
}

var AllWeapons = map[string][]Weapon{
	"simple melee":   SimpleMelee,
	"simple ranged":  SimpleRanged,
	"martial melee":  MartialMelee,
	"martial ranged": MartialRanged,
}

func GetAllWeapons() []Weapon {
	all := []Weapon{}
	for _, group := range AllWeapons {
		all = append(all, group...)
	}
	return all
}

func GetSimpleWeapons() []Weapon {
	all := []Weapon{}
	all = append(all, SimpleMelee...)
	all = append(all, SimpleRanged...)
	return all
}

func GetMartialWeapons() []Weapon {
	all := []Weapon{}
	all = append(all, MartialMelee...)
	all = append(all, MartialRanged...)
	return all
}
