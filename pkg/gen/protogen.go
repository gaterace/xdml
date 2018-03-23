package gen

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"unicode"
	"unicode/utf8"

	"github.com/gaterace/xdml/pkg/compiler"
	"github.com/gaterace/xdml/pkg/dml"
)

type GenBase struct {
	BaseName      string
	JavaPackage   string
	CsharpPackage string
	GoPackage     string

	Imports  []string
	Enums    []*GenEnum
	Classes  []*GenClass
	Services []*GenService
}

type GenEnum struct {
	EnumName      string
	Documentation []string
	Fields        []*GenEnumField
}

type GenEnumField struct {
	FieldName     string
	Documentation []string
	FieldVal      int32
}

type GenClass struct {
	ClassName     string
	Documentation []string
	Fields        []*GenClassField
}

type GenClassField struct {
	FieldName     string
	FieldType     string
	FieldModifier string
	Documentation []string
	FieldVal      int32
}

type GenService struct {
	ServiceName   string
	Documentation []string
	Methods       []*GenMethod
}

type GenMethod struct {
	MethodName    string
	Documentation []string
	RequestClass  string
	ResponseClass string
}

func NewGenBase() *GenBase {
	s := GenBase{}
	s.Imports = make([]string, 0)
	s.Enums = make([]*GenEnum, 0)
	s.Classes = make([]*GenClass, 0)
	s.Services = make([]*GenService, 0)

	return &s
}

func NewGenEnum() *GenEnum {
	s := GenEnum{}
	s.Documentation = make([]string, 0)
	s.Fields = make([]*GenEnumField, 0)
	return &s
}

func NewGenEnumField() *GenEnumField {
	s := GenEnumField{}
	s.Documentation = make([]string, 0)
	return &s
}

func NewGenClass() *GenClass {
	s := GenClass{}
	s.Documentation = make([]string, 0)
	s.Fields = make([]*GenClassField, 0)
	return &s
}

func NewGenClassField() *GenClassField {
	s := GenClassField{}
	s.Documentation = make([]string, 0)
	return &s
}

func NewGenService() *GenService {
	s := GenService{}
	s.Documentation = make([]string, 0)
	s.Methods = make([]*GenMethod, 0)
	return &s
}

func NewGenMethod() *GenMethod {
	s := GenMethod{}
	s.Documentation = make([]string, 0)
	return &s
}

