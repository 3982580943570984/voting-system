package database

import (
	"log"
	"shared/ent/generated"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

var (
	Client *generated.Client
)

func init() {
	var err error

	dsn := "postgresql://postgres@postgres:5432/database?sslmode=disable"

	Client, err = generated.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("Error during database open: %v", err)
	}

	// err = Client.Schema.Create(context.Background())
	// if err != nil {
	// 	log.Fatalf("Error during schema creation: %v", err)
	// }
}
