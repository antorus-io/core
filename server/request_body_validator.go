package server

type Validateable interface {
	Validate() []error
}

func ValidateRequestBody(data Validateable) []error {
	validationErrors := data.Validate()

	return validationErrors
}
