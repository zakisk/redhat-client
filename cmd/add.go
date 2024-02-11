package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/zakisk/redhat-client/network"
)

var files []string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add uploads files to server",
	Long: `
examples of command. For example:
	
// upload two files to server
store add -f file1.txt -f file2.txt
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(files) == 0 {
			cmd.Help()
			return fmt.Errorf("arguments can't be empty, pass file names")
		}

		nc := network.NewNetworkCaller(&http.Client{}, &bytes.Buffer{})
		success, failed := 0, 0
		for _, file := range files {
			if !isFileExist(file) {
				fmt.Printf("There is no such file `%s`\n", file)
				failed++
				continue
			}

			resp, err := nc.CreateFile(file)
			if err != nil {
				fmt.Printf("Error while uploading file\nerror: `%s`", err.Error())
				failed++
				continue
			}
			fmt.Println(resp.Message)
			success++
		}
		fmt.Printf("Success: %d\tFailed: %d\n", success, failed)
		return nil
	},
}

func init() {
	addCmd.Flags().StringArrayVarP(&files, "file", "f", []string{}, "Files to be uploaded on server")
}

func isFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
