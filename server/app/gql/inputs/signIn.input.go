package gqlInputs

type SignInInput struct {
	Input SignInInputForm
}

type SignInInputForm struct {
	Phone    *string
	Password *string
}
