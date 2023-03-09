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
	errorPath   typePath = "error"
	sessionPath typePath = "session"
	regionPath  typePath = "region"
	relayPath   typePath = "relay"
)

type TXDBReader interface {
}

type TXDBWriter interface {
	CreateError(types.Error) (types.Error, error)
	CreateSession(types.PocketSession) (types.PocketSession, error)
	CreateRegion(types.PortalRegion) (types.PortalRegion, error)
}

type TXDBClient interface {
	TXDBReader
	TXDBWriter
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

func (db TXClient) CreateError(errPayload types.Error) (types.Error, error) {
	body, err := json.Marshal(errPayload)
	if err != nil {
		return types.Error{}, err
	}

	return performHttpReq[types.Error](http.MethodPost, db.versionedBasePath(errorPath), db.headers, body, db.httpClient)
}

func (db TXClient) CreateSession(session types.PocketSession) (types.PocketSession, error) {
	body, err := json.Marshal(session)
	if err != nil {
		return types.PocketSession{}, err
	}

	return performHttpReq[types.PocketSession](http.MethodPost, db.versionedBasePath(sessionPath), db.headers, body, db.httpClient)
}

func (db TXClient) CreateRegion(region types.PortalRegion) (types.PortalRegion, error) {
	body, err := json.Marshal(region)
	if err != nil {
		return types.PortalRegion{}, err
	}

	return performHttpReq[types.PortalRegion](http.MethodPost, db.versionedBasePath(regionPath), db.headers, body, db.httpClient)
}

func (db TXClient) CreateRelay(region types.Relay) (types.Relay, error) {
	body, err := json.Marshal(region)
	if err != nil {
		return types.Relay{}, err
	}

	return performHttpReq[types.Relay](http.MethodPost, db.versionedBasePath(relayPath), db.headers, body, db.httpClient)
}

func (db TXClient) GetRelay(id string) (types.Relay, error) {
	path := fmt.Sprintf("%s/%s", db.versionedBasePath(relayPath), id)
	return performHttpReq[types.Relay](http.MethodPost, path, db.headers, nil, db.httpClient)
}
