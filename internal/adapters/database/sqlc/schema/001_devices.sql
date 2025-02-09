-- +goose Up

CREATE TABLE devices(
                      id UUID PRIMARY KEY,
                      created_at TIMESTAMP NOT NULL,
                      updated_at TIMESTAMP NOT NULL,
                      uid TEXT NOT NULL,
                      serial TEXT NOT NULL
);

-- +goose Down

DROP TABLE devices;
