package cmd

import (
	"bytes"

	"github.com/spf13/cobra"
	"github.com/zakisk/redhat-client/network"
	"github.com/zakisk/redhat-client/utils"
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "lists all files of server",
	Long: `
example:
	
store ls
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nc := network.NewNetworkCaller(&bytes.Buffer{})
		resp, err := nc.ListFiles()
		if err != nil {
			return err
		}

		rows := [][]string{}
		for _, file := range resp.Files {
			// output will be similar to `ls -lh` command on unix-system
			rows = append(rows, []string{file.Mode, file.ModifiedAt, file.Name, file.Size})
		}

		utils.PrintToTable([]string{}, rows)
		return nil
	},
}

func init() {

}
