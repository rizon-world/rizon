package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rest"
)

// RegisterHandlers registers treasury's rest handlers to a router
func RegisterHandlers(clientCtx client.Context, r *mux.Router) {
	dr := rest.WithHTTPDeprecationHeaders(r)
	registerQueryRoutes(clientCtx, dr)
	registerTxHandlers(clientCtx, dr)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/treasury/currencies", getCurrenciesHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/treasury/currencies/{denom}", getCurrencyHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/treasury/parameters", getParamsHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/treasury/mint", mintHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/treasury/burn", burnHandler(clientCtx)).Methods("POST")

}
