package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func EnrichArmor(inputPath, outputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	header = append(header, "base_ac", "dex_bonus", "max_bonus")

	if err := ensureDir(filepath.Dir(outputPath)); err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create uotput: %w", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()
	writer.Write(header)

	var armorIndexes []string
	var records [][]string

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		name := strings.TrimSpace(record[0])
		itemType := strings.ToLower(strings.TrimSpace(record[1]))

		if itemType == "armor" {
			apiIndex := SanitizeApiKey(name)
			armorIndexes = append(armorIndexes, apiIndex)
		}

		records = append(records, record)
	}

	results := fetchArmorBatchFn(armorIndexes)

	var processed, missing int

	for _, record := range records {
		name := strings.TrimSpace(record[0])
		itemType := strings.ToLower(strings.TrimSpace(record[1]))

		if itemType != "armor" {
			writer.Write(append(record, "N/A", "N/A", "N/A"))
			continue
		}

		apiIndex := SanitizeApiKey(name)
		apiResp := results[apiIndex]

		if apiResp == nil {
			fmt.Printf("Skipping %s: not found\n", name)
			writer.Write(append(record, "N/A", "N/A", "N/A"))
			missing++
			continue
		}

		domainArmor := ToDomainArmor(apiResp)
		record = append(
			record,
			fmt.Sprintf("%d", domainArmor.BaseAC),
			fmt.Sprintf("%t", domainArmor.DexBonus),
			fmt.Sprintf("%d", domainArmor.MaxBonus),
		)
		writer.Write(record)
		processed++
	}

	writer.Flush()
	fmt.Printf("Finished; processed %d, missing %d\n", processed, missing)

	return nil
}
