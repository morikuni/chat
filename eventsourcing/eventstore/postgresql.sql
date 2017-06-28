CREATE TABLE IF NOT EXISTS event (
    id BIGSERIAL PRIMARY KEY,
    aggregate_id VARCHAR(255) NOT NULL,
    version SMALLSERIAL NOT NULL,
    package VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    occured_at TIMESTAMP NOT NULL,
    data TEXT NOT NULL,

    UNIQUE(aggregate_id, version)
);
