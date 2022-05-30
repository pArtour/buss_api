CREATE TABLE "stop" (
    "stop_id" int NOT NULL,
    "code" varchar,
    "name" varchar NOT NULL,
    "lat" float8 NOT NULL,
    "lon" float8 NOT NULL,
    "zone_id" int NOT NULL,
    "alias" varchar,
    "area" varchar NOT NULL,
    "desc" varchar,
    "lest_x" float8 NOT NULL,
    "lest_y" float8 NOT NULL,
    "zone_name" varchar,
    "authority" varchar
);

CREATE TABLE "place" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "name" varchar NOT NULL
);

CREATE TABLE "route" (
    "route_id" varchar NOT NULL,
    "agency_id" int,
    "route_short_name" varchar,
    "route_long_name" varchar,
    "route_type" varchar,
    "route_color" varchar
);

CREATE TABLE "trip" (
    "route_id" varchar NOT NULL,
    "service_id" int NOT NULL,
    "trip_id" int NOT NULL,
    "trip_headsign" varchar,
    "trip_long_name" varchar,
    "direction_code" varchar,
    "wheelchair_accessible" int
);

CREATE TABLE "stop_time" (
     "trip_id" int NOT NULL,
     "arrival_time" varchar NOT NULL,
     "departure_time" varchar NOT NULL,
     "stop_id" int NOT NULL,
     "stop_sequence" int,
     "pickup_type" int,
     "drop_off_type" int
);