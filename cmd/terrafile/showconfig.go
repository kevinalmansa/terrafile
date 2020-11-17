package terrafile

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showconfigCmd represents the command
var showconfigCmd = &cobra.Command{
	Use:   "show-config",
	Short: "show loaded configuration",
	Long: `Shows the loaded configuration, taking into account command
parameters, configuration file, and default values.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%+v", configuration)
	},
}

func init() {
	rootCmd.AddCommand(showconfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showconfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showconfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
