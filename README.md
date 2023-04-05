# graphql-transport-ws

This Go package is a mix of two packages:

  * [github.com/graph-gophers/graphql-transport-ws](https://github.com/graph-gophers/graphql-transport-ws) v0.0.2

    Handler and GraphQLService API

  * [github.com/99designs/gqlgen](https://github.com/99designs/gqlgen) v0.17.28

    Implementation of WebSocket transport for GraphQL that supports two protocol versions:
        
    * [Apollo protocol](https://github.com/apollographql/subscriptions-transport-ws/blob/master/PROTOCOL.md)
    * [graphql-ws protocol](https://github.com/enisdenjo/graphql-ws/blob/master/PROTOCOL.md)

### Use with graph-gophers/graphql-go

To use this library with [github.com/graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go) you can wrap
the `relay` handler it provides the following way:

```go
package main

import (
	"fmt"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
)

const schema = `
	schema {
		subscription: Subscription
	}

	type Subscription {
		...
	}
`

type resolver struct {
	// ...
}

func main() {
	// init graphQL schema
	s, err := graphql.ParseSchema(schema, &resolver{})
	if err != nil {
		panic(err)
	}

	// graphQL handler
	graphQLHandler := graphqlws.NewHandlerFunc(s, &relay.Handler{Schema: s})
	http.HandleFunc("/graphql", graphQLHandler)

	// start HTTP server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil); err != nil {
		panic(err)
	}
}
```

For a more in depth example see [this repo](https://github.com/matiasanaya/go-graphql-subscription-example).

### Client

Check [graphql-ws Get Started guide](https://the-guild.dev/graphql/ws/get-started#use-the-client) for details
on how to use WebSockets on the client side.
