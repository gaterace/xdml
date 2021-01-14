// Copyright 2017-2021 Demian Harvill
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/gaterace/xdml/pkg/compiler"
	"github.com/gaterace/xdml/pkg/dml"
)

type GenTable struct {
	TableName     string
	DatabaseName  string
	Documentation []string
	Columns       []*GenColumn
	Indices       []*GenIndex
}

type GenColumn struct {
	ColumnName     string
	ColumnType     string
	ColumnModifier string
	IsNullable     bool
	Documentation  []string
	HasComma       bool
}

type GenIndex struct {
	IndexName string
	TableName string
	IsPrimary bool
	IsUnique  bool
	FieldList string
	HasComma  bool
}

func NewGenTable() *GenTable {
	s := GenTable{}
	s.Documentation = make([]string, 0)
	s.Columns = make([]*GenColumn, 0)
	s.Indices = make([]*GenIndex, 0)

	return &s
}

func NewGenColumn() *GenColumn {
	s := GenColumn{}
	s.Documentation = make([]string, 0)
	s.HasComma = true
	return &s
}

func NewGenIndex() *GenIndex {
	s := GenIndex{}
	s.HasComma = true

	return &s
}

func DbNameAndTypeDml(dmlField *dml.DmlField, dbengine string) (string, string, string) {
	fieldName := dmlField.GetFieldName()
	fieldType := dmlField.GetFieldType()
	fieldTypeQual := int(dmlField.GetFieldTypeQualifier())

	if dmlField.GetFieldTypePackage() != "builtin" {
		enumHelper := compiler.GetEnumInPackage(dmlField.GetFieldTypePackage(), fieldType)
		if enumHelper != nil {
			fieldType = "int32"
		} else {
			// cannot be a simple db column or param
			return "", "", fieldType
		}
	}

	dbName, dbType := DbNameAndType(fieldName, fieldType, fieldTypeQual, dbengine)

	return dbName, dbType, fieldType
}

func DbNameAndType(fieldName string, fieldType string, fieldTypeQual int, dbengine string) (string, string) {
	dbName := fieldName
	dbType := fieldType

	prefix := ""

	switch fieldType {
	case "int32":
		prefix = "int"
		dbType = "INT"
	case "uint32":
		prefix = "int"
		dbType = "INT UNSIGNED"
	case "int64":
		prefix = "inb"
		dbType = "BIGINT"
	case "uint64":
		prefix = "inb"
		dbType = "BIGINT UNSIGNED"
	case "string":
		prefix = "chv"
		dbType = "VARCHAR(" + strconv.Itoa(fieldTypeQual) + ")"
	case "chararray":
		prefix = "chv"
		dbType = "CHAR(" + strconv.Itoa(fieldTypeQual) + ")"
	case "bytes":
		prefix = "bin"
		dbType = "VARBINARY(" + strconv.Itoa(fieldTypeQual) + ")"
	case "bytearray":
		prefix = "bin"
		dbType = "BINARY(" + strconv.Itoa(fieldTypeQual) + ")"
	case "datetime":
		prefix = "dtm"
		dbType = "DATETIME"
	case "guid":
		prefix = "uid"
		dbType = "BINARY(16)"
	case "decimal":
		prefix = "dec"
		dbType = "DECIMAL(19,2)"
	case "bool":
		prefix = "bit"
		dbType = "BOOL"
	case "float":
		prefix = "flt"
		dbType = "FLOAT"
	case "double":
		prefix = "dbl"
		dbType = "DOUBLE"
	}

	if dbengine == "sqlserver" {

		switch fieldType {
		case "bool":
			dbType = "BIT"
		case "guid":
			dbType = "UNIQUEIDENTIFIER"
		case "string":
			dbType = "NVARCHAR(" + strconv.Itoa(fieldTypeQual) + ")"
		case "chararray":
			dbType = "CHAR(" + strconv.Itoa(fieldTypeQual) + ")"
		}
	}

	if dbengine == "postgres" {
		switch fieldType {
		case "uint32":
			prefix = "int"
			dbType = "INT"
		case "uint64":
			prefix = "inb"
			dbType = "BIGINT"
		case "bytes":
			prefix = "bin"
			dbType = "BYTEA"
		case "bytearray":
			prefix = "bin"
			dbType = "BYTEA"
		case "datetime":
			dbType = "TIMESTAMP"
		case "guid":
			prefix = "uid"
			dbType = "BYTEA"
		case "float":
			prefix = "flt"
			dbType = "REAL"
		case "double":
			prefix = "dbl"
			dbType = "FLOAT8"
		}
	}
	if prefix == "" {
		dbName = ""
		dbType = ""
	} else {
		if dbengine == "postgres" {
			// dbName = prefix + "_" + fieldName
			dbName = fieldName
		} else {
			dbName = prefix + UnderscoreToCamelcase(fieldName)
		}
	}

	return dbName, dbType
}

