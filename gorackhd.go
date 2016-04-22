package gorackhd

import (
	// import go-swagger so it can generate the bindings
	_ "github.com/go-swagger/go-swagger"

	// import the generated go-swagger bindings
	"github.com/emccode/gorackhd/client"
	_ "github.com/emccode/gorackhd/models"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func GetClient(host, basePath string, schemes []string) *client.Monorail {
	return client.New(httptransport.New(host, basePath, schemes), strfmt.Default)
}
