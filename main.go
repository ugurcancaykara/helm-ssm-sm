package main

import (
	"github.com/spf13/cobra"
	"log"
	"text/template"
)

var ssmFlag bool
var smFlag bool
var fileFlag bool
var verbose bool
var tpl *template.Template
var valueFile string

var rootCmd = &cobra.Command{
	Use:   "ssm",
	Short: "Fetch parameter value from AWS SSM Parameter Store",
	Run: func(cmd *cobra.Command, args []string) {
		processTemplate(tpl, valueFile, ssmFlag, smFlag, verbose)
	},
}

func init() {
	rootCmd.Flags().BoolVar(&ssmFlag, "ssm", false, "Enable SSM Parameter store")
	rootCmd.Flags().BoolVar(&smFlag, "sm", false, "Enable Secrets Manager")
	rootCmd.Flags().BoolVar(&verbose, "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVarP(&valueFile, "file", "f", "", "YAML file to process")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
