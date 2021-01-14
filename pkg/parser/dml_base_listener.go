// Code generated from DML.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // DML

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDMLListener is a complete listener for a parse tree produced by DMLParser.
type BaseDMLListener struct{}

var _ DMLListener = &BaseDMLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDMLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDMLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDMLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDMLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModel is called when production model is entered.
func (s *BaseDMLListener) EnterModel(ctx *ModelContext) {}

// ExitModel is called when production model is exited.
func (s *BaseDMLListener) ExitModel(ctx *ModelContext) {}

// EnterModelParts is called when production modelParts is entered.
func (s *BaseDMLListener) EnterModelParts(ctx *ModelPartsContext) {}

// ExitModelParts is called when production modelParts is exited.
func (s *BaseDMLListener) ExitModelParts(ctx *ModelPartsContext) {}

// EnterSourceElements is called when production sourceElements is entered.
func (s *BaseDMLListener) EnterSourceElements(ctx *SourceElementsContext) {}

// ExitSourceElements is called when production sourceElements is exited.
func (s *BaseDMLListener) ExitSourceElements(ctx *SourceElementsContext) {}

// EnterSourceElement is called when production sourceElement is entered.
func (s *BaseDMLListener) EnterSourceElement(ctx *SourceElementContext) {}

// ExitSourceElement is called when production sourceElement is exited.
func (s *BaseDMLListener) ExitSourceElement(ctx *SourceElementContext) {}

// EnterEnumElement is called when production enumElement is entered.
func (s *BaseDMLListener) EnterEnumElement(ctx *EnumElementContext) {}

// ExitEnumElement is called when production enumElement is exited.
func (s *BaseDMLListener) ExitEnumElement(ctx *EnumElementContext) {}

// EnterEnumValues is called when production enumValues is entered.
func (s *BaseDMLListener) EnterEnumValues(ctx *EnumValuesContext) {}

// ExitEnumValues is called when production enumValues is exited.
func (s *BaseDMLListener) ExitEnumValues(ctx *EnumValuesContext) {}

// EnterIntegerConst is called when production integerConst is entered.
func (s *BaseDMLListener) EnterIntegerConst(ctx *IntegerConstContext) {}