func TableGen(helper *compiler.DmlHelper, outdir string, dbname string, dbengine string, externalTemplate string) error {
	if (helper == nil) || (outdir == "") {
		return fmt.Errorf("TableGen invalid parameters")
	}

	if dbname == "" {
		return fmt.Errorf("TableGen requires dbname")
	}

	if (dbengine != "mysql") && (dbengine != "sqlserver") && (dbengine != "postgres") {
		return fmt.Errorf("TableGen supports only dbengine of mysql, postgres and sqlserver")
	}

	for _, class := range helper.AstRoot.GetClassList() {
		if class.GetTableName() == "" {
			continue
		}

		fieldSetName := class.GetFieldsetName()

		genTable := NewGenTable()
		genTable.TableName = class.GetTableName()
		if dbengine == "postgres" {
			genTable.TableName = strings.ToLower(class.GetTableName())
		}
		genTable.DatabaseName = dbname
		for _, doc := range class.GetDocumentation() {
			genTable.Documentation = append(genTable.Documentation, doc)
		}

		for _, field := range class.GetFields() {
			if (field.GetModifier() == dml.DmlFieldModifier_AUTO) || (field.GetModifier() == dml.DmlFieldModifier_ID) ||
				(field.GetModifier() == dml.DmlFieldModifier_IDGEN) || (field.GetModifier() == dml.DmlFieldModifier_REQUIRED) ||
				(field.GetModifier() == dml.DmlFieldModifier_OPTIONAL) {
				genColumn := NewGenColumn()
				fieldName := field.GetFieldName()
				dmlField := helper.GetField(fieldSetName, fieldName)
				if dmlField != nil {
					for _, doc := range dmlField.GetDocumentation() {
						genColumn.Documentation = append(genColumn.Documentation, doc)
					}

					fieldType := dmlField.FieldType
					// TODO: enum value
					fieldTypeQual := dmlField.GetFieldTypeQualifier()
					dbName, dbType := DbNameAndType(fieldName, fieldType, int(fieldTypeQual), dbengine)
					if dbType != "" {
						// fmt.Printf("dbName: %s, dbType: %s\n", dbName, dbType)
						genColumn.ColumnName = dbName
						genColumn.ColumnType = dbType
						genColumn.IsNullable = field.GetModifier() == dml.DmlFieldModifier_OPTIONAL
						if field.GetModifier() == dml.DmlFieldModifier_IDGEN {
							genColumn.ColumnModifier = "AUTO_INCREMENT"
							if dbengine == "postgres" {
								if dbType == "INT" {
									genColumn.ColumnType = "SERIAL"
								} else if dbType == "BIGINT" {
									genColumn.ColumnType = "BIGSERIAL"
								}
								genColumn.ColumnModifier = ""
							}
						}

						genTable.Columns = append(genTable.Columns, genColumn)
					}
				}

			}
		}

		for _, index := range class.GetIndexes() {
			genIndex := NewGenIndex()
			genIndex.TableName = genTable.TableName

			indexDbNames := make([]string, 0)
			for _, indexField := range index.GetIndexFields() {
				dmlField := helper.GetField(fieldSetName, indexField)
				if dmlField != nil {
					dbName, _ := DbNameAndType(indexField, dmlField.GetFieldType(), int(dmlField.GetFieldTypeQualifier()), dbengine)
					indexDbNames = append(indexDbNames, dbName)
				}

			}

			commaList := strings.Join(indexDbNames, ",")

			genIndex.FieldList = commaList

			if index.GetIndexType() == dml.DmlIndexType_PRIMARY {
				genIndex.IsPrimary = true
				genIndex.IndexName = "PK_" + genTable.TableName
				if dbengine == "postgres" {
					genIndex.IndexName = strings.ToLower(genIndex.IndexName)
				}
			} else {
				underscoreList := strings.Join(indexDbNames, "_")
				genIndex.IndexName = "IX_" + genTable.TableName + "_" + underscoreList
				if dbengine == "postgres" {
					genIndex.IndexName = strings.ToLower(genIndex.IndexName)
				}
				if index.GetIndexType() == dml.DmlIndexType_UNIQUE {
					genIndex.IsUnique = true

				}
			}

			genTable.Indices = append(genTable.Indices, genIndex)

		}

		// fixup comma list
		numIndices := len(genTable.Indices)
		numColumns := len(genTable.Columns)
		if numIndices > 0 {
			genTable.Indices[numIndices-1].HasComma = false
		}
		if numColumns > 0 {
			if (dbengine == "postgres") || (numIndices > 0) {
				genTable.Columns[numColumns-1].HasComma = false
			}
		}

		sqlName := outdir + string(os.PathSeparator) + genTable.TableName + ".sql"

		sqlFile, err := os.Create(sqlName)
		if err != nil {
			return err
		}

		defer sqlFile.Close()

		var tmpl string

		if externalTemplate != "" {
			tmpl = externalTemplate
		} else if dbengine == "mysql" {
			tmpl = TableMysqlTemplate
		} else if dbengine == "sqlserver" {
			tmpl = TableSqlServerlTemplate
		} else if dbengine == "postgres" {
			tmpl = TablePostgresTemplate
		}

		var t = template.Must(template.New("t").Parse(tmpl))

		if err := t.Execute(sqlFile, genTable); err != nil {
			log.Fatal(err)
		}

	}

	return nil
}
