/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "A command to interact with github",
	Long: `This is a command to interact with github
				in a easy way to configure the differente repos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("github called")
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// githubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
