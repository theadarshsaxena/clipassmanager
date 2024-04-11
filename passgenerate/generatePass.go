package passgenerate

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(n int, nums bool, specialChar bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if nums {
		letterRunes = append(letterRunes, []rune("0123456789")...)
	}
	if specialChar {
		letterRunes = append(letterRunes, []rune("!@#$%^&*()_+{}[]:;<>,.?/~")...)
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}

	return string(b)
}
func GetPasswordStrengthLevel(password string) string {
	// Define the criteria for each strength level
	weakCriteria := []string{"abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	strongCriteria := []string{"0123456789"}
	veryStrongCriteria := []string{"!@#$%^&*()_+{}[]:;<>,.?/~"}

	// Check if the password meets the criteria for each strength level
	if containsAll(password, weakCriteria, strongCriteria, veryStrongCriteria) {
		if len(password) < 5 {
			return fmt.Sprintf("\033[31m%s\033[0m", "Weak")
		} else if len(password) < 8 {
			return fmt.Sprintf("\033[33m%s\033[0m", "Moderate")
		} else if len(password) < 14 {
			return fmt.Sprintf("\033[32m%s\033[0m", "Strong")
		} else {
			return fmt.Sprintf("\033[32m%s\033[0m", "Very Strong")
		}
	} else if containsAll(password, weakCriteria, strongCriteria) {
		if len(password) < 6 {
			return fmt.Sprintf("\033[31m%s\033[0m", "Weak")
		} else if len(password) < 12 {
			return fmt.Sprintf("\033[33m%s\033[0m", "Moderate")
		} else if len(password) < 18 {
			return fmt.Sprintf("\033[32m%s\033[0m", "Strong")
		} else {
			return fmt.Sprintf("\033[32m%s\033[0m", "Very Strong")
		}
	} else if containsAll(password, weakCriteria, veryStrongCriteria) {
		if len(password) < 6 {
			return fmt.Sprintf("\033[31m%s\033[0m", "Weak")
		} else if len(password) < 12 {
			return fmt.Sprintf("\033[33m%s\033[0m", "Moderate")
		} else if len(password) < 18 {
			return fmt.Sprintf("\033[32m%s\033[0m", "Strong")
		} else {
			return fmt.Sprintf("\033[32m%s\033[0m", "Very Strong")
		}
	} else if containsAll(password, weakCriteria) {
		if len(password) < 6 {
			return fmt.Sprintf("\033[31m%s\033[0m", "Very Weak")
		} else if len(password) < 10 {
			return fmt.Sprintf("\033[31m%s\033[0m", "Weak")
		} else if len(password) < 14 {
			return fmt.Sprintf("\033[33m%s\033[0m", "Moderate")
		} else if len(password) < 18 {
			return fmt.Sprintf("\033[32m%s\033[0m", "Strong")
		} else {
			return fmt.Sprintf("\033[32m%s\033[0m", "Very Strong")
		}
	}

	// If the password doesn't meet any of the criteria, consider it as "Unknown"
	return "Unknown"
}

func containsAll(password string, criteria ...[]string) bool {
	// fmt.Println("checking for ", criteria, " in ", password)
	for _, c := range criteria {
		found := false
		// fmt.Println("checking for ", c, " in ", password)
		for _, charset := range c {
			for _, char := range charset {
				if strings.Contains(password, string(char)) {
					found = true
					break
				}
			}
		}
		if !found {
			return false
		}
	}
	return true
}
