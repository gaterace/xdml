package gen

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"gaterace.org/xdml/pkg/dml"

	"gaterace.org/xdml/pkg/compiler"
)

type GenServiceCode struct {
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
}

type GenResultCode struct {
	SqlGetter  string
	ColumnName string
	JavaSetter string
	ResultObj  string
	Converter  string
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

func MapSqlAccessor(fieldType string) string {

	accessor := ""

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

	return accessor

}

func MapSqlType(fieldType string) string {

	sqlType := ""

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

	return sqlType

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

func ServiceGen(helper *compiler.DmlHelper, outdir string, dbname string, dbengine string, language string) error {
	if (helper == nil) || (outdir == "") {
		return fmt.Errorf("ServiceGen invalid parameters")
	}

	if dbname == "" {
		return fmt.Errorf("ServiceGen requires dbname")
	}

	if (dbengine != "mysql") && (dbengine != "sqlserver") {
		return fmt.Errorf("ServiceGen supports only dbengine of mysql and sqlserver")
	}

	if language != "java" {
		return fmt.Errorf("ServiceGen supports only java language")
	}

	itfPackage := helper.AstRoot.GetPackageName() + ".api"
	implPackage := helper.AstRoot.GetPackageName() + ".impl"
	packageName := helper.AstRoot.GetPackageName()
	serviceBaseName := helper.BaseName

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
			requestClass := UnderscoreToCamelcase(methodName) + "Request"
			responseClass := UnderscoreToCamelcase(methodName) + "Response"
			genMethod.MethodName = javaMethodName
			genMethod.RequestClass = requestClass
			genMethod.ResponseClass = responseClass
			procName := dbname + "." + methodName
			genMethod.ProcName = procName

			importPath := packageName + "." + serviceBaseName + "." + requestClass
			genService.ItfImports = append(genService.ItfImports, importPath)
			importPath = packageName + "." + serviceBaseName + "." + responseClass
			genService.ItfImports = append(genService.ItfImports, importPath)

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
					dbName, _, baseType := DbNameAndTypeDml(dmlField, dbengine)
					genParam.SqlName = "in_" + dbName
					genParam.ParamIndex = index

					// fmt.Printf("dbName: %s, dbType: %s, baseType: %s\n", dbName, dbType, baseType)

					accessor := MapSqlAccessor(baseType)
					if accessor != "" {
						genParam.SqlGetter = "get" + accessor
						genParam.SqlSetter = "set" + accessor
					}

					sqlType := MapSqlType(baseType)
					genParam.SqlType = sqlType

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
							genMethod.ResultSetter = "add" + UnderscoreToCamelcase(field.GetParameterField())
						} else {
							genMethod.HasRepeated = false
							genMethod.ResultSetter = "set" + UnderscoreToCamelcase(field.GetParameterField())
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

					genParam.SqlName = "out_" + dbName
					genParam.ParamIndex = index
					genParam.IsOutput = true

					accessor := MapSqlAccessor(baseType)
					if accessor != "" {
						genParam.SqlGetter = "get" + accessor
						genParam.SqlSetter = "set" + accessor
					}

					sqlType := MapSqlType(baseType)
					genParam.SqlType = sqlType

					genParam.JavaGetter = "request.get" + UnderscoreToCamelcase(genParam.ParamName) + "()"
					genParam.JavaSetter = "response.set" + UnderscoreToCamelcase(genParam.ParamName)

					genParam.Converter = MapConverter(baseType)

					genMethod.Params = append(genMethod.Params, genParam)

				}
			}

			genMethod.IdResult = methodReturnsIds
			for _, classField := range classHelper.Class.Fields {
				fieldName := classField.GetFieldName()
				// fieldModifier := classField.GetModifier()
				if (fieldName != "deleted") && (fieldName != "is_deleted") {
					dmlField := helper.GetField(fieldSetName, fieldName)
					if !methodReturnsIds || (fieldName == responseField) {
						dbName, _, baseType := DbNameAndTypeDml(dmlField, dbengine)
						genResult := NewGenResultCode()
						accessor := MapSqlAccessor(baseType)
						genResult.SqlGetter = "get" + accessor
						genResult.ColumnName = dbName
						if methodReturnsIds {
							genResult.JavaSetter = "add" + UnderscoreToCamelcase(fieldName)
						} else {
							genResult.JavaSetter = "set" + UnderscoreToCamelcase(fieldName)
						}
						genResult.ResultObj = genMethod.ResultObj
						genResult.Converter = MapConverter(baseType)
						genMethod.Results = append(genMethod.Results, genResult)
					}
				}
			}

			genService.Methods = append(genService.Methods, genMethod)
		}

		for classImport, _ := range classImportMap {
			genService.ImplImports = append(genService.ImplImports, classImport)
		}

		// create interface
		itfFileName := outdir + string(os.PathSeparator) + "I" + serviceName + ".java"

		itfFile, err := os.Create(itfFileName)
		if err != nil {
			return err
		}

		defer itfFile.Close()

		var tmpl string

		tmpl = ServiceJavaTemplateItf

		var t = template.Must(template.New("t").Parse(tmpl))

		if err := t.Execute(itfFile, genService); err != nil {
			log.Fatal(err)
		}

		implFileName := outdir + string(os.PathSeparator) + serviceName + ".java"
		implFile, err := os.Create(implFileName)
		if err != nil {
			return err
		}

		defer implFile.Close()

		tmpl = ServiceJavaTemplateImpl

		var t2 = template.Must(template.New("t2").Parse(tmpl))

		if err := t2.Execute(implFile, genService); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
