package cmd

import (
	"TCPScan/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	h     string
	start int
	end   int
)

var host = &cobra.Command{
	Use:   "host",
	Short: "端口扫描",
	Long:  "针对单一主机的端口扫描",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("address:", h)
		address := fmt.Sprintf("%s:%%d", h)
		pkg.StartPort(address)
	},
}

func init() {
	host.Flags().StringVarP(&h, "address", "a", "127.0.0.1", "请输入主机地址")
}
