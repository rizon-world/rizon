# Deploy own network

{% hint style="info" %}
Before setting up your node, you should [install Rizon Platform](install-rizon-platform.md) first.
{% endhint %}

## Setup your node

Follow the steps below to create genesis setting and initial validator.

### init node

```bash
$ rizond init my_node --chain-id my_testnet
{"app_message":{"auth":{"accounts":[],"params":{"max_memo_characters":"256","sig_verify_cost_ed25519":"590","sig_verify_cost_secp256k1":"1000","tx_sig_limit":"7","tx_size_cost_per_byte":"10"}},"bank":{"balances":[],"denom_metadata":[],"params":{"default_send_enabled":true,"send_enabled":[]},"supply":[]},"capability":{"index":"1","owners":[]},"crisis":{"constant_fee":{"amount":"1000","denom":"stake"}},"distribution":{"delegator_starting_infos":[],"delegator_withdraw_infos":[],"fee_pool":{"community_pool":[]},"outstanding_rewards":[],"params":{"base_proposer_reward":"0.010000000000000000","bonus_proposer_reward":"0.040000000000000000","community_tax":"0.020000000000000000","withdraw_addr_enabled":true},"previous_proposer":"","validator_accumulated_commissions":[],"validator_current_rewards":[],"validator_historical_rewards":[],"validator_slash_events":[]},"evidence":{"evidence":[]},"genutil":{"gen_txs":[]},"gov":{"deposit_params":{"max_deposit_period":"172800s","min_deposit":[{"amount":"10000000","denom":"stake"}]},"deposits":[],"proposals":[],"starting_proposal_id":"1","tally_params":{"quorum":"0.334000000000000000","threshold":"0.500000000000000000","veto_threshold":"0.334000000000000000"},"votes":[],"voting_params":{"voting_period":"172800s"}},"ibc":{"channel_genesis":{"ack_sequences":[],"acknowledgements":[],"channels":[],"commitments":[],"next_channel_sequence":"0","receipts":[],"recv_sequences":[],"send_sequences":[]},"client_genesis":{"clients":[],"clients_consensus":[],"clients_metadata":[],"create_localhost":false,"next_client_sequence":"0","params":{"allowed_clients":["06-solomachine","07-tendermint"]}},"connection_genesis":{"client_connection_paths":[],"connections":[],"next_connection_sequence":"0"}},"mint":{"minter":{"annual_provisions":"0.000000000000000000","inflation":"0.130000000000000000"},"params":{"blocks_per_year":"6311520","goal_bonded":"0.670000000000000000","inflation_max":"0.200000000000000000","inflation_min":"0.070000000000000000","inflation_rate_change":"0.130000000000000000","mint_denom":"stake"}},"params":null,"slashing":{"missed_blocks":[],"params":{"downtime_jail_duration":"600s","min_signed_per_window":"0.500000000000000000","signed_blocks_window":"100","slash_fraction_double_sign":"0.050000000000000000","slash_fraction_downtime":"0.010000000000000000"},"signing_infos":[]},"staking":{"delegations":[],"exported":false,"last_total_power":"0","last_validator_powers":[],"params":{"bond_denom":"stake","historical_entries":10000,"max_entries":7,"max_validators":100,"unbonding_time":"1814400s"},"redelegations":[],"unbonding_delegations":[],"validators":[]},"tokenswap":{"params":{"signer":"rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7","swappable":true}},"transfer":{"denom_traces":[],"params":{"receive_enabled":true,"send_enabled":true},"port_id":"transfer"},"treasury":{"params":{"currency_list":[{"denom":"skrw","desc":"stable coin of KRW","mintable":true,"owner":"rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7"},{"denom":"susd","desc":"stable coin of USD","mintable":true,"owner":"rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7"}],"mintable":true,"sequence":"1"},"seq":{"number":"0"}},"upgrade":{},"vesting":{}},"chain_id":"my_testnet","gentxs_dir":"","moniker":"my_node","node_id":"091067d86e57c2a4e87cf4562845751292528dc0"}
```

### create a local wallet key

```bash
$ rizond keys add my_key --keyring-backend test

- name: my_key
  type: local
  address: rizon1rtusr0r6pa9jp02k9ejzprn40qnssgjypzgg2c
  pubkey: rizonpub1addwnpepqv5wlez3aq35xrp0c8tayaf5e77wry80ncv4wfn038qqyl89et93z2wa6hr
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

jacket tissue ritual orient crater fame lyrics cost palace cook room oblige digital body napkin parent wine keen drink baby crystal artist edge grocery
```

### add genesis account

