// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Teeth/app/ent/foodmenu"
	"github.com/Teeth/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Foodmenu is the model entity for the Foodmenu schema.
type Foodmenu struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FoodmenuName holds the value of the "foodmenu_name" field.
	FoodmenuName string `json:"foodmenu_name,omitempty"`
	// FoodmenuType holds the value of the "foodmenu_type" field.
	FoodmenuType string `json:"foodmenu_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FoodmenuQuery when eager-loading is set.
	Edges    FoodmenuEdges `json:"edges"`
	owner_id *int
}

// FoodmenuEdges holds the relations/edges for other nodes in the graph.
type FoodmenuEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Eatinghistory holds the value of the eatinghistory edge.
	Eatinghistory []*Eatinghistory
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FoodmenuEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// EatinghistoryOrErr returns the Eatinghistory value or an error if the edge
// was not loaded in eager-loading.
func (e FoodmenuEdges) EatinghistoryOrErr() ([]*Eatinghistory, error) {
	if e.loadedTypes[1] {
		return e.Eatinghistory, nil
	}
	return nil, &NotLoadedError{edge: "eatinghistory"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Foodmenu) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // foodmenu_name
		&sql.NullString{}, // foodmenu_type
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Foodmenu) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // owner_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Foodmenu fields.
func (f *Foodmenu) assignValues(values ...interface{}) error {
	if m, n := len(values), len(foodmenu.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	f.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field foodmenu_name", values[0])
	} else if value.Valid {
		f.FoodmenuName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field foodmenu_type", values[1])
	} else if value.Valid {
		f.FoodmenuType = value.String
	}
	values = values[2:]
	if len(values) == len(foodmenu.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
		} else if value.Valid {
			f.owner_id = new(int)
			*f.owner_id = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Foodmenu.
func (f *Foodmenu) QueryOwner() *UserQuery {
	return (&FoodmenuClient{config: f.config}).QueryOwner(f)
}

// QueryEatinghistory queries the eatinghistory edge of the Foodmenu.
func (f *Foodmenu) QueryEatinghistory() *EatinghistoryQuery {
	return (&FoodmenuClient{config: f.config}).QueryEatinghistory(f)
}

// Update returns a builder for updating this Foodmenu.
// Note that, you need to call Foodmenu.Unwrap() before calling this method, if this Foodmenu
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Foodmenu) Update() *FoodmenuUpdateOne {
	return (&FoodmenuClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *Foodmenu) Unwrap() *Foodmenu {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Foodmenu is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Foodmenu) String() string {
	var builder strings.Builder
	builder.WriteString("Foodmenu(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", foodmenu_name=")
	builder.WriteString(f.FoodmenuName)
	builder.WriteString(", foodmenu_type=")
	builder.WriteString(f.FoodmenuType)
	builder.WriteByte(')')
	return builder.String()
}

// Foodmenus is a parsable slice of Foodmenu.
type Foodmenus []*Foodmenu

func (f Foodmenus) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
