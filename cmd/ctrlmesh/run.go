package main

import (
	"github.com/openziti-incubator/ctrlmesh"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run <config.yml>",
	Short: "Run ctrlmesh agent",
	Args:  cobra.ExactArgs(1),
	Run:   run,
}

func run(_ *cobra.Command, args []string) {
	cfg, err := ctrlmesh.Load(args[0])
	if err != nil {
		panic(err)
	}
	agent, err := ctrlmesh.NewAgent(cfg)
	if err != nil {
		panic(err)
	}
	if err := agent.Join(); err != nil {
		panic(err)
	}
	for {
		time.Sleep(5 * time.Second)
		agent.Status()
	}
}
