// Copyright 2017-2019 Demian Harvill
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
	"strings"
	"text/template"
	"unicode/utf8"

	"github.com/gaterace/xdml/pkg/compiler"
	"github.com/gaterace/xdml/pkg/dml"
)

type GenProc struct {
	ProcName      string
	TableName     string
	DatabaseName  string
	Documentation []string
	Pattern       string
	ErrCodeParam  string
	InputPrefix   string
	OutputPrefix  string
	DtmNow        string
	Now           string
	TableAlias    *GenNameVal
	Params        []*GenParam
	Declares      []*GenNameVal
	Sets          []*GenNameVal
	Body          []*GenNameVal
	Where         []*GenNameVal
	PostSets      []*GenNameVal
}

type GenParam struct {
	ParamName     string
	ParamType     string
	IsOutput      bool
	Documentation []string
	HasComma      bool
}

type GenNameVal struct {
	Name     string
	Val      string
	HasComma bool
}

func NewGenProc() *GenProc {
	s := GenProc{}
	s.Documentation = make([]string, 0)
	s.Params = make([]*GenParam, 0)
	s.Declares = make([]*GenNameVal, 0)
	s.Sets = make([]*GenNameVal, 0)
	s.Body = make([]*GenNameVal, 0)
	s.Where = make([]*GenNameVal, 0)
	s.PostSets = make([]*GenNameVal, 0)

	return &s
}

func NewGenParam() *GenParam {
	s := GenParam{}
	s.Documentation = make([]string, 0)
	s.HasComma = true
	return &s
}

func NewGenNameVal() *GenNameVal {
	s := GenNameVal{}
	s.HasComma = true
	return &s
}

func rpadText(length int, text string) string {
	strlen := utf8.RuneCountInString(text)
	if length <= strlen {
		return text
	}

	return text + strings.Repeat(" ", length-strlen)

}

var funcMap = template.FuncMap{"rpadText": rpadText}

func ProcGenInit(helper *compiler.DmlHelper, dbname string, dbengine string, fieldSetName string, method *dml.DmlServiceMethod) *GenProc {
	patternType := method.GetPatternType()
	patternClass := method.GetPatternClass()
	classPackage := method.GetPatternClassPackage()
	classHelper, _ := compiler.GetClassInPackage(classPackage, patternClass)
	if classHelper == nil {
		fmt.Printf("cannot find %s in %s\n", patternClass, classPackage)
		return nil
	}

	if classHelper.Class.GetTableName() == "" {
		return nil
	}

	genProc := NewGenProc()

	if dbengine == "sqlserver" {
		genProc.ProcName = "usp_" + method.GetMethodName()
	} else {
		genProc.ProcName = dbname + "." + method.GetMethodName()
	}
	genProc.TableName = classHelper.Class.GetTableName()
	genProc.DatabaseName = dbname
	genProc.Pattern = patternType.String()

	tableAlias := "tbMain"

	nameVal := NewGenNameVal()
	nameVal.HasComma = false
	if dbengine == "sqlserver" {
		nameVal.Name = genProc.TableName
	} else {
		nameVal.Name = genProc.DatabaseName + "." + genProc.TableName
	}

	nameVal.Val = tableAlias
	genProc.TableAlias = nameVal

	genProc.InputPrefix = "in_"
	genProc.OutputPrefix = "out_"
	genProc.DtmNow = "dtmNow"

	if dbengine == "sqlserver" {
		genProc.InputPrefix = "@"
		genProc.OutputPrefix = "@out_"
		genProc.DtmNow = "@dtmNow"
	}

	genProc.Now = "NOW()"
	if dbengine == "sqlserver" {
		genProc.Now = "getdate()"
	}

	for _, doc := range method.GetDocumentation() {
		genProc.Documentation = append(genProc.Documentation, doc)
	}

	for _, param := range method.GetInputFields() {
		fieldName := param.ParameterField
		dmlField := helper.GetField(fieldSetName, fieldName)
		if dmlField == nil {
			continue
		}

		dbName, dbType, _ := DbNameAndTypeDml(dmlField, dbengine)
		if dbName == "" {
			continue
		}

		genParam := NewGenParam()
		for _, doc := range dmlField.GetDocumentation() {
			genParam.Documentation = append(genParam.Documentation, doc)
		}

		genParam.ParamName = genProc.InputPrefix + dbName
		genParam.ParamType = dbType
		genProc.Params = append(genProc.Params, genParam)

	}

	for _, param := range method.GetOutputFields() {
		fieldName := param.ParameterField
		dmlField := helper.GetField(fieldSetName, fieldName)
		if dmlField == nil {
			continue
		}

		dbName, dbType, _ := DbNameAndTypeDml(dmlField, dbengine)
		if dbName == "" {
			continue
		}

		genParam := NewGenParam()
		for _, doc := range dmlField.GetDocumentation() {
			genParam.Documentation = append(genParam.Documentation, doc)
		}

		if fieldName == "error_code" {
			genProc.ErrCodeParam = genProc.OutputPrefix + dbName
		}

		genParam.ParamName = genProc.OutputPrefix + dbName
		genParam.ParamType = dbType
		genParam.IsOutput = true
		genProc.Params = append(genProc.Params, genParam)

	}

	numParams := len(genProc.Params)
	if numParams > 0 {
		genProc.Params[numParams-1].HasComma = false
	}

	return genProc

}

