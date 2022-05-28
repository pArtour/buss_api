CREATE TABLE "stop" (
    "id" SERIAL PRIMARY KEY,
    "code" varchar,
    "name" varchar,
    "lat" float8,
    "lon" float8,
    "zone_id" int NOT NULL,
    "alias" varchar,
    "area" varchar,
    "desc" varchar,
    "lest_x" float8,
    "lest_y" float8,
    "zone_name" varchar,
    "authority" varchar
);

CREATE TABLE "place" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar
);

CREATE TABLE "agency" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "url" varchar,
  "timezone" varchar,
  "phone" varchar,
  "lang" varchar
);

CREATE TABLE "route" (
    "id" SERIAL PRIMARY KEY,
    "agency_id" int NOT NULL,
    "short_name" varchar,
    "long_name" varchar,
    "color" varchar,
    "competent_authority" varchar,
    "desc" varchar
);

CREATE TABLE "trip" (
    "id" SERIAL PRIMARY KEY,
    "route_id" int NOT NULL,
    "headsign" varchar,
    "long_name" varchar,
    "direction_code" varchar,
    "wheelchair_accessible" boolean DEFAULT false
);

CREATE TABLE "stop_time" (
     "id" SERIAL PRIMARY KEY,
     "trip_id" int NOT NULL,
     "arrival_time" varchar,
     "departure_time" varchar,
     "stop_id" int,
     "stop_sequence" int
);

CREATE INDEX ON "stop" ("name");

CREATE INDEX ON "place" ("name");

CREATE INDEX ON "route" ("short_name");

CREATE INDEX ON "route" ("long_name");

CREATE INDEX ON "trip" ("long_name");

CREATE INDEX ON "trip" ("direction_code");

CREATE INDEX ON "stop_time" ("trip_id");

CREATE INDEX ON "stop_time" ("stop_id");

CREATE INDEX ON "stop_time" ("arrival_time", "departure_time");

ALTER TABLE "route" ADD FOREIGN KEY ("agency_id") REFERENCES "agency" ("id");

ALTER TABLE "trip" ADD FOREIGN KEY ("route_id") REFERENCES "route" ("id");

ALTER TABLE "stop_time" ADD FOREIGN KEY ("trip_id") REFERENCES "trip" ("id");

ALTER TABLE "stop_time" ADD FOREIGN KEY ("stop_id") REFERENCES "stop" ("id");
