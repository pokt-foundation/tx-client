package txclient

import (
	"github.com/pokt-foundation/transaction-db/types"
)

// TODO: Write failure tests
func (ts *txClientTestSuite) TestClient_WriteServiceRecord() {
	tests := []struct {
		name string
		sr   types.ServiceRecord
		err  error
	}{
		{
			name: "success writing a single service record",
			sr: types.ServiceRecord{
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
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		ts.Equal(ts.client.CreateServiceRecord(tt.sr), tt.err)
	}
}

func (ts *txClientTestSuite) TestClient_WriteServiceRecords() {
	tests := []struct {
		name string
		srs  []*types.ServiceRecord
		err  error
	}{
		{
			name: "success writing multiple service records",
			srs: []*types.ServiceRecord{
				{
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
				},
				{
					NodePublicKey:          "123",
					PoktChainID:            "0001",
					SessionKey:             "abc",
					RequestID:              "123",
					PortalRegionName:       "Los Praditos",
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
				},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		ts.Equal(ts.client.CreateServiceRecords(tt.srs), tt.err)
	}
}

func (ts *txClientTestSuite) TestClient_ReadServiceRecord() {
	tests := []struct {
		name            string
		serviceRecordID int
		sr              types.ServiceRecord
		err             error
	}{
		{
			name:            "success reading a single relay",
			serviceRecordID: ts.sr.ServiceRecordID,
			sr:              ts.sr,
			err:             nil,
		},
	}

	for _, tt := range tests {
		sr, err := ts.client.GetServiceRecord(tt.serviceRecordID)
		ts.Equal(err, tt.err)
		ts.Equal(sr, tt.sr)
	}
}
