package characterClasses

type ClassSkills struct {
	MaxAllowed int
	Skills     []string
}

var Classes = map[string]ClassSkills{
	"barbarian": {
		MaxAllowed: 2,
		Skills: []string{
			"animal handling", "athletics", "intimidation", "nature", "perception", "survival",
		},
	},
	"bard": {
		MaxAllowed: 3,
		Skills: []string{
			"acrobatics", "animal handling", "arcana", "athletics", "deception",
			"history", "insight", "intimidation", "investigation", "medicine", "nature",
			"perception", "performance", "persuasion", "religion", "sleight of hand", "stealth", "survival",
		},
	},
	"cleric": {
		MaxAllowed: 2,
		Skills: []string{
			"history", "insight", "medicine", "persuasion", "religion",
		},
	},
	"druid": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "animal handling", "insight", "medicine", "nature", "perception", "religion", "survival",
		},
	},
	"fighter": {
		MaxAllowed: 2,
		Skills: []string{
			"acrobatics", "animal handling", "athletics", "history", "insight", "intimidation", "perception", "survival",
		},
	},
	"monk": {
		MaxAllowed: 2,
		Skills: []string{
			"acrobatics", "athletics", "history", "insight", "religion", "stealth",
		},
	},
	"paladin": {
		MaxAllowed: 2,
		Skills: []string{
			"athletics", "insight", "intimidation", "medicine", "persuasion", "religion",
		},
	},
	"ranger": {
		MaxAllowed: 3,
		Skills: []string{
			"animal handling", "athletics", "insight", "investigation", "nature", "perception", "stealth", "survival",
		},
	},
	"rogue": {
		MaxAllowed: 4,
		Skills: []string{
			"acrobatics", "athletcis", "deception", "insight", "intimidation", "investigation", "perception", "performance", "persuasion", "sleight of hand", "stealth",
		},
	},
	"sorcerer": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "deception", "insight", "intimidation", "persuasion", "religion",
		},
	},
	"warlock": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "deception", "history", "intimidation", "investigation", "nature", "religion",
		},
	},
	"wizard": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "history", "insight", "investigation", "medicine", "religion",
		},
	},
}
