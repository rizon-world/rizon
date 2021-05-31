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

#### sample

```bash
$ rizond tx tokenswap create a4aa35ffe32f5d2d0fbb5a3b2273695c44576c11e529760f0fbf7ed16d90e99f 3 rizon1rkxs2xeq93qlwhvcp2wp8zpstrp8vver3ap4y5 --from swap_owner --chain-id adora
```

