package main

import (
	"github.com/hashicorp/terraform/plugin"
	kubernetes "github.com/lawrencegripper/terraform-provider-kubernetes-yaml/kubernetes"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kubernetes.Provider})
}
