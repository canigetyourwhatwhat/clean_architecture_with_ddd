package request

type CreateUser struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
}
