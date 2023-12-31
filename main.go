package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"terraform-provider-dptech-demo/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string
)

func main() {

	fmt.Println("test content print")

	fmt.Print("aaaaaaaaaaaaaaa")
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like ")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "xieguihua123/dptech-demo",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
