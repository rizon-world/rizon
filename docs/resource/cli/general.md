# general

## Keys

These keys may be in any format supported by the Tendermint crypto library and can be used by light-clients, full nodes, or any other application that needs to sign with a private key.

The keyring supports the following backends:

| backend | desc |
| :--- | :--- |
| os | Uses the operating system's default credentials store. |
| file | Uses encrypted file-based keystore within the app's configuration directory. This keyring will request a password each time it is accessed, which may occur multiple times in a single command resulting in repeated password prompts. |
| test | Stores keys insecurely to disk. It does not prompt for a password to be unlocked and it should be use only for testing purposes. |

### Add keys

#### Create new wallet

```bash
$ rizond keys add [wallet_name] [flags]
```

This command returns wallet address and mnemonic. You must remember the results.

#### Recover existing account from mnemonic or private key.

```bash
# recover from mnemonic
$ rizond keys add [wallet_name] --recover

or

# import from private_key
$ rizond keys import [wallet_name] [private_key_file]
```

### Displaying Keys

#### show

Display keys details. If multiple names or addresses are provided, then an ephemeral multisig key will be created under the name "multi" consisting of all the keys provided by name and multisig threshold.

```bash
$  rizond keys show [name_or_address [name_or_address...]] [flags]
```

#### list

Return a list of all public keys stored by this key manager along with their associated name and address.

```bash
$ rizond keys list [flags]
```

## Query

You can query all relevant information from the blockchain using `rizond` binary.

If you want to know about the list of query commands, just type

```bash
$ rizond query
```

### Commands

These are frequently used commands. For each command, you can use `--help` flag to get more information about.

#### account

Query for account by address

```text
$ rizond query account [address] [flags]
```

#### balance

Query the total balance of an account or of a specific denomination.

```text
$ rizond query bank balances [address] --denom=[denom]
```

#### block

Get verified data for a the block at given height

```text
$ rizond query block [height] [flags]
```

#### delegations

Query delegations for an individual delegator

```bash
$ rizond query staking delegations [delegator-addr] [flags]
```

#### validator

Query details about an individual validator.

```text
$ rizond query staking validator [validator-addr] [flags]
```

#### rewards

Query all rewards earned by a delegator, optionally restrict to rewards from a single validator

```text
$ rizond query distribution rewards [delegator-addr] [validator-addr] [flags]
```

#### tx

Query for a transaction by hash in a committed block

```text
$ rizond query tx [hash] [flags]
```

## Transaction

Other than querying blockchain data, you also can interact with the blockchain, sending transactions containing messages.

```text
$ rizond tx
```

Please check each module's tx commands to learn how generate and broadcast txs with which message parameters.

### Send

Most frequently used command of tx is `send`. Through this `send` command, see how to use tx commands.

#### send

Send funds from one account to another.

```bash
$ rizond tx bank send [from_key_or_address] [to_address] [amount] [flags]
```

The tx command processes through `generate` - `sign` - `broadcast` steps. So the command above works internally generate unsigned transaction first, and ask password of `from_key_or_address` to sign,  sign the generated transaction with key, and finally broadcast the signed transaction.

#### simulating a transaction

You can simulate a transaction without actually broadcasting it by appending the `--dry-run` flag to the command line.

```text
$ rizond tx bank send [from_key_or_address] [to_address] [amount] [flags] --dry-run
```

This returns gas estimation if success or error result when fail.

#### generating a transaction <a id="generating-a-transaction-without-sending"></a>

You can build a transaction with `--generate-only` flag. This allows you to separate the creation and signing of a transaction without broadcasting.

```text
$ rizond tx bank send [from_address] [to_address] [amount] [flags] --generate-only
```

Note that `from_key` is not supported in generate-only mode. You need to use `from_address`.

This command returns unsigned transaction as standard output. You'd better redirect the stdout to file to sign it easily.

#### sign

Sign a transaction created with the `--generate-only` flag. It will read a transaction from `[file]`, sign it, and print its JSON encoding.

If the `--signature-only` flag is set, it will output the signature parts only.

The `--offline` flag makes sure that the client will not reach out to full node. As a result, the account and sequence number queries will not be performed and it is required to set such parameters manually. Note, invalid values will cause the transaction to fail.

