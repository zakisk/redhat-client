package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zakisk/redhat-client/network"
)

var file string

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove file froms server",
	Long: `
example:
	
store rm file.txt
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(file) == 0 {
			return fmt.Errorf("provide file name")
		}

		nc := network.NewNetworkCaller(&bytes.Buffer{})
		resp, err := nc.RemoveFile(file)
		if err != nil {
			return err
		}

		fmt.Println(resp.Message)

		return nil
	},
}

func init() {
	removeCmd.Flags().StringVarP(&file, "file", "f", "", "File to be updated")
}
