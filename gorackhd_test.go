package gorackhd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/emccode/gorackhd/client"
	"github.com/emccode/gorackhd/client/lookups"
	"github.com/emccode/gorackhd/client/nodes"
)

func TestNodeGetOperation(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/api/1.1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	//use any function to do REST operations
	resp, err := client.Nodes.GetNodes(nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

func TestNodeLookupOperation(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/api/1.1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	nodeId := "56dde3441722c192796e3a38"
	//use any function to do REST operations
	resp, err := client.Lookups.GetLookups(&lookups.GetLookupsParams{Q: &nodeId}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

func TestNodePostOperation(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/api/1.1", []string{"http"})

	// configure the host
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	c := &Node{
		ID:   "1234abcd1234abcd1234abcd",
		Name: "somename",
		Type: "compute",
		ObmSettings: []*ObmSettings{&ObmSettings{
			Service: "ipmi-obm-service",
			Config: &ObmConfig{
				Host:     "1.2.3.4",
				User:     "myuser",
				Password: "mypass",
			},
		}},
	}

	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	resp, err := client.Nodes.PostNodes(&nodes.PostNodesParams{Identifiers: c}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

func TestNodeDeleteOperation(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/api/1.1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	//use any function to do REST operations
	resp, err := client.Nodes.DeleteNodesIdentifier(&nodes.DeleteNodesIdentifierParams{Identifier: "1234abcd1234abcd1234abcd"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp)
}

type ObmConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ObmSettings struct {
	Service string     `json:"service"`
	Config  *ObmConfig `json:"config"`
}

type Node struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	ObmSettings []*ObmSettings `json:"obmSettings"`
}
