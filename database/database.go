package database

import (
	"errors"

	"github.com/nolwn/important-server/types"
)

type table struct {
	FieldCounter int
	Table        []interface{}
}

// TODO: Make table arrays specific to their resource
var greetingTable = table{FieldCounter: 0, Table: make([]interface{}, 0, 10)}
var cardTable table = table{FieldCounter: 0, Table: make([]interface{}, 0, 10)}

func makeNewGreeting(pg *types.PostGreeting, gr *types.GreetingRecord) {
	gr.ID = greetingTable.FieldCounter
	gr.Greeting = pg.Greeting
	gr.Language = pg.Language
	gr.IsFormal = pg.IsFormal
	gr.NumWords = pg.NumWords

	greetingTable.FieldCounter++
}

func initCardRecord(oldCard *types.Card, newCard *types.CardRecord) {
	newCard.ID = cardTable.FieldCounter
	newCard.Shape = oldCard.Shape
	newCard.Number = oldCard.Number
	newCard.Color = oldCard.Color
	newCard.Pattern = oldCard.Pattern

	cardTable.FieldCounter++
}

// Add adds a new greeting object to the database.
func Add(resource interface{}) (types.CreatedResponse, error) {

	switch resource.(type) {
	case types.PostGreeting:
		var newRecord types.GreetingRecord
		value, ok := resource.(types.PostGreeting)

		if !ok {
			return types.CreatedResponse{}, errors.New("something broke when trying to convert type")
		}

		makeNewGreeting(&value, &newRecord)
		greetingTable.Table = append(greetingTable.Table, newRecord)
		return types.CreatedResponse{ID: newRecord.ID}, nil

	case types.Card:
		var newRecord types.CardRecord
		value, ok := resource.(types.Card)

		if !ok {
			return types.CreatedResponse{}, errors.New("something broke when trying to convert type")
		}

		initCardRecord(&value, &newRecord)
		cardTable.Table = append(cardTable.Table, newRecord)
		return types.CreatedResponse{ID: newRecord.ID}, nil

	default:
		return types.CreatedResponse{}, errors.New("could not find a corresponding type")
	}
}

// Get returns all greetings in the database.
func Get(i interface{}) (interface{}, error) {
	switch i.(type) {
	case []types.GreetingRecord:
		i = greetingTable
		return i, nil

	case []types.CardRecord:
		i = cardTable.Table
		return i, nil

	default:
		return nil, errors.New("could not find a corresponding type")
	}
}
