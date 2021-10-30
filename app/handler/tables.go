package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/stanyx/catalog/internal/storage"
)

type ColumnParams struct {
	Name       string
	Type       string
	PrimaryKey bool
}

type CreateTableParams struct {
	Name    string
	Columns []ColumnParams
}

func (h *Handler) CreateTable(w http.ResponseWriter, r *http.Request) {

	var params CreateTableParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	db := storage.GetDB()

	var columns []string
	for _, column := range params.Columns {
		columnExpr := fmt.Sprintf("%s %s", column.Name, column.Type)
		if column.PrimaryKey {
			columnExpr += " PRIMARY KEY"
		}
		columns = append(columns, columnExpr)
	}

	createTableQuery := fmt.Sprintf(`
	create table %s (
		%s
	)`, params.Name, strings.Join(columns, ",\n"))

	if _, err := db.Exec(createTableQuery); err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	fmt.Fprint(w, "ok")
}
