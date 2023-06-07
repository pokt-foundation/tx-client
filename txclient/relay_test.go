package txclient

import (
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

// TODO: Write failure tests
func (ts *txClientTestSuite) TestClient_WriteRelay() {
	tests := []struct {
		name  string
		relay types.Relay
		err   error
	}{
		{
			name: "success writing a single relay",
			relay: types.Relay{
				PoktChainID:              "0001",
				ProtocolAppPublicKey:     "unit",
				EndpointID:               "21",
				SessionKey:               ts.relay.SessionKey,
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  false,
				RelayRoundtripTime:       1,
				RelayChainMethodIDs:      []string{"eth_getLogs"},
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalRegionName:         ts.relay.PortalRegionName,
				IsAltruistRelay:          false,
				RelaySourceURL:           "example.com",
				PoktNodeDomain:           "node.com",
				PoktNodePublicKey:        "1234",
				RequestID:                "1234",
				PoktTxID:                 "1234",
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		ts.Equal(ts.client.CreateRelay(tt.relay), tt.err)
	}
}

func (ts *txClientTestSuite) TestClient_WriteRelays() {
	tests := []struct {
		name   string
		relays []*types.Relay
		err    error
	}{
		{
			name: "success writing multiple relays",
			relays: []*types.Relay{{
				PoktChainID:              "0001",
				ProtocolAppPublicKey:     "22",
				EndpointID:               "21",
				SessionKey:               ts.relay.SessionKey,
				PoktNodeAddress:          "23",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  false,
				RelayRoundtripTime:       1,
				RelayChainMethodIDs:      []string{"eth_getLogs"},
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalRegionName:         ts.relay.PortalRegionName,
				IsAltruistRelay:          false,
				RelaySourceURL:           "example.com",
				PoktNodeDomain:           "node.com",
				PoktNodePublicKey:        "1234",
				RequestID:                "1234",
				PoktTxID:                 "1234",
			},
				{
					PoktChainID:              "0001",
					ProtocolAppPublicKey:     "22",
					EndpointID:               "21",
					SessionKey:               ts.relay.SessionKey,
					PoktNodeAddress:          "24",
					RelayStartDatetime:       time.Now(),
					RelayReturnDatetime:      time.Now(),
					IsError:                  false,
					RelayRoundtripTime:       1,
					RelayChainMethodIDs:      []string{"eth_getLogs"},
					RelayDataSize:            21,
					RelayPortalTripTime:      21,
					RelayNodeTripTime:        21,
					RelayURLIsPublicEndpoint: false,
					PortalRegionName:         ts.relay.PortalRegionName,
					IsAltruistRelay:          false,
					RelaySourceURL:           "example.com",
					PoktNodeDomain:           "node.com",
					PoktNodePublicKey:        "1234",
					RequestID:                "1234",
					PoktTxID:                 "1234",
				}},
			err: nil,
		},
	}

	for _, tt := range tests {
		ts.Equal(ts.client.CreateRelays(tt.relays), tt.err)
	}
}

func (ts *txClientTestSuite) TestClient_ReadRelay() {
	tests := []struct {
		name    string
		relayID int
		relay   types.Relay
		err     error
	}{
		{
			name:    "success reading a single relay",
			relayID: ts.relay.RelayID,
			relay:   ts.relay,
			err:     nil,
		},
	}

	for _, tt := range tests {
		relay, err := ts.client.GetRelay(tt.relayID)
		ts.Equal(err, tt.err)
		ts.Equal(relay, tt.relay)
	}
}
