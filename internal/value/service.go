package value

type GetUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int8   `json:"age"`
}
