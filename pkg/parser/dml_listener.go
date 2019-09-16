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

// Generated from DML.g4 by ANTLR 4.6.

package parser // DML

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DMLListener is a complete listener for a parse tree produced by DMLParser.
type DMLListener interface {
	antlr.ParseTreeListener

	// EnterModel is called when entering the model production.
	EnterModel(c *ModelContext)

	// EnterModelParts is called when entering the modelParts production.
	EnterModelParts(c *ModelPartsContext)

	// EnterSourceElements is called when entering the sourceElements production.
	EnterSourceElements(c *SourceElementsContext)

	// EnterSourceElement is called when entering the sourceElement production.
	EnterSourceElement(c *SourceElementContext)

	// EnterEnumElement is called when entering the enumElement production.
	EnterEnumElement(c *EnumElementContext)

	// EnterEnumValues is called when entering the enumValues production.
	EnterEnumValues(c *EnumValuesContext)

	// EnterIntegerConst is called when entering the integerConst production.
	EnterIntegerConst(c *IntegerConstContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// EnterFieldsetElement is called when entering the fieldsetElement production.
	EnterFieldsetElement(c *FieldsetElementContext)

	// EnterExtendsQualifier is called when entering the extendsQualifier production.
	EnterExtendsQualifier(c *ExtendsQualifierContext)

	// EnterFieldsetValues is called when entering the fieldsetValues production.
	EnterFieldsetValues(c *FieldsetValuesContext)

	// EnterFieldsetValue is called when entering the fieldsetValue production.
	EnterFieldsetValue(c *FieldsetValueContext)

	// EnterQualifiedId is called when entering the qualifiedId production.
	EnterQualifiedId(c *QualifiedIdContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterIndexType is called when entering the indexType production.
	EnterIndexType(c *IndexTypeContext)

	// EnterIndexElement is called when entering the indexElement production.
	EnterIndexElement(c *IndexElementContext)

	// EnterClassElement is called when entering the classElement production.
	EnterClassElement(c *ClassElementContext)

	// EnterClassOptions is called when entering the classOptions production.
	EnterClassOptions(c *ClassOptionsContext)

	// EnterClassOption is called when entering the classOption production.
	EnterClassOption(c *ClassOptionContext)

	// EnterFieldsetOption is called when entering the fieldsetOption production.
	EnterFieldsetOption(c *FieldsetOptionContext)

	// EnterTableOption is called when entering the tableOption production.
	EnterTableOption(c *TableOptionContext)

	// EnterMediatesOption is called when entering the mediatesOption production.
	EnterMediatesOption(c *MediatesOptionContext)

	// EnterClassFields is called when entering the classFields production.
	EnterClassFields(c *ClassFieldsContext)

	// EnterReferencesModifier is called when entering the referencesModifier production.
	EnterReferencesModifier(c *ReferencesModifierContext)

	// EnterClassField is called when entering the classField production.
	EnterClassField(c *ClassFieldContext)

	// EnterParameterModifier is called when entering the parameterModifier production.
	EnterParameterModifier(c *ParameterModifierContext)

	// EnterFieldModifier is called when entering the fieldModifier production.
	EnterFieldModifier(c *FieldModifierContext)

	// EnterAssociationElement is called when entering the associationElement production.
	EnterAssociationElement(c *AssociationElementContext)

	// EnterAssociationOptions is called when entering the associationOptions production.
	EnterAssociationOptions(c *AssociationOptionsContext)

	// EnterAssociationOption is called when entering the associationOption production.
	EnterAssociationOption(c *AssociationOptionContext)

	// EnterServiceElement is called when entering the serviceElement production.
	EnterServiceElement(c *ServiceElementContext)

	// EnterPatternType is called when entering the patternType production.
	EnterPatternType(c *PatternTypeContext)

	// EnterMethodPattern is called when entering the methodPattern production.
	EnterMethodPattern(c *MethodPatternContext)

	// EnterMethods is called when entering the methods production.
	EnterMethods(c *MethodsContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterInputList is called when entering the inputList production.
	EnterInputList(c *InputListContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterFieldInstance is called when entering the fieldInstance production.
	EnterFieldInstance(c *FieldInstanceContext)

	// EnterMethodReturn is called when entering the methodReturn production.
	EnterMethodReturn(c *MethodReturnContext)

	// EnterPackageElement is called when entering the packageElement production.
	EnterPackageElement(c *PackageElementContext)

	// EnterImportElements is called when entering the importElements production.
	EnterImportElements(c *ImportElementsContext)

	// ExitModel is called when exiting the model production.
	ExitModel(c *ModelContext)

	// ExitModelParts is called when exiting the modelParts production.
	ExitModelParts(c *ModelPartsContext)

	// ExitSourceElements is called when exiting the sourceElements production.
	ExitSourceElements(c *SourceElementsContext)

	// ExitSourceElement is called when exiting the sourceElement production.
	ExitSourceElement(c *SourceElementContext)

	// ExitEnumElement is called when exiting the enumElement production.
	ExitEnumElement(c *EnumElementContext)

	// ExitEnumValues is called when exiting the enumValues production.
	ExitEnumValues(c *EnumValuesContext)

	// ExitIntegerConst is called when exiting the integerConst production.
	ExitIntegerConst(c *IntegerConstContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)

	// ExitFieldsetElement is called when exiting the fieldsetElement production.
	ExitFieldsetElement(c *FieldsetElementContext)

	// ExitExtendsQualifier is called when exiting the extendsQualifier production.
	ExitExtendsQualifier(c *ExtendsQualifierContext)

	// ExitFieldsetValues is called when exiting the fieldsetValues production.
	ExitFieldsetValues(c *FieldsetValuesContext)

	// ExitFieldsetValue is called when exiting the fieldsetValue production.
	ExitFieldsetValue(c *FieldsetValueContext)

	// ExitQualifiedId is called when exiting the qualifiedId production.
	ExitQualifiedId(c *QualifiedIdContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitIndexType is called when exiting the indexType production.
	ExitIndexType(c *IndexTypeContext)

	// ExitIndexElement is called when exiting the indexElement production.
	ExitIndexElement(c *IndexElementContext)

	// ExitClassElement is called when exiting the classElement production.
	ExitClassElement(c *ClassElementContext)

	// ExitClassOptions is called when exiting the classOptions production.
	ExitClassOptions(c *ClassOptionsContext)

	// ExitClassOption is called when exiting the classOption production.
	ExitClassOption(c *ClassOptionContext)

	// ExitFieldsetOption is called when exiting the fieldsetOption production.
	ExitFieldsetOption(c *FieldsetOptionContext)

	// ExitTableOption is called when exiting the tableOption production.
	ExitTableOption(c *TableOptionContext)

	// ExitMediatesOption is called when exiting the mediatesOption production.
	ExitMediatesOption(c *MediatesOptionContext)

	// ExitClassFields is called when exiting the classFields production.
	ExitClassFields(c *ClassFieldsContext)

	// ExitReferencesModifier is called when exiting the referencesModifier production.
	ExitReferencesModifier(c *ReferencesModifierContext)

	// ExitClassField is called when exiting the classField production.
	ExitClassField(c *ClassFieldContext)

	// ExitParameterModifier is called when exiting the parameterModifier production.
	ExitParameterModifier(c *ParameterModifierContext)

	// ExitFieldModifier is called when exiting the fieldModifier production.
	ExitFieldModifier(c *FieldModifierContext)

	// ExitAssociationElement is called when exiting the associationElement production.
	ExitAssociationElement(c *AssociationElementContext)

	// ExitAssociationOptions is called when exiting the associationOptions production.
	ExitAssociationOptions(c *AssociationOptionsContext)

	// ExitAssociationOption is called when exiting the associationOption production.
	ExitAssociationOption(c *AssociationOptionContext)

	// ExitServiceElement is called when exiting the serviceElement production.
	ExitServiceElement(c *ServiceElementContext)

	// ExitPatternType is called when exiting the patternType production.
	ExitPatternType(c *PatternTypeContext)

	// ExitMethodPattern is called when exiting the methodPattern production.
	ExitMethodPattern(c *MethodPatternContext)

	// ExitMethods is called when exiting the methods production.
	ExitMethods(c *MethodsContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitInputList is called when exiting the inputList production.
	ExitInputList(c *InputListContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitFieldInstance is called when exiting the fieldInstance production.
	ExitFieldInstance(c *FieldInstanceContext)

	// ExitMethodReturn is called when exiting the methodReturn production.
	ExitMethodReturn(c *MethodReturnContext)

	// ExitPackageElement is called when exiting the packageElement production.
	ExitPackageElement(c *PackageElementContext)

	// ExitImportElements is called when exiting the importElements production.
	ExitImportElements(c *ImportElementsContext)
}