```text
$ rizond tx sign [file] [flags]
```

This command also returns signed result as standard output. So you'd better redirect the stdout to file to broadcast, too.

#### validate-signatures

Before broadcast signed transaction, you can validate the transaction's signatures.

This prints the addresses that must sign the transaction, those who have already signed it, and make sure that signatures are in the correct order.

The command would check whether all required signers have signed the transactions, whether the signatures were collected in the right order, and if the signature is valid over the given transaction. If the `--offline` flag is also set, signature validation over the transaction will be not be performed as that will require RPC communication with a full node.

```text
$ rizond tx validate-signatures [file] [flags]
```

#### broadcast

You can broadcast transactions created with the `--generate-only` flag and signed with the `sign` command. This command reads the transaction from `[file]` and broadcast it to a node.

```text
$ rizond tx broadcast [file] [flags]
```

### Create validator

Create new validator initialized with a self-delegation to it

```bash
$ rizond tx staking create-validator [flags]
```

`create-validator` command needs some mandatory flags.

| flag | desc |
| :--- | :--- |
| amount | Amount of coins to bond |
| pubkey | The Bech32 encoded PubKey of the validator |
| moniker | The validator's name |
| commission-rate | The initial commission rate percentage |
| commission-max-rate | The maximum commission rate percentage |
| commission-max-change-rate | The maximum commission change rate percentage \(per day\) |
| min-self-delegation | The minimum self delegation required on the validator |
| from | Name or address of private key with which to sign |
| chain-id | The network chain ID to connect |
| fees\(or gas-prices\) | Fees \(or gas-prices\) to pay along with transaction |

## Gas and Fees

In the Cosmos based platforms, `gas` is a special unit that is used to track the consumption of resources during execution. `gas` is typically consumed whenever read and writes are made to the store, but it can also be consumed if expensive computation needs to be done. It serves two main purposes:

* Make sure blocks are not consuming too many resources and will be finalized. This is implemented by default in the SDK via the [block gas meter](https://docs.cosmos.network/v0.42/basics/gas-fees.html#block-gas-meter).
* Prevent spam and abuse from end-user. To this end, `gas` consumed during [`message`](https://docs.cosmos.network/v0.42/building-modules/messages-and-queries.html#messages) execution is typically priced, resulting in a `fee` \(`fees = gas * gas-prices`\). `fees` generally have to be paid by the sender of the `message`. Note that the Cosmos platform does not enforce `gas` pricing by default, as there may be other ways to prevent spam \(e.g. bandwidth schemes\). Still, most applications will implement `fee` mechanisms to prevent spam. This is done via the [`AnteHandler`](https://docs.cosmos.network/v0.42/basics/gas-fees.html#antehandler).

Transactions on the network need to include a transaction fee in order to be processed. This fee pays for the gas required to run the transaction. There are three flags to set transaction fee.

The `gas` is dependent on the transaction. Different transaction require different amount of `gas`. 

The `gas-prices` is the price of each unit of `gas`. Each validator sets a `minimum-gas-prices` value, and will only include transactions that have a `gas-prices` greater than their `minimum-gas-prices`.

The `fees` are the product of `gas` and `gas-prices`. 

### Set Fees

As a user, you must input `gas-prices` or `fees`. The validators specify a `minimum-gas-prices` at `app.toml` file that they use to determine whether to include a transaction, where `gas-prices >= min-gas-prices`. The higher `gas-prices` or `fees` , the higher chance that your transaction will get included in a block. 

For setting fees, you first can use `--gas-prices` flag.

```text
$ rizond tx send ... --gas-prices=0.0001uatolo
```

You don't need to use `--gas` flag every time, the `gas` amount for a transaction is calculated as it is being processed automatically as default. Of course, this only gives an estimate, so it returns failures. So you can adjust this estimate with the flag `--gas-adjustment` if you want to be sure you provide enough `gas` for the transaction.

You also use `--fees` sure.

```text
$ rizond tx send ... --fees=10000uatolo
```

If you use fees, validators will calculate the implied `min-gas-prices` by dividing your fee with the estimated `gas` consumption, to properly assign the right priority to your transaction.

