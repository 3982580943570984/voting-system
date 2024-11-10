// Code generated by ent, DO NOT EDIT.

package ent

import (
	"voting-system/ent/election"
	"voting-system/ent/profile"
	"voting-system/ent/schema"
	"voting-system/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	electionFields := schema.Election{}.Fields()
	_ = electionFields
	// electionDescTitle is the schema descriptor for title field.
	electionDescTitle := electionFields[0].Descriptor()
	// election.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	election.TitleValidator = func() func(string) error {
		validators := electionDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// electionDescIsActive is the schema descriptor for is_active field.
	electionDescIsActive := electionFields[4].Descriptor()
	// election.DefaultIsActive holds the default value on creation for the is_active field.
	election.DefaultIsActive = electionDescIsActive.Default.(bool)
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescFirstName is the schema descriptor for first_name field.
	profileDescFirstName := profileFields[0].Descriptor()
	// profile.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	profile.FirstNameValidator = func() func(string) error {
		validators := profileDescFirstName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(first_name string) error {
			for _, fn := range fns {
				if err := fn(first_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// profileDescLastName is the schema descriptor for last_name field.
	profileDescLastName := profileFields[1].Descriptor()
	// profile.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	profile.LastNameValidator = func() func(string) error {
		validators := profileDescLastName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(last_name string) error {
			for _, fn := range fns {
				if err := fn(last_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescIsActive is the schema descriptor for is_active field.
	userDescIsActive := userFields[2].Descriptor()
	// user.DefaultIsActive holds the default value on creation for the is_active field.
	user.DefaultIsActive = userDescIsActive.Default.(bool)
}
