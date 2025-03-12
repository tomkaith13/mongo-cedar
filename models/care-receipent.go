package models

type CareReceipentModel struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`

	AuthorizedCareGiverIds map[string]bool `bson:"authorizedCareGiverIds"`

	// CareGiverCapabilityPermissionMap Schema look like this:
	// We want to assign each care-receipent to assign perms for each caregiver in their system for each capability.
	// aka map[caregiverId][capabilityId]PermSet
	CareGiverCapabilityPermissionMap map[string]map[string]CapabilityPermissionSet `bson:"capabilityPermissionMap"`
}
