package gqlInputs

type AddMemberToFamilyInput struct {
	Input AddMemberToFamilyInputForm
}

type AddMemberToFamilyInputForm struct {
	Phone    *string
	Name     *string
	Password *string
}
