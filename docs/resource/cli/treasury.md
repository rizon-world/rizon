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

#### sample

```bash
$ rizond tx treasury burn 100susd --from susd_owner --chain-id adora
```

### mint

Create a tx of mint coin request

```bash
$ rizond tx treasury mint [amount] [receiver-address] --from [address] [flags]
```

#### sample

```bash
$ rizond tx treasury mint 100skrw rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7 --from skrw_owner --chain-id adora
```

