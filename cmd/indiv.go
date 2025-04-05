/*
Copyright © 2025 Kdaito
*/
package cmd

import (
	"log"

	"github.com/Kdaito/blogen/internal/blog"
	"github.com/spf13/cobra"
)

var workspace string

// indivCmd represents the indiv command
var indivCmd = &cobra.Command{
	Use:   "indiv",
	Short: "idを指定して、個別の記事を取得することができます",
	Long:  "`indiv`コマンドは、指定したidを持つ個別の記事を取得するためのコマンドです。idは複数指定することができ、各記事の詳細情報を取得することができます",
	Run: func(cmd *cobra.Command, ids []string) {
		completeItems := make([]string, 0, len(ids))
		unCompleteItems := make([]string, 0)

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

	indivCmd.Flags().StringVarP(&workspace, "workspace", "w", "", "htmlを出力したいディレクトリのパスを指定します")
}
