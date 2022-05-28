// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
)

type Agency struct {
	ID       int32          `json:"id"`
	Name     sql.NullString `json:"name"`
	Url      sql.NullString `json:"url"`
	Timezone sql.NullString `json:"timezone"`
	Phone    sql.NullString `json:"phone"`
	Lang     sql.NullString `json:"lang"`
}

type Place struct {
	ID   int32          `json:"id"`
	Name sql.NullString `json:"name"`
}

type Route struct {
	ID                 int32          `json:"id"`
	AgencyID           int32          `json:"agency_id"`
	ShortName          sql.NullString `json:"short_name"`
	LongName           sql.NullString `json:"long_name"`
	Color              sql.NullString `json:"color"`
	CompetentAuthority sql.NullString `json:"competent_authority"`
	Desc               sql.NullString `json:"desc"`
}

type Stop struct {
	ID        int32           `json:"id"`
	Code      sql.NullString  `json:"code"`
	Name      sql.NullString  `json:"name"`
	Lat       sql.NullFloat64 `json:"lat"`
	Lon       sql.NullFloat64 `json:"lon"`
	ZoneID    int32           `json:"zone_id"`
	Alias     sql.NullString  `json:"alias"`
	Area      sql.NullString  `json:"area"`
	Desc      sql.NullString  `json:"desc"`
	LestX     sql.NullFloat64 `json:"lest_x"`
	LestY     sql.NullFloat64 `json:"lest_y"`
	ZoneName  sql.NullString  `json:"zone_name"`
	Authority sql.NullString  `json:"authority"`
}

type StopTime struct {
	ID            int32          `json:"id"`
	TripID        int32          `json:"trip_id"`
	ArrivalTime   sql.NullString `json:"arrival_time"`
	DepartureTime sql.NullString `json:"departure_time"`
	StopID        sql.NullInt32  `json:"stop_id"`
	StopSequence  sql.NullInt32  `json:"stop_sequence"`
}

type Trip struct {
	ID                   int32          `json:"id"`
	RouteID              int32          `json:"route_id"`
	Headsign             sql.NullString `json:"headsign"`
	LongName             sql.NullString `json:"long_name"`
	DirectionCode        sql.NullString `json:"direction_code"`
	WheelchairAccessible sql.NullBool   `json:"wheelchair_accessible"`
}