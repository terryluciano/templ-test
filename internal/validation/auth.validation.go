package validation

type SignupSchema struct {
	Email     string `json:"email" validate:"required,email,max=255"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
	FirstName string `json:"fname" validate:"max=255"`
	LastName  string `json:"lname" validate:"max=255"`
}
