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
	"html/template"
	"log"
	"os"

	"github.com/gaterace/xdml/pkg/compiler"
	"github.com/gaterace/xdml/pkg/dml"
)

type GenServiceCode struct {
	Action        string
	ItfPackage    string
	ImplPackage   string
	Name          string
	Documentation []string
	ItfImports    []string
	ImplImports   []string
	Methods       []*GenMethodCode
}

type GenMethodCode struct {
	MethodName    string
	RequestClass  string
	ResponseClass string
	ResultClass   string
	ResultObj     string
	Documentation []string
	HasErrorCode  bool
	ProcName      string
	ParamPattern  string
	ResultSetter  string
	HasResultSet  bool
	HasRepeated   bool
	IdResult      bool
	Params        []*GenParamCode
	Results       []*GenResultCode
}

type GenParamCode struct {
	ParamName  string
	SqlName    string
	ParamIndex int32
	IsOutput   bool
	SqlType    string
	SqlGetter  string
	SqlSetter  string
	JavaGetter string
	JavaSetter string
	Converter  string
	ParamType  string
}

type GenResultCode struct {
	SqlGetter   string
	ColumnName  string
	ColumnIndex int32
	JavaSetter  string
	ResultObj   string
	ResultField string
	Converter   string
}

func NewGenServiceCode() *GenServiceCode {
	s := GenServiceCode{}
	s.Documentation = make([]string, 0)
	s.ItfImports = make([]string, 0)
	s.ImplImports = make([]string, 0)
	s.Methods = make([]*GenMethodCode, 0)

	return &s
}

func NewGenMethodCode() *GenMethodCode {
	s := GenMethodCode{}
	s.Documentation = make([]string, 0)
	s.Params = make([]*GenParamCode, 0)
	s.Results = make([]*GenResultCode, 0)

	return &s
}

func NewGenParamCode() *GenParamCode {
	s := GenParamCode{}

	return &s
}

func NewGenResultCode() *GenResultCode {
	s := GenResultCode{}

	return &s
}

func MapSqlAccessor(fieldType string, language string) string {

	accessor := ""

	if language == "java" {

		switch fieldType {
		case "int32":
			accessor = "Int"
		case "uint32":
			accessor = "Int"
		case "int64":
			accessor = "Long"
		case "uint64":
			accessor = "Long"
		case "string":
			accessor = "String"
		case "chararray":
			accessor = "String"
		case "bytes":
			accessor = "Bytes"
		case "bytearray":
			accessor = "Bytes"
		case "datetime":
			accessor = "Timestamp"
		case "guid":
			accessor = "Bytes"
		case "decimal":
			accessor = "BigDecimal"
		case "bool":
			accessor = "Boolean"
		case "float":
			accessor = "Float"
		case "double":
			accessor = "Double"
		}
	} else if language == "csharp" {
		switch fieldType {
		case "int32":
			accessor = "Int32"
		case "uint32":
			accessor = "Int32"
		case "int64":
			accessor = "Int64"
		case "uint64":
			accessor = "Int64"
		case "string":
			accessor = "String"
		case "chararray":
			accessor = "String"
		case "bytes":
			accessor = "Bytes"
		case "bytearray":
			accessor = "Bytes"
		case "datetime":
			accessor = "DateTime"
		case "guid":
			accessor = "Guid"
		case "decimal":
			accessor = "Decimal"
		case "bool":
			accessor = "Boolean"
		case "float":
			accessor = "Float"
		case "double":
			accessor = "Double"
		}
	}

	return accessor

}

func MapSqlType(fieldType string, language string) string {

	sqlType := ""

	if language == "java" {

		switch fieldType {
		case "int32":
			sqlType = "INTEGER"
		case "uint32":
			sqlType = "INTEGER"
		case "int64":
			sqlType = "BIGINT"
		case "uint64":
			sqlType = "BIGINT"
		case "string":
			sqlType = "VARCHAR"
		case "chararray":
			sqlType = "CHAR"
		case "bytes":
			sqlType = "VARBINARY"
		case "bytearray":
			sqlType = "BINARY"
		case "datetime":
			sqlType = "DATE"
		case "guid":
			sqlType = "BINARY"
		case "decimal":
			sqlType = "DECIMAL"
		case "bool":
			sqlType = "BIT"
		case "float":
			sqlType = "FLOAT"
		case "double":
			sqlType = "DOUBLE"
		}
	} else if language == "csharp" {
		switch fieldType {
		case "int32":
			sqlType = "Int"
		case "uint32":
			sqlType = "Int"
		case "int64":
			sqlType = "BigInt"
		case "uint64":
			sqlType = "BigInt"
		case "string":
			sqlType = "NVarChar"
		case "chararray":
			sqlType = "Char"
		case "bytes":
			sqlType = "VarBinary"
		case "bytearray":
			sqlType = "Binary"
		case "datetime":
			sqlType = "DateTime"
		case "guid":
			sqlType = "UniqueIdentifier"
		case "decimal":
			sqlType = "Decimal"
		case "bool":
			sqlType = "Bit"
		case "float":
			sqlType = "Float"
		case "double":
			sqlType = "Double"
		}
	}

	return sqlType

}

