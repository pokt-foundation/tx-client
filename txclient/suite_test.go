package txclient

import (
	"testing"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
	"github.com/stretchr/testify/suite"
)

type txClientTestSuite struct {
	suite.Suite
	client TxClientRW
	// Use to reference primary keys constraints
	relay  types.Relay
	sr     types.ServiceRecord
	region string
}

func Test_RunTXClientTestSuite(t *testing.T) {
	suite.Run(t, new(txClientTestSuite))
}

// SetupSuite runs before each test suite run
func (ts *txClientTestSuite) SetupSuite() {
	ts.NoError(ts.initClient())

	ts.region = "region"

	ts.NoError(ts.client.CreateRegion(types.PortalRegion{
		PortalRegionName: ts.region,
	}))

	ts.NoError(ts.client.CreateSession(types.PocketSession{
		SessionKey:       "abc",
		SessionHeight:    22,
		PortalRegionName: ts.region,
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
		PortalRegionName:         ts.region,
		IsAltruistRelay:          false,
		RelaySourceURL:           "example.com",
		PoktNodeDomain:           "node.com",
		PoktNodePublicKey:        "1234",
		RequestID:                "1234",
		PoktTxID:                 "1234",
	}))

	ts.NoError(ts.client.CreateServiceRecord(types.ServiceRecord{
		NodePublicKey:          "123",
		PoktChainID:            "0001",
		SessionKey:             "abc",
		RequestID:              "123",
		PortalRegionName:       ts.region,
		Latency:                0.112,
		Tickets:                10,
		Result:                 "result",
		Available:              true,
		Successes:              100,
		Failures:               1,
		P90SuccessLatency:      0.05,
		MedianSuccessLatency:   0.09,
		WeightedSuccessLatency: 0.1,
		SuccessRate:            0.9,
	}))

	// write endpoints for relays and service records in transaction-http-db return
	// before saving to the db, which may end up in the test failing due to fecthing
	// an item before being stored.
	time.Sleep(500 * time.Millisecond)

	dbRelay, err := ts.client.GetRelay(1)
	ts.NoError(err)
	ts.relay = dbRelay

	sr, err := ts.client.GetServiceRecord(1)
	ts.NoError(err)
	ts.sr = sr
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
