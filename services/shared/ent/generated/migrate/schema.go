// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CandidatesColumns holds the columns for the "candidates" table.
	CandidatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 100},
		{Name: "description", Type: field.TypeString, Size: 1000},
		{Name: "photo_url", Type: field.TypeString, Nullable: true},
		{Name: "votes_count", Type: field.TypeInt, Default: 0},
		{Name: "election_candidates", Type: field.TypeInt, Nullable: true},
	}
	// CandidatesTable holds the schema information for the "candidates" table.
	CandidatesTable = &schema.Table{
		Name:       "candidates",
		Columns:    CandidatesColumns,
		PrimaryKey: []*schema.Column{CandidatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "candidates_elections_candidates",
				Columns:    []*schema.Column{CandidatesColumns[7]},
				RefColumns: []*schema.Column{ElectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "contents", Type: field.TypeString, Size: 2000},
		{Name: "comment_children", Type: field.TypeInt, Nullable: true},
		{Name: "election_comments", Type: field.TypeInt, Nullable: true},
		{Name: "user_comments", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_comments_children",
				Columns:    []*schema.Column{CommentsColumns[4]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_elections_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{ElectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ElectionsColumns holds the columns for the "elections" table.
	ElectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Size: 64},
		{Name: "description", Type: field.TypeString, Size: 1000},
		{Name: "completed", Type: field.TypeBool, Default: false},
		{Name: "user_elections", Type: field.TypeInt, Nullable: true},
	}
	// ElectionsTable holds the schema information for the "elections" table.
	ElectionsTable = &schema.Table{
		Name:       "elections",
		Columns:    ElectionsColumns,
		PrimaryKey: []*schema.Column{ElectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "elections_users_elections",
				Columns:    []*schema.Column{ElectionsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ElectionFiltersColumns holds the columns for the "election_filters" table.
	ElectionFiltersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "has_first_name", Type: field.TypeBool, Default: false},
		{Name: "has_last_name", Type: field.TypeBool, Default: false},
		{Name: "has_birthdate", Type: field.TypeBool, Default: false},
		{Name: "has_phone_number", Type: field.TypeBool, Default: false},
		{Name: "has_bio", Type: field.TypeBool, Default: false},
		{Name: "has_address", Type: field.TypeBool, Default: false},
		{Name: "has_photo_url", Type: field.TypeBool, Default: false},
		{Name: "election_filters", Type: field.TypeInt, Unique: true},
	}
	// ElectionFiltersTable holds the schema information for the "election_filters" table.
	ElectionFiltersTable = &schema.Table{
		Name:       "election_filters",
		Columns:    ElectionFiltersColumns,
		PrimaryKey: []*schema.Column{ElectionFiltersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "election_filters_elections_filters",
				Columns:    []*schema.Column{ElectionFiltersColumns[8]},
				RefColumns: []*schema.Column{ElectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ElectionSettingsColumns holds the columns for the "election_settings" table.
	ElectionSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "is_anonymous", Type: field.TypeBool, Default: false},
		{Name: "allow_comments", Type: field.TypeBool, Default: true},
		{Name: "max_votes", Type: field.TypeInt, Default: 1},
		{Name: "duration", Type: field.TypeTime},
		{Name: "election_settings", Type: field.TypeInt, Unique: true},
	}
	// ElectionSettingsTable holds the schema information for the "election_settings" table.
	ElectionSettingsTable = &schema.Table{
		Name:       "election_settings",
		Columns:    ElectionSettingsColumns,
		PrimaryKey: []*schema.Column{ElectionSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "election_settings_elections_settings",
				Columns:    []*schema.Column{ElectionSettingsColumns[7]},
				RefColumns: []*schema.Column{ElectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "birthdate", Type: field.TypeTime, Nullable: true},
		{Name: "phone_number", Type: field.TypeString, Nullable: true},
		{Name: "bio", Type: field.TypeString, Nullable: true, Size: 500},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "photo_url", Type: field.TypeString, Nullable: true},
		{Name: "user_profile", Type: field.TypeInt, Unique: true},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:       "profiles",
		Columns:    ProfilesColumns,
		PrimaryKey: []*schema.Column{ProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "profiles_users_profile",
				Columns:    []*schema.Column{ProfilesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "last_login", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "is_organizer", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VotesColumns holds the columns for the "votes" table.
	VotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "candidate_votes", Type: field.TypeInt, Nullable: true},
		{Name: "user_votes", Type: field.TypeInt, Nullable: true},
	}
	// VotesTable holds the schema information for the "votes" table.
	VotesTable = &schema.Table{
		Name:       "votes",
		Columns:    VotesColumns,
		PrimaryKey: []*schema.Column{VotesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "votes_candidates_votes",
				Columns:    []*schema.Column{VotesColumns[3]},
				RefColumns: []*schema.Column{CandidatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "votes_users_votes",
				Columns:    []*schema.Column{VotesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagElectionsColumns holds the columns for the "tag_elections" table.
	TagElectionsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeInt},
		{Name: "election_id", Type: field.TypeInt},
	}
	// TagElectionsTable holds the schema information for the "tag_elections" table.
	TagElectionsTable = &schema.Table{
		Name:       "tag_elections",
		Columns:    TagElectionsColumns,
		PrimaryKey: []*schema.Column{TagElectionsColumns[0], TagElectionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_elections_tag_id",
				Columns:    []*schema.Column{TagElectionsColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_elections_election_id",
				Columns:    []*schema.Column{TagElectionsColumns[1]},
				RefColumns: []*schema.Column{ElectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CandidatesTable,
		CommentsTable,
		ElectionsTable,
		ElectionFiltersTable,
		ElectionSettingsTable,
		ProfilesTable,
		TagsTable,
		UsersTable,
		VotesTable,
		TagElectionsTable,
	}
)

func init() {
	CandidatesTable.ForeignKeys[0].RefTable = ElectionsTable
	CommentsTable.ForeignKeys[0].RefTable = CommentsTable
	CommentsTable.ForeignKeys[1].RefTable = ElectionsTable
	CommentsTable.ForeignKeys[2].RefTable = UsersTable
	ElectionsTable.ForeignKeys[0].RefTable = UsersTable
	ElectionFiltersTable.ForeignKeys[0].RefTable = ElectionsTable
	ElectionSettingsTable.ForeignKeys[0].RefTable = ElectionsTable
	ProfilesTable.ForeignKeys[0].RefTable = UsersTable
	VotesTable.ForeignKeys[0].RefTable = CandidatesTable
	VotesTable.ForeignKeys[1].RefTable = UsersTable
	TagElectionsTable.ForeignKeys[0].RefTable = TagsTable
	TagElectionsTable.ForeignKeys[1].RefTable = ElectionsTable
}
