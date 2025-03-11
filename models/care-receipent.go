package models

type CareReceipentModel struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`

	AuthorizedCareGiverIds map[string]bool `bson:"authorizedCareGiverIds"`
	// map[caregiverId][capabilityId]PermSet
	CapabilityPermissionMap map[string]map[string]CapabilityPermissionSet `bson:"capabilityPermissionMap"`
}
