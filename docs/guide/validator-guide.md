# Validator Guide

## Join testnet

### Install Rizon Platform

You need to install rizon binary to join our network.

```bash
$ git clone https://github.com/rizon-world/rizon.git && cd rizon
$ make install
```

And check the binary has installed.

```bash
$ rizond version
```

Please refer [Install Rizon Platform](../getting-started/install-rizon-platform.md) page for more details about prerequisites or supported OS, something like that.

### Setup initial node

After you finish installing the Rizon binary, you can setup your initial node.

If you had setup node, you'd better clean existing configurations.

```bash
# remove old configurations
$ rm -rf ~/.rizon/config

# reset old blockchain data
$ rizond unsafe-reset-all
```

And init your node.

```bash
$ rizond init <moniker> --chain-id <chain-id>
```

* `moniker` : the displayed name of your node over Rizon network.
* `chain-id` : the chain-id of the network that you want to connect. For testnet, the chain-id is noticed at [https://github.com/rizon-world/testnet](https://github.com/rizon-world/testnet).

The genesis file is stored in `~/.rizon/config/genesis.json`. You will update this file next.

### Prepare genesis.json

The `genesis.json` file is which defines the initial states of your blockchain, acting as a genesis block. The state which defined in the `genesis.json` contains all the necessary information, like initial token allocation, genesis time, default parameters, and more.

You have to download the genesis.json file from [https://raw.githubusercontent.com/rizon-world/testnet/master/genesis.json](https://raw.githubusercontent.com/rizon-world/testnet/master/genesis.json) and replace `~/.rizon/config/genesis.json` to join testnet.

```bash
$ wget https://raw.githubusercontent.com/rizon-world/testnet/master/genesis.json
$ cp -f genesis.json ~/.rizon/config/genesis.json
```

### Configure the node

For configuring the node to connect testnet, we will modify two main configuration files of Rizon platform, `app.toml` and `config.toml`.

First edit the `~/.rizon/config/app.toml` file to prevent spamming. It rejects incoming transactions with less than the minimum gas prices.

```bash
$ sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.0001uatolo"/g' ~/.rizon/config/app.toml
```

Then you should add at least one or more seed nodes to `~/.rizon/config/config.toml` to tell your node how find other peers. So open `~/.rizon/config/config.toml` file and edit `seeds` of `[p2p]` section.

```text
#######################################################
###           P2P Configuration Options             ###
#######################################################
[p2p]

...

# Comma separated list of seed nodes to connect to
seeds = "bad95d276696706c4df2272d45395d7da4a13f43@13.124.50.97:26656"
```

You can get proper list of seed nodes on [https://github.com/rizon-world/testnet](https://github.com/rizon-world/testnet).

#### Optional configurations

By default, the REST API endpoint is disabled. If you want to enable the endpoint, you need to open `~/.rizon/config/app.toml` file and edit `enable` to `true` of the `[api]` section. You also can configure _laddr_ by edting `address`.

```text
###############################################################################
###                           API Configuration                             ###
###############################################################################

[api]

# Enable defines if the API server should be enabled.
enable = true

# Address defines the API server to listen on.
address = "tcp://0.0.0.0:1317"
```

### Connect to the network

Now we finish preparing to join the testnet. Start node with command below.

```bash
$ rizond start
```

If there is no failure, node will launch and start to sync blockchain states from genesis.

```bash
4:06PM INF starting ABCI with Tendermint
4:06PM INF Starting multiAppConn service impl=multiAppConn module=proxy
...
4:06PM INF minted coins from module account amount=45519uatolo from=mint module=x/bank
4:06PM INF executed block height=1 module=state num_invalid_txs=0 num_valid_txs=0
4:06PM INF commit synced commit=436F6D6D697449447B5B3131332032343420333920343420313635203235203935203133352034302031383720393420393720373220383120323434203237203137382031383020313938203137302031373620333920333520393220323234203139382031393220313220313230203836203732203234335D3A317D
4:06PM INF committed state app_hash=71F4272CA5195F8728BB5E614851F41BB2B4C6AAB027235CE0C6C00C785648F3 height=1 module=state num_txs=0
4:06PM INF indexed block height=1 module=txindex
```

You need to wait until blocks are fully synced. You can check the sync is finished by command below.

```bash
$ rizond status 2>&1 | awk -F'catching_up":' '{print $2}' | cut -c -5
```

If command returns `false`, sync has finished. Otherwise, you need to wait more.

### Create an account

Before creating validator, you should have an account for your node and validator.

#### Keyring

Rizon platform uses Cosmos keyring implementation. The keyring holds the private/public keypairs used to interact with a node. For instance, a validator key needs to be set up before running the blockchain node, so that blocks can be correctly signed. The private key can be stored in different locations, called "backends", such as a file or the operating system's own key storage.

If you want to know more information of keyring, please see [https://docs.cosmos.network/v0.42/run-node/keyring.html](https://docs.cosmos.network/v0.42/run-node/keyring.html).

#### Create new wallet

```bash
$ rizond keys add <wallet_name>
```

This command returns wallet address and mnemonic. You must remember the results.

#### Recover existing account from mnemonic or private key.

```bash
# recover from mnemonic
$ rizond keys add <wallet_name> --recover

or

# import from private_key
$ rizond keys import <wallet_name> <private_key_file>
```

#### Note

The command may require system password. If you don't want to use any password, you can use `--keyring-backend test` flag, but we do not recommend this for use in production environments.

### Get some coins

You can get some test coins via _faucet_ or requesting to administrator.

### Create validator

After the node is fully synced, you can create validator via send `create-validator` tx to join testnet.

See the sample command below.

```bash
$ rizond tx staking create-validator \
--amount="10000000000uatolo" \
--pubkey=$(rizond tendermint show-validator) \
--moniker="my_node" \
--commission-rate="0.10" \
--commission-max-rate="0.20" \
--commission-max-change-rate="0.01" \
--min-self-delegation="1" \
--from after \
--chain-id=adora \
--fees="1000uatolo"
```

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
| fees | Fees to pay along with transaction |

Once the `create-validator` transaction has completed, you can check whether your validator has been properly added via command below.

```bash
$ rizond q tendermint-validator-set | grep `rizond tendermint show-address`
```

You can modify your validator information whenever you want. Please refer this help.

```bash
$ rizond tx staking edit-validator --help
```

