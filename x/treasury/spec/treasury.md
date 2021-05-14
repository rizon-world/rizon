# treasury

## Abstract

The treasury manages all stable currencies which are supported by Rizon platform.

## State

### Currency

Currency defines a single currency info

```go
type Currency struct {
        // denom is the name of the currency
        Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
        // desc is a description of the currency
        Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc"`
        // owner is who can mint this currency
        Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
        // mintable indicates whether this currency can be minted or not
        Mintable bool `protobuf:"varint,4,opt,name=mintable,proto3" json:"mintable,omitempty" yaml:"mintable"`
}
```

### Currencies

Currencies is the list of all supported currencies' name

```go
type Currencies struct {
        // denoms is the denom list of all currencies
        Denoms []string `protobuf:"bytes,1,rep,name=denoms,proto3" json:"denoms,omitempty" yaml:"denoms"`
}
```

### Sequence

Sequence describes currency sequence number of current state

```go
type Sequence struct {
        // number is the currency sequence number of current state
        Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty" yaml:"number"`
}
```

### Params

Params defines the parameters for the treasury module

```go
type Params struct {
        // mintable indicates whether every currencie of treasury module are able to mint or not
        Mintable bool `protobuf:"varint,1,opt,name=mintable,proto3" json:"mintable,omitempty" yaml:"mintable"`
        // sequence of currency state
        Sequence int64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty" yaml:"sequence"`
        // currency_list is the list of supported currencies
        CurrencyList []Currency `protobuf:"bytes,3,rep,name=currency_list,json=currencyList,proto3" json:"currency_list" yaml:"currency_list"`
}
```

## Message

### MsgMintRequest

MsgMintRequest defines a SDK message for minting currency

```go
type MsgMintRequest struct {
        // receiver is the target address of minted coins
        Receiver string `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty" yaml:"receiver"`
        // signer is who requests minting
        Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
        // amount is the amount to mint
        Amount types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}
```

### MsgBurnRequest

MsgBurnRequest defines a SDK message for burning currency

```go
type MsgBurnRequest struct {
        // signer is who requests burning
        Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
        // amount is the amount to burn
        Amount types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}
```

## Event

### EndBlocker

TBD

### msg\_server

#### Mint

| Type | AttributeKey | AttributeValue |
| :--- | :--- | :--- |
| mint | denom | {msg.Amount.Denom} |
| mint | receiver | {msg.Receiver} |
| mint | amount | {msg.Amount.Amount} |
| message | module | treasury |
| message | sender | {msg.Signer} |

#### Burn

| Type | AttributeKey | AttributeValue |
| :--- | :--- | :--- |
| burn | denom | {msg.Amount.Denom} |
| burn | amount | {msg.Amount.Amount} |
| message | module | treasury |
| message | sender | {msg.Signer} |

## Parameter

| Key | Type | Example |
| :--- | :--- | :--- |
| Mintable | bool | true |
| Sequence | int64 | 1 |
| CurrencyList | \[\]Currency | \[{"denom":"skrw","desc":"stable coin of KRW","owner":"rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7","mintable":true}\] |



