package txclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/pokt-foundation/transaction-db/types"
)

// APIVersion is a version of the API, usage not currently implemented, declared to avoid breaking
// Changes when a version bump is needed.
type APIVersion string

const (
	V0 APIVersion = "v0"
)

var validAPIVersions = map[APIVersion]bool{
	V0: true,
}

type typePath string

const (
	sessionPath typePath = "session"
	regionPath  typePath = "region"
	relayPath   typePath = "relay"
	relaysPath  typePath = "relays"
)

type TXDBClient interface {
	CreateSession(types.PocketSession) error
	CreateRegion(types.PortalRegion) error
	CreateRelay(types.Relay) error
	CreateRelays([]types.Relay) error
	GetRelay(int) (types.Relay, error)
}

var _ TXDBClient = TXClient{}

type TXClient struct {
	httpClient *httpclient.Client
	config     Config
	headers    http.Header
}

func NewTXClient(config Config) (TXDBClient, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	return &TXClient{
		httpClient: newHTTPClient(config),
		config:     config,
		headers:    http.Header{"Authorization": {config.APIKey}, "Content-Type": {"application/json"}},
	}, nil
}

// versionedBasePath returns the base path for a given data type eg. `https://pocket.http-db-url.com/v1/application`
func (db TXClient) versionedBasePath(dataTypePath typePath) string {
	return fmt.Sprintf("%s/%s/%s", db.config.BaseURL, db.config.Version, dataTypePath)
}

func (db TXClient) CreateSession(session types.PocketSession) error {
	body, err := json.Marshal(session)
	if err != nil {
		return err
	}

	_, err = performHttpReq[any](http.MethodPost, db.versionedBasePath(sessionPath), db.headers, body, db.httpClient)
	return err
}

func (db TXClient) CreateRegion(region types.PortalRegion) error {
	body, err := json.Marshal(region)
	if err != nil {
		return err
	}

	_, err = performHttpReq[any](http.MethodPost, db.versionedBasePath(regionPath), db.headers, body, db.httpClient)
	return err
}

func (db TXClient) CreateRelay(relay types.Relay) error {
	if err := relay.Validate(); err != nil {
		return err
	}

	body, err := json.Marshal(relay)
	if err != nil {
		return err
	}

	_, err = performHttpReq[any](http.MethodPost, db.versionedBasePath(relayPath), db.headers, body, db.httpClient)
	return err
}

func (db TXClient) CreateRelays(relays []types.Relay) error {
	for _, relay := range relays {
		if err := relay.Validate(); err != nil {
			return err
		}
	}

	body, err := json.Marshal(relays)
	if err != nil {
		return err
	}

	_, err = performHttpReq[any](http.MethodPost, db.versionedBasePath(relaysPath), db.headers, body, db.httpClient)
	return err
}

func (db TXClient) GetRelay(id int) (types.Relay, error) {
	path := fmt.Sprintf("%s/%d", db.versionedBasePath(relayPath), id)
	return performHttpReq[types.Relay](http.MethodGet, path, db.headers, nil, db.httpClient)
}
