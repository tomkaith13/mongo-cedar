package models

type Permission int

const (
	Create Permission = iota + 1
	Read
	Update
	Delete
	All
)

type CapabilityPermissionSet struct {
	ID          string       `bson:"_id"`
	Name        string       `bson:"name"`
	Description string       `bson:"description"`
	Permissions []Permission `bson:"permissions"`
}
