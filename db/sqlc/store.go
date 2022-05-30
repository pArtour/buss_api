package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (s *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); err != nil {
			return fmt.Errorf("tx: %v, rbErr: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type ExtendedTime struct {
	StopTime
	Trips  []Trip  `json:"trips"`
	Routes []Route `json:"routes"`
}

type StopRouteInfo struct {
	Times []ExtendedTime `json:"times"`
}

// GetStopRouteInfo get detailed info about route
func (s *Store) GetStopRouteInfo(ctx context.Context, stopId int32) (StopRouteInfo, error) {
	var result StopRouteInfo

	err := s.execTx(ctx, func(queries *Queries) error {
		var err error
		times, err := queries.ListTimesByStop(ctx, stopId)
		if err != nil {
			return err
		}

		if len(times) == 0 {
			return fmt.Errorf("no times found")
		}

		for _, time := range times {
			listedTrips, err := queries.ListTripsById(ctx, time.TripID)
			if err != nil {
				return err
			}

			routes := []Route{}
			for _, trip := range listedTrips {
				route, err := queries.GetRouteById(ctx, trip.RouteID)
				if err != nil {
					return err
				}
				routes = append(routes, route)
			}

			result.Times = append(result.Times, ExtendedTime{time, listedTrips, routes})
		}
		return err
	})

	return result, err
}
