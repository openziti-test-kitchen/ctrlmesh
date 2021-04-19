package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	pfxlog.Global(logrus.InfoLevel)
	pfxlog.SetPrefix("github.com/openziti-incubator/")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])),
	Short: "Distributed Control Plane",
}