func MapLanguageType(fieldType string, language string) string {
	languageType := ""

	if language == "java" {
		switch fieldType {
		case "int32":
			languageType = "int"
		case "uint32":
			languageType = "int"
		case "int64":
			languageType = "long"
		case "uint64":
			languageType = "long"
		case "string":
			languageType = "String"
		case "chararray":
			languageType = "String"
		case "bytes":
			languageType = "Bytes"
		case "bytearray":
			languageType = "Bytes"
		case "datetime":
			languageType = "Timestamp"
		case "guid":
			languageType = "Bytes"
		case "decimal":
			languageType = "BigDecimal"
		case "bool":
			languageType = "boolean"
		case "float":
			languageType = "float"
		case "double":
			languageType = "double"
		}
	} else if language == "csharp" {
		switch fieldType {
		case "int32":
			languageType = "int"
		case "uint32":
			languageType = "int"
		case "int64":
			languageType = "long"
		case "uint64":
			languageType = "long"
		case "string":
			languageType = "string"
		case "chararray":
			languageType = "string"
		case "bytes":
			languageType = "byte[]"
		case "bytearray":
			languageType = "byte[]"
		case "datetime":
			languageType = "System.DateTime"
		case "guid":
			languageType = "Guid"
		case "decimal":
			languageType = "decimal"
		case "bool":
			languageType = "boolean"
		case "float":
			languageType = "float"
		case "double":
			languageType = "double"
		}
	}

	return languageType
}

func MapConverter(fieldType string) string {

	converter := ""

	switch fieldType {

	case "datetime":
		converter = "DmlUtil.ConvertDateTime"
	case "guid":
		converter = "DmlUtil.ConvertGuid"
	case "decimal":
		converter = "DmlUtil.ConvertDecimal"
	}

	return converter

}