func ProcGenDelete(helper *compiler.DmlHelper, dbname string, dbengine string, fieldSetName string, method *dml.DmlServiceMethod) *GenProc {
	genProc := ProcGenInit(helper, dbname, dbengine, fieldSetName, method)
	if genProc == nil {
		return nil
	}

	genProc.Where = genProc.Where[:0]

	patternClass := method.GetPatternClass()
	classPackage := method.GetPatternClassPackage()
	classHelper, _ := compiler.GetClassInPackage(classPackage, patternClass)
	if classHelper == nil {
		fmt.Printf("cannot find %s in %s\n", patternClass, classPackage)
		return nil
	}

	if classHelper.Class.GetTableName() == "" {
		return nil
	}

	hasVersion := false
	hasDeleted := false
	hasErrorCode := false

	for _, param := range method.GetOutputFields() {
		fieldName := param.GetParameterField()
		if fieldName == "version" {
			hasVersion = true
		} else if fieldName == "deleted" {
			hasDeleted = true
		} else if fieldName == "error_code" {
			hasErrorCode = true
		}
	}

	if !hasErrorCode {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "intErrorCode"
		nameVal.Val = "INT"
		genProc.Declares = append(genProc.Declares, nameVal)
	}

	nameVal := NewGenNameVal()
	nameVal.Name = genProc.DtmNow
	nameVal.Val = "DATETIME"
	genProc.Declares = append(genProc.Declares, nameVal)

	nameVal = NewGenNameVal()
	nameVal.Name = genProc.DtmNow
	nameVal.Val = genProc.Now
	genProc.Sets = append(genProc.Sets, nameVal)

	if hasDeleted {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "dtmDeleted"
		nameVal.Val = genProc.DtmNow
		genProc.Sets = append(genProc.Sets, nameVal)
	}

	if hasVersion {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "intVersion"
		nameVal.Val = genProc.InputPrefix + "intVersion + 1"
		genProc.Sets = append(genProc.Sets, nameVal)
	}

	inputFieldMap := make(map[string]int, 0)

	for _, param := range method.GetInputFields() {
		fieldName := param.GetParameterField()
		inputFieldMap[fieldName] = 1
	}

	hasIsDeleted := false

	for _, classField := range classHelper.Class.GetFields() {
		fieldName := classField.GetFieldName()
		if fieldName == "is_deleted" {
			hasIsDeleted = true
			break
		}
	}

	for _, classField := range classHelper.Class.GetFields() {
		fieldName := classField.GetFieldName()
		fieldModifier := classField.GetModifier()
		dmlField := helper.GetField(fieldSetName, fieldName)
		dbName, _, _ := DbNameAndTypeDml(dmlField, dbengine)

		_, ok := inputFieldMap[fieldName]
		if ok {
			if fieldName != "dummy_param" {
				if (fieldModifier == dml.DmlFieldModifier_IDGEN) || (fieldModifier == dml.DmlFieldModifier_ID) || (fieldModifier == dml.DmlFieldModifier_AUTO) {
					// generate where list
					nameVal := NewGenNameVal()
					nameVal.Name = dbName
					nameVal.Val = genProc.InputPrefix + dbName
					genProc.Where = append(genProc.Where, nameVal)

				}
			}
		} else {
			if fieldName == "is_deleted" {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = "1"
				genProc.Body = append(genProc.Body, nameVal)

				nameVal = NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = "0"
				genProc.Where = append(genProc.Where, nameVal)
			}
		}

		if hasIsDeleted {
			if fieldName == "deleted" {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = genProc.DtmNow
				genProc.Body = append(genProc.Body, nameVal)
			}

			if fieldName == "version" {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = dbName + " + 1"
				genProc.Body = append(genProc.Body, nameVal)
			}
		}

	}

	numElems := len(genProc.Body)
	if numElems > 0 {
		genProc.Body[numElems-1].HasComma = false
	}

	numElems = len(genProc.Where)
	if numElems > 0 {
		genProc.Where[numElems-1].HasComma = false
	}

	return genProc
}

