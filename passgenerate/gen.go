/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package passgenerate

import (

	// "pass/passgenerate"

	"github.com/spf13/cobra"
)

var numberOfChars int
var nums bool
var specialchar bool
var use bool

var GenCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate the password of specified length with optional special characters and numbers.",
	Long: `Generate the password of specified length with optional special characters and numbers.
  Example: To generate a random password of length 10 with special characters and numbers, run:
  pass gen -l 10 -s -n`,
	Run: func(cmd *cobra.Command, args []string) {
		GeneratePassPrompt()
		// if use {
		// 	passwordGenerated := GenerateRandomString(numberOfChars, nums, specialchar)
		// 	functions.CopyToClipboard(passwordGenerated)
		// 	passwordFile := "/root/projects/pass/password.txt"
		// 	GeneratePassPrompt()
		// 	// if err != nil {
		// 	// 	fmt.Println("Failed to save password:", err)
		// 	// 	return
		// 	// }
		// 	fmt.Println("Password saved to", passwordFile)
		// 	return
		// }
		// passwordGenerated := algos.GenerateRandomString(numberOfChars, nums, specialchar)
		// fmt.Println(passwordGenerated+"  ", algos.GetPasswordStrengthLevel(passwordGenerated))
		// passwordGenerated = algos.GenerateRandomString(numberOfChars, nums, specialchar)
		// fmt.Println(passwordGenerated+"  ", algos.GetPasswordStrengthLevel(passwordGenerated))
		// passwordGenerated = algos.GenerateRandomString(numberOfChars, nums, specialchar)
		// fmt.Println(passwordGenerated+"  ", algos.GetPasswordStrengthLevel(passwordGenerated))
	},
}

func init() {
	GenCmd.Flags().IntVarP(&numberOfChars, "length", "l", 10, "Length of the password")
	GenCmd.Flags().BoolVarP(&nums, "nums", "n", false, "Include numbers in the password")
	GenCmd.Flags().BoolVarP(&specialchar, "specialchar", "s", false, "Include special characters in the password")
	GenCmd.Flags().BoolVarP(&use, "use", "u", false, "Copy the password to clipboard and don't print in screen")

}
