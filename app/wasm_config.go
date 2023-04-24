package rizon

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultRizonInstanceCost is initially set the same as in wasmd
	DefaultRizonInstanceCost uint64 = 60_000
	// DefaultRizonCompileCost set to a large number for testing
	DefaultRizonCompileCost uint64 = 3
)

// RizonGasRegisterConfig is defaults plus a custom compile amount
func RizonGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultRizonInstanceCost
	gasConfig.CompileCost = DefaultRizonCompileCost

	return gasConfig
}

func NewRizonWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(RizonGasRegisterConfig())
}
