package migration

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func migrate(db *sql.DB, tableName string, model any) {
	val := reflect.TypeOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		log.Fatalf("Model must be a struct or pointer to struct")
	}

	columns := map[string]string{}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		colName := field.Tag.Get("db")
		colType := field.Tag.Get("sql")
		if colName != "" && colType != "" {
			columns[colName] = colType
		}
	}

	if tableExists(db, tableName) {
		existingCols := getExistingColumns(db, tableName)

		for col, colType := range columns {
			if _, found := existingCols[col]; !found {
				query := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", tableName, col, colType)
				_, err := db.Exec(query)
				if err != nil {
					log.Printf("Warning: Failed to add column %s: %v", col, err)
				} else {
					fmt.Printf("Column added: %s %s\n", col, colType)
				}
			}
		}

		for existingCol := range existingCols {
			if _, found := columns[existingCol]; !found {
				if existingCol == "created_at" || existingCol == "updated_at" || existingCol == "deleted_at" {
					continue
				}
				query := fmt.Sprintf("ALTER TABLE %s DROP COLUMN IF EXISTS %s;", tableName, existingCol)
				_, err := db.Exec(query)
				if err != nil {
					log.Printf("Warning: Failed to drop column %s: %v", existingCol, err)
				} else {
					fmt.Printf("Column dropped: %s\n", existingCol)
				}
			}
		}

		fmt.Printf("AutoMigration completed for existing table: %s\n", tableName)
	} else {
		var fieldDefs []string
		for col, colType := range columns {
			fieldDefs = append(fieldDefs, fmt.Sprintf("%s %s", col, colType))
		}
		query := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(fieldDefs, ", "))
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
		fmt.Printf("Table created: %s\n", tableName)
	}
}

func tableExists(db *sql.DB, tableName string) bool {
	query := `SELECT EXISTS (
		SELECT 1 FROM information_schema.tables 
		WHERE table_schema = 'public' AND table_name = $1
	);`
	var exists bool
	err := db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check table existence: %v", err)
	}
	return exists
}

func getExistingColumns(db *sql.DB, tableName string) map[string]bool {
	query := `
		SELECT column_name, data_type, is_nullable, column_default
		FROM information_schema.columns 
		WHERE table_schema = 'public' AND table_name = $1;
	`
	rows, err := db.Query(query, tableName)
	if err != nil {
		log.Fatalf("Failed to fetch existing columns: %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	columns := make(map[string]bool)
	for rows.Next() {
		var colName, dataType, isNullable string
		var columnDefault sql.NullString
		if err := rows.Scan(&colName, &dataType, &isNullable, &columnDefault); err != nil {
			log.Fatal(err)
		}
		columns[colName] = true

		// Log column details for debugging
		fmt.Printf("Existing column: %s (type: %s, nullable: %s)\n", colName, dataType, isNullable)
	}
	return columns
}