func ProcGenUpdate(helper *compiler.DmlHelper, dbname string, dbengine string, fieldSetName string, method *dml.DmlServiceMethod) *GenProc {
	genProc := ProcGenInit(helper, dbname, dbengine, fieldSetName, method)
	if genProc == nil {
		return nil
	}

	genProc.Where = genProc.Where[:0]

	patternClass := method.GetPatternClass()
	classPackage := method.GetPatternClassPackage()
	classHelper, _ := compiler.GetClassInPackage(classPackage, patternClass)
	if classHelper == nil {
		fmt.Printf("cannot find %s in %s\n", patternClass, classPackage)
		return nil
	}

	if classHelper.Class.GetTableName() == "" {
		return nil
	}

	hasVersion := false
	hasModified := false
	hasErrorCode := false

	for _, param := range method.GetOutputFields() {
		fieldName := param.GetParameterField()
		if fieldName == "version" {
			hasVersion = true
		} else if fieldName == "modified" {
			hasModified = true
		} else if fieldName == "error_code" {
			hasErrorCode = true
		}
	}

	if !hasErrorCode {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "intErrorCode"
		nameVal.Val = "INT"
		genProc.Declares = append(genProc.Declares, nameVal)
	}

	nameVal := NewGenNameVal()
	nameVal.Name = genProc.DtmNow
	nameVal.Val = "DATETIME"
	genProc.Declares = append(genProc.Declares, nameVal)

	nameVal = NewGenNameVal()
	nameVal.Name = genProc.DtmNow
	nameVal.Val = genProc.Now
	genProc.Sets = append(genProc.Sets, nameVal)

	if hasModified {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "dtmModified"
		nameVal.Val = genProc.DtmNow
		genProc.Sets = append(genProc.Sets, nameVal)
	}

	if hasVersion {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "intVersion"
		nameVal.Val = genProc.InputPrefix + "intVersion + 1"
		genProc.Sets = append(genProc.Sets, nameVal)
	}

	inputFieldMap := make(map[string]int, 0)

	for _, param := range method.GetInputFields() {
		fieldName := param.GetParameterField()
		inputFieldMap[fieldName] = 1
	}

	for _, classField := range classHelper.Class.GetFields() {
		fieldName := classField.GetFieldName()
		fieldModifier := classField.GetModifier()
		dmlField := helper.GetField(fieldSetName, fieldName)
		dbName, _, _ := DbNameAndTypeDml(dmlField, dbengine)

		_, ok := inputFieldMap[fieldName]
		if ok {
			if (fieldName == "is_deleted") || ((fieldModifier != dml.DmlFieldModifier_IDGEN) && (fieldModifier != dml.DmlFieldModifier_ID) && (fieldModifier != dml.DmlFieldModifier_AUTO)) {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = genProc.InputPrefix + dbName
				genProc.Body = append(genProc.Body, nameVal)
			}
			if (fieldModifier == dml.DmlFieldModifier_IDGEN) || (fieldModifier == dml.DmlFieldModifier_ID) || (fieldModifier == dml.DmlFieldModifier_AUTO) {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = genProc.InputPrefix + dbName
				genProc.Where = append(genProc.Where, nameVal)

				if fieldName == "version" {
					nameVal := NewGenNameVal()
					nameVal.Name = dbName
					nameVal.Val = dbName + " +  1"
					genProc.Body = append(genProc.Body, nameVal)
				}
			}
		} else {
			// not an input field
			if fieldName == "modified" {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = genProc.DtmNow
				genProc.Body = append(genProc.Body, nameVal)
			} else if fieldName == "is_deleted" {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = "0"
				genProc.Where = append(genProc.Where, nameVal)
			} else if fieldName == "version" {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = dbName + " +  1"
				genProc.Body = append(genProc.Body, nameVal)
			}
		}
	}

	numElems := len(genProc.Body)
	if numElems > 0 {
		genProc.Body[numElems-1].HasComma = false
	}

	numElems = len(genProc.Where)
	if numElems > 0 {
		genProc.Where[numElems-1].HasComma = false
	}

	return genProc
}

