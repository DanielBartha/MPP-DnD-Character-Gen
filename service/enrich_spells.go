package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func EnrichSpellsCSV(inputPath, outputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open csv: %w", err)
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
		return fmt.Errorf("failed to create output csv: %w", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()
	writer.Write(header)

	var processed, missing int

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		name := strings.TrimSpace(record[0])
		apiIndex := sanitizeAPIIndex(name)

		apiResp, fetchErr := FetchSpell(apiIndex)
		if fetchErr != nil {
			fmt.Printf("skipping %s: %v\n", name, fetchErr)
			writer.Write(append(record, "N/A", "N/A"))
			missing++
			continue
		}

		domainSpell := ToDomainSpell(apiResp)

		record = append(record, domainSpell.School, domainSpell.Range)
		writer.Write(record)
		processed++

		time.Sleep(150 * time.Millisecond)
	}

	fmt.Printf("Finished; processed: %d, missing: %d\n", processed, missing)
	return nil
}

func sanitizeAPIIndex(name string) string {
	index := strings.ToLower(name)
	index = strings.ReplaceAll(index, "'", "")
	index = strings.ReplaceAll(index, "’", "")
	index = strings.ReplaceAll(index, "(", "")
	index = strings.ReplaceAll(index, ")", "")
	index = strings.ReplaceAll(index, ",", "")
	index = strings.ReplaceAll(index, "/", "-")
	index = strings.ReplaceAll(index, ":", "")
	index = strings.ReplaceAll(index, ".", "")
	index = strings.ReplaceAll(index, " ", "-")
	index = strings.TrimSpace(index)
	return index
}
