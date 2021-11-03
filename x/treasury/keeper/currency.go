package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/treasury/types"
)

// CurrencyMintable
func (k Keeper) CurrencyMintable(ctx sdk.Context, denom string) bool {
	currency := k.GetCurrency(ctx, denom)

	return currency.Mintable
}

// Owner
func (k Keeper) Owner(ctx sdk.Context, denom string) string {
	currency := k.GetCurrency(ctx, denom)

	return currency.Owner
}

// SetCurrencies stores all supported currency denom list
func (k Keeper) SetCurrencies(ctx sdk.Context, currencies types.Currencies) {
	store := k.Store(ctx)

	store.Set(types.CurrenciesKey, k.cdc.MustMarshalBinaryBare(&currencies))
}

// GetCurrencies returns all supported currency denom list
func (k Keeper) GetCurrencies(ctx sdk.Context) types.Currencies {
	store := k.Store(ctx)

	bz := store.Get(types.CurrenciesKey)

	var currencies types.Currencies
	k.cdc.MustUnmarshalBinaryBare(bz, &currencies)

	return currencies
}

// ClearCurrencies clears all currencies
func (k Keeper) ClearCurrencies(ctx sdk.Context) {
	store := k.Store(ctx)
	iter := sdk.KVStorePrefixIterator(store, types.CurrencyPrefix)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// SetCurrency stores an item of a currency
func (k Keeper) SetCurrency(ctx sdk.Context, currency types.Currency) {
	store := k.Store(ctx)
	currencyStore := prefix.NewStore(store, types.CurrencyPrefix)
	key := types.GetCurrencyKey(currency.Denom)

	currencyStore.Set(key, k.cdc.MustMarshalBinaryBare(&currency))
}

// GetCurrency returns an information of a currency
func (k Keeper) GetCurrency(ctx sdk.Context, denom string) types.Currency {
	store := k.Store(ctx)
	currencyStore := prefix.NewStore(store, types.CurrencyPrefix)

	var currency types.Currency
	key := types.GetCurrencyKey(denom)
	bz := currencyStore.Get(key)
	k.cdc.MustUnmarshalBinaryBare(bz, &currency)

	return currency
}

// SetSequence stores a sequence number of current state
func (k Keeper) SetSequence(ctx sdk.Context, seq types.Sequence) {
	store := k.Store(ctx)

	store.Set(types.SequenceKey, k.cdc.MustMarshalBinaryBare(&seq))
}

// GetSequence returns a sequence number of current state
func (k Keeper) GetSequence(ctx sdk.Context) int64 {
	store := k.Store(ctx)

	var seq types.Sequence
	bz := store.Get(types.SequenceKey)
	k.cdc.MustUnmarshalBinaryBare(bz, &seq)

	return seq.Number
}