func ProcGenInsert(helper *compiler.DmlHelper, dbname string, dbengine string, fieldSetName string, method *dml.DmlServiceMethod) *GenProc {
	genProc := ProcGenInit(helper, dbname, dbengine, fieldSetName, method)
	if genProc == nil {
		return nil
	}

	genProc.Where = genProc.Where[:0]

	patternClass := method.GetPatternClass()
	classPackage := method.GetPatternClassPackage()
	classHelper, _ := compiler.GetClassInPackage(classPackage, patternClass)
	if classHelper == nil {
		fmt.Printf("cannot find %s in %s\n", patternClass, classPackage)
		return nil
	}

	if classHelper.Class.GetTableName() == "" {
		return nil
	}

	hasVersion := false
	hasCreated := false
	hasErrorCode := false

	for _, param := range method.GetOutputFields() {
		fieldName := param.GetParameterField()
		if fieldName == "version" {
			hasVersion = true
		} else if fieldName == "created" {
			hasCreated = true
		} else if fieldName == "error_code" {
			hasErrorCode = true
		}
	}

	if !hasErrorCode {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "intErrorCode"
		nameVal.Val = "INT"
		genProc.Declares = append(genProc.Declares, nameVal)
	}

	nameVal := NewGenNameVal()
	nameVal.Name = genProc.DtmNow
	nameVal.Val = "DATETIME"
	genProc.Declares = append(genProc.Declares, nameVal)

	nameVal = NewGenNameVal()
	nameVal.Name = genProc.DtmNow
	nameVal.Val = genProc.Now
	genProc.Sets = append(genProc.Sets, nameVal)

	if hasCreated {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "dtmCreated"
		nameVal.Val = genProc.DtmNow
		genProc.Sets = append(genProc.Sets, nameVal)
	}

	if hasVersion {
		nameVal := NewGenNameVal()
		nameVal.Name = genProc.OutputPrefix + "intVersion"
		nameVal.Val = "0"
		genProc.Sets = append(genProc.Sets, nameVal)
	}

	inputFieldMap := make(map[string]int, 0)

	for _, param := range method.GetInputFields() {
		fieldName := param.GetParameterField()
		inputFieldMap[fieldName] = 1
	}

	for _, classField := range classHelper.Class.GetFields() {
		fieldName := classField.GetFieldName()
		fieldModifier := classField.GetModifier()
		dmlField := helper.GetField(fieldSetName, fieldName)
		dbName, _, _ := DbNameAndTypeDml(dmlField, dbengine)
		if fieldModifier == dml.DmlFieldModifier_IDGEN {
			nameVal := NewGenNameVal()
			nameVal.Name = genProc.OutputPrefix + dbName
			nameVal.Val = "LAST_INSERT_ID()"
			genProc.PostSets = append(genProc.PostSets, nameVal)
		} else if fieldModifier == dml.DmlFieldModifier_AUTO {
			nameVal := NewGenNameVal()
			nameVal.Name = dbName
			switch fieldName {
			case "created":
				nameVal.Val = genProc.DtmNow
			case "modified":
				nameVal.Val = genProc.DtmNow
			case "deleted":
				nameVal.Val = genProc.DtmNow
			case "is_deleted":
				nameVal.Val = "0"
			case "version":
				nameVal.Val = "0"
			default:
				nameVal.Val = "?"
			}
			genProc.Body = append(genProc.Body, nameVal)

		} else {
			_, ok := inputFieldMap[fieldName]
			if ok {
				nameVal := NewGenNameVal()
				nameVal.Name = dbName
				nameVal.Val = genProc.InputPrefix + dbName
				genProc.Body = append(genProc.Body, nameVal)
			}
		}
	}

	numElems := len(genProc.Body)
	if numElems > 0 {
		genProc.Body[numElems-1].HasComma = false
	}

	return genProc
}

