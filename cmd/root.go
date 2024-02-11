/*
Copyright © 2024 Mohammed Zaki
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "store",
	Short: "A text files storage utility tool",
	Long: `A CLI app to perform CRUD operations on text files alongwith other textual operations
examples of application. For example:
	
// upload two files to server
store add file1.txt file2.txt

// list all files
store ls

// remove a file
store rm file.txt

// update a file
store update file.txt

// count all words in all files
store wc

// list N most frequent words in order [asc | dsc]
store freq-words [--limit|-n 10] [--order=dsc|asc]
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(addCmd, deployCmd, frequentWordsCmd, listCmd, removeCmd, updateCmd)
}