func ServiceGen(helper *compiler.DmlHelper, outdir string, dbname string, dbengine string, language string, externalTemplate string) error {
	if (helper == nil) || (outdir == "") {
		return fmt.Errorf("ServiceGen invalid parameters")
	}

	if dbname == "" {
		return fmt.Errorf("ServiceGen requires dbname")
	}

	if (dbengine != "mysql") && (dbengine != "sqlserver") {
		return fmt.Errorf("ServiceGen supports only dbengine of mysql and sqlserver")
	}

	if (language != "java") && (language != "csharp") {
		return fmt.Errorf("ServiceGen supports only java and csharp language")
	}

	if (language == "csharp") && (dbengine == "mysql") {
		return fmt.Errorf("ServiceGen does not support csharp language for mysql")
	}

	packageName := helper.AstRoot.GetPackageName()
	serviceBaseName := helper.BaseName

	itfPackage := packageName + ".api"
	implPackage := packageName + ".impl"

	if language == "csharp" {
		itfPackage = serviceBaseName + ".API"
		implPackage = serviceBaseName + ".Impl"
	}

	for _, service := range helper.AstRoot.GetServiceList() {
		fieldSetName := service.GetFieldsetName()

		genService := NewGenServiceCode()
		serviceName := service.GetServiceName()
		genService.Name = serviceName
		genService.ItfPackage = itfPackage
		genService.ImplPackage = implPackage

		for _, doc := range service.GetDocumentation() {
			genService.Documentation = append(genService.Documentation, doc)
		}

		classImportMap := make(map[string]bool)

		if language == "csharp" {
			genService.ItfImports = append(genService.ItfImports, serviceBaseName)
		}

		for _, method := range service.GetMethods() {
			genMethod := NewGenMethodCode()

			patternClass := method.GetPatternClass()
			classPackage := method.GetPatternClassPackage()
			classHelper, baseName := compiler.GetClassInPackage(classPackage, patternClass)

			methodReturnsIds := false
			responseField := ""

			// fmt.Printf("patternClass: %s\n", patternClass)
			patternType := method.GetPatternType()

			if patternType == dml.DmlPatternType_SELECT {
				genMethod.HasResultSet = true
				classImport := classPackage + "." + baseName + "." + patternClass
				if language == "csharp" {
					classImport = baseName
				}
				// fmt.Printf("classImport: %s\n", classImport)
				_, ok := classImportMap[classImport]
				if !ok {
					classImportMap[classImport] = true
				}

				genMethod.ResultClass = patternClass
				genMethod.ResultObj = firstLetterLower(patternClass)
			}

			methodName := method.GetMethodName()
			javaMethodName := UnderscoreToCamelcaseLow(methodName)
			csharpMethodName := UnderscoreToCamelcase(methodName)
			requestClass := UnderscoreToCamelcase(methodName) + "Request"
			responseClass := UnderscoreToCamelcase(methodName) + "Response"
			genMethod.MethodName = javaMethodName
			if language == "csharp" {
				genMethod.MethodName = csharpMethodName
			}
			genMethod.RequestClass = requestClass
			genMethod.ResponseClass = responseClass
			procName := dbname + "." + methodName
			if dbengine == "sqlserver" {
				procName = "usp_" + methodName
			}

			genMethod.ProcName = procName

			if language == "java" {
				importPath := packageName + "." + serviceBaseName + "." + requestClass
				genService.ItfImports = append(genService.ItfImports, importPath)
				importPath = packageName + "." + serviceBaseName + "." + responseClass
				genService.ItfImports = append(genService.ItfImports, importPath)
			}

			for _, doc := range method.GetDocumentation() {
				genMethod.Documentation = append(genMethod.Documentation, doc)
			}

			var index int32
			index = 0

			for _, field := range method.GetInputFields() {
				dmlField := helper.GetField(fieldSetName, field.GetParameterField())
				if dmlField != nil {
					// TODO
					if len(genMethod.ParamPattern) == 0 {
						genMethod.ParamPattern = "?"
					} else {
						genMethod.ParamPattern = genMethod.ParamPattern + ",?"
					}

					genParam := NewGenParamCode()
					index = index + 1
					genParam.ParamName = field.GetParameterField()
					if language == "csharp" {
						genParam.ParamName = UnderscoreToCamelcase(field.GetParameterField())
					}
					dbName, _, baseType := DbNameAndTypeDml(dmlField, dbengine)
					if dbengine == "sqlserver" {
						genParam.SqlName = "@" + dbName
					} else {
						genParam.SqlName = "in_" + dbName
					}
					genParam.ParamIndex = index

					// fmt.Printf("dbName: %s, dbType: %s, baseType: %s\n", dbName, dbType, baseType)

					accessor := MapSqlAccessor(baseType, language)
					if accessor != "" {
						if language == "java" {
							genParam.SqlGetter = "get" + accessor
							genParam.SqlSetter = "set" + accessor
						} else if language == "csharp" {
							genParam.SqlGetter = "Get" + accessor
							// genParam.SqlSetter = "set" + accessor
						}
					}

					sqlType := MapSqlType(baseType, language)
					genParam.SqlType = sqlType

					languageType := MapLanguageType(baseType, language)
					genParam.ParamType = languageType

					genParam.JavaGetter = "request.get" + UnderscoreToCamelcase(genParam.ParamName) + "()"
					genParam.JavaSetter = "response.set" + UnderscoreToCamelcase(genParam.ParamName)

					genParam.Converter = MapConverter(baseType)

					genMethod.Params = append(genMethod.Params, genParam)
				}
			}

			// get the names of ID and IDGEN fields

			idFieldNames := make(map[string]bool)

			for _, classField := range classHelper.Class.Fields {
				fieldName := classField.GetFieldName()
				if (classField.GetModifier() == dml.DmlFieldModifier_ID) || (classField.GetModifier() == dml.DmlFieldModifier_IDGEN) {
					idFieldNames[fieldName] = true
				}

			}

			for _, field := range method.GetOutputFields() {
				dmlField := helper.GetField(fieldSetName, field.GetParameterField())
				if dmlField != nil {
					dbName, dbType, baseType := DbNameAndTypeDml(dmlField, dbengine)

					if dbType == "" {

						if field.GetModifier() == dml.DmlParameterModifier_PARAM_REPEATED {
							genMethod.HasRepeated = true
							if language == "java" {
								genMethod.ResultSetter = "add" + UnderscoreToCamelcase(field.GetParameterField())
							} else if language == "csharp" {
								genMethod.ResultSetter = UnderscoreToCamelcase(field.GetParameterField())
							}
						} else {
							genMethod.HasRepeated = false
							if language == "java" {
								genMethod.ResultSetter = "set" + UnderscoreToCamelcase(field.GetParameterField())
							} else if language == "csharp" {
								genMethod.ResultSetter = UnderscoreToCamelcase(field.GetParameterField())
							}
						}

						continue
					}

					if field.GetModifier() == dml.DmlParameterModifier_PARAM_REPEATED {
						_, ok := idFieldNames[field.GetParameterField()]
						if ok {
							methodReturnsIds = true
							genMethod.HasRepeated = true
							responseField = field.GetParameterField()
							continue
						}
					}

					if len(genMethod.ParamPattern) == 0 {
						genMethod.ParamPattern = "?"
					} else {
						genMethod.ParamPattern = genMethod.ParamPattern + ",?"
					}

					if dmlField.GetFieldName() == "error_code" {
						genMethod.HasErrorCode = true
					}

					genParam := NewGenParamCode()
					index = index + 1
					genParam.ParamName = field.GetParameterField()
					if language == "csharp" {
						genParam.ParamName = UnderscoreToCamelcase(field.GetParameterField())
					}

					if dbengine == "sqlserver" {
						genParam.SqlName = "@out_" + dbName
					} else {
						genParam.SqlName = "out_" + dbName
					}
					genParam.ParamIndex = index
					genParam.IsOutput = true

					accessor := MapSqlAccessor(baseType, language)
					if accessor != "" {
						if language == "java" {
							genParam.SqlGetter = "get" + accessor
							genParam.SqlSetter = "set" + accessor
						} else if language == "chsarp" {
							genParam.SqlGetter = "Get" + accessor
							// genParam.SqlSetter = "Set" + accessor
						}
					}

					sqlType := MapSqlType(baseType, language)
					genParam.SqlType = sqlType

					languageType := MapLanguageType(baseType, language)
					genParam.ParamType = languageType

					genParam.JavaGetter = "request.get" + UnderscoreToCamelcase(genParam.ParamName) + "()"
					genParam.JavaSetter = "response.set" + UnderscoreToCamelcase(genParam.ParamName)

					genParam.Converter = MapConverter(baseType)

					genMethod.Params = append(genMethod.Params, genParam)

				}
			}

			genMethod.IdResult = methodReturnsIds
			if genMethod.HasResultSet {
				var columnIndex int32
				for _, classField := range classHelper.Class.Fields {
					fieldName := classField.GetFieldName()
					// fieldModifier := classField.GetModifier()
					if (fieldName != "deleted") && (fieldName != "is_deleted") {
						dmlField := helper.GetField(fieldSetName, fieldName)

						if !methodReturnsIds || (fieldName == responseField) {
							dbName, _, baseType := DbNameAndTypeDml(dmlField, dbengine)
							genResult := NewGenResultCode()
							accessor := MapSqlAccessor(baseType, language)
							if language == "java" {
								genResult.SqlGetter = "get" + accessor
							} else if language == "csharp" {
								genResult.SqlGetter = "Get" + accessor
							}
							genResult.ColumnName = dbName
							useIndex := columnIndex
							genResult.ColumnIndex = useIndex
							columnIndex = columnIndex + 1
							if methodReturnsIds {
								genResult.JavaSetter = "add" + UnderscoreToCamelcase(fieldName)
							} else {
								genResult.JavaSetter = "set" + UnderscoreToCamelcase(fieldName)
							}
							genResult.ResultObj = genMethod.ResultObj
							genResult.ResultField = UnderscoreToCamelcase(fieldName)
							genResult.Converter = MapConverter(baseType)
							genMethod.Results = append(genMethod.Results, genResult)
						}
					}
				}
			}

			genService.Methods = append(genService.Methods, genMethod)
		}

		for classImport, _ := range classImportMap {
			genService.ImplImports = append(genService.ImplImports, classImport)
		}

		fileExt := ".java"

		if language == "csharp" {
			fileExt = ".cs"
		}

		// create interface
		itfFileName := outdir + string(os.PathSeparator) + "I" + serviceName + fileExt

		itfFile, err := os.Create(itfFileName)
		if err != nil {
			return err
		}

		defer itfFile.Close()

		var tmpl string

		if externalTemplate != "" {
			tmpl = externalTemplate
		} else if language == "java" {
			tmpl = ServiceJavaTemplate
		} else if language == "csharp" {
			tmpl = ServiceCsharpTemplate
		}

		var t = template.Must(template.New("t").Parse(tmpl))

		genService.Action = "itf"

		if err := t.Execute(itfFile, genService); err != nil {
			log.Fatal(err)
		}

		implFileName := outdir + string(os.PathSeparator) + serviceName + fileExt
		implFile, err := os.Create(implFileName)
		if err != nil {
			return err
		}

		defer implFile.Close()

		genService.Action = "impl"

		if err := t.Execute(implFile, genService); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
