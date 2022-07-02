-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "event" (
    "id" uuid PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "description" text NOT NULL,
    "date"  timestamp NOT NULL,
    "location" varchar(255) NOT NULL,
    "create_at" timestamp NOT NULL DEFAULT (now()),
    "update_at" timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "event";
-- +goose StatementEnd
