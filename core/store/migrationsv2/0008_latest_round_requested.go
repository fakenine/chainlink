package migrationsv2

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const (
	up8 = `
CREATE TABLE offchainreporting_latest_round_requested (
	offchainreporting_oracle_spec_id integer PRIMARY KEY REFERENCES offchainreporting_oracle_specs (id) DEFERRABLE INITIALLY IMMEDIATE,
	requester bytea not null CHECK (octet_length(requester) = 20),
	config_digest bytea not null CHECK (octet_length(config_digest) = 16),
	epoch bigint not null,
	round bigint not null,
	raw jsonb not null
);
`
	down8 = `
DROP TABLE offchainreporting_latest_round_requested;
`
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "0008_latest_round_requested",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up8).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down8).Error
		},
	})
}
