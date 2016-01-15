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
| `GOPATH`                 | it's required to set your [$GOPATH](https://golang.org/doc/code.html#GOPATH)    |
| `GORACKHD_ENDPOINT`      | the API endpoint, `localhost:9090` or `192.168.90.100`               |


## Using the client
Below is a sample script used to retrieve all Skus. 
```
package main

import (
    "fmt"
    apiclient "github.com/emccode/gorackhd/client"
    httptransport "github.com/go-swagger/go-swagger/httpkit/client"
    "github.com/go-swagger/go-swagger/spec"
    "github.com/go-swagger/go-swagger/strfmt"
    "log"
    "os"
)

func main() {

    // load the swagger spec from local $GOPATH and binding
    goPath := os.Getenv("GOPATH")
    doc, err := spec.Load(goPath + "/src/github.com/emccode/gorackhd/swagger-spec/monorail.yml")
    if err != nil {
        log.Fatal(err)
    }

    //Retrieve everything over http vs https. Https is defaulted
    spec := doc.Spec()
    spec.Schemes = []string{"http"}

    // create the transport
    transport := httptransport.New(doc)

    // configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
    if os.Getenv("GORACKHD_ENDPOINT") != "" {
        transport.Host = os.Getenv("GORACKHD_ENDPOINT")
    }

    // create the API client, with the transport
    client := apiclient.New(transport, strfmt.Default)

    //use any function to do REST operations
    resp, err := client.Skus.GetSkus(nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%#v\n", resp.Payload)

}

```

### Using the client to get fetch individual items

Import the library for what you are trying to retrieve:
```
import (
    ...
    "github.com/emccode/gorackhd/client/skus"
    ...
)
```

Lookup the params that are required in a struct:
```
resp, err := client.Skus.GetSkusIdentifier(&skus.GetSkusIdentifierParams{Identifier: "568e8b76c3354ff04bab27e0"})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%#v\n", resp.Payload)

```


## Future
- Add documentation to show examples of the main API calls needed for administration. This includes GET/PUT/POST/DELETE for Skus, Workflows, & Catalogs.

## Contribution
Create a fork of the project into your own reposity. Make all necessary changes and create a pull request with a description on what was added or removed and details explaining the changes in lines of code. If approved, project owners will merge it.

## Licensing
gorackhd is freely distributed under the [MIT License](http://emccode.github.io/sampledocs/LICENSE "LICENSE"). See LICENSE for details.

##Support
Please file bugs and issues on the Github issues page for this project. This is to help keep track and document everything related to this repo. For general discussions and further support you can join the [EMC {code} Community slack team](http://community.emccode.com/) and join the **#rackhd** channel. Lastly, for questions asked on [Stackoverflow.com](https://stackoverflow.com) please tag them with **EMC**. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.






