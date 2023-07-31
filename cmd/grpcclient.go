/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"mycaserver/pkg/grpc/client"
)

// grpcclientCmd represents the grpcclient command
var grpcclientCmd = &cobra.Command{
	Use:   "grpcclient",
	Short: "start a gRPC client to test CA web server",
	Long: `Start a gRPC client to execute a test-call on CA web server, the server should running in gRPC mode - not http mode，
	and must enable the mTLS`,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	client.Run(*certId)
}

var certId *string

func init() {
	rootCmd.AddCommand(grpcclientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcclientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcclientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	certId = grpcclientCmd.LocalFlags().String("certid", "", "certid id of the client certificate")
	grpcclientCmd.MarkFlagRequired("certid")
}
