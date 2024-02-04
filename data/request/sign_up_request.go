package request

type UserSignUpRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Password string `validate:"required,min=6" json:"password"`
}
