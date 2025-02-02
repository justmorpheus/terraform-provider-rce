package main

import (
    "context"
    "flag"
    "log"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"
    "terraform-provider-custom/internal/provider"
)

func main() {
    var debug bool
    flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
    flag.Parse()

    opts := providerserver.ServeOpts{
        Address: "example.com/yourusername/custom",
        Debug:   debug,
    }

    err := providerserver.Serve(context.Background(), provider.New, opts)
    if err != nil {
        log.Fatal(err.Error())
    }
}
