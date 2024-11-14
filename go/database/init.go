package database

import (
	"context"
	"log"
	"voting-system/ent"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

var (
	Client     *ent.Client
	Candidates *ent.CandidateClient
	Comments   *ent.CommentClient
	Complaints *ent.ComplaintClient
	Elections  *ent.ElectionClient
	Feedbacks  *ent.FeedbackClient
	Organizers *ent.OrganizerClient
	Profiles   *ent.ProfileClient
	Results    *ent.ResultClient
	Settings   *ent.SettingClient
	Tags       *ent.TagClient
	Users      *ent.UserClient
	Votes      *ent.VoteClient
	Voters     *ent.VoterClient
)

func init() {
	var err error

	dsn := "postgresql://postgres@database:5432/database?sslmode=disable"

	log.Print("Opening database connection")
	Client, err = ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("Error during database open: %v", err)
	}

	log.Print("Creating database schema")
	err = Client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("Error during schema creation: %v", err)
	}

	Candidates = Client.Candidate
	Comments = Client.Comment
	Complaints = Client.Complaint
	Elections = Client.Election
	Feedbacks = Client.Feedback
	Organizers = Client.Organizer
	Profiles = Client.Profile
	Results = Client.Result
	Settings = Client.Setting
	Tags = Client.Tag
	Users = Client.User
	Votes = Client.Vote
	Voters = Client.Voter
}
