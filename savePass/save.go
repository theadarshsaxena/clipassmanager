/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package savePass

import (
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var SaveCmd = &cobra.Command{
	Use:   "save",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SavePrompt()
		// err := SaveEntry("aliasname", []string{"alternateNames"}, "password", "email", "username", "url", []string{"tags"}, "desc", "strength")
		// if err != nil {
		// 	panic(err)
		// }
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
