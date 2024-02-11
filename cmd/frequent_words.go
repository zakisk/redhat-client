package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zakisk/redhat-client/network"
	"github.com/zakisk/redhat-client/utils"
)

var (
	n     int
	order string
)

var frequentWordsCmd = &cobra.Command{
	Use:   "freq-words",
	Short: "Prints N most frequent words on server",
	Long: `
examples of command. For example:
	
// example
store freq-words -n 10
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		spinner := utils.CreateDefaultSpinner("Loading", "")
		spinner.Start()
		nc := network.NewNetworkCaller(&bytes.Buffer{})
		resp, err := nc.GetFrequentWords(n, order)
		if err != nil {
			spinner.Stop()
			return err
		}

		spinner.Stop()
		for k, v := range resp.Words {
			fmt.Printf("Word: %s\tOccurrence: %d\n", k, v)
		}

		return nil
	},
}

func init() {
	frequentWordsCmd.Flags().IntVarP(&n, "no", "n", 10, "Number of frequent words")
	frequentWordsCmd.Flags().StringVarP(&order, "order", "o", "asc", "Order of words")
}
