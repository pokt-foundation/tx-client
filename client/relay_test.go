package txclient

import (
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

func (ts *txClientTestSuite) TestClient_WriteRelay() {
	tests := []struct {
		name  string
		relay types.Relay
		err   error
	}{
		{
			name: "success writing a single relay",
			relay: types.Relay{
				ChainID:                  21,
				EndpointID:               21,
				SessionKey:               ts.relay.SessionKey,
				PoktNodeAddress:          "21",
				RelayStartDatetime:       time.Now(),
				RelayReturnDatetime:      time.Now(),
				IsError:                  true,
				RelayRoundtripTime:       1,
				RelayChainMethodID:       21,
				RelayDataSize:            21,
				RelayPortalTripTime:      21,
				RelayNodeTripTime:        21,
				RelayURLIsPublicEndpoint: false,
				PortalOriginRegionID:     ts.relay.PortalOriginRegionID,
				IsAltruistRelay:          false,
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		ts.Equal(ts.client.CreateRelay(tt.relay), tt.err)
	}
}

func (ts *txClientTestSuite) TestClient_ReadRelay() {
	tests := []struct {
		name    string
		relayID int64
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
