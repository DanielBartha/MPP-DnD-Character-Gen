package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type apiSpellResp struct {
	Index  string `json:"index"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	School struct {
		Name string `json:"name"`
	}
	Range string `json:"range"`
}

type apiWeaponResp struct {
	Index          string `json:"index"`
	Name           string `json:"name"`
	WeaponCategory string `json:"weapon_category"`
	Range          struct {
		Normal int `json:"normal"`
		Long   int `json:"long"`
	} `json:"range"`
	TwoHanded bool `json:"two_handed"`
}

type apiArmorResp struct {
	Index      string `json:"index"`
	Name       string `json:"name"`
	ArmorClass struct {
		Base     int  `json:"base"`
		DexBonus bool `json:"dex_bonus"`
		MaxBonus int  `json:"max_bonus"`
	} `json:"armor_class"`
}

var requestsPerSecond = 6

var requestTimeout = 8 * time.Second

var baseURL = "https://www.dnd5eapi.co/api"

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func ensureDir(path string) error {
	return os.MkdirAll(path, 0o755)
}

// fetching single URL w/ context + return bytes
func fetchURL(ctx context.Context, url string) ([]byte, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return nil, fmt.Errorf("http %d: %s", resp.StatusCode, string(body))
	}
	return io.ReadAll(resp.Body)
}

// helper functions for spells/armor/weapons
func FetchSpell(index string) (*apiSpellResp, error) {
	index = strings.ToLower(index)

	url := fmt.Sprintf("%s/spells/%s", baseURL, index)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	b, err := fetchURL(ctx, url)
	if err != nil {
		return nil, err
	}

	var r apiSpellResp
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func FetchWeapon(index string) (*apiWeaponResp, error) {
	index = strings.ToLower(index)

	url := fmt.Sprintf("%s/equipment/%s", baseURL, index)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	b, err := fetchURL(ctx, url)
	if err != nil {
		return nil, err
	}

	var r apiWeaponResp
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func FetchArmor(index string) (*apiArmorResp, error) {
	index = strings.ToLower(index)

	url := fmt.Sprintf("%s/equipment/%s", baseURL, index)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	b, err := fetchURL(ctx, url)
	if err != nil {
		return nil, err
	}

	var r apiArmorResp
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func FetchSpellsBatch(indexes []string) map[string]*apiSpellResp {
	results := make(map[string]*apiSpellResp)

	type result struct {
		idx string
		res *apiSpellResp
		err error
	}
	chIn := make(chan string)
	chOut := make(chan result)

	worker := func() {
		for idx := range chIn {
			r, err := FetchSpell(idx)
			chOut <- result{idx: idx, res: r, err: err}
		}
	}

	// start goroutines (parallel)
	workers := 5
	for i := 0; i < workers; i++ {
		go worker()
	}

	interval := time.Second / time.Duration(requestsPerSecond)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// goroutine
	go func() {
		for _, idx := range indexes {
			<-ticker.C
			chIn <- idx
		}
		close(chIn)
	}()

	//collect
	for i := 0; i < len(indexes); i++ {
		r := <-chOut
		if r.err != nil {
			fmt.Printf("error fetching %s: %v\n", r.idx, r.err)
			results[r.idx] = nil
			continue
		}
		results[r.idx] = r.res
	}

	cahceDir := filepath.Join("data", "api_cache")
	if err := ensureDir(cahceDir); err != nil {
		fmt.Println("failed ensuring cache directory: ", err)
		return results
	}

	cachePath := filepath.Join(cahceDir, "spells.json")
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal spells.json: ", err)
		return results
	}

	if err := os.WriteFile(cachePath, data, 0o644); err != nil {
		fmt.Println("failed to write spells.json: ", err)
		return results
	}

	fmt.Printf("Saved %d spells to %s\n", len(results), cachePath)

	return results
}

func FetchWeaponsBatch(indexes []string) map[string]*apiWeaponResp {
	results := make(map[string]*apiWeaponResp)

	type result struct {
		idx string
		res *apiWeaponResp
		err error
	}

	chIn := make(chan string)
	chOut := make(chan result)

	worker := func() {
		for idx := range chIn {
			r, err := FetchWeapon(idx)
			chOut <- result{idx: idx, res: r, err: err}
		}
	}

	workers := 5
	for i := 0; i < workers; i++ {
		go worker()
	}

	interval := time.Second / time.Duration(requestsPerSecond)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	go func() {
		for _, idx := range indexes {
			<-ticker.C
			chIn <- idx
		}
		close(chIn)
	}()

	for i := 0; i < len(indexes); i++ {
		r := <-chOut

		if r.err != nil {
			fmt.Printf("error fetching %s: %v\n", r.idx, r.err)
			results[r.idx] = nil
			continue
		}
		results[r.idx] = r.res
	}

	cacheDir := filepath.Join("data", "api_cache")
	if err := ensureDir(cacheDir); err != nil {
		fmt.Println("failed to ensure cache directory:", err)
		return results
	}

	cachePath := filepath.Join(cacheDir, "weapons.json")
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal weapons.json: ", err)
		return results
	}

	if err := os.WriteFile(cachePath, data, 0o644); err != nil {
		fmt.Println("failed to write weapons.json: ", err)
		return results
	}

	fmt.Printf("Saved %d weapons to %s\n", len(results), cachePath)

	return results
}

func FetchArmorBatch(indexes []string) map[string]*apiArmorResp {
	results := make(map[string]*apiArmorResp)

	type result struct {
		idx string
		res *apiArmorResp
		err error
	}

	chIn := make(chan string)
	chOut := make(chan result)

	worker := func() {
		for idx := range chIn {
			r, err := FetchArmor(idx)
			chOut <- result{idx: idx, res: r, err: err}
		}
	}

	workers := 5
	for i := 0; i < workers; i++ {
		go worker()
	}

	interval := time.Second / time.Duration(requestsPerSecond)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	go func() {
		for _, idx := range indexes {
			<-ticker.C
			chIn <- idx
		}
		close(chIn)
	}()

	for i := 0; i < len(indexes); i++ {
		r := <-chOut
		if r.err != nil {
			fmt.Printf("error fetching %s: %v\n", r.idx, r.err)
			results[r.idx] = nil
			continue
		}
		results[r.idx] = r.res
	}

	cacheDir := filepath.Join("data", "api_cache")
	if err := ensureDir(cacheDir); err != nil {
		fmt.Println("failed to ensure cache directory:", err)
		return results
	}

	cachePath := filepath.Join(cacheDir, "armor.json")
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal armor.json: ", err)
		return results
	}

	if err := os.WriteFile(cachePath, data, 0o644); err != nil {
		fmt.Println("failed to write armor.json: ", err)
		return results
	}

	fmt.Printf("Saved %d armor to %s\n", len(results), cachePath)

	return results
}
