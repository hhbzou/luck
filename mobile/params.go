// Copyright 2020 The go-luck Authors
// This file is part of the go-luck library.
//
// The go-luck library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-luck library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-luck library. If not, see <http://www.gnu.org/licenses/>.

// Contains all the wrappers from the params package.

package luck

import (
	"encoding/json"

	"github.com/luck/go-luck/core"
	"github.com/luck/go-luck/p2p/discv5"
	"github.com/luck/go-luck/params"
)

// MainnetGenesis returns the JSON spec to use for the main Luck network. It
// is actually empty since that defaults to the hard coded binary genesis block.
func MainnetGenesis() string {
	return ""
}

// RopstenGenesis returns the JSON spec to use for the Ropsten test network.
func RopstenGenesis() string {
	enc, err := json.Marshal(core.DefaultRopstenGenesisBlock())
	if err != nil {
		panic(err)
	}
	return string(enc)
}

// TestnetGenesis returns the JSON spec to use for the Testnet test network
func TestnetGenesis() string {
	enc, err := json.Marshal(core.DefaultTestnetGenesisBlock())
	if err != nil {
		panic(err)
	}
	return string(enc)
}

// GoerliGenesis returns the JSON spec to use for the Goerli test network
func GoerliGenesis() string {
	enc, err := json.Marshal(core.DefaultGoerliGenesisBlock())
	if err != nil {
		panic(err)
	}
	return string(enc)
}

// FoundationBootnodes returns the lnode URLs of the P2P bootstrap nodes operated
// by the foundation running the V5 discovery protocol.
func FoundationBootnodes() *Enodes {
	nodes := &Enodes{nodes: make([]*discv5.Node, len(params.DiscoveryV5Bootnodes))}
	for i, url := range params.DiscoveryV5Bootnodes {
		nodes.nodes[i] = discv5.MustParseNode(url)
	}
	return nodes
}
