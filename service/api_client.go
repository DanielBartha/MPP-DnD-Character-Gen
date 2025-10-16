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
	ArmorClass []struct {
		Value int `json:"value"`
	} `json:"armor_class"`
	ArmorClassDexBonus struct {
		Type string `json:"type"`
	} `json:"-"`
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

// cached response path helpers
func spellCachePath(index string) string {
	return filepath.Join("data", "api_cache", "spells", index+".json")
}

func weaponCachePath(index string) string {
	return filepath.Join("data", "api_cache", "weapons", index+".json")
}

func armorCachePath(index string) string {
	return filepath.Join("data", "api_cache", "armor", index+".json")
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

// public helper functions: fetch Spell/Weapon/Armor
// will check cache first and return chached file if present; else request and cache
func FetchSpell(index string) (*apiSpellResp, error) {
	index = strings.ToLower(index)
	cachePath := spellCachePath(index)

	if data, err := os.ReadFile(cachePath); err == nil {
		var r apiSpellResp
		if err := json.Unmarshal(data, &r); err == nil {
			return &r, nil
		}
		// when corrupt, falls back to re-fetch
	}

	// build url + fetch
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

	// ensure cache dir
	_ = ensureDir(filepath.Dir(cachePath))
	_ = os.WriteFile(cachePath, b, 0o644)

	return &r, nil
}

func FetchWeapon(index string) (*apiWeaponResp, error) {
	index = strings.ToLower(index)
	cachePath := weaponCachePath(index)

	if data, err := os.ReadFile(cachePath); err == nil {
		var r apiWeaponResp
		if err := json.Unmarshal(data, &r); err == nil {
			return &r, nil
		}
	}

	url := fmt.Sprintf("%s/weapons/%s", baseURL, index)
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

	_ = ensureDir(filepath.Dir(cachePath))
	_ = os.WriteFile(cachePath, b, 0o644)

	return &r, nil
}

func FetchArmor(index string) (*apiArmorResp, error) {
	index = strings.ToLower(index)
	cachePath := armorCachePath(index)

	if data, err := os.ReadFile(cachePath); err == nil {
		var r apiArmorResp
		if err := json.Unmarshal(data, &r); err == nil {
			return &r, nil
		}
	}

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

	_ = ensureDir(filepath.Dir(cachePath))
	_ = os.WriteFile(cachePath, b, 0o644)

	return &r, nil
}

// many many moons ago, many many spells were created which is why this batching is needed
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

	// rate limiter
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

	return results
}
