vpsAdmin Client in Go
---------------------
The client is generated using `haveapi-go-client` from
[HaveAPI](https://github.com/vpsfreecz/haveapi).

## Usage
```go
package main

import (
	"fmt"
	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

func main() {
	api := client.New("https://api.vpsfree.cz")

	response, err := api.Cluster.PublicStats.Call()

	if err != nil {
		fmt.Println(err)
		return
	} else if !response.Status {
		fmt.Printf("Action failed: %s\n", response.Message)
	}

	// Response envelope
	fmt.Printf("%+v\n", resp)

	// Actual action output, namespaced
	fmt.Printf("%+v\n", resp.Response)
	fmt.Printf("%+v\n", resp.Response.Cluster)

	// Action output, directly accessed
	fmt.Printf("%+v\n", resp.Output)
}
```

### Authentication
Both HTTP basic and token authentication is supported.

```go
// HTTP basic
api.SetBasicAuthentication("user", "password")

// New token
err := api.SetNewTokenAuth("user", "password")

if err != nil {
	fmt.Printf("Token is: %s\n", api.Authentication.(client.TokenAuth).Token)
}

// Existing token
api.SetExistingTokenAuth("mytoken")
```

### Path parameters
Path parameters are used to identify resources, e.g. by their ID.

```go
action := api.Vps.Show.Prepare()
action.SetPathParamInt("vps_id", 123)
response, err := action.Call()
```

### Input parameters
The client can be configured either to send just the configured parameters,
or to send all parameters, even unset ones.

```go
action := api.Vps.Update.Prepare()
action.SetPathParamInt("vps_id", 123)

input := action.NewInput()
input.SetHostname("new-hostname")

// Although Vps.Update has many more input parameters, only Hostname was set
// and so no other parameter will be sent. This is what we want in order to
// change just the hostname.
response, err := action.Call()
```

To send all input parameters, do not use the setter functions, but access
the input struct's fields directly. Or create the struct manually
(e.g. `ActionVpsUpdateInput{}`) and then use `action.SetInput()`.

## Updating the client
```
$ haveapi-go-client https://api.vpsfree.cz <path-to-this-repository>/client
```
