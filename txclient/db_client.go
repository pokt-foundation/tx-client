package txclient

import (
	"context"
	"errors"
	"time"

	postgresdriver "github.com/pokt-foundation/transaction-db/postgres-driver"
	"github.com/pokt-foundation/transaction-db/types"
)

type DBClient struct {
	driver  *postgresdriver.PostgresDriver
	Timeout time.Duration
}

func NewDBClient(connectionString string, timeout time.Duration) (*DBClient, error) {
	driver, err := postgresdriver.NewPostgresDriver(connectionString)
	if err != nil {
		return nil, err
	}

	return &DBClient{
		driver:  driver,
		Timeout: timeout,
	}, nil
}

func (db DBClient) CreateSession(session types.PocketSession) error {
	if err := session.Validate(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.Timeout)
	defer cancel()

	return db.driver.WriteSession(ctx, session)
}

func (db DBClient) CreateRegion(region types.PortalRegion) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.Timeout)
	defer cancel()

	return db.driver.WriteRegion(ctx, region)
}

func (db DBClient) CreateRelay(relay types.Relay) error {
	if err := relay.Validate(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.Timeout)
	defer cancel()

	return db.driver.WriteRelay(ctx, relay)
}

func (db DBClient) CreateRelays(relays []types.Relay) error {
	var relayErrors []error
	var validRelays []*types.Relay

	for _, relay := range relays {
		if err := relay.Validate(); err != nil {
			relayErrors = append(relayErrors, err)
		}
		validRelays = append(validRelays, &relay)
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.Timeout)
	defer cancel()

	err := db.driver.WriteRelays(ctx, validRelays)

	if relayErrors != nil || err != nil {
		relayErrors = append(relayErrors, err)
		return errors.Join(relayErrors...)
	}

	return nil
}

func (db DBClient) GetRelay(id int) (types.Relay, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.Timeout)
	defer cancel()

	return db.driver.ReadRelay(ctx, id)
}

func (db DBClient) CreateServiceRecord(record types.ServiceRecord) error {
	if err := record.Validate(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.Timeout)
	defer cancel()

	return db.driver.WriteServiceRecord(ctx, record)
}

func (db DBClient) CreateServiceRecords(records []types.ServiceRecord) error {
	var recordrrors []error
	var validRecords []*types.ServiceRecord

	for _, relay := range records {
		if err := relay.Validate(); err != nil {
			recordrrors = append(recordrrors, err)
		}
		validRecords = append(validRecords, &relay)
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.Timeout)
	defer cancel()

	err := db.driver.WriteServiceRecords(ctx, validRecords)

	if recordrrors != nil || err != nil {
		recordrrors = append(recordrrors, err)
		return errors.Join(recordrrors...)
	}

	return nil
}

func (db DBClient) GetServiceRecord(id int) (types.ServiceRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.Timeout)
	defer cancel()

	return db.driver.ReadServiceRecord(ctx, id)
}
