package awf

type Forecast struct {
	Date        string  `json:"date"`
	Temperature float32 `json:"temperature"`
	Condition   string  `json:"Condition"`
	Settlement  string  `json:"settlement"`
}

type Alert struct {
	Event       string `json:"event"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}
