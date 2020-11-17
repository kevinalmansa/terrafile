package terrafile

import (
	"log"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install/update modules based on config file",
	Long: `Install and or update modules defined in config file. Any missing
modules will be downloaded, but previously cached modules will not be removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Installing modules...")
		if err := terraCache.Create(); err != nil {
			log.Printf("Error installing modules: %s", err)
		} else {
			log.Printf("Modules successfully installed.")
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
