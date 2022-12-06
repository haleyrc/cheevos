package password

import "fmt"

func ValidateLength(min, max uint) Validation {
	return func(p Password) error {
		if len(p.s) < int(min) {
			return fmt.Errorf("password must be between %d and %d characters", min, max)
		}
		if len(p.s) > int(max) {
			return fmt.Errorf("password must be between %d and %d characters", min, max)
		}
		return nil
	}
}
