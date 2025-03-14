package models

type InviteStatus string

const (
	Accepted InviteStatus = "ACCEPTED"
	Pending  InviteStatus = "PENDING"
	Revoked  InviteStatus = "REVOKED"
)

type CareGiverModel struct {
	ID string `bson:"_id"`

	PlatformUID string `bson:"pid"`

	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`

	// CRs set
	CareReceipentIds map[string]bool `bson:"crIds"`

	//inviteStatusMap per map[CrID]
	CareReceipentInviteMap map[string]InviteStatus `bson:"crInviteStatus"`

	// Care givers can have their own resources when not impersonating
	AllowedResourceIds map[string]bool `bson:"allowed_resources"`
}
