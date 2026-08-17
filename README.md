# D&D 5e Character Sheet Generator

A command-line tool for creating, managing, and equipping Dungeons & Dragons 5th Edition characters, with a lightweight browser-based viewer for displaying saved characters on a printable-style sheet.

Character data (spells, weapons, and armor) is fetched from the [D&D 5th Edition API](https://www.dnd5eapi.co/), enriched, and cached locally for fast, repeatable use.

## Features

- **Character creation**: race, class, background, level, and ability scores via CLI flags.
- **Equipment management**: equip weapons (by slot), armor, and shields, with slot-occupancy validation.
- **Spellcasting**: learn and prepare spells per character.
- **Persistent storage**: characters are saved to a local JSON file (`data/characters.json`).
- **Concurrent data enrichment**: spells, weapons, and armor are fetched from the D&D 5e API through a rate-limited worker-pool pipeline (goroutines + channels, ticker-throttled) and cached to disk, avoiding repeated API calls on subsequent runs.
- **Web viewer**: a simple static page (`charactersheet.html`) that reads saved characters and renders them on a fillable, printable-style 5e character sheet in the browser.
- **Supported classes**: Barbarian, Bard, Cleric, Druid, Fighter, Monk, Paladin, Ranger, Rogue, Sorcerer, Warlock, Wizard.

## Tech Stack

- **Go** (module: `github.com/DanielBartha/MPP-DnD-Character-Gen`, go 1.24)
- Standard library only for the CLI/backend (`net/http`, `encoding/json`, `flag`)
- Vanilla HTML/CSS/JS for the browser viewer

## Prerequisites

- Go 1.24 or later

## Installation

```bash
git clone https://github.com/DanielBartha/MPP-DnD-Character-Gen.git
cd MPP-DnD-Character-Gen
go build -o dndgen .
```

Or run directly with `go run main.go <command>` without building a binary.

## Usage

### Create a character

```bash
go run main.go create -name Aria -race Elf -class Wizard -level 3 \
  -str 8 -dex 14 -con 12 -intel 16 -wis 10 -cha 12
```

### View a character

```bash
go run main.go view -name Aria
```

### List all saved characters

```bash
go run main.go list
```

### Delete a character

```bash
go run main.go delete -name Aria
```

### Equip items

```bash
go run main.go equip -name Aria -weapon Dagger -slot mainhand
go run main.go equip -name Aria -armor "Leather Armor"
go run main.go equip -name Aria -shield "Shield"
```

### Manage spells

```bash
go run main.go learn-spell -name Aria -spell "Magic Missile"
go run main.go prepare-spell -name Aria -spell "Magic Missile"
```

### Enrich local data caches

Pre-fetches and caches spell/weapon/armor data from the D&D 5e API:

```bash
go run main.go enrich-spells
go run main.go enrich-weapons
go run main.go enrich-armor
```

### Web viewer

Serve the current directory and open the character sheet viewer in your browser:

```bash
go run main.go serve
```

Then visit `http://localhost:8080/charactersheet.html`, pick a saved character from the dropdown, and their stats, equipment, and spells will populate the sheet. Note the viewer is read-only, characters must be created via the CLI first.

## Project Structure

```
domain/       Core types and business rules (Character, Stats, Equipment, Spellcasting, classes)
repository/   JSON-backed persistence layer
service/      Application logic: API client/fetcher, enrichment, character facade, formatting
data/         Saved characters and cached API data (created at runtime)
charactersheet.html, script.js, style.css   Static browser-based character sheet viewer
```

## Testing

```bash
go test ./...
```

The project includes unit tests across the domain, repository, and service layers (~73% coverage).

## Data & Attribution

Spell and equipment reference data (`5e-SRD-Spells.csv`, `5e-SRD-Equipment.csv`) is sourced from the D&D 5th Edition SRD via the [D&D 5e API](https://www.dnd5eapi.co/), released under the [Open Gaming License](https://dnd.wizards.com/resources/systems-reference-document). See `SRD_CC_v5.1.pdf` for the full license text.

## License

No license currently specified for the project's own source code.
