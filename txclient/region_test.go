package txclient

import "github.com/pokt-foundation/transaction-db/types"

func (ts *txClientTestSuite) TestClient_WriteRegion() {
	tests := []struct {
		name   string
		region types.PortalRegion
		err    error
	}{
		{
			name: "success writing a region",
			region: types.PortalRegion{
				PortalRegionName: "La 42",
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		ts.Equal(ts.client.CreateRegion(tt.region), tt.err)
	}
}
