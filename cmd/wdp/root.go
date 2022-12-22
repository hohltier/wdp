package wdp

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/sho7a/wdp/internal/wdp"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "wdp",
	Short:   "wdp - Web Development Proxy",
	Long:    "A Web Development Proxy with live reload capabilities.",
	Version: "0.1.2",
	Run:     run,
}

func Execute() {
	cmd.PersistentFlags().IntVarP(&wdp.Listen, "listen", "l", 0, "listen port (default open port)")
	cmd.PersistentFlags().IntVarP(&wdp.Port, "port", "p", 80, "server port")
	cmd.PersistentFlags().StringVarP(&wdp.Watch, "watch", "w", ".", "watch path")
	if err := cmd.Execute(); err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	figure.NewFigure("wdp", "small", true).Print()
	fmt.Println()
	go wdp.Watcher()
	wdp.Server()
}
