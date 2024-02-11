package cmd

import (
	"github.com/spf13/cobra"
)

var file string

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove file froms server",
	Long: `
example:
	
store rm file.txt
`,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	
}
