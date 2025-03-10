// Code generated by ent, DO NOT EDIT.

package profile

import (
	"time"
	"voting-system/ent/generated/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldID, id))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldLastName, v))
}

// Birthdate applies equality check predicate on the "birthdate" field. It's identical to BirthdateEQ.
func Birthdate(v time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldBirthdate, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldPhoneNumber, v))
}

// Bio applies equality check predicate on the "bio" field. It's identical to BioEQ.
func Bio(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldBio, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldAddress, v))
}

// PhotoURL applies equality check predicate on the "photo_url" field. It's identical to PhotoURLEQ.
func PhotoURL(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldPhotoURL, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameIsNil applies the IsNil predicate on the "first_name" field.
func FirstNameIsNil() predicate.Profile {
	return predicate.Profile(sql.FieldIsNull(FieldFirstName))
}

// FirstNameNotNil applies the NotNil predicate on the "first_name" field.
func FirstNameNotNil() predicate.Profile {
	return predicate.Profile(sql.FieldNotNull(FieldFirstName))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameIsNil applies the IsNil predicate on the "last_name" field.
func LastNameIsNil() predicate.Profile {
	return predicate.Profile(sql.FieldIsNull(FieldLastName))
}

// LastNameNotNil applies the NotNil predicate on the "last_name" field.
func LastNameNotNil() predicate.Profile {
	return predicate.Profile(sql.FieldNotNull(FieldLastName))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldLastName, v))
}

// BirthdateEQ applies the EQ predicate on the "birthdate" field.
func BirthdateEQ(v time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldBirthdate, v))
}

// BirthdateNEQ applies the NEQ predicate on the "birthdate" field.
func BirthdateNEQ(v time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldBirthdate, v))
}

// BirthdateIn applies the In predicate on the "birthdate" field.
func BirthdateIn(vs ...time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldBirthdate, vs...))
}

// BirthdateNotIn applies the NotIn predicate on the "birthdate" field.
func BirthdateNotIn(vs ...time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldBirthdate, vs...))
}

// BirthdateGT applies the GT predicate on the "birthdate" field.
func BirthdateGT(v time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldBirthdate, v))
}

// BirthdateGTE applies the GTE predicate on the "birthdate" field.
func BirthdateGTE(v time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldBirthdate, v))
}

// BirthdateLT applies the LT predicate on the "birthdate" field.
func BirthdateLT(v time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldBirthdate, v))
}

// BirthdateLTE applies the LTE predicate on the "birthdate" field.
func BirthdateLTE(v time.Time) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldBirthdate, v))
}

// BirthdateIsNil applies the IsNil predicate on the "birthdate" field.
func BirthdateIsNil() predicate.Profile {
	return predicate.Profile(sql.FieldIsNull(FieldBirthdate))
}

// BirthdateNotNil applies the NotNil predicate on the "birthdate" field.
func BirthdateNotNil() predicate.Profile {
	return predicate.Profile(sql.FieldNotNull(FieldBirthdate))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberIsNil applies the IsNil predicate on the "phone_number" field.
func PhoneNumberIsNil() predicate.Profile {
	return predicate.Profile(sql.FieldIsNull(FieldPhoneNumber))
}

// PhoneNumberNotNil applies the NotNil predicate on the "phone_number" field.
func PhoneNumberNotNil() predicate.Profile {
	return predicate.Profile(sql.FieldNotNull(FieldPhoneNumber))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// BioEQ applies the EQ predicate on the "bio" field.
func BioEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldBio, v))
}

// BioNEQ applies the NEQ predicate on the "bio" field.
func BioNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldBio, v))
}

// BioIn applies the In predicate on the "bio" field.
func BioIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldBio, vs...))
}

// BioNotIn applies the NotIn predicate on the "bio" field.
func BioNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldBio, vs...))
}

// BioGT applies the GT predicate on the "bio" field.
func BioGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldBio, v))
}

// BioGTE applies the GTE predicate on the "bio" field.
func BioGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldBio, v))
}

// BioLT applies the LT predicate on the "bio" field.
func BioLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldBio, v))
}

// BioLTE applies the LTE predicate on the "bio" field.
func BioLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldBio, v))
}

// BioContains applies the Contains predicate on the "bio" field.
func BioContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldBio, v))
}

// BioHasPrefix applies the HasPrefix predicate on the "bio" field.
func BioHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldBio, v))
}

// BioHasSuffix applies the HasSuffix predicate on the "bio" field.
func BioHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldBio, v))
}

// BioIsNil applies the IsNil predicate on the "bio" field.
func BioIsNil() predicate.Profile {
	return predicate.Profile(sql.FieldIsNull(FieldBio))
}

// BioNotNil applies the NotNil predicate on the "bio" field.
func BioNotNil() predicate.Profile {
	return predicate.Profile(sql.FieldNotNull(FieldBio))
}

// BioEqualFold applies the EqualFold predicate on the "bio" field.
func BioEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldBio, v))
}

// BioContainsFold applies the ContainsFold predicate on the "bio" field.
func BioContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldBio, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.Profile {
	return predicate.Profile(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.Profile {
	return predicate.Profile(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldAddress, v))
}

// PhotoURLEQ applies the EQ predicate on the "photo_url" field.
func PhotoURLEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldPhotoURL, v))
}

// PhotoURLNEQ applies the NEQ predicate on the "photo_url" field.
func PhotoURLNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldPhotoURL, v))
}

// PhotoURLIn applies the In predicate on the "photo_url" field.
func PhotoURLIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldPhotoURL, vs...))
}

// PhotoURLNotIn applies the NotIn predicate on the "photo_url" field.
func PhotoURLNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldPhotoURL, vs...))
}

// PhotoURLGT applies the GT predicate on the "photo_url" field.
func PhotoURLGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldPhotoURL, v))
}

// PhotoURLGTE applies the GTE predicate on the "photo_url" field.
func PhotoURLGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldPhotoURL, v))
}

// PhotoURLLT applies the LT predicate on the "photo_url" field.
func PhotoURLLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldPhotoURL, v))
}

// PhotoURLLTE applies the LTE predicate on the "photo_url" field.
func PhotoURLLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldPhotoURL, v))
}

// PhotoURLContains applies the Contains predicate on the "photo_url" field.
func PhotoURLContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldPhotoURL, v))
}

// PhotoURLHasPrefix applies the HasPrefix predicate on the "photo_url" field.
func PhotoURLHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldPhotoURL, v))
}

// PhotoURLHasSuffix applies the HasSuffix predicate on the "photo_url" field.
func PhotoURLHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldPhotoURL, v))
}

// PhotoURLIsNil applies the IsNil predicate on the "photo_url" field.
func PhotoURLIsNil() predicate.Profile {
	return predicate.Profile(sql.FieldIsNull(FieldPhotoURL))
}

// PhotoURLNotNil applies the NotNil predicate on the "photo_url" field.
func PhotoURLNotNil() predicate.Profile {
	return predicate.Profile(sql.FieldNotNull(FieldPhotoURL))
}

// PhotoURLEqualFold applies the EqualFold predicate on the "photo_url" field.
func PhotoURLEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldPhotoURL, v))
}

// PhotoURLContainsFold applies the ContainsFold predicate on the "photo_url" field.
func PhotoURLContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldPhotoURL, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Profile) predicate.Profile {
	return predicate.Profile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Profile) predicate.Profile {
	return predicate.Profile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Profile) predicate.Profile {
	return predicate.Profile(sql.NotPredicates(p))
}
