package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func EnrichSpells(inputPath, outputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input csv: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	header = append(header, "school", "range")

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
	writer.Write(header)

	var records [][]string
	var indexes []string

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		records = append(records, record)
		indexes = append(indexes, SanitizeApiKey(record[0]))
	}

	fmt.Printf("Fetching data for %d spells...\n", len(indexes))
	spellMap := FetchSpellsBatch(indexes)

	var processed, missing int

	for _, record := range records {
		name := record[0]
		apiIndex := SanitizeApiKey(name)

		apiResp := spellMap[apiIndex]
		if apiResp == nil {
			fmt.Printf("missing %s\n", name)
			record = append(record, "N/A", "N/A")
			missing++
		} else {
			domainSpell := ToDomainSpell(apiResp)
			record = append(record, domainSpell.School, domainSpell.Range)
			processed++
		}

		writer.Write(record)
	}

	fmt.Printf("Enriched spells; processed %d, missing: %d\n", processed, missing)

	return nil
}
