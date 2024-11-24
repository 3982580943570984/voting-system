package database

import (
	"context"
	"log"
	ent "voting-system/ent/generated"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

var (
	Client *ent.Client
	Users  *ent.UserClient
)

func init() {
	var err error

	dsn := "postgresql://postgres@database:5432/database?sslmode=disable"

	Client, err = ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("Error during database open: %v", err)
	}

	err = Client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("Error during schema creation: %v", err)
	}

	Users = Client.User
}
