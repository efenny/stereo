package cmd

import (
	"fmt"
	"os"
	stereo_data "stereo-server/internal/stereo-data"

	"github.com/spf13/cobra"
)

var (
	json     string
	filePath string
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Will process input JSON and return markdown output.",
	Long: `Will process input JSON and return markdown output.

	This supports:
	- stdin
	- raw JSON strings
	- JSON files`,
	Args: cobra.NoArgs,
	Run:  processInput,
}

func init() {
	rootCmd.AddCommand(processCmd)
	processCmd.Flags().StringVar(&json, "json", "", "raw JSON string")
	processCmd.Flags().StringVarP(&filePath, "file", "f", "", "JSON file")
}

func processInput(cmd *cobra.Command, args []string) {
	var jsonString string

	if filePath == "" && json == "" {
		fmt.Print("Nothing to process. Please check inputs\n")
		os.Exit(1)
	}

	if filePath != "" {
		data, err := os.ReadFile(filePath)

		if err != nil {
			fmt.Printf("There was an error reading the file: %s", err)
			os.Exit(1)
		}

		jsonString = string(data)
	}

	if len(jsonString) > 0 {
		data := stereo_data.NewData(jsonString)

		data.ProcessJsonFromInputString()
		data.GetInputData()
		markdown := data.ConvertToMarkdown()

		_, err := os.Stdout.WriteString(markdown)
		if err != nil {
			fmt.Printf("Error writing to Stdout: %s", err)
		}
	} else {
		fmt.Print("Nothing to process. Please check inputs\n")
	}

	os.Exit(0)
}
