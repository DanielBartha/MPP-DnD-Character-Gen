package main

type Character struct {
	name      string
	race      string
	class     string
	level     string
	skills    Stats
	equipment Equipment
}

type Stats struct {
	STR int
	DEX int
	CON int
	INT int
	WIS int
	CHA int
}

type Equipment struct {
	armaments string
	armor     string
	gear      string
	tools     string
	mounts    string
}
