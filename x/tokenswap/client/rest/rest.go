package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rest"
)

// RegisterHandlers registers tokenswap's rest handlers to a router
func RegisterHandlers(clientCtx client.Context, r *mux.Router) {
	dr := rest.WithHTTPDeprecationHeaders(r)
	registerQueryRoutes(clientCtx, dr)
	registerTxHandlers(clientCtx, dr)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/tokenswap/tokenswaps/{tx_hash}", getTokenswapHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/tokenswap/parameters", getParamsHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/tokenswap/create", createTokenswapHandler(clientCtx)).Methods("POST")

}
