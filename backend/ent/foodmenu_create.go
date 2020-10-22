// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Teeth/app/ent/eatinghistory"
	"github.com/Teeth/app/ent/foodmenu"
	"github.com/Teeth/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// FoodmenuCreate is the builder for creating a Foodmenu entity.
type FoodmenuCreate struct {
	config
	mutation *FoodmenuMutation
	hooks    []Hook
}

// SetFoodmenuName sets the foodmenu_name field.
func (fc *FoodmenuCreate) SetFoodmenuName(s string) *FoodmenuCreate {
	fc.mutation.SetFoodmenuName(s)
	return fc
}

// SetFoodmenuType sets the foodmenu_type field.
func (fc *FoodmenuCreate) SetFoodmenuType(s string) *FoodmenuCreate {
	fc.mutation.SetFoodmenuType(s)
	return fc
}

// SetOwnerID sets the owner edge to User by id.
func (fc *FoodmenuCreate) SetOwnerID(id int) *FoodmenuCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fc *FoodmenuCreate) SetNillableOwnerID(id *int) *FoodmenuCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the owner edge to User.
func (fc *FoodmenuCreate) SetOwner(u *User) *FoodmenuCreate {
	return fc.SetOwnerID(u.ID)
}

// AddEatinghistoryIDs adds the eatinghistory edge to Eatinghistory by ids.
func (fc *FoodmenuCreate) AddEatinghistoryIDs(ids ...int) *FoodmenuCreate {
	fc.mutation.AddEatinghistoryIDs(ids...)
	return fc
}

// AddEatinghistory adds the eatinghistory edges to Eatinghistory.
func (fc *FoodmenuCreate) AddEatinghistory(e ...*Eatinghistory) *FoodmenuCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fc.AddEatinghistoryIDs(ids...)
}

// Mutation returns the FoodmenuMutation object of the builder.
func (fc *FoodmenuCreate) Mutation() *FoodmenuMutation {
	return fc.mutation
}

// Save creates the Foodmenu in the database.
func (fc *FoodmenuCreate) Save(ctx context.Context) (*Foodmenu, error) {
	if _, ok := fc.mutation.FoodmenuName(); !ok {
		return nil, &ValidationError{Name: "foodmenu_name", err: errors.New("ent: missing required field \"foodmenu_name\"")}
	}
	if v, ok := fc.mutation.FoodmenuName(); ok {
		if err := foodmenu.FoodmenuNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "foodmenu_name", err: fmt.Errorf("ent: validator failed for field \"foodmenu_name\": %w", err)}
		}
	}
	if _, ok := fc.mutation.FoodmenuType(); !ok {
		return nil, &ValidationError{Name: "foodmenu_type", err: errors.New("ent: missing required field \"foodmenu_type\"")}
	}
	if v, ok := fc.mutation.FoodmenuType(); ok {
		if err := foodmenu.FoodmenuTypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "foodmenu_type", err: fmt.Errorf("ent: validator failed for field \"foodmenu_type\": %w", err)}
		}
	}
	var (
		err  error
		node *Foodmenu
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FoodmenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FoodmenuCreate) SaveX(ctx context.Context) *Foodmenu {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FoodmenuCreate) sqlSave(ctx context.Context) (*Foodmenu, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FoodmenuCreate) createSpec() (*Foodmenu, *sqlgraph.CreateSpec) {
	var (
		f     = &Foodmenu{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: foodmenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: foodmenu.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.FoodmenuName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: foodmenu.FieldFoodmenuName,
		})
		f.FoodmenuName = value
	}
	if value, ok := fc.mutation.FoodmenuType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: foodmenu.FieldFoodmenuType,
		})
		f.FoodmenuType = value
	}
	if nodes := fc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodmenu.OwnerTable,
			Columns: []string{foodmenu.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.EatinghistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   foodmenu.EatinghistoryTable,
			Columns: []string{foodmenu.EatinghistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eatinghistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return f, _spec
}
