#gorackhd

### Go bindings for RackHD

## Description
gorackhd represents API bindings for Go that allow you to manage [RackHD](https://github.com/RackHD/RackHD). The intended functions are a direct implementation of what's available within the RackHD API.

## Compatibility
gorackhd was created using [go-swagger](https://github.com/go-swagger/go-swagger) client generator. 

The client relies on a swagger spec file that is taken from RackHD's [on-http](https://github.com/RackHD/on-http) library and packaged at `/swagger-spec/monorail.yml`. As the swagger spec changes, this client will update as needed. 

## Installation
Install the client into your `$GOPATH`
```
go get github.com/emccode/gorackhd
```


## Environment Variables
| Name        | Description           |
| ------------- |:-------------:|
| `GORACKHD_ENDPOINT`      | the API endpoint. `localhost:9090` is used by default if not set             |


## Using the client

This sample script will retrieve (GET) all nodes 

```
package main

import (
    "fmt"
    "log"
    "os"

    //swagclient "github.com/go-swagger/go-swagger/client"
    httptransport "github.com/go-swagger/go-swagger/httpkit/client"
    "github.com/go-swagger/go-swagger/strfmt"

    apiclient "github.com/emccode/gorackhd/client"
    "github.com/emccode/gorackhd/client/nodes"
)

func main() {

    // create the transport
    transport := httptransport.New("localhost:9090", "/api/1.1", []string{"http"})

    // If not using the Vagrant image, set this environment variable to something other than localhost:9090
    if os.Getenv("GORACKHD_ENDPOINT") != "" {
        transport.Host = os.Getenv("GORACKHD_ENDPOINT")
    }

    // create the API client, with the transport
    client := apiclient.New(transport, strfmt.Default)

    //use any function to do REST operations
    resp, err := client.Nodes.GetNodes(nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%#v\n", resp.Payload)

}

```

### Use the client to create (POST) a new item
Import the library for what you are trying to post:
```
import (
    ...
    "github.com/emccode/gorackhd/client/nodes"
    ...
)
```

Create the JSON struct
```
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
```


```
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

    resp, err := client.Nodes.PostNodes(&nodes.PostNodesParams{Identifiers: c})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%#v\n", resp.Payload)
```

### Use the client to fetch (GET) individual items

Import the library for what you are trying to retrieve:
```
import (
    ...
    "github.com/emccode/gorackhd/client/nodes"
    ...
)
```

Lookup the params that are required in a struct:
```
resp, err := client.Skus.GetSkusIdentifier(&nodes.GetNodesIdentifierParams{Identifier: "1234abcd1234abcd1234abcd"})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%#v\n", resp.Payload)

```

### Use the client to remove (DELETE) individual items
Import the library for what you are trying to retrieve:
```
import (
    ...
    "github.com/emccode/gorackhd/client/nodes"
    ...
)
```

Lookup the params that are required in a struct:
```
resp, err := client.Nodes.DeleteNodesIdentifier(&nodes.DeleteNodesIdentifierParams{Identifier: "1234abcd1234abcd1234abcd"})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%#v\n", resp)
```

## Future
- Add documentation to show examples of the main API calls needed for administration. This includes GET/PUT/POST/DELETE for Skus, Workflows, & Catalogs.

## Contribution
Create a fork of the project into your own repository. 

Use the latest version of `go-swagger` to validate the swagger spec.
```
go get github.com/emccode/gorackhd
cd $GOPATH/src/github.com/emccode/gorackhd
make deps
make install
```

Run the tests provided in `gorackhd_test.go` to verify GET/POST/DELETE still function:
```
env DEBUG=1 go test -run TestNodePostOperation -v
env DEBUG=1 go test -run TestNodeGetOperation -v
env DEBUG=1 go test -run TestNodeDeleteOperation -v
```

If tests do not pass, please create an issue so it can be addressed or fix and create a PR.

If all tests pass, make changes and create a pull request with a description on what was added or removed and details explaining the changes in lines of code. If approved, project owners will merge it.

## Licensing
gorackhd is freely distributed under the [MIT License](http://emccode.github.io/sampledocs/LICENSE "LICENSE"). See LICENSE for details.

##Support
Please file bugs and issues on the Github issues page for this project. This is to help keep track and document everything related to this repo. For general discussions and further support you can join the [EMC {code} Community slack team](http://community.emccode.com/) and join the **#rackhd** channel. Lastly, for questions asked on [Stackoverflow.com](https://stackoverflow.com) please tag them with **EMC**. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.






