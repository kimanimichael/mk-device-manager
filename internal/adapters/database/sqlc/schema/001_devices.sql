-- +goose Up

CREATE TABLE devices(
                      id UUID PRIMARY KEY,
                      created_at TIMESTAMP NOT NULL,
                      updated_at TIMESTAMP NOT NULL,
                      uid TEXT  UNIQUE NOT NULL,
                      serial TEXT UNIQUE NOT NULL
);

-- +goose Down

DROP TABLE devices;
