package models

type CareReceipentModel struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`

	AuthorizedCareGiverIds  map[string]bool                    `bson:"authorizedCareGiverIds"`
	CapabilityPermissionMap map[string]CapabilityPermissionSet `bson:"capabilityPermissionMap"`
}
