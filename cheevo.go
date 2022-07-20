package cheevos

// Cheevo represents an achievement that can be awarded to a user by authorized
// members of an organization.
type Cheevo struct {
	// A unique identifer for the cheevo.
	ID string

	// The short name for the cheevo.
	Name string

	// A description of the act that the achievement is recognizing.
	Description string

	// The parent Organization that owns the Cheevo.
	Organization string
}
