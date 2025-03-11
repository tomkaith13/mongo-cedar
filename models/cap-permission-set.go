package models

type Permission string

const (
	Create Permission = "CREATE"
	Read   Permission = "READ"
	Update Permission = "UPDATE"
	Delete Permission = "DELETE"
	All    Permission = "ALL"
)

type CapabilityPermissionSet struct {
	ID          string       `bson:"_id"`
	Name        string       `bson:"name"`
	Description string       `bson:"description"`
	Permissions []Permission `bson:"permissions"`
}
