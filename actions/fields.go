package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/mataram/genos/models"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Field)
// DB Table: Plural (fields)
// Resource: Plural (Fields)
// Path: Plural (/fields)
// View Template Folder: Plural (/templates/fields/)

// FieldsResource is the resource for the Field model
type FieldsResource struct {
	buffalo.Resource
}

// List gets all Fields. This function is mapped to the path
// GET /fields
func (v FieldsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	fields := &models.Fields{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Fields from the DB
	if err := q.All(fields); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, fields))
}

// Show gets the data for one Field. This function is mapped to
// the path GET /fields/{field_id}
func (v FieldsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Field
	field := &models.Field{}

	// To find the Field the parameter field_id is used.
	if err := tx.Find(field, c.Param("field_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, field))
}

// New renders the form for creating a new Field.
// This function is mapped to the path GET /fields/new
func (v FieldsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Field{}))
}

// Create adds a Field to the DB. This function is mapped to the
// path POST /fields
func (v FieldsResource) Create(c buffalo.Context) error {
	// Allocate an empty Field
	field := &models.Field{}

	// Bind field to the html form elements
	if err := c.Bind(field); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(field)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, field))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Field was created successfully")

	// and redirect to the fields index page
	return c.Render(201, r.Auto(c, field))
}

// Edit renders a edit form for a Field. This function is
// mapped to the path GET /fields/{field_id}/edit
func (v FieldsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Field
	field := &models.Field{}

	if err := tx.Find(field, c.Param("field_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, field))
}

// Update changes a Field in the DB. This function is mapped to
// the path PUT /fields/{field_id}
func (v FieldsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Field
	field := &models.Field{}

	if err := tx.Find(field, c.Param("field_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Field to the html form elements
	if err := c.Bind(field); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(field)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, field))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Field was updated successfully")

	// and redirect to the fields index page
	return c.Render(200, r.Auto(c, field))
}

// Destroy deletes a Field from the DB. This function is mapped
// to the path DELETE /fields/{field_id}
func (v FieldsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Field
	field := &models.Field{}

	// To find the Field the parameter field_id is used.
	if err := tx.Find(field, c.Param("field_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(field); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Field was destroyed successfully")

	// Redirect to the fields index page
	return c.Render(200, r.Auto(c, field))
}