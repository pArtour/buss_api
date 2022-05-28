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

type RouteSearchParams struct {
	stopId sql.NullInt32
}

type StopRouteInfo struct {
	Times  []StopTime `json:"times"`
	Trip   Trip       `json:"trip"`
	Route  Route      `json:"route"`
	Agency Agency     `json:"agency"`
}

// GetStopRouteInfo get detailed info about route
func (s *Store) GetStopRouteInfo(ctx context.Context, searchParams RouteSearchParams) (StopRouteInfo, error) {
	var result StopRouteInfo

	err := s.execTx(ctx, func(queries *Queries) error {
		var err error
		result.Times, err = queries.ListTimesByStop(ctx, searchParams.stopId)
		if err != nil {
			return err
		}
		if len(result.Times) == 0 {
			return fmt.Errorf("no times found")
		}
		result.Trip, err = queries.GetTripById(ctx, result.Times[0].ID)
		if err != nil {
			return err
		}
		result.Route, err = queries.GetRouteById(ctx, result.Trip.RouteID)
		if err != nil {
			return err
		}
		result.Agency, err = queries.GetAgencyById(ctx, result.Route.AgencyID)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
