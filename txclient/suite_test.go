package txclient

import (
	"testing"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
	"github.com/stretchr/testify/suite"
)

type txClientTestSuite struct {
	suite.Suite
	client TXDBClient
	// Use to reference primary keys constraints
	relay types.Relay
}

func Test_RunTXClientTestSuite(t *testing.T) {
	suite.Run(t, new(txClientTestSuite))
}

// SetupSuite runs before each test suite run
func (ts *txClientTestSuite) SetupSuite() {
	ts.NoError(ts.initClient())

	ts.NoError(ts.client.CreateSession(types.PocketSession{
		SessionKey:    "abc",
		SessionHeight: 22,
	}))

	ts.NoError(ts.client.CreateRegion(types.PortalRegion{
		PortalRegionName: "Los Praditos",
	}))

	ts.NoError(ts.client.CreateRelay(types.Relay{
		PoktChainID:              "0001",
		ProtocolAppPublicKey:     "22",
		EndpointID:               "21",
		SessionKey:               "abc",
		PoktNodeAddress:          "21",
		RelayStartDatetime:       time.Date(199, time.July, 21, 0, 0, 0, 0, time.Local),
		RelayReturnDatetime:      time.Date(199, time.July, 21, 0, 0, 0, 0, time.Local),
		IsError:                  false,
		RelayRoundtripTime:       1,
		RelayChainMethodIDs:      []string{"eth_getLogs"},
		RelayDataSize:            21,
		RelayPortalTripTime:      21,
		RelayNodeTripTime:        21,
		RelayURLIsPublicEndpoint: false,
		PortalRegionName:         "Los Praditos",
		IsAltruistRelay:          false,
		RelaySourceURL:           "example.com",
		PoktNodeDomain:           "node.com",
		PoktNodePublicKey:        "1234",
		RequestID:                "1234",
		PoktTxID:                 "1234",
	}))

	dbRelay, err := ts.client.GetRelay(1)
	ts.NoError(err)

	ts.relay = dbRelay
}

func (ts *txClientTestSuite) initClient() error {
	client, err := NewTXClient(Config{
		BaseURL: "http://localhost:8080",
		APIKey:  "test_api_key",
		Version: V0,
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	ts.client = client

	return nil
}
