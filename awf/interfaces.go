package awf

type Forecast struct {
	Date        string  `json:"date"`
	Temperature float32 `json:"temperature"`
	Condition   string  `json:"Condition"`
	Settlement  string  `json:"settlement"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Alert struct {
	Event       string `json:"event"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

type Preferences struct {
	Location string `json:"location"`
	Unit     string `json:"unit"`
	Language string `json:"language"`
}
