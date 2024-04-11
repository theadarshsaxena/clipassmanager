package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type SearchEntry struct {
	Name           string   `json:"aliasname"`
	AlternateNames []string `json:"alternateNames"`
}

func SearchAlias(name string) error {
	// Open the file for read only
	file, err := os.Open("/root/projects/pass/functions/entries.json")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal the JSON data
	var entries []SearchEntry
	err = json.Unmarshal(content, &entries)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	foundPositions, err := FuzzySearchPositions(name, entries)
	if err != nil {
		return fmt.Errorf("entry not found")
	}
	PrintFinalResult(name, entries, foundPositions)
	return nil
}

// FuzzySearchAlias searches for a name using fuzzy matching
func FuzzySearchPositions(name string, entries []SearchEntry) ([]int, error) {
	var foundNames []int
Outer:
	for foundPosition, entry := range entries {
		if fuzzy.Match(name, entry.Name) {
			foundNames = append(foundNames, foundPosition)
			continue Outer
		}
		if len(fuzzy.Find(name, entry.AlternateNames)) > 0 {
			foundNames = append(foundNames, foundPosition)
			continue Outer
		}
	}
	return foundNames, nil
}

func PrintFuzzyColoredResult(name string, entry string) string {
	var result string
	for _, t1 := range name {
		for i, t2 := range entry {
			if t1 == t2 {
				entry = entry[i+utf8.RuneLen(t2):]
				result += fmt.Sprintf("\033[1;31m%c\033[0m", t1)
				break
			}
			result += fmt.Sprintf("%c", t2)
			if i == len(entry)-1 {
				return result
			}
		}
	}
	result += entry
	return result
}

func PrintFinalResult(name string, entries []SearchEntry, foundNames []int) {
	maxNameLength := 0
	for _, foundPosition := range foundNames {
		entry := entries[foundPosition]
		nameLength := len(entry.Name)
		if nameLength > maxNameLength {
			maxNameLength = nameLength
		}
	}
	fmt.Printf("\033[36m#  %-*s  %-s\033[0m\n", maxNameLength, "Name", "Alternate Names")

	for i, foundPosition := range foundNames {
		entry := entries[foundPosition]
		fmt.Printf("%-2d %-*s ", i+1, maxNameLength, entry.Name)
		for _, altName := range entry.AlternateNames {
			fmt.Printf(" %s", PrintFuzzyColoredResult(name, altName))
		}
		fmt.Println()
	}
}
