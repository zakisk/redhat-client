package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zakisk/redhat-client/network"
	"github.com/zakisk/redhat-client/utils"
)

var wordCountCmd = &cobra.Command{
	Use:   "wc",
	Short: "Counts words in all files on server",
	Long: `
example:
	
store wc
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		spinner := utils.CreateDefaultSpinner("Loading", "")
		spinner.Start()
		nc := network.NewNetworkCaller(&bytes.Buffer{})
		resp, err := nc.GetWordsCount()
		if err != nil {
			spinner.Stop()
			return err
		}
		
		spinner.Stop()
		fmt.Printf("\n%d unique words are found in %d files\n", resp.AllWordsCount, resp.AllFilesProcessed)
		return nil
	},
}

func init() {

}
