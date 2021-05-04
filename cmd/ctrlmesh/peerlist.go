package main

import (
	"github.com/openziti-incubator/ctrlmesh"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	data, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	dataMap := make(map[string]interface{})
	if err := yaml.Unmarshal(data, dataMap); err != nil {
		panic(err)
	}
	if _, err = ctrlmesh.LoadPeerListConfig(dataMap); err != nil {
		panic(err)
	}
}
