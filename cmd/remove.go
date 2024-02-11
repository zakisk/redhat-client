package cmd

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "add",
	Short: "Add uploads files to server",
	Long: `
examples of command. For example:
	
// upload two files to server
store add file1.txt file2.txt
`,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	
}
