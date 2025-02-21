-- +goose Up

CREATE TABLE messages(
                        id UUID PRIMARY KEY,
                        created_at TIMESTAMP NOT NULL,
                        payload jsonb NOT NULL,
                        device_uid TEXT NOT NULL REFERENCES devices(uid)
);

-- +goose Down

DROP TABLE messages;