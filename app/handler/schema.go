package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stanyx/catalog/internal/storage"
)

type SchemaColumnParams struct {
	Name string
	Type string
}

type CreateSchemaParams struct {
	Name    string
	Version string
	Columns []SchemaColumnParams
}

func (h *Handler) CreateSchema(w http.ResponseWriter, r *http.Request) {
	var params CreateSchemaParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	db := storage.GetDB()

	schemaExpr, _ := json.Marshal(params.Columns)

	createSchemaQuery := fmt.Sprintf(`
	insert into _schema (name, schema, version) values ($1, $2, $3)
	`)

	if _, err := db.Exec(createSchemaQuery, params.Name, schemaExpr, params.Version); err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	fmt.Fprint(w, "ok")
}
