package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configDirName string
var consulDC string
var consulToken string
var consulAddr string
var consulPrefix string
var consulCAFile string
var consulCAPath string
var consulCertFile string
var consulKeyFile string
var consulInsecure bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "skewer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&configDirName, "config", "/etc/skewer", "configuration directory")
	RootCmd.PersistentFlags().StringVar(&consulAddr, "consul-addr", "", "Consul address (ex: http://127.0.0.1:8500)")
	RootCmd.PersistentFlags().StringVar(&consulDC, "consul-dc", "", "Consul datacenter")
	RootCmd.PersistentFlags().StringVar(&consulToken, "consul-token", "", "Consul token")
	RootCmd.PersistentFlags().StringVar(&consulPrefix, "consul-prefix", "skewer", "Where to find configuration in Consul KV")
	RootCmd.PersistentFlags().StringVar(&consulCAFile, "consul-ca-file", "", "optional path to the CA certificate used for Consul")
	RootCmd.PersistentFlags().StringVar(&consulCAPath, "consul-ca-path", "", "optional path to a directory of CA certificates to use for Consul")
	RootCmd.PersistentFlags().StringVar(&consulCertFile, "consul-cert-file", "", "optional path to the client certificate for Consul")
	RootCmd.PersistentFlags().StringVar(&consulKeyFile, "consul-key-file", "", "optional path to the client private key for Consul")
	RootCmd.PersistentFlags().BoolVar(&consulInsecure, "consul-insecure", false, "if set to true will disable TLS host verification")
}
