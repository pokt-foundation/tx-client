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
	sessionPath        typePath = "session"
	regionPath         typePath = "region"
	relayPath          typePath = "relay"
	relaysPath         typePath = "relays"
	serviceRecordPath  typePath = "service-record"
	serviceRecordsPath typePath = "service-records"
)

type TxClientWrite interface {
	CreateSession(types.PocketSession) error
	CreateRegion(types.PortalRegion) error
	CreateRelay(types.Relay) error
	CreateRelays([]types.Relay) error
	CreateServiceRecord(types.ServiceRecord) error
	CreateServiceRecords([]types.ServiceRecord) error
}

type TxClientRead interface {
	GetRelay(int) (types.Relay, error)
	GetServiceRecord(int) (types.ServiceRecord, error)
}

type TxClientRW interface {
	TxClientRead
	TxClientWrite
}

var _ TxClientRW = HttpClient{}

type HttpClient struct {
	httpClient *httpclient.Client
	config     Config
	headers    http.Header
}

func NewHttpClient(config Config) (*HttpClient, error) {
	if err := config.validate(); err != nil {
		return &HttpClient{}, err
	}

	return &HttpClient{
		httpClient: newHTTPClient(config),
		config:     config,
		headers:    http.Header{"Authorization": {config.APIKey}, "Content-Type": {"application/json"}},
	}, nil
}

// versionedBasePath returns the base path for a given data type eg. `https://pocket.http-db-url.com/v1/application`
func (db HttpClient) versionedBasePath(dataTypePath typePath) string {
	return fmt.Sprintf("%s/%s/%s", db.config.BaseURL, db.config.Version, dataTypePath)
}

func (db HttpClient) CreateSession(session types.PocketSession) error {
	body, err := json.Marshal(session)
	if err != nil {
		return err
	}

	_, err = performHttpReq[any](http.MethodPost, db.versionedBasePath(sessionPath), db.headers, body, db.httpClient)
	return err
}

func (db HttpClient) CreateRegion(region types.PortalRegion) error {
	body, err := json.Marshal(region)
	if err != nil {
		return err
	}

	_, err = performHttpReq[any](http.MethodPost, db.versionedBasePath(regionPath), db.headers, body, db.httpClient)
	return err
}

func (db HttpClient) CreateRelay(relay types.Relay) error {
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

func (db HttpClient) CreateRelays(relays []types.Relay) error {
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

func (db HttpClient) CreateServiceRecord(sr types.ServiceRecord) error {
	if err := sr.Validate(); err != nil {
		return err
	}

	body, err := json.Marshal(sr)
	if err != nil {
		return err
	}

	_, err = performHttpReq[any](http.MethodPost, db.versionedBasePath(serviceRecordPath), db.headers, body, db.httpClient)
	return err
}

func (db HttpClient) CreateServiceRecords(srs []types.ServiceRecord) error {
	for _, record := range srs {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	body, err := json.Marshal(srs)
	if err != nil {
		return err
	}

	_, err = performHttpReq[any](http.MethodPost, db.versionedBasePath(serviceRecordsPath), db.headers, body, db.httpClient)
	return err
}

func (db HttpClient) GetRelay(id int) (types.Relay, error) {
	path := fmt.Sprintf("%s/%d", db.versionedBasePath(relayPath), id)
	return performHttpReq[types.Relay](http.MethodGet, path, db.headers, nil, db.httpClient)
}

func (db HttpClient) GetServiceRecord(id int) (types.ServiceRecord, error) {
	path := fmt.Sprintf("%s/%d", db.versionedBasePath(serviceRecordPath), id)
	fmt.Println("path", path)
	return performHttpReq[types.ServiceRecord](http.MethodGet, path, db.headers, nil, db.httpClient)
}
