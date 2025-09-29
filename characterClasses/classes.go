package characterClasses

type ClassSkills struct {
	MaxAllowed int
	Skills     []string
	Armor      []string
	Shields    string
	// weapons
	MainHand string
	OffHand  string
}

var Classes = map[string]ClassSkills{
	"barbarian": {
		MaxAllowed: 2,
		Skills: []string{
			"animal handling", "athletics", "intimidation", "nature", "perception", "survival",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
		},
		Shields:  "shield",
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"bard": {
		MaxAllowed: 3,
		Skills: []string{
			"acrobatics", "animal handling", "arcana", "athletics", "deception",
			"history", "insight", "intimidation", "investigation", "medicine", "nature",
			"perception", "performance", "persuasion", "religion", "sleight of hand", "stealth", "survival",
		},
		Armor: []string{
			"padded", "leather", "studded leather",
		},
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"cleric": {
		MaxAllowed: 2,
		Skills: []string{
			"history", "insight", "medicine", "persuasion", "religion",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
		},
		Shields:  "shield",
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"druid": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "animal handling", "insight", "medicine", "nature", "perception", "religion", "survival",
		},
		// druids don't wear armor or shields made of metal, fucking animals
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
		},
		Shields:  "shield",
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"fighter": {
		MaxAllowed: 2,
		Skills: []string{
			"acrobatics", "animal handling", "athletics", "history", "insight", "intimidation", "perception", "survival",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate", "ring mail", "chain mail", "splint", "plate",
		},
		Shields:  "shield",
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"monk": {
		MaxAllowed: 2,
		Skills: []string{
			"acrobatics", "athletics", "history", "insight", "religion", "stealth",
		},
		// NO ARMOR
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"paladin": {
		MaxAllowed: 2,
		Skills: []string{
			"athletics", "insight", "intimidation", "medicine", "persuasion", "religion",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate", "ring mail", "chain mail", "splint", "plate",
		},
		Shields:  "shield",
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"ranger": {
		MaxAllowed: 3,
		Skills: []string{
			"animal handling", "athletics", "insight", "investigation", "nature", "perception", "stealth", "survival",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
		},
		Shields:  "shield",
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"rogue": {
		MaxAllowed: 4,
		Skills: []string{
			"acrobatics", "athletcis", "deception", "insight", "intimidation", "investigation", "perception", "performance", "persuasion", "sleight of hand", "stealth",
		},
		Armor: []string{
			"padded", "leather", "studded leather",
		},
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"sorcerer": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "deception", "insight", "intimidation", "persuasion", "religion",
		},
		// NO ARMOR
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"warlock": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "deception", "history", "intimidation", "investigation", "nature", "religion",
		},
		Armor: []string{
			"padded", "leather", "studded leather",
		},
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"wizard": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "history", "insight", "investigation", "medicine", "religion",
		},
		// NO ARMOR
		MainHand: "main hand",
		OffHand:  "off hand",
	},
}
