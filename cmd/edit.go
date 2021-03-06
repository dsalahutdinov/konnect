package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// EditCmd - Open the konnect.yml config file with the default editor.
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the config file with the default editor",
	Long:  "Open the config file with the default editor",
	Run: func(cmd *cobra.Command, args []string) {
		// Resolve filename from flags.
		filename, err := resolveFilename(cmd)
		handleErr(err)

		fmt.Printf("Opening config file at %v\n", filename)

		if err := exec.Command("open", filename).Run(); err != nil {
			log.Fatal("Error when opening the config file.")
		}
	},
}
