package terrafile

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/kevinalmansa/terrafile/pkg/cache"
	terrafile "github.com/kevinalmansa/terrafile/pkg/config"
	"github.com/spf13/viper"
)

var cfgFile string
var configuration terrafile.Config
var terraCache cache.Cache

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terrafile",
	Short: "Version control terraform modules",
	Long: `terrafile is a CLI command to enable dynamic versioning of terraform
modules stored in git without modifying the terraform code. The aim is to
simplify development of terraform modules and integration within CI/CD
solutions.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is ./terrafile.yaml)")
	rootCmd.PersistentFlags().StringVar(&configuration.CacheDir, "cache", "modules",
		"cache directory (default is ./terrafile/modules)")
	rootCmd.PersistentFlags().StringVar(&configuration.Branch, "branch", "master",
		"branch to checkout for modules (default is main)")
	rootCmd.PersistentFlags().StringVar(&configuration.Tag, "tag", "",
		"tag to checkout for modules (default is unset - overrides branch)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//Set Viper flags from Cobra
	viper.Set("Cache", configuration.CacheDir)
	viper.Set("Branch", configuration.Branch)
	viper.Set("Tag", configuration.Tag)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("terrafile")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Error decoding config file: %s", err)
	}
	terraCache = &cache.LocalModuleCache{Configuration: &configuration}
}
