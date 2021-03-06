# tokenswap

## Abstract

This module is for coinswap from HDAC chain to RIZON chain.

## State

See proto files on [https://github.com/rizon-world/rizon/tree/master/proto/tokenswap](https://github.com/rizon-world/rizon/tree/master/proto/tokenswap)

### Tokenswap

Tokenswap defines the tokenswap state

```go
type Tokenswap struct {
        // tx_hash is the tx hash of burn tx from legacy chain
        // tx_hash is used for store key
        TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
        // receiver is the target of tokenswap
        Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
        // signer is who confirms the swap process
        Signer string `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
        // amount is the amount of swap process
        Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}
```

### Params

Params defines the parameters for the tokenswap module

```go
type Params struct {
        // swappable indicates whether tokenswap module is enabled or not
        Swappable bool `protobuf:"varint,1,opt,name=swappable,proto3" json:"swappable,omitempty" yaml:"swappable"`
        // signer is someone who can request tokenswap
        Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
}
```

## Message

### MsgCreateTokenswapRequest

MsgCreateTokenswapRequest defines a SDK message for creating a new tokenswap

#### structure

```go
type MsgCreateTokenswapRequest struct {
        // tx_hash is the tx hash of burn tx from legacy chain
        // tx_hash is used for store key
        TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" yaml:"tx_hash"`
        // receiver is the target of tokenswap
        Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty" yaml:"receiver"`
        // signer is who confirms the swap process
        Signer string `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
        // amount is the amount of swap process
        Amount *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"amount,omitempty"`
}
```

#### request sample

```javascript
{
    "body": {
        "messages": [
            {
                "@type": "/rizonworld.rizon.tokenswap.MsgCreateTokenswapRequest",
                "amount": "3.000000000000000000",
                "receiver": "rizon1rkxs2xeq93qlwhvcp2wp8zpstrp8vver3ap4y5",
                "signer": "rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7",
                "tx_hash": "a4aa35ffe32f5d2d0fbb5a3b2273695c44576c11e529760f0fbf7ed16d90e99f"
            }
        ],
    },
}
```

## Event

### msg\_server

#### CreateTokenswap

| Type | AttributeKey | AttributeValue |
| :--- | :--- | :--- |
| create\_tokenswap | tx\_hash | {msg.TxHash} |
| create\_tokenswap | receiver | {msg.Receiver} |
| create\_tokenswap | amount | {msg.Amount} |
| message | module | tokenswap |
| message | sender | {msg.Signer} |

## Parameter

| Key | Type | Example |
| :--- | :--- | :--- |
| Swappable | bool | true |
| Signer | string | "rizon1cafygu0kppg46tq9kkzz9z9nrf5v8zwwnf5t9l" |

