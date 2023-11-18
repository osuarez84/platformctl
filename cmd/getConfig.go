/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v56/github"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// getConfigCmd represents the getConfig command
var getConfigCmd = &cobra.Command{
	Use:   "get-rules",
	Short: "Get the rules of a github repo branch",
	Long: `This command will get the rules of a github repo branch and store it in a file`,
	Run: func(cmd *cobra.Command, args []string) {
		token := cmd.Flag("token").Value.String()
		repo := cmd.Flag("repo").Value.String()
		branch := cmd.Flag("branch").Value.String()

		rules, err := getRules(repo, token, branch)
		if err != nil {
			fmt.Println("Error getting branch protection rules: ", err)
			return
		}

		err = writeToYaml(rules)
		if err != nil {
			fmt.Println("Error writing to YAML: ", err)
			return
		}

	},
}

func init() {
	githubCmd.AddCommand(getConfigCmd)
	getConfigCmd.Flags().StringP("repo", "r", "", "Name of the repo")
	getConfigCmd.Flags().StringP("token", "t", "", "PAT token")
	getConfigCmd.Flags().StringP("branch", "b", "main", "Repo branch")
}

func getRules(repo, token, branch string) (*github.Protection, error) {
	client := github.NewClient(nil).WithAuthToken(token)

	owner, repoName, err := parseRepo(repo)
	if err != nil {
		return nil, err
	}

	rules, _, err := client.Repositories.GetBranchProtection(context.Background(), owner, repoName, branch)
	if err != nil {
		return nil, err
	}

	return rules, nil
}


func parseRepo(repo string) (string, string, error) {
	parts := strings.Split(repo, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("repo should be in the format 'onwer/repo'")
	}
	return parts[0], parts[1], nil
}


func writeToYaml(rules interface{}) error {
	yamlProtection, err := yaml.Marshal(&rules)
	if err != nil {
		return err
	}

	// Write YAML to a file
	file, err := os.Create("protection.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(yamlProtection)
	if err != nil {
		return err
	}

	fmt.Println("Configuration saved to protection.yaml")
	return nil
}