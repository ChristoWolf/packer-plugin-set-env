// package main registers and runs the plugin.
package main

import (
	"log"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	v "github.com/hashicorp/packer-plugin-sdk/version"
	"github.com/rgl/packer-plugin-windows-update/update"
)

var (
	version = "0.0.0"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	log.Println("Starting packer-plugin-set-env")
	pps := plugin.NewSet()
	pps.RegisterProvisioner(plugin.DEFAULT_NAME, new(update.Provisioner))
	pps.SetVersion(v.InitializePluginVersion(version, ""))
	err := pps.Run()
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
