/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v56/github"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// setRulesCmd represents the setRules command
var setRulesCmd = &cobra.Command{
	Use:   "set-rules",
	Short: "Set a list of rules in a repo branch",
	Long: `This command uses a YAML file with rule definitions and applies them to a repo branch`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setRules called")
		repo := cmd.Flag("repo").Value.String()
		token := cmd.Flag("token").Value.String()
		branch := cmd.Flag("branch").Value.String()
		fileName := cmd.Flag("file").Value.String()
		
		owner, repoName, err := parseRepo(repo)
		if err != nil {
			fmt.Println("Error parsing repo: ", err)
			return
		}

		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading YAML file: ", err)
			return
		}

		var protectionRequest github.ProtectionRequest
		err = yaml.Unmarshal(data, &protectionRequest)
		if err != nil {
			fmt.Println("Error unmarshaling YAML file: ", err)
			return
		}

		client := github.NewClient(nil).WithAuthToken(token)

		_, _, err = client.Repositories.UpdateBranchProtection(context.Background(), owner, repoName, branch, &protectionRequest)
		if err != nil {
			fmt.Println("Error updating branch protection rules: ", err)
			return
		}

		fmt.Println("Branch protection rules updated successfully")

	},
}

func init() {
	githubCmd.AddCommand(setRulesCmd)
	setRulesCmd.Flags().StringP("repo", "r", "", "Name of the repo")
	setRulesCmd.Flags().StringP("token", "t", "", "PAT token")
	setRulesCmd.Flags().StringP("branch", "b", "main", "Repo branch")
	setRulesCmd.Flags().StringP("file", "f", "", "YAML file with rules")

	setRulesCmd.MarkFlagRequired("repo")
	setRulesCmd.MarkFlagRequired("token")
	setRulesCmd.MarkFlagRequired("file")
	setRulesCmd.MarkFlagRequired("branch")
}
