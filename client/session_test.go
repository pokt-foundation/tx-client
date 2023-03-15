package txclient

import "github.com/pokt-foundation/transaction-db/types"

func (ts *txClientTestSuite) TestClient_WriteSession() {
	tests := []struct {
		name    string
		session types.PocketSession
		err     error
	}{
		{
			name: "success writing a session",
			session: types.PocketSession{
				SessionKey:            "abcd",
				SessionHeight:         10,
				ProtocolApplicationID: 10,
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		ts.Equal(ts.client.CreateSession(tt.session), tt.err)
	}
}
