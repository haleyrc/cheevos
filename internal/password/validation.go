package password

type Validation func(p Password) error

var defaultValidator = &Validator{}

func AddValidation(v Validation) {
	defaultValidator.Validations = append(defaultValidator.Validations, v)
}

func Validate(p Password) error {
	return defaultValidator.Validate(p)
}

func NewValidator(validations ...Validation) Validator {
	return Validator{Validations: validations}
}

type Validator struct {
	Validations []Validation
}

func (v *Validator) Validate(p Password) error {
	errors := []error{}
	for _, validator := range v.Validations {
		if err := validator(p); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errors: errors}
	}
	return nil
}