// ExitIntegerConst is called when production integerConst is exited.
func (s *BaseDMLListener) ExitIntegerConst(ctx *IntegerConstContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseDMLListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseDMLListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterFieldsetElement is called when production fieldsetElement is entered.
func (s *BaseDMLListener) EnterFieldsetElement(ctx *FieldsetElementContext) {}

// ExitFieldsetElement is called when production fieldsetElement is exited.
func (s *BaseDMLListener) ExitFieldsetElement(ctx *FieldsetElementContext) {}

// EnterExtendsQualifier is called when production extendsQualifier is entered.
func (s *BaseDMLListener) EnterExtendsQualifier(ctx *ExtendsQualifierContext) {}

// ExitExtendsQualifier is called when production extendsQualifier is exited.
func (s *BaseDMLListener) ExitExtendsQualifier(ctx *ExtendsQualifierContext) {}

// EnterFieldsetValues is called when production fieldsetValues is entered.
func (s *BaseDMLListener) EnterFieldsetValues(ctx *FieldsetValuesContext) {}

// ExitFieldsetValues is called when production fieldsetValues is exited.
func (s *BaseDMLListener) ExitFieldsetValues(ctx *FieldsetValuesContext) {}

// EnterFieldsetValue is called when production fieldsetValue is entered.
func (s *BaseDMLListener) EnterFieldsetValue(ctx *FieldsetValueContext) {}

// ExitFieldsetValue is called when production fieldsetValue is exited.
func (s *BaseDMLListener) ExitFieldsetValue(ctx *FieldsetValueContext) {}

// EnterQualifiedId is called when production qualifiedId is entered.
func (s *BaseDMLListener) EnterQualifiedId(ctx *QualifiedIdContext) {}

// ExitQualifiedId is called when production qualifiedId is exited.
func (s *BaseDMLListener) ExitQualifiedId(ctx *QualifiedIdContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseDMLListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseDMLListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseDMLListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseDMLListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexElement is called when production indexElement is entered.
func (s *BaseDMLListener) EnterIndexElement(ctx *IndexElementContext) {}

// ExitIndexElement is called when production indexElement is exited.
func (s *BaseDMLListener) ExitIndexElement(ctx *IndexElementContext) {}

// EnterClassElement is called when production classElement is entered.
func (s *BaseDMLListener) EnterClassElement(ctx *ClassElementContext) {}

// ExitClassElement is called when production classElement is exited.
func (s *BaseDMLListener) ExitClassElement(ctx *ClassElementContext) {}

// EnterClassOptions is called when production classOptions is entered.
func (s *BaseDMLListener) EnterClassOptions(ctx *ClassOptionsContext) {}

// ExitClassOptions is called when production classOptions is exited.
func (s *BaseDMLListener) ExitClassOptions(ctx *ClassOptionsContext) {}

// EnterClassOption is called when production classOption is entered.
func (s *BaseDMLListener) EnterClassOption(ctx *ClassOptionContext) {}

// ExitClassOption is called when production classOption is exited.
func (s *BaseDMLListener) ExitClassOption(ctx *ClassOptionContext) {}

// EnterFieldsetOption is called when production fieldsetOption is entered.
func (s *BaseDMLListener) EnterFieldsetOption(ctx *FieldsetOptionContext) {}

// ExitFieldsetOption is called when production fieldsetOption is exited.
func (s *BaseDMLListener) ExitFieldsetOption(ctx *FieldsetOptionContext) {}

// EnterTableOption is called when production tableOption is entered.
func (s *BaseDMLListener) EnterTableOption(ctx *TableOptionContext) {}

// ExitTableOption is called when production tableOption is exited.
func (s *BaseDMLListener) ExitTableOption(ctx *TableOptionContext) {}

// EnterMediatesOption is called when production mediatesOption is entered.
func (s *BaseDMLListener) EnterMediatesOption(ctx *MediatesOptionContext) {}

// ExitMediatesOption is called when production mediatesOption is exited.
func (s *BaseDMLListener) ExitMediatesOption(ctx *MediatesOptionContext) {}

// EnterClassFields is called when production classFields is entered.
func (s *BaseDMLListener) EnterClassFields(ctx *ClassFieldsContext) {}

// ExitClassFields is called when production classFields is exited.
func (s *BaseDMLListener) ExitClassFields(ctx *ClassFieldsContext) {}

// EnterReferencesModifier is called when production referencesModifier is entered.
func (s *BaseDMLListener) EnterReferencesModifier(ctx *ReferencesModifierContext) {}

// ExitReferencesModifier is called when production referencesModifier is exited.
func (s *BaseDMLListener) ExitReferencesModifier(ctx *ReferencesModifierContext) {}

// EnterClassField is called when production classField is entered.
func (s *BaseDMLListener) EnterClassField(ctx *ClassFieldContext) {}

// ExitClassField is called when production classField is exited.
func (s *BaseDMLListener) ExitClassField(ctx *ClassFieldContext) {}

// EnterParameterModifier is called when production parameterModifier is entered.
func (s *BaseDMLListener) EnterParameterModifier(ctx *ParameterModifierContext) {}

// ExitParameterModifier is called when production parameterModifier is exited.
func (s *BaseDMLListener) ExitParameterModifier(ctx *ParameterModifierContext) {}

// EnterFieldModifier is called when production fieldModifier is entered.
func (s *BaseDMLListener) EnterFieldModifier(ctx *FieldModifierContext) {}

// ExitFieldModifier is called when production fieldModifier is exited.
func (s *BaseDMLListener) ExitFieldModifier(ctx *FieldModifierContext) {}

// EnterAssociationElement is called when production associationElement is entered.
func (s *BaseDMLListener) EnterAssociationElement(ctx *AssociationElementContext) {}

// ExitAssociationElement is called when production associationElement is exited.
func (s *BaseDMLListener) ExitAssociationElement(ctx *AssociationElementContext) {}

// EnterAssociationOptions is called when production associationOptions is entered.
func (s *BaseDMLListener) EnterAssociationOptions(ctx *AssociationOptionsContext) {}

// ExitAssociationOptions is called when production associationOptions is exited.
func (s *BaseDMLListener) ExitAssociationOptions(ctx *AssociationOptionsContext) {}

// EnterAssociationOption is called when production associationOption is entered.
func (s *BaseDMLListener) EnterAssociationOption(ctx *AssociationOptionContext) {}

// ExitAssociationOption is called when production associationOption is exited.
func (s *BaseDMLListener) ExitAssociationOption(ctx *AssociationOptionContext) {}

// EnterServiceElement is called when production serviceElement is entered.
func (s *BaseDMLListener) EnterServiceElement(ctx *ServiceElementContext) {}

// ExitServiceElement is called when production serviceElement is exited.
func (s *BaseDMLListener) ExitServiceElement(ctx *ServiceElementContext) {}

// EnterPatternType is called when production patternType is entered.
func (s *BaseDMLListener) EnterPatternType(ctx *PatternTypeContext) {}

// ExitPatternType is called when production patternType is exited.
func (s *BaseDMLListener) ExitPatternType(ctx *PatternTypeContext) {}

// EnterMethodPattern is called when production methodPattern is entered.
func (s *BaseDMLListener) EnterMethodPattern(ctx *MethodPatternContext) {}

// ExitMethodPattern is called when production methodPattern is exited.
func (s *BaseDMLListener) ExitMethodPattern(ctx *MethodPatternContext) {}

// EnterMethods is called when production methods is entered.
func (s *BaseDMLListener) EnterMethods(ctx *MethodsContext) {}

// ExitMethods is called when production methods is exited.
func (s *BaseDMLListener) ExitMethods(ctx *MethodsContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseDMLListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseDMLListener) ExitMethod(ctx *MethodContext) {}

// EnterInputList is called when production inputList is entered.
func (s *BaseDMLListener) EnterInputList(ctx *InputListContext) {}

// ExitInputList is called when production inputList is exited.
func (s *BaseDMLListener) ExitInputList(ctx *InputListContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseDMLListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseDMLListener) ExitFieldList(ctx *FieldListContext) {}

// EnterFieldInstance is called when production fieldInstance is entered.
func (s *BaseDMLListener) EnterFieldInstance(ctx *FieldInstanceContext) {}

// ExitFieldInstance is called when production fieldInstance is exited.
func (s *BaseDMLListener) ExitFieldInstance(ctx *FieldInstanceContext) {}

// EnterMethodReturn is called when production methodReturn is entered.
func (s *BaseDMLListener) EnterMethodReturn(ctx *MethodReturnContext) {}

// ExitMethodReturn is called when production methodReturn is exited.
func (s *BaseDMLListener) ExitMethodReturn(ctx *MethodReturnContext) {}

// EnterPackageElement is called when production packageElement is entered.
func (s *BaseDMLListener) EnterPackageElement(ctx *PackageElementContext) {}

// ExitPackageElement is called when production packageElement is exited.
func (s *BaseDMLListener) ExitPackageElement(ctx *PackageElementContext) {}

// EnterImportElements is called when production importElements is entered.
func (s *BaseDMLListener) EnterImportElements(ctx *ImportElementsContext) {}

// ExitImportElements is called when production importElements is exited.
func (s *BaseDMLListener) ExitImportElements(ctx *ImportElementsContext) {}
