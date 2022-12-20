package wdp

import (
	"fmt"
	"os"

	"github.com/sho7a/wdp/pkg/wdp"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "wdp",
	Short:   "wdp - Web Development Proxy",
	Long:    "A Web Development Proxy with live reload capabilities.",
	Version: "0.0.1",
	Run:     run,
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	wdp.Start()
}
