package types

// PostGreeting is the body of a post request for a new greeting.
type PostGreeting struct {
	Greeting string `json:"greeting"`
	Language string `json:"language"`
	IsFormal bool   `json:"isFormal"`
	NumWords int    `json:"numWords"`
}

// GreetingRecord is a greeting retrieved from the database.
type GreetingRecord struct {
	ID int `json:"id"`
	PostGreeting
}

// CreatedResponse is the response given when a new item is created successfully.
type CreatedResponse struct {
	ID int `json:"id"`
}

// Card is a struct to represent a Set® card
type Card struct {
	Shape   string `json:"shape"`
	Number  int    `json:"number"`
	Color   string `color:"color"`
	Pattern string `pattern:"pattern"`
}

// CardRecord is a record retrieved from the database.
type CardRecord struct {
	ID int `json:"id"`
	Card
}

// CardSet set is a struct to represent three Set® cards
type CardSet struct {
	Cards [3]Card
}

// IsValidSet is a method to determine whether three Set® cards
// comprise a valid set
func (set *CardSet) IsValidSet() bool {
	return true
}
