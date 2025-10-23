package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func LoadBaseWeapons(path string) ([]domain.WeaponInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var weapons []domain.WeaponInfo

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) < 2 || !strings.EqualFold(record[1], "Weapon") {
			continue
		}

		name := strings.TrimSpace(record[0])
		weapons = append(weapons, domain.WeaponInfo{Name: name})
	}

	return weapons, nil
}

func EnrichWeapons(inputPath, outputPath string) error {
	baseWeapons, err := LoadBaseWeapons(inputPath)
	if err != nil {
		return fmt.Errorf("failed to load base weapons: %w", err)
	}

	var indexes []string
	for _, w := range baseWeapons {
		indexes = append(indexes, SanitizeApiKey(w.Name))
	}

	fmt.Printf("Fetching weapon data for %d weapons...\n", len(indexes))
	weaponMap := FetchWeaponsBatch(indexes)

	if err := ensureDir(filepath.Dir(outputPath)); err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output: %w", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	writer.Write([]string{"name", "category", "range", "two_handed"})

	for _, w := range baseWeapons {
		idx := SanitizeApiKey(w.Name)
		apiData := weaponMap[idx]

		if apiData == nil {
			writer.Write([]string{w.Name, "N/A", "N/A", "N/A"})
			continue
		}

		domainW := ToDomainWeapon(apiData)
		writer.Write([]string{
			domainW.Name,
			domainW.Category,
			fmt.Sprintf("%d", domainW.Range),
			fmt.Sprintf("%t", domainW.TwoHanded),
		})
	}

	fmt.Printf("Finished enriching %d weapons\n", len(baseWeapons))
	return nil
}

func LoadEnrichedWeapons(path string) ([]domain.WeaponInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	_, _ = reader.Read()

	var weapons []domain.WeaponInfo
	for {
		rec, err := reader.Read()
		if err != nil {
			break
		}
		if len(rec) < 4 {
			continue
		}

		name := strings.TrimSpace(rec[0])
		category := rec[1]
		rng, _ := strconv.Atoi(rec[2])
		twoHanded := strings.EqualFold(rec[3], "true")

		weapons = append(weapons, domain.WeaponInfo{
			Name:      name,
			Category:  category,
			Range:     rng,
			TwoHanded: twoHanded,
		})
	}
	return weapons, nil
}

func FilterWeapons(weapons []domain.WeaponInfo, fn func(domain.WeaponInfo) bool) []domain.WeaponInfo {
	var result []domain.WeaponInfo

	for _, w := range weapons {
		if fn(w) {
			result = append(result, w)
		}
	}
	return result
}

func SimpleWeapons(all []domain.WeaponInfo) []domain.WeaponInfo {
	return FilterWeapons(all, func(w domain.WeaponInfo) bool {
		return strings.EqualFold(w.Category, "Simple")
	})
}

func MartialWeapons(all []domain.WeaponInfo) []domain.WeaponInfo {
	return FilterWeapons(all, func(w domain.WeaponInfo) bool {
		return strings.EqualFold(w.Category, "Martial")
	})
}
