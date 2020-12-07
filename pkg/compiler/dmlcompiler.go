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

package compiler

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gaterace/xdml/pkg/dml"
	"github.com/gaterace/xdml/pkg/parser"
)

type DmlErrorListener struct {
	*antlr.DefaultErrorListener
	NumErrors int
}

func (d *DmlErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.NumErrors = d.NumErrors + 1
}

func NewDmlErrorListener() *DmlErrorListener {
	listener := DmlErrorListener{}
	return &listener
}

type TreeListener struct {
	*parser.BaseDMLListener
	AstRoot       *dml.DmlTree
	dmlEnum       *dml.DmlEnum
	dmlEnumVal    *dml.DmlEnumValue
	dmlFieldSet   *dml.DmlFieldSet
	dmlField      *dml.DmlField
	dmlClass      *dml.DmlClass
	dmlClassField *dml.DmlClassField
	dmlIndex      *dml.DmlClassIndex
	dmlService    *dml.DmlService
	dmlMethod     *dml.DmlServiceMethod
	ctxMap        map[interface{}]interface{}
}

func NewTreeListener() *TreeListener {
	listener := TreeListener{ctxMap: make(map[interface{}]interface{})}
	return &listener
}

func (s *TreeListener) SetValue(ctx interface{}, val interface{}) {
	if ctx != nil {
		s.ctxMap[ctx] = val
	}
}

func (s *TreeListener) EnterModel(ctx *parser.ModelContext) {
	s.AstRoot = &dml.DmlTree{}
}

func (s *TreeListener) ExitPackageElement(ctx *parser.PackageElementContext) {
	qctx := ctx.QualifiedId()
	packageName, ok := s.ctxMap[qctx]
	if ok {
		s.AstRoot.PackageName = packageName.(string)
	}

}

func (s *TreeListener) EnterQualifiedId(ctx *parser.QualifiedIdContext) {
	var buffer bytes.Buffer
	for _, child := range ctx.GetChildren() {
		var payload = child.GetPayload()
		token, ok := payload.(*antlr.CommonToken)
		if ok {
			buffer.WriteString(token.GetText())
		}
	}

	s.SetValue(ctx, buffer.String())
}

func (s *TreeListener) ExitImportElements(ctx *parser.ImportElementsContext) {
	if ctx.GetChildCount() == 3 {
		child := ctx.GetChild(1)
		var payload = child.GetPayload()
		token, ok := payload.(*antlr.CommonToken)
		if ok {
			quotedText := token.GetText()
			importName := quotedText[1 : len(quotedText)-1]
			s.AstRoot.ImportedPackages = append(s.AstRoot.ImportedPackages, importName)
		}

	}
}

func (s *TreeListener) EnterEnumElement(ctx *parser.EnumElementContext) {
	s.dmlEnum = &dml.DmlEnum{}
	s.dmlEnum.EnumName = ctx.ID().GetText()
	for _, child := range ctx.AllKEEP_COMMENT() {
		nodeText := child.GetText()
		docString := strings.TrimSpace(nodeText[3:])
		s.dmlEnum.Documentation = append(s.dmlEnum.Documentation, docString)
	}
}

func (s *TreeListener) ExitEnumElement(ctx *parser.EnumElementContext) {
	if (s.AstRoot != nil) && (s.dmlEnum != nil) {
		s.AstRoot.EnumList = append(s.AstRoot.EnumList, s.dmlEnum)

		s.dmlEnum = nil
	}
}

func (s *TreeListener) EnterIntegerConst(ctx *parser.IntegerConstContext) {
	text := ctx.GetText()
	k, err := strconv.Atoi(text)
	if err == nil {
		s.SetValue(ctx, k)
	}

}

func (s *TreeListener) EnterEnumValue(ctx *parser.EnumValueContext) {
	s.dmlEnumVal = &dml.DmlEnumValue{}
	s.dmlEnumVal.EnumValueName = ctx.ID().GetText()
	for _, child := range ctx.AllKEEP_COMMENT() {
		nodeText := child.GetText()
		docString := strings.TrimSpace(nodeText[3:])
		s.dmlEnumVal.Documentation = append(s.dmlEnumVal.Documentation, docString)
	}

}

