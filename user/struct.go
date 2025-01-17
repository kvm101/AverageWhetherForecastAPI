package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Preferences struct {
	Location string `json:"location"`
	Unit     string `json:"unit"`
	Language string `json:"language"`
}
