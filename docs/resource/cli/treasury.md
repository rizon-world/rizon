# treasury

## Query

### currencies

Query all supported currency's denoms

```bash
$ rizond query treasury currencies [flags]
```

### currency

Query an information of single currency

```bash
$ rizond query treasury currency [denom] [flags]
```

### params

Query the current parameters of treasury

```bash
$ rizond query treasury params [flags]
```

## Transaction

### burn

Create a tx of burn coin request

```bash
$ rizond tx treasury burn [amount] --from [address] [flags]
```

### mint

Create a tx of mint coin request

```bash
$ rizond tx treasury mint [amount] [receiver-address] --from [address] [flags]
```

