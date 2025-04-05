/*
Copyright © 2025 Kdaito
*/
package cmd

import (
	"log"

	"github.com/Kdaito/blogen/internal/blog"
	"github.com/Kdaito/blogen/internal/utils"
	"github.com/spf13/cobra"
)

var workspace string

const MAX_ID_LENGTH = 50

// indivCmd represents the indiv command
var indivCmd = &cobra.Command{
	Use:   "indiv",
	Short: "Fetch articles by specifying their IDs",
	Long:  "The `indiv` command allows you to fetch one or more articles by specifying their IDs. You can retrieve multiple articles at once by passing multiple IDs.",
	Run: func(cmd *cobra.Command, ids []string) {
		completeItems := make([]string, 0, len(ids))
		unCompleteItems := make([]string, 0)

		// Check if workspace is provided
		if err := utils.CheckWorkspaceExist(workspace); err != nil {
			log.Fatal(err)
		}

		// Check id length
		if len(ids) > MAX_ID_LENGTH {
			log.Fatalf("You can only fetch up to %d articles at a time.", MAX_ID_LENGTH)
		}

		for _, id := range ids {
			err := blog.GenerateHtmlIndiv(id, workspace)

			if err != nil {
				log.Print(err)
				unCompleteItems = append(unCompleteItems, id)
				continue
			}

			completeItems = append(completeItems, id)
		}

		blog.Output(completeItems, unCompleteItems)
	},
}

func init() {
	rootCmd.AddCommand(indivCmd)

	indivCmd.Flags().StringVarP(&workspace, "workspace", "w", "", "Path to the workspace directory for output files")
}
