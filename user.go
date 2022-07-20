package cheevos

// User represents a single user of the system. A user can be a member of many
// organizations, but they are all tied to the same underlying account.
type User struct {
	// A unique identifier for the user.
	ID string

	// The username for display.
	Username string
}
