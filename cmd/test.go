package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "TCPScan",
	Short: "简易网络嗅探器",
}

func init() {
	rootCmd.AddCommand(host)
}

func Excute() error {
	return rootCmd.Execute()
}