func (s *TreeListener) ExitEnumValue(ctx *parser.EnumValueContext) {
	integerCtx := ctx.IntegerConst()
	var k int
	k = -1
	intVal, ok := s.ctxMap[integerCtx]
	if ok {
		k = intVal.(int)
	}

	if s.dmlEnum != nil {
		if k < 0 {
			// need to fix up
			size := len(s.dmlEnum.GetValues())
			if size == 0 {
				k = 0
			} else {
				lastVal := s.dmlEnum.GetValues()[size-1].GetEnumValue()
				k = int(lastVal + 1)
			}
		}

		if s.dmlEnumVal != nil {
			s.dmlEnumVal.EnumValue = int32(k)
			s.dmlEnum.Values = append(s.dmlEnum.Values, s.dmlEnumVal)
			s.dmlEnumVal = nil
		}
	}
}

func (s *TreeListener) EnterExtendsQualifier(ctx *parser.ExtendsQualifierContext) {
	extendsName := ctx.ID().GetText()
	s.SetValue(ctx, extendsName)
}

func (s *TreeListener) EnterFieldsetElement(ctx *parser.FieldsetElementContext) {
	s.dmlFieldSet = &dml.DmlFieldSet{}
	s.dmlFieldSet.FieldsetName = ctx.ID().GetText()
	for _, child := range ctx.AllKEEP_COMMENT() {
		nodeText := child.GetText()
		docString := strings.TrimSpace(nodeText[3:])
		s.dmlFieldSet.Documentation = append(s.dmlFieldSet.Documentation, docString)
	}
}

func (s *TreeListener) ExitFieldsetElement(ctx *parser.FieldsetElementContext) {
	if s.dmlFieldSet != nil {
		for _, child := range ctx.GetChildren() {
			qctx, ok := child.(*parser.ExtendsQualifierContext)
			if ok {
				qctxVal, ok := s.ctxMap[qctx]
				if ok {
					extendsQual := qctxVal.(string)
					s.dmlFieldSet.ExtendsFieldsetName = extendsQual
				}
			}
		}

		if s.AstRoot != nil {
			s.AstRoot.FieldsetList = append(s.AstRoot.FieldsetList, s.dmlFieldSet)
		}
		s.dmlFieldSet = nil
	}
}

func (s *TreeListener) EnterFieldsetValue(ctx *parser.FieldsetValueContext) {
	s.dmlField = &dml.DmlField{}
	s.dmlField.FieldName = ctx.ID().GetText()
	for _, child := range ctx.AllKEEP_COMMENT() {
		nodeText := child.GetText()
		docString := strings.TrimSpace(nodeText[3:])
		s.dmlField.Documentation = append(s.dmlField.Documentation, docString)
	}

}

func (s *TreeListener) ExitFieldsetValue(ctx *parser.FieldsetValueContext) {
	if s.dmlField != nil {
		var k int
		k = -1
		for _, child := range ctx.GetChildren() {
			integerCtx, ok := child.(*parser.IntegerConstContext)
			if ok {
				intVal, ok := s.ctxMap[integerCtx]
				if ok {
					k = intVal.(int)
				}
			}
		}

		if k < 0 {
			// need to fix up
			size := len(s.dmlFieldSet.GetFields())
			if size == 0 {
				k = 0
			} else {
				lastVal := s.dmlFieldSet.GetFields()[size-1].GetFieldValue()
				k = int(lastVal + 1)
			}
		}

		s.dmlField.FieldValue = int32(k)

		if s.dmlFieldSet != nil {
			s.dmlFieldSet.Fields = append(s.dmlFieldSet.Fields, s.dmlField)
		}
		s.dmlField = nil
	}

}

func (s *TreeListener) ExitFieldType(ctx *parser.FieldTypeContext) {
	if s.dmlField != nil {
		var fieldType string
		var fieldQualifier int32

		for _, child := range ctx.GetChildren() {
			qidctx, ok := child.(*parser.QualifiedIdContext)
			if ok {
				qidVal, ok := s.ctxMap[qidctx]
				if ok {
					if fieldType == "" {
						qualId := qidVal.(string)
						fieldType = qualId
					}

				}
			}
			termNode, ok := child.(*antlr.TerminalNodeImpl)
			if ok {
				nodeText := termNode.GetText()
				if fieldType == "" {
					fieldType = nodeText
				}

			}

			integerCtx, ok := child.(*parser.IntegerConstContext)
			if ok {
				intVal, ok := s.ctxMap[integerCtx]
				if ok {
					k := intVal.(int)
					fieldQualifier = int32(k)
				}
			}

		}

		s.dmlField.FieldType = fieldType
		s.dmlField.FieldTypeQualifier = fieldQualifier
	}

}

func (s *TreeListener) EnterClassElement(ctx *parser.ClassElementContext) {
	s.dmlClass = &dml.DmlClass{}
	s.dmlClass.ClassName = ctx.ID().GetText()
	for _, child := range ctx.AllKEEP_COMMENT() {
		nodeText := child.GetText()
		docString := strings.TrimSpace(nodeText[3:])
		s.dmlClass.Documentation = append(s.dmlClass.Documentation, docString)
	}

}

