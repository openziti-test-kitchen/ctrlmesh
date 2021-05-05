package main

import (
	"github.com/openziti-incubator/ctrlmesh"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(peerlistCmd)
}

var peerlistCmd = &cobra.Command{
	Use:   "peerlist <config.yml>",
	Short: "Run peerlist agent",
	Args:  cobra.ExactArgs(1),
	Run:   peerlist,
}

func peerlist(_ *cobra.Command, args []string) {
	if _, err := ctrlmesh.LoadPeerListConfigYaml(args[0]); err != nil {
		panic(err)
	}
}
