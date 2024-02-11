package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zakisk/redhat-client/network"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates files",
	Long: `
example:

store update file.txt
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(file) == 0 {
			return fmt.Errorf("provide file name")
		}

		nc := network.NewNetworkCaller(&bytes.Buffer{})
		resp, err := nc.UpdateFile(file)
		if err != nil {
			return err
		}

		fmt.Println(resp.Message)

		return nil
	},
}

func init() {
	updateCmd.Flags().StringVarP(&file, "file", "f", "", "File to be updated")
}
