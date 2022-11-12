package request

type StudentPatchRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
