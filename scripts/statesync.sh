#!/bin/bash
# microtick and bitcanna contributed significantly here.
# Pebbledb state sync script.
set -uxe

# Set Golang environment variables.
export GOPATH=~/go
export PATH=$PATH:~/go/bin

# Install with pebbledb 
#go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
#go mod tidy
#go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...

# NOTE: ABOVE YOU CAN USE ALTERNATIVE DATABASES, HERE ARE THE EXACT COMMANDS
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb' -tags rocksdb ./...
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb' -tags badgerdb ./...
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb' -tags boltdb ./...
# Initialize chain.
qubeticsd init test --chain-id qubetics_9000-2

# Get Genesis
wget https://archive.qubetics.org/mainnet/genesis.json
mv genesis.json ~/.qubeticsd/config/

wget -O ~/.qubeticsd/config/adrbook.json https://snapshot.notional.ventures/qubetics/addrbook.json

# Get "trust_hash" and "trust_height".
INTERVAL=1000
LATEST_HEIGHT=$(curl -s https://qubetics-rpc.polkachu.com/block | jq -r .result.block.header.height)
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL)) 
TRUST_HASH=$(curl -s "https://qubetics-rpc.polkachu.com/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

# Print out block and transaction hash from which to sync state.
echo "trust_height: $BLOCK_HEIGHT"
echo "trust_hash: $TRUST_HASH"

# Export state sync variables.
export qubeticsd_STATESYNC_ENABLE=true
export qubeticsd_P2P_MAX_NUM_INBOUND_PEERS=200
export qubeticsd_P2P_MAX_NUM_OUTBOUND_PEERS=200
export qubeticsd_STATESYNC_RPC_SERVERS="https://qubetics-rpc.polkachu.com:443,https://rpc-qubetics-ia.notional.ventures:443"
export qubeticsd_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export qubeticsd_STATESYNC_TRUST_HASH=$TRUST_HASH

# Fetch and set list of seeds from chain registry.
export qubeticsd_P2P_SEEDS=$(curl -s https://raw.githubusercontent.com/cosmos/chain-registry/master/qubetics/chain.json | jq -r '[foreach .peers.seeds[] as $item (""; "\($item.id)@\($item.address)")] | join(",")')

# Start chain.
# Add the flag --db_backend=pebbledb if you want to use pebble.

qubeticsd start --x-crisis-skip-assert-invariants --p2p.persistent_peers 0fb7c8cbf6b92c0e1002d799c99177e45dfad3bc@15.165.48.105:26656,a6a5a90522d461ea6db9468c043fb9cfb8b82b20@138.201.81.22:26656,90be8866f4714619ad6623fe590512df0a2d7a09@52.201.75.247:26656,348f98b9dcf21025c5946ebff1d8278a2af1c3c2@140.82.60.127:26656,2be90ecc844a071ff56b1a94dc48ff0cc386bf08@135.181.20.164:26656,0188876ca9927965ce7af86dbe0505080434f0b7@142.132.206.174:26656,c7589a3d0be2b6324f4fb55785500e7e20aae977@65.21.200.142:34656,7a8457bb9bd53b06f1b0b876bf423db1e78cbd2d@95.216.45.124:30004,d3abb4dbf82f5f0d1d30f0706c077b2a90379601@178.63.86.221:26656,ca711eb0847c1adf26c2193b0e759e3e8cab8000@188.40.122.160:26656