```bash
$ rizond add-genesis-account $(rizond keys show my_key -a --keyring-backend test) 10000000000stake
```

### create genesis tx

```bash
$ rizond gentx my_key 1000000000stake --chain-id my_testnet --keyring-backend test
Genesis transaction written to "/Users/yoosah/.rizon/config/gentx/gentx-091067d86e57c2a4e87cf4562845751292528dc0.json"
```

### collect genesis tx

```bash
$ rizond collect-gentxs
{"app_message":{"auth":{"accounts":[{"@type":"/cosmos.auth.v1beta1.BaseAccount","account_number":"0","address":"rizon1rtusr0r6pa9jp02k9ejzprn40qnssgjypzgg2c","pub_key":null,"sequence":"0"}],"params":{"max_memo_characters":"256","sig_verify_cost_ed25519":"590","sig_verify_cost_secp256k1":"1000","tx_sig_limit":"7","tx_size_cost_per_byte":"10"}},"bank":{"balances":[{"address":"rizon1rtusr0r6pa9jp02k9ejzprn40qnssgjypzgg2c","coins":[{"amount":"10000000000","denom":"stake"}]}],"denom_metadata":[],"params":{"default_send_enabled":true,"send_enabled":[]},"supply":[{"amount":"10000000000","denom":"stake"}]},"capability":{"index":"1","owners":[]},"crisis":{"constant_fee":{"amount":"1000","denom":"stake"}},"distribution":{"delegator_starting_infos":[],"delegator_withdraw_infos":[],"fee_pool":{"community_pool":[]},"outstanding_rewards":[],"params":{"base_proposer_reward":"0.010000000000000000","bonus_proposer_reward":"0.040000000000000000","community_tax":"0.020000000000000000","withdraw_addr_enabled":true},"previous_proposer":"","validator_accumulated_commissions":[],"validator_current_rewards":[],"validator_historical_rewards":[],"validator_slash_events":[]},"evidence":{"evidence":[]},"genutil":{"gen_txs":[{"auth_info":{"fee":{"amount":[],"gas_limit":"200000","granter":"","payer":""},"signer_infos":[{"mode_info":{"single":{"mode":"SIGN_MODE_DIRECT"}},"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Ayjv5FHoI0MML8HX0nU0z7zhkO+eGVcmb4nAAnzlyssR"},"sequence":"0"}]},"body":{"extension_options":[],"memo":"091067d86e57c2a4e87cf4562845751292528dc0@192.168.0.37:26656","messages":[{"@type":"/cosmos.staking.v1beta1.MsgCreateValidator","commission":{"max_change_rate":"0.010000000000000000","max_rate":"0.200000000000000000","rate":"0.100000000000000000"},"delegator_address":"rizon1rtusr0r6pa9jp02k9ejzprn40qnssgjypzgg2c","description":{"details":"","identity":"","moniker":"my_node","security_contact":"","website":""},"min_self_delegation":"1","pubkey":{"@type":"/cosmos.crypto.ed25519.PubKey","key":"ggkWhyCA8lll4X57VvTekTTwUwP46G5gukjR2xrjfF4="},"validator_address":"rizonvaloper1rtusr0r6pa9jp02k9ejzprn40qnssgjyjx2fgg","value":{"amount":"1000000000","denom":"stake"}}],"non_critical_extension_options":[],"timeout_height":"0"},"signatures":["oVBdo91FUjLY067ZD5YpnvP1BQz46dDdESPry8H3VDQ6vO09y1eyoXY9vFnyCVcns+lZKjOAaIFuXhcPV8oW+g=="]}]},"gov":{"deposit_params":{"max_deposit_period":"172800s","min_deposit":[{"amount":"10000000","denom":"stake"}]},"deposits":[],"proposals":[],"starting_proposal_id":"1","tally_params":{"quorum":"0.334000000000000000","threshold":"0.500000000000000000","veto_threshold":"0.334000000000000000"},"votes":[],"voting_params":{"voting_period":"172800s"}},"ibc":{"channel_genesis":{"ack_sequences":[],"acknowledgements":[],"channels":[],"commitments":[],"next_channel_sequence":"0","receipts":[],"recv_sequences":[],"send_sequences":[]},"client_genesis":{"clients":[],"clients_consensus":[],"clients_metadata":[],"create_localhost":false,"next_client_sequence":"0","params":{"allowed_clients":["06-solomachine","07-tendermint"]}},"connection_genesis":{"client_connection_paths":[],"connections":[],"next_connection_sequence":"0"}},"mint":{"minter":{"annual_provisions":"0.000000000000000000","inflation":"0.130000000000000000"},"params":{"blocks_per_year":"6311520","goal_bonded":"0.670000000000000000","inflation_max":"0.200000000000000000","inflation_min":"0.070000000000000000","inflation_rate_change":"0.130000000000000000","mint_denom":"stake"}},"params":null,"slashing":{"missed_blocks":[],"params":{"downtime_jail_duration":"600s","min_signed_per_window":"0.500000000000000000","signed_blocks_window":"100","slash_fraction_double_sign":"0.050000000000000000","slash_fraction_downtime":"0.010000000000000000"},"signing_infos":[]},"staking":{"delegations":[],"exported":false,"last_total_power":"0","last_validator_powers":[],"params":{"bond_denom":"stake","historical_entries":10000,"max_entries":7,"max_validators":100,"unbonding_time":"1814400s"},"redelegations":[],"unbonding_delegations":[],"validators":[]},"tokenswap":{"params":{"signer":"rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7","swappable":true}},"transfer":{"denom_traces":[],"params":{"receive_enabled":true,"send_enabled":true},"port_id":"transfer"},"treasury":{"params":{"currency_list":[{"denom":"skrw","desc":"stable coin of KRW","mintable":true,"owner":"rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7"},{"denom":"susd","desc":"stable coin of USD","mintable":true,"owner":"rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7"}],"mintable":true,"sequence":"1"},"seq":{"number":"0"}},"upgrade":{},"vesting":{}},"chain_id":"my_testnet","gentxs_dir":"/Users/yoosah/.rizon/config/gentx","moniker":"my_node","node_id":"091067d86e57c2a4e87cf4562845751292528dc0"}

$ rizond validate-genesis
File at /Users/yoosah/.rizon/config/genesis.json is a valid genesis file
```

