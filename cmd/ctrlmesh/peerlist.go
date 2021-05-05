package main

import (
	"github.com/openziti-incubator/ctrlmesh"
	"github.com/spf13/cobra"
	"time"
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
	cfg, err := ctrlmesh.LoadPeerListConfigYaml(args[0])
	if err != nil {
		panic(err)
	}
	_, err = ctrlmesh.NewPeerList(cfg)
	for {
		time.Sleep(30 * time.Second)
	}
}
