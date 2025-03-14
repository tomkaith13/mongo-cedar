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

	// PII store for guests; non-guests PII can be looked up using PUID
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`

	// CRs set
	CareReceipentIds map[string]bool `bson:"crIds"`

	//inviteStatusMap per CR
	CareReceipentInviteMap map[string]InviteStatus `bson:"crInviteStatus"`
}
