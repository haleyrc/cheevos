package cheevos

// Organization represents a group of users belonging to a related parent
// entity. This may be an actual organization or simply a group of friends who
// want to recognize each other for significant events. An organization also
// acts as a boundary for managing cheevos, since every cheevo is "owned" by an
// organization and can only be granted to members of that organization.
type Organization struct {
	// A unique identifier for the organization.
	ID string

	// The name of the organization.
	Name string

	// The owner of the organization.
	Owner string
}
