/*
 * Copyright © 2018. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package common

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This test requires to start byfn fabric network using "byfn.sh up -a -s couchdb"
//   and set the fabPath below to the absolute path of "fabric-samples"
const (
	connectorName = "test"
	fabPath       = "/Users/yxu/work/dovetail-lab2/hyperledger/fabric-samples"
	configFile    = "${HOME}/work/dovetail-lab2/fabric-samples/testdata/config_test.yaml"
	matcherFile   = "${HOME}/work/dovetail-lab2/fabric-samples/testdata/local_entity_matchers.yaml"
	channelID     = "mychannel"
	org           = "org1"
	user          = "User1"
	ccID          = "basic"
)

func TestClient(t *testing.T) {
	os.Setenv("FAB_PATH", fabPath)
	networkConfig, err := ReadFile(configFile)
	require.NoError(t, err, "failed to read config file %s", configFile)

	entityMatcherOverride, err := ReadFile(matcherFile)
	require.NoError(t, err, "failed to read entity matcher file %s", matcherFile)

	fbClient, err := NewFabricClient(ConnectorSpec{
		Name:           connectorName,
		NetworkConfig:  networkConfig,
		EntityMatchers: entityMatcherOverride,
		OrgName:        org,
		UserName:       user,
		ChannelID:      channelID,
	})
	require.NoError(t, err, "failed to create fabric client %s", connectorName)
	fmt.Printf("created fabric client %+v\n", fbClient)

	// initialize ledger
	result, _, err := fbClient.ExecuteChaincode(ccID, "InitLedger", [][]byte{}, nil)
	require.NoError(t, err, "failed to invoke %s", ccID)
	fmt.Printf("InitLedger result: %s\n", string(result))

	// query original
	result, _, err = fbClient.QueryChaincode(ccID, "ReadAsset", [][]byte{[]byte("asset6")}, nil)
	require.NoError(t, err, "failed to query %s", ccID)
	fmt.Printf("Query asset6 result: %s\n", string(result))
	origValue := result

	// update
	result, _, err = fbClient.ExecuteChaincode(ccID, "TransferAsset", [][]byte{[]byte("asset6"), []byte("Jose")}, nil)
	require.NoError(t, err, "failed to invoke %s", ccID)
	fmt.Printf("Transfer asset6 result: %s\n", string(result))

	// query after update
	result, _, err = fbClient.QueryChaincode(ccID, "ReadAsset", [][]byte{[]byte("asset6")}, nil)
	require.NoError(t, err, "failed to query %s", ccID)
	fmt.Printf("Query asset6 result: %s\n", string(result))
	assert.NotEqual(t, origValue, result, "original %s should different from %s", string(origValue), string(result))
}
