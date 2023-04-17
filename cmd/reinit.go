// Package cmd ...
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// reinitCmd represents the reinit command
var reinitCmd = &cobra.Command{
	Use:   "reinit",
	Short: "Overwrites the .am file with a default/blank version",
	Long: `Reinit (am reinit) will overwrite the .am file located in the home directory
with the default version of it. The user will be asked for confirmation before
the .am file is replaced if it already exists.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reinit called")
	},
}

func init() {
	rootCmd.AddCommand(reinitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reinitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reinitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
