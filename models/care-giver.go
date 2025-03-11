package models

type CareGiverModel struct {
	ID        string `bson:"_id"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`

	// CRs
	CareReceipentIds map[string]bool `bson:"crIds"`
}
