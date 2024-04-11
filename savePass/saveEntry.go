package savePass

import (
	"encoding/json"
	"os"
	"time"
)

type Entry struct {
	Aliasname      string   `json:"aliasname"`
	AlternateNames []string `json:"alternateNames"`
	Password       string   `json:"password"`
	Email          string   `json:"email"`
	Username       string   `json:"username"`
	URL            string   `json:"url"`
	Tags           []string `json:"tags"`
	Desc           string   `json:"desc"`
	Creation       string   `json:"creation"`
	Modified       string   `json:"modified"`
	Strength       string   `json:"strength"`
	Expiry         string   `json:"expiry"`
}

func SaveEntry(aliasname string, alternateNames []string, password string, email string, username string, url string, tags []string, desc string, strength string) error {
	// Read existing entries from file
	entries, err := readEntriesFromFile()
	if err != nil {
		return err
	}

	// Create new entry
	newEntry := Entry{
		Aliasname:      aliasname,
		AlternateNames: alternateNames,
		Password:       password,
		Email:          email,
		Username:       username,
		URL:            url,
		Tags:           tags,
		Desc:           desc,
		Creation:       time.Now().Format(time.RFC3339),
		Modified:       time.Now().Format(time.RFC3339),
		Strength:       strength,
		Expiry:         "",
	}

	// Append new entry to existing entries
	entries = append(entries, newEntry)

	// Write updated entries to file
	err = writeEntriesToFile(entries)
	if err != nil {
		return err
	}

	return nil
}

func readEntriesFromFile() ([]Entry, error) {
	// Open file for reading
	file, err := os.OpenFile("/root/projects/pass/functions/entries.json", os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode JSON content into entries
	var entries []Entry
	err = json.NewDecoder(file).Decode(&entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func writeEntriesToFile(entries []Entry) error {
	// Open file for writing
	file, err := os.OpenFile("/root/projects/pass/functions/entries.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode entries into JSON content with indentation and line breaks
	content, err := json.MarshalIndent(entries, "", "    ")
	if err != nil {
		return err
	}

	// Write JSON content to file
	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}
