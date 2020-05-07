package database

import "github.com/nolwn/important-server/types"

var fieldCounter = 0

// MakeNewRecord is current index of the database.
func MakeNewRecord(pg *types.PostGreeting, gr *types.GreetingRecord) {
	gr.ID = fieldCounter
	gr.Greeting = pg.Greeting
	gr.Language = pg.Language
	gr.IsFormal = pg.IsFormal
	gr.NumWords = pg.NumWords

	fieldCounter++
}

// AddGreeting adds a new greeting object to the database.
func AddGreeting(gt types.PostGreeting) types.GreetingRecord {
	var newRecord types.GreetingRecord

	MakeNewRecord(&gt, &newRecord)
	Database = append(Database, newRecord)

	return newRecord
}

// GetGreetings returns all greetings in the database.
func GetGreetings() []types.GreetingRecord {
	return Database
}

// Database remove me.
var Database = make([]types.GreetingRecord, 0, 10)
