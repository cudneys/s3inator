/*
Copyright © 2020 Scott Cudney <scott@cudneys.net>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var awsSecretKey string
var awsAuthKey string
var awsRegion string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "s3inator",
	Short: "S3Inator Is An AWS S3 Bucket Reporting Tool",
	Long: `The s3inator is a Doofensmirtzian tool for AWS S3 that 
could help you take over the tri-state area, if it did everything
that it's supposed to.

Authentication is set up in one of two ways.  You can either use the 
built-in CLI flags (--authkey & --secretkey) to define your dreentials
or you can set the AWS_SECRET_ACCESS_KEY and AWS_ACCESS_KEY_ID envirronment
variables.

You can set your default region using the --region flag or you can set the 
AWS_REGION environment variable.

`,
	Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        if len(awsSecretKey) > 3 {
            os.Setenv("AWS_SECRET_ACCESS_KEY",awsSecretKey)
        }

        if len(awsAuthKey) > 3 {
            os.Setenv("AWS_ACCESS_KEY_ID",awsAuthKey)
        }

        if len(awsRegion) > 3 {
           os.Setenv("AWS_REGION",awsRegion)
        }
    },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.s3inator.yaml)")
    rootCmd.PersistentFlags().StringVar(&awsAuthKey,  "authkey","", "AWS Auth Key")
    rootCmd.PersistentFlags().StringVar(&awsSecretKey,  "secretkey","", "AWS Auth Key")
    rootCmd.PersistentFlags().StringVar(&awsRegion,  "region","us-east-1", "AWS Region")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".s3inator")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
