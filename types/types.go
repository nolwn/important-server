package types

var fieldCounter = 0

// PostGreeting is the body of a post request for a new greeting.
type PostGreeting struct {
	Greeting string `json:"greeting"`
	Language string `json:"language"`
	IsFormal bool   `json:"isFormal"`
	NumWords int    `json:"numWords"`
}

// GreetingRecord is a greeting retrieved from the database.
type GreetingRecord struct {
	ID       int    `json:"id"`
	Greeting string `json:"greeting"`
	Language string `json:"language"`
	IsFormal bool   `json:"isFormal"`
	NumWords int    `json:"numWords"`
}

// CreatedResponse is the response given when a new item is created successfully.
type CreatedResponse struct {
	ID int `json:"id"`
}
