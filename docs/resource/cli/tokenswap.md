# tokenswap

## Query

### get

Query swap request by tx hash

```bash
$ rizond query tokenswap get [tx-hash] [flags]
```

### params

Query the current parameters of tokenswap

```bash
$ rizond query tokenswap params [flags]
```

## Transaction

### create

Create a token swap request

```bash
$ rizond tx tokenswap create [tx-hash] [amount] [receiver-address] --from [address] [flags]
```

Note that the only permitted signer \(from address\) can request swap.

