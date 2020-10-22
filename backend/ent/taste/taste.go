// Code generated by entc, DO NOT EDIT.

package taste

const (
	// Label holds the string label denoting the taste type in the database.
	Label = "taste"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTasteName holds the string denoting the taste_name field in the database.
	FieldTasteName = "taste_name"

	// EdgeEatinghistory holds the string denoting the eatinghistory edge name in mutations.
	EdgeEatinghistory = "eatinghistory"

	// Table holds the table name of the taste in the database.
	Table = "tastes"
	// EatinghistoryTable is the table the holds the eatinghistory relation/edge.
	EatinghistoryTable = "eatinghistories"
	// EatinghistoryInverseTable is the table name for the Eatinghistory entity.
	// It exists in this package in order to avoid circular dependency with the "eatinghistory" package.
	EatinghistoryInverseTable = "eatinghistories"
	// EatinghistoryColumn is the table column denoting the eatinghistory relation/edge.
	EatinghistoryColumn = "taste_id"
)

// Columns holds all SQL columns for taste fields.
var Columns = []string{
	FieldID,
	FieldTasteName,
}

var (
	// TasteNameValidator is a validator for the "taste_name" field. It is called by the builders before save.
	TasteNameValidator func(string) error
)
