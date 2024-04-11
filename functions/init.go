package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitialisePass(key string) error {
	// Create the ~/.pass directory if it doesn't exist
	configDir, err := GetPassConfigDir()
	if err != nil {
		return fmt.Errorf("Error getting the home dir, %w", err)
	}
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("Error creating the key directory, %w", err)
	}

	// Create the ~/.pass/passdb directory if it doesn't exist
	passDBDir, err := GetPassDBPath()
	if err != nil {
		return fmt.Errorf("Error getting the passdb dir, %w", err)
	}
	err = os.MkdirAll(filepath.Dir(passDBDir), 0755)
	if err != nil {
		return fmt.Errorf("Error creating the passdb directory, %w", err)
	}

	// Create the ~/.passdb/keys.json file if it doesn't exist
	_, err = os.Create(passDBDir)
	if err != nil {
		return fmt.Errorf("Error creating the keys.json file, %w", err)
	}
	return nil
	// Create the ~/.pass/config.json file if it doesn't exist
	// Create the ~/.pass/keys.json file if it doesn't exist

}

func GetPassConfigDir() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting the home dir")
		return "", err
	}
	return filepath.Join(homedir, ".pass"), nil
}

func GetPassDBPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting the home dir")
		return "", err
	}
	return filepath.Join(homedir, ".pass", "db.json"), nil
}

func GetEncryptionKeyStoragePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting the home dir")
		return "", err
	}
	return filepath.Join(homedir, ".passdb", "keys.json"), nil
}
