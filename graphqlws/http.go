package graphqlws

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/graph-gophers/graphql-transport-ws/transport"
)

type GraphQLService = transport.GraphQLService

var defaultUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var defaultTransport = transport.Websocket{
	Upgrader:              defaultUpgrader,
	InitTimeout:           5 * time.Second,
	KeepAlivePingInterval: 10 * time.Second,
}

// Option applies configuration when a graphql websocket connection is handled
type Option func(*handlerConfig)

func WithWebsocketTransport(transport *transport.Websocket) Option {
	return func(cfg *handlerConfig) {
		cfg.Transport = transport
	}
}

// NewHandlerFunc returns an http.HandlerFunc that supports GraphQL over websockets
func NewHandlerFunc(svc GraphQLService, httpHandler http.Handler, opts ...Option) http.HandlerFunc {
	cfg := handlerConfig{
		Transport: &defaultTransport,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if cfg.Transport.Supports(r) {
			cfg.Transport.Do(w, r, svc)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}
}

type handlerConfig struct {
	Transport *transport.Websocket
}
