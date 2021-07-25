package cmd

import (
	"daikincli/cmd/get"
	"daikincli/cmd/set"
	"daikincli/internal/dclilog"
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

var (
	hash    string
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "dcli",
		Short: "dcli",
		Long: `

    ____          _  __    _                      __ _ 
   / __ \ ____ _ (_)/ /__ (_)____          _____ / /(_)
  / / / // __  // // //_// // __ \ ______ / ___// // / 
 / /_/ // /_/ // // ,<  / // / / //_____// /__ / // /  
/_____/ \__,_//_//_/|_|/_//_/ /_/        \___//_//_/   
                                                       
=======================================================
	cli tool for Daikin air controller`,
	}
)

// Execute adds all child commands to the root command.
func Execute(version, commit string) {
	rootCmd.Version = version
	hash = commit

	setVersion()

	if err := rootCmd.Execute(); err != nil {
		dclilog.GetInstance().Error(err)
		os.Exit(1)
	}
}

func setVersion() {
	template := fmt.Sprintf("dcli version: %s, commit: %s \n", rootCmd.Version, hash)
	rootCmd.SetVersionTemplate(template)
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "set logging level to verbose")

	rootCmd.AddCommand(set.New())
	rootCmd.AddCommand(get.New())

}