## Start your node

```bash
$ rizond start
5:14PM INF starting ABCI with Tendermint
5:14PM INF Starting multiAppConn service impl=multiAppConn module=proxy
5:14PM INF Starting localClient service connection=query impl=localClient module=abci-client
5:14PM INF Starting localClient service connection=snapshot impl=localClient module=abci-client
5:14PM INF Starting localClient service connection=mempool impl=localClient module=abci-client
5:14PM INF Starting localClient service connection=consensus impl=localClient module=abci-client
5:14PM INF Starting EventBus service impl=EventBus module=events
5:14PM INF Starting PubSub service impl=PubSub module=pubsub
5:14PM INF Starting IndexerService service impl=IndexerService module=txindex
5:14PM INF ABCI Handshake App Info hash= height=0 module=consensus protocol-version=0 software-version=
5:14PM INF ABCI Replay Blocks appHeight=0 module=consensus stateHeight=0 storeHeight=0
5:14PM INF asserting crisis invariants inv=0/11 module=x/crisis name=gov/module-account
5:14PM INF asserting crisis invariants inv=1/11 module=x/crisis name=distribution/nonnegative-outstanding
5:14PM INF asserting crisis invariants inv=2/11 module=x/crisis name=distribution/can-withdraw
5:14PM INF asserting crisis invariants inv=3/11 module=x/crisis name=distribution/reference-count
5:14PM INF asserting crisis invariants inv=4/11 module=x/crisis name=distribution/module-account
5:14PM INF asserting crisis invariants inv=5/11 module=x/crisis name=bank/nonnegative-outstanding
5:14PM INF asserting crisis invariants inv=6/11 module=x/crisis name=bank/total-supply
5:14PM INF asserting crisis invariants inv=7/11 module=x/crisis name=staking/module-accounts
5:14PM INF asserting crisis invariants inv=8/11 module=x/crisis name=staking/nonnegative-power
5:14PM INF asserting crisis invariants inv=9/11 module=x/crisis name=staking/positive-delegation
5:14PM INF asserting crisis invariants inv=10/11 module=x/crisis name=staking/delegator-shares
5:14PM INF asserted all invariants duration=1.60729 height=0 module=x/crisis
5:14PM INF created new capability module=ibc name=ports/transfer
5:14PM INF port binded module=x/ibc/port port=transfer
5:14PM INF claimed capability capability=1 module=transfer name=ports/transfer
5:14PM INF Completed ABCI Handshake - Tendermint and App are synced appHash= appHeight=0 module=consensus
5:14PM INF Version info block=11 p2p=8 software=v0.34.9
5:14PM INF This node is a validator addr=2EC4F6D58A4A4A396F75CA692CB267BEED75125A module=consensus pubKey=ggkWhyCA8lll4X57VvTekTTwUwP46G5gukjR2xrjfF4=
5:14PM INF P2P Node ID ID=091067d86e57c2a4e87cf4562845751292528dc0 file=/Users/yoosah/.rizon/config/node_key.json module=p2p
5:14PM INF Adding persistent peers addrs=[] module=p2p
5:14PM INF Adding unconditional peer ids ids=[] module=p2p
5:14PM INF Add our address to book addr={"id":"091067d86e57c2a4e87cf4562845751292528dc0","ip":"0.0.0.0","port":26656} book=/Users/yoosah/.rizon/config/addrbook.json module=p2p
5:14PM INF Starting Node service impl=Node
5:14PM INF Starting pprof server laddr=localhost:6060
5:14PM INF Starting P2P Switch service impl="P2P Switch" module=p2p
5:14PM INF Starting PEX service impl=PEX module=pex
5:14PM INF Starting AddrBook service book=/Users/yoosah/.rizon/config/addrbook.json impl=AddrBook module=p2p
5:14PM INF Starting Mempool service impl=Mempool module=mempool
5:14PM INF Starting BlockchainReactor service impl=BlockchainReactor module=blockchain
5:14PM INF Starting Consensus service impl=ConsensusReactor module=consensus
5:14PM INF Reactor  module=consensus waitSync=false
5:14PM INF Starting State service impl=ConsensusState module=consensus
5:14PM INF Ensure peers module=pex numDialing=0 numInPeers=0 numOutPeers=0 numToDial=10
5:14PM INF Starting RPC HTTP server on 127.0.0.1:26657 module=rpc-server
5:14PM INF No addresses to dial. Falling back to seeds module=pex
5:14PM INF Starting baseWAL service impl=baseWAL module=consensus wal=/Users/yoosah/.rizon/data/cs.wal/wal
5:14PM INF Starting Group service impl=Group module=consensus wal=/Users/yoosah/.rizon/data/cs.wal/wal
5:14PM INF Searching for height height=1 max=0 min=0 module=consensus wal=/Users/yoosah/.rizon/data/cs.wal/wal
5:14PM INF Searching for height height=0 max=0 min=0 module=consensus wal=/Users/yoosah/.rizon/data/cs.wal/wal
5:14PM INF Found height=0 index=0 module=consensus wal=/Users/yoosah/.rizon/data/cs.wal/wal
5:14PM INF Catchup by replaying consensus messages height=1 module=consensus
5:14PM INF Replay: Done module=consensus
5:14PM INF Starting TimeoutTicker service impl=TimeoutTicker module=consensus
5:14PM INF Starting Evidence service impl=Evidence module=evidence
5:14PM INF Starting StateSync service impl=StateSync module=statesync
5:14PM INF Saving AddrBook to file book=/Users/yoosah/.rizon/config/addrbook.json module=p2p size=0
5:14PM INF Timed out dur=4964.853 height=1 module=consensus round=0 step=1
5:14PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"5D10855668904C360D7387ECF8701189F8C2DD48F0BB0448C24827ACFAAD6329","parts":{"hash":"0742F82AC474BB85AE01BD3DDEC6DA12221F8D3FA5E025E81445B9D0CFD923E7","total":1}},"height":1,"pol_round":-1,"round":0,"signature":"PsldeP7KVDdO5mN0Tf/33LQmXPG9TGNZWBRKMRkvhsA368ZIgoE5740SF4f+k1rL/NwxQXw8LWE+zC/+85NTDQ==","timestamp":"2021-05-25T08:14:22.515323Z"}
5:14PM INF received complete proposal block hash=5D10855668904C360D7387ECF8701189F8C2DD48F0BB0448C24827ACFAAD6329 height=1 module=consensus
5:14PM INF finalizing commit of block hash=5D10855668904C360D7387ECF8701189F8C2DD48F0BB0448C24827ACFAAD6329 height=1 module=consensus num_txs=0 root=E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855
5:14PM INF minted coins from module account amount=205stake from=mint module=x/bank
5:14PM INF executed block height=1 module=state num_invalid_txs=0 num_valid_txs=0
5:14PM INF commit synced commit=436F6D6D697449447B5B363020323239203633203133382031303620313935203131312032392031373720313136203137352039322032303020313735203133352031323520343520323437203234362031353120313438203233342031303620313635203134382031333720353020313320353920313836203232322039305D3A317D
5:14PM INF committed state app_hash=3CE53F8A6AC36F1DB174AF5CC8AF877D2DF7F69794EA6AA59489320D3BBADE5A height=1 module=state num_txs=0
5:14PM INF indexed block height=1 module=txindex
```



**Congratulations!** Now you have your own Rizon network.