func ProtoFieldType(dmlField *dml.DmlField, helper *compiler.DmlHelper) string {
	if dmlField == nil {
		return ""
	}

	fieldType := dmlField.GetFieldType()
	if dmlField.GetFieldTypePackage() == "builtin" {
		switch fieldType {
		case "datetime":
			fieldType = "dml.DateTime"
		case "decimal":
			fieldType = "dml.Decimal"
		case "guid":
			fieldType = "dml.Guid"
		case "chararray":
			fieldType = "string"
		case "bytearray":
			fieldType = "bytes"
		}
	} else {
		fieldTypePackage := dmlField.GetFieldTypePackage()
		if fieldTypePackage != helper.AstRoot.GetPackageName() {
			fieldType = dmlField.GetFieldTypePackage() + "." + fieldType
		}

	}

	return fieldType
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

func UnderscoreToCamelcase(s string) string {
	doCap := true
	b := buffer{r: make([]byte, 0, len(s))}

	for _, ch := range s {
		if doCap {
			m := unicode.ToUpper(ch)
			b.write(m)
			doCap = false
		} else if ch == '_' {
			doCap = true
		} else {
			b.write(ch)
		}
	}

	return string(b.r)

}

func UnderscoreToCamelcaseLow(s string) string {
	doCap := false
	b := buffer{r: make([]byte, 0, len(s))}

	for _, ch := range s {
		if doCap {
			m := unicode.ToUpper(ch)
			b.write(m)
			doCap = false
		} else if ch == '_' {
			doCap = true
		} else {
			b.write(ch)
		}
	}

	return string(b.r)

}

func firstLetterLower(s string) string {
	forceLower := true
	b := buffer{r: make([]byte, 0, len(s))}

	for _, ch := range s {
		if forceLower {
			m := unicode.ToLower(ch)
			b.write(m)
			forceLower = false
		} else {
			b.write(ch)
		}
	}

	return string(b.r)
}

func ProtoGen(helper *compiler.DmlHelper, outdir string, resequence bool, externalTemplate string) error {
	if (helper == nil) || (outdir == "") {
		return fmt.Errorf("ProtoGen invalid parameters")
	}

	genBase := NewGenBase()
	baseName := helper.BaseName
	genBase.JavaPackage = helper.AstRoot.GetPackageName()
	genBase.CsharpPackage = baseName
	genBase.GoPackage = strings.ToLower(baseName)

	needDmlExtension := true

	for _, importName := range helper.AstRoot.GetImportedPackages() {
		importProto := importName + ".proto"
		genBase.Imports = append(genBase.Imports, importProto)
	}

	if needDmlExtension {
		genBase.Imports = append(genBase.Imports, "DmlExtension.proto")
	}

	for _, enm := range helper.AstRoot.GetEnumList() {
		genEnum := NewGenEnum()
		genEnum.EnumName = enm.GetEnumName()
		for _, doc := range enm.GetDocumentation() {
			genEnum.Documentation = append(genEnum.Documentation, doc)
		}

		genBase.Enums = append(genBase.Enums, genEnum)

		for _, enumValue := range enm.GetValues() {
			genEnumField := NewGenEnumField()
			genEnumField.FieldName = enumValue.GetEnumValueName()
			genEnumField.FieldVal = enumValue.GetEnumValue()
			for _, doc := range enumValue.GetDocumentation() {
				genEnumField.Documentation = append(genEnumField.Documentation, doc)
			}

			genEnum.Fields = append(genEnum.Fields, genEnumField)

		}
	}

	for _, class := range helper.AstRoot.GetClassList() {
		fieldSetName := class.GetFieldsetName()
		// fieldSetHelper, ok := helper.FieldsetMap[fieldSetName]

		genClass := NewGenClass()
		genClass.ClassName = class.GetClassName()
		for _, doc := range class.GetDocumentation() {
			genClass.Documentation = append(genClass.Documentation, doc)
		}

		genBase.Classes = append(genBase.Classes, genClass)

		var fieldSequence int32
		fieldSequence = 0

		for _, field := range class.GetFields() {
			genClassField := NewGenClassField()
			genClassField.FieldName = field.GetFieldName()
			if (field.GetModifier() == dml.DmlFieldModifier_REPEATED) || (field.GetModifier() == dml.DmlFieldModifier_LIST) {
				genClassField.FieldModifier = "repeated "
			}

			dmlField := helper.GetField(fieldSetName, field.GetFieldName())
			if dmlField != nil {
				fieldSequence = fieldSequence + 1
				genClassField.FieldType = ProtoFieldType(dmlField, helper)
				if resequence {
					genClassField.FieldVal = fieldSequence

				} else {
					genClassField.FieldVal = dmlField.GetFieldValue()
				}
				for _, doc := range dmlField.GetDocumentation() {
					genClassField.Documentation = append(genClassField.Documentation, doc)
				}
			}

			genClass.Fields = append(genClass.Fields, genClassField)
		}
	}

	for _, service := range helper.AstRoot.GetServiceList() {
		fieldSetName := service.GetFieldsetName()

		genService := NewGenService()
		genService.ServiceName = service.GetServiceName()
		for _, doc := range service.GetDocumentation() {
			genService.Documentation = append(genService.Documentation, doc)
		}

		genBase.Services = append(genBase.Services, genService)

		for _, method := range service.GetMethods() {
			genMethod := NewGenMethod()
			methodName := method.GetMethodName()
			genMethod.MethodName = methodName
			requestClass := UnderscoreToCamelcase(methodName) + "Request"
			responseClass := UnderscoreToCamelcase(methodName) + "Response"
			genMethod.RequestClass = requestClass
			genMethod.ResponseClass = responseClass

			for _, doc := range method.GetDocumentation() {
				genMethod.Documentation = append(genMethod.Documentation, doc)
			}

			genService.Methods = append(genService.Methods, genMethod)

			genClass := NewGenClass()
			genClass.ClassName = requestClass
			genClass.Documentation = append(genClass.Documentation, "request parameters for method "+methodName)
			genBase.Classes = append(genBase.Classes, genClass)

			var fieldSequence int32
			fieldSequence = 0

			for _, field := range method.GetInputFields() {
				genClassField := NewGenClassField()
				genClassField.FieldName = field.GetParameterField()
				if field.GetModifier() == dml.DmlParameterModifier_PARAM_REPEATED {
					genClassField.FieldModifier = "repeated "
				}

				dmlField := helper.GetField(fieldSetName, field.GetParameterField())
				if dmlField != nil {
					fieldSequence = fieldSequence + 1
					genClassField.FieldType = ProtoFieldType(dmlField, helper)
					if resequence {
						genClassField.FieldVal = fieldSequence
					} else {
						genClassField.FieldVal = dmlField.GetFieldValue()
					}
					for _, doc := range dmlField.GetDocumentation() {
						genClassField.Documentation = append(genClassField.Documentation, doc)
					}
				}

				genClass.Fields = append(genClass.Fields, genClassField)
			}

			genClass = NewGenClass()
			genClass.ClassName = responseClass
			genClass.Documentation = append(genClass.Documentation, "response parameters for method "+methodName)
			genBase.Classes = append(genBase.Classes, genClass)
			fieldSequence = 0
			for _, field := range method.GetOutputFields() {
				genClassField := NewGenClassField()
				genClassField.FieldName = field.GetParameterField()
				if field.GetModifier() == dml.DmlParameterModifier_PARAM_REPEATED {
					genClassField.FieldModifier = "repeated "
				}

				dmlField := helper.GetField(fieldSetName, field.GetParameterField())
				if dmlField != nil {
					fieldSequence = fieldSequence + 1
					genClassField.FieldType = ProtoFieldType(dmlField, helper)
					if resequence {
						genClassField.FieldVal = fieldSequence
					} else {
						genClassField.FieldVal = dmlField.GetFieldValue()
					}
					for _, doc := range dmlField.GetDocumentation() {
						genClassField.Documentation = append(genClassField.Documentation, doc)
					}
				}

				genClass.Fields = append(genClass.Fields, genClassField)
			}

		}
	}

	protoName := outdir + string(os.PathSeparator) + baseName + ".proto"

	protoFile, err := os.Create(protoName)
	if err != nil {
		return err
	}

	defer protoFile.Close()

	var tmpl string

	if externalTemplate != "" {
		tmpl = externalTemplate
	} else {
		tmpl = ProtoTemplate
	}

	// run the template
	var t = template.Must(template.New("t").Parse(tmpl))

	if err := t.Execute(protoFile, genBase); err != nil {
		log.Fatal(err)
	}

	return nil
}