func ProcGenSelect(helper *compiler.DmlHelper, dbname string, dbengine string, fieldSetName string, method *dml.DmlServiceMethod) *GenProc {
	genProc := ProcGenInit(helper, dbname, dbengine, fieldSetName, method)
	if genProc == nil {
		return nil
	}

	// patternType := method.GetPatternType()
	patternClass := method.GetPatternClass()
	classPackage := method.GetPatternClassPackage()
	classHelper, _ := compiler.GetClassInPackage(classPackage, patternClass)
	if classHelper == nil {
		fmt.Printf("cannot find %s in %s\n", patternClass, classPackage)
		return nil
	}

	if classHelper.Class.GetTableName() == "" {
		return nil
	}

	for _, param := range method.GetInputFields() {
		fieldName := param.ParameterField
		if fieldName != "dummy_param" {
			dmlField := helper.GetField(fieldSetName, fieldName)
			if dmlField == nil {
				continue
			}

			dbName, _, _ := DbNameAndTypeDml(dmlField, dbengine)
			if dbName == "" {
				continue
			}

			nameVal := NewGenNameVal()
			nameVal.Name = genProc.TableAlias.Val + "." + dbName
			nameVal.Val = genProc.InputPrefix + dbName
			genProc.Where = append(genProc.Where, nameVal)
		}

	}

	idOnlyResult := false
	idFieldName := ""

	for _, param := range method.GetOutputFields() {
		fieldName := param.ParameterField

		classField, ok := classHelper.FieldMap[fieldName]
		if ok {
			if (classField.Modifier == dml.DmlFieldModifier_ID) || (classField.Modifier == dml.DmlFieldModifier_IDGEN) {
				idOnlyResult = true
				idFieldName = fieldName
				continue
			}
		}

	}

	var deletedFieldName string

	tableAlias := genProc.TableAlias.Val

	for _, classField := range classHelper.Class.GetFields() {
		fieldName := classField.GetFieldName()
		fieldModifier := classField.GetModifier()
		keep := false
		if idOnlyResult {
			keep = fieldName == idFieldName
		} else {
			switch fieldModifier {
			case dml.DmlFieldModifier_AUTO:
				keep = (fieldName != "is_deleted") && (fieldName != "deleted")
				if fieldName == "is_deleted" {
					deletedFieldName = fieldName
				}
			case dml.DmlFieldModifier_ID:
				keep = true
			case dml.DmlFieldModifier_IDGEN:
				keep = true
			case dml.DmlFieldModifier_REQUIRED:
				keep = true
			case dml.DmlFieldModifier_OPTIONAL:
				keep = true
			}
		}

		if keep {
			dmlField := helper.GetField(fieldSetName, fieldName)
			dbName, _, _ := DbNameAndTypeDml(dmlField, dbengine)
			nameVal := NewGenNameVal()
			nameVal.Name = tableAlias + "." + dbName
			nameVal.Val = dbName
			genProc.Body = append(genProc.Body, nameVal)

		}

	}

	numBodyElems := len(genProc.Body)

	if numBodyElems > 0 {
		genProc.Body[numBodyElems-1].HasComma = false
	}

	if deletedFieldName != "" {
		dmlField := helper.GetField(fieldSetName, deletedFieldName)
		dbName, _, _ := DbNameAndTypeDml(dmlField, dbengine)
		nameVal := NewGenNameVal()
		nameVal.Name = tableAlias + "." + dbName
		nameVal.Val = "0"
		genProc.Where = append(genProc.Where, nameVal)

	}

	numWheres := len(genProc.Where)
	if numWheres > 0 {
		genProc.Where[numWheres-1].HasComma = false
	}

	return genProc

}

func ProcGen(helper *compiler.DmlHelper, outdir string, dbname string, dbengine string, externalTemplate string) error {
	if (helper == nil) || (outdir == "") {
		return fmt.Errorf("ProcGen invalid parameters")
	}

	if dbname == "" {
		return fmt.Errorf("ProcGen requires dbname")
	}

	if (dbengine != "mysql") && (dbengine != "sqlserver") {
		return fmt.Errorf("ProcGen supports only dbengine of mysql and sqlserver")
	}

	for _, service := range helper.AstRoot.GetServiceList() {
		serviceName := service.GetServiceName()
		if serviceName == "" {
			continue
		}

		fieldSetName := service.GetFieldsetName()
		for _, method := range service.GetMethods() {
			patternType := method.GetPatternType()

			var genProc *GenProc

			if patternType == dml.DmlPatternType_SELECT {
				genProc = ProcGenSelect(helper, dbname, dbengine, fieldSetName, method)
			} else if patternType == dml.DmlPatternType_INSERT {
				genProc = ProcGenInsert(helper, dbname, dbengine, fieldSetName, method)
			} else if patternType == dml.DmlPatternType_UPDATE {
				genProc = ProcGenUpdate(helper, dbname, dbengine, fieldSetName, method)
			} else if patternType == dml.DmlPatternType_DELETE {
				genProc = ProcGenDelete(helper, dbname, dbengine, fieldSetName, method)
			}

			if genProc == nil {
				continue
			}

			sqlName := outdir + string(os.PathSeparator) + genProc.ProcName + ".sql"

			sqlFile, err := os.Create(sqlName)
			if err != nil {
				return err
			}

			defer sqlFile.Close()

			var tmpl string

			if externalTemplate != "" {
				tmpl = externalTemplate
			} else if dbengine == "mysql" {

				tmpl = ProcMysqlTemplate
			} else if dbengine == "sqlserver" {
				tmpl = ProcSqlServerTemplate
			}

			var t = template.Must(template.New("t").Funcs(funcMap).Parse(tmpl))

			if err := t.Execute(sqlFile, genProc); err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}
