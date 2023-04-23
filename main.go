package main

import (
	"github.com/spf13/cobra"
	hssm "github.com/ugurcancaykara/helm-ssm-sm/internal"
	"log"
	"text/template"
)

var ssmFlag bool
var smFlag bool
var dryRun bool
var fileFlag bool
var verbose bool
var tpl *template.Template
var valueFile string

var rootCmd = &cobra.Command{
	Use:   "ssm",
	Short: "Fetch parameter value from AWS SSM Parameter Store",
	Run: func(cmd *cobra.Command, args []string) {
		hssm.ProcessTemplate(tpl, valueFile, ssmFlag, smFlag, verbose, dryRun)

	},
}

func init() {
	rootCmd.Flags().BoolVar(&ssmFlag, "ssm", false, "Enable SSM Parameter store")
	rootCmd.Flags().BoolVar(&smFlag, "sm", false, "Enable Secrets Manager")
	rootCmd.Flags().BoolVar(&verbose, "v", false, "Enable verbose output")
	rootCmd.Flags().BoolVar(&dryRun, "d", false, "Don't replace the file content")
	rootCmd.PersistentFlags().StringVarP(&valueFile, "file", "f", "", "YAML file to process")

	rootCmd.MarkFlagRequired("file")
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