func (s *TreeListener) ExitClassElement(ctx *parser.ClassElementContext) {
	if s.dmlClass != nil {
		if s.AstRoot != nil {
			s.AstRoot.ClassList = append(s.AstRoot.ClassList, s.dmlClass)
		}
		s.dmlClass = nil
	}
}

func (s *TreeListener) EnterAssociationElement(ctx *parser.AssociationElementContext) {
	s.dmlClass = &dml.DmlClass{}
	s.dmlClass.IsAssociation = true
	s.dmlClass.ClassName = ctx.ID().GetText()
	for _, child := range ctx.AllKEEP_COMMENT() {
		nodeText := child.GetText()
		docString := strings.TrimSpace(nodeText[3:])
		s.dmlClass.Documentation = append(s.dmlClass.Documentation, docString)
	}
}

func (s *TreeListener) ExitAssociationElement(ctx *parser.AssociationElementContext) {
	if s.dmlClass != nil {

		// TODO

		if s.AstRoot != nil {
			s.AstRoot.ClassList = append(s.AstRoot.ClassList, s.dmlClass)
		}
		s.dmlClass = nil
	}
}

func (s *TreeListener) EnterFieldsetOption(ctx *parser.FieldsetOptionContext) {
	termNode := ctx.ID()
	fieldSet := termNode.GetText()
	if s.dmlClass != nil {
		s.dmlClass.FieldsetName = fieldSet
	}

	if s.dmlService != nil {
		s.dmlService.FieldsetName = fieldSet
	}
}

func (s *TreeListener) EnterTableOption(ctx *parser.TableOptionContext) {
	termNode := ctx.ID()
	tableName := termNode.GetText()
	if s.dmlClass != nil {
		s.dmlClass.TableName = tableName
	}

}

func (s *TreeListener) EnterClassField(ctx *parser.ClassFieldContext) {
	s.dmlClassField = &dml.DmlClassField{}

	s.dmlClassField.FieldName = ctx.ID().GetText()

}

func (s *TreeListener) ExitClassField(ctx *parser.ClassFieldContext) {
	if s.dmlClassField != nil {
		fieldModifierCtx := ctx.FieldModifier()
		intVal, ok := s.ctxMap[fieldModifierCtx]
		if ok {
			ival := intVal.(int32)
			s.dmlClassField.Modifier = dml.DmlFieldModifier(ival)
		}

		if s.dmlClass != nil {
			s.dmlClass.Fields = append(s.dmlClass.Fields, s.dmlClassField)
		}

		s.dmlClassField = nil
	}
}

func (s *TreeListener) EnterFieldModifier(ctx *parser.FieldModifierContext) {
	nodeText := ctx.GetText()
	nodeText = strings.ToUpper(nodeText)
	modifierEnum, ok := dml.DmlFieldModifier_value[nodeText]
	if ok {
		s.SetValue(ctx, modifierEnum)
	}

}

func (s *TreeListener) EnterReferencesModifier(ctx *parser.ReferencesModifierContext) {
	allId := ctx.AllID()
	if len(allId) == 2 {
		refClass := allId[0].GetText()
		refField := allId[1].GetText()
		if s.dmlClassField != nil {
			s.dmlClassField.ReferencesClass = refClass
			s.dmlClassField.ReferencesField = refField
		}
	}
}

func (s *TreeListener) EnterIndexElement(ctx *parser.IndexElementContext) {
	s.dmlIndex = &dml.DmlClassIndex{}
	for _, child := range ctx.AllID() {
		nodeText := child.GetText()
		s.dmlIndex.IndexFields = append(s.dmlIndex.IndexFields, nodeText)
	}

}

func (s *TreeListener) ExitIndexElement(ctx *parser.IndexElementContext) {
	if s.dmlIndex != nil {
		if s.dmlClass != nil {
			s.dmlClass.Indexes = append(s.dmlClass.Indexes, s.dmlIndex)
		}

		s.dmlIndex = nil
	}
}

func (s *TreeListener) EnterIndexType(ctx *parser.IndexTypeContext) {
	nodeText := ctx.GetText()
	nodeText = strings.ToUpper(nodeText)
	indexType, ok := dml.DmlIndexType_value[nodeText]
	if ok {
		if s.dmlIndex != nil {
			s.dmlIndex.IndexType = dml.DmlIndexType(indexType)
		}
	}
}

func (s *TreeListener) EnterServiceElement(ctx *parser.ServiceElementContext) {
	s.dmlService = &dml.DmlService{}

	s.dmlService.ServiceName = ctx.ID().GetText()
	for _, child := range ctx.AllKEEP_COMMENT() {
		nodeText := child.GetText()
		docString := strings.TrimSpace(nodeText[3:])
		s.dmlService.Documentation = append(s.dmlService.Documentation, docString)
	}

}

func (s *TreeListener) ExitServiceElement(ctx *parser.ServiceElementContext) {
	if s.dmlService != nil {

		if s.AstRoot != nil {
			s.AstRoot.ServiceList = append(s.AstRoot.ServiceList, s.dmlService)
		}

		s.dmlService = nil
	}

}

func (s *TreeListener) EnterMethod(ctx *parser.MethodContext) {
	s.dmlMethod = &dml.DmlServiceMethod{}

	s.dmlMethod.MethodName = ctx.ID().GetText()
	for _, child := range ctx.AllKEEP_COMMENT() {
		nodeText := child.GetText()
		docString := strings.TrimSpace(nodeText[3:])
		s.dmlMethod.Documentation = append(s.dmlMethod.Documentation, docString)
	}

}

func (s *TreeListener) ExitMethod(ctx *parser.MethodContext) {
	if s.dmlMethod != nil {

		if s.dmlService != nil {
			s.dmlService.Methods = append(s.dmlService.Methods, s.dmlMethod)
		}

		s.dmlMethod = nil
	}
}

func (s *TreeListener) EnterPatternType(ctx *parser.PatternTypeContext) {
	if s.dmlMethod != nil {
		nodeText := ctx.GetText()
		nodeText = strings.ToUpper(nodeText)
		patternType, ok := dml.DmlPatternType_value[nodeText]
		if ok {
			s.dmlMethod.PatternType = dml.DmlPatternType(patternType)
		}
	}
}

func (s *TreeListener) ExitMethodPattern(ctx *parser.MethodPatternContext) {
	if s.dmlMethod != nil {
		qctx := ctx.QualifiedId()
		qualIdItf, ok := s.ctxMap[qctx]
		if ok {
			qualId := qualIdItf.(string)
			index := strings.LastIndex(qualId, ".")
			if index > 0 {
				s.dmlMethod.PatternClassPackage = qualId[:index]
				s.dmlMethod.PatternClass = qualId[index+1:]
			} else {
				s.dmlMethod.PatternClass = qualId
				// need to fix up package later
			}
		}
	}
}

func (s *TreeListener) EnterFieldInstance(ctx *parser.FieldInstanceContext) {
	methodParam := &dml.DmlMethodParam{}
	methodParam.ParameterField = ctx.ID().GetText()
	nodeText := ctx.ParameterModifier().GetText()
	nodeText = strings.ToUpper(nodeText)
	switch nodeText {
	case "REQUIRED":
		methodParam.Modifier = dml.DmlParameterModifier_PARAM_REQUIRED
	case "OPTIONAL":
		methodParam.Modifier = dml.DmlParameterModifier_PARAM_OPTIONAL
	case "REPEATED":
		methodParam.Modifier = dml.DmlParameterModifier_PARAM_REPEATED
	default:
		methodParam.Modifier = dml.DmlParameterModifier_PARAM_UNKNOWN
	}

	// now see if this goes on the input or returns list
	placeOnInput := false
	parent := ctx.GetParent()
	if parent != nil {
		grandparent := parent.GetParent()
		if grandparent != nil {
			_, ok := grandparent.(*parser.InputListContext)
			if ok {
				placeOnInput = true
			}
		}
	}

	if s.dmlMethod != nil {
		if placeOnInput {
			s.dmlMethod.InputFields = append(s.dmlMethod.InputFields, methodParam)
		} else {
			s.dmlMethod.OutputFields = append(s.dmlMethod.OutputFields, methodParam)
		}
	}

}

func CompileAst(infile string) (*dml.DmlTree, error) {
	_, err := os.Stat(infile)
	if err != nil {
		return nil, err
	}

	input, err := antlr.NewFileStream(infile)
	if err != nil {
		return nil, err
	}

	lexer := parser.NewDMLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewDMLParser(stream)
	// p.RemoveErrorListeners()
	errListener := NewDmlErrorListener()
	p.AddErrorListener(errListener)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Model()
	if errListener.NumErrors > 0 {
		return nil, fmt.Errorf("error parsing file: %s", infile)
	}

	fmt.Printf("Compiled %s, numErrors: %d\n", infile, errListener.NumErrors)

	listener := NewTreeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.AstRoot, nil

}
