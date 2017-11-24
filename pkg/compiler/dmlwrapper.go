package compiler

import "fmt"

func CompileTest(infile string, outfile string) error {
	astRoot, err := CompileAst(infile)
	if err != nil {
		return err
	}

	fmt.Printf("packageName: %s\n", astRoot.GetPackageName())

	for _, importName := range astRoot.GetImportedPackages() {
		fmt.Printf("imports: %s\n", importName)
	}

	for _, enumElem := range astRoot.GetEnumList() {
		fmt.Printf("enum name: %s\n", enumElem.GetEnumName())
		for _, doc := range enumElem.GetDocumentation() {
			fmt.Printf("    %s\n", doc)
		}

		for _, enumVal := range enumElem.GetValues() {
			fmt.Printf("  enum value: %s %d\n", enumVal.GetEnumValueName(), enumVal.GetEnumValue())
			for _, doc := range enumVal.GetDocumentation() {
				fmt.Printf("    %s\n", doc)
			}
		}
	}

	for _, fieldSet := range astRoot.GetFieldsetList() {
		fmt.Printf("fieldset: %s extends %s\n", fieldSet.GetFieldsetName(), fieldSet.GetExtendsFieldsetName())
		for _, doc := range fieldSet.GetDocumentation() {
			fmt.Printf("    %s\n", doc)
		}

		for _, field := range fieldSet.GetFields() {
			fmt.Printf("  field: %s, type: %s, qual: %d, number: %d\n", field.GetFieldName(), field.GetFieldType(), field.GetFieldTypeQualifier(), field.GetFieldValue())
			for _, doc := range field.GetDocumentation() {
				fmt.Printf("    %s\n", doc)
			}
		}
	}

	for _, class := range astRoot.GetClassList() {
		fmt.Printf("class: %s, fieldset: %s, table: %s\n", class.ClassName, class.FieldsetName, class.TableName)
		for _, doc := range class.GetDocumentation() {
			fmt.Printf("    %s\n", doc)
		}

		for _, field := range class.GetFields() {
			fmt.Printf("  field: %s, modifier: %s, refclass: %s, reffield: %s\n", field.GetFieldName(), field.GetModifier().String(), field.GetReferencesClass(), field.GetReferencesField())
		}

		for _, index := range class.GetIndexes() {
			fmt.Printf("  index %s\n", index.GetIndexType().String())
			for _, indexField := range index.GetIndexFields() {
				fmt.Printf("    index field: %s\n", indexField)
			}
		}
	}

	for _, service := range astRoot.GetServiceList() {
		fmt.Printf("service: %s, fieldset: %s\n ", service.GetServiceName(), service.GetFieldsetName())
		for _, doc := range service.GetDocumentation() {
			fmt.Printf("    %s\n", doc)
		}

		for _, method := range service.GetMethods() {
			fmt.Printf("  method: %s, patternType: %s, patternClass: %s, patternPackage: %s\n", method.GetMethodName(), method.GetPatternType().String(), method.GetPatternClass(), method.GetPatternClassPackage())
			for _, doc := range method.GetDocumentation() {
				fmt.Printf("      %s\n", doc)
			}

			for _, param := range method.GetInputFields() {
				fmt.Printf("      input field: %s, type: %s\n", param.GetParameterField(), param.GetModifier().String())
			}

			for _, param := range method.GetOutputFields() {
				fmt.Printf("      output field: %s, type: %s\n", param.GetParameterField(), param.GetModifier().String())
			}

		}

	}

	return nil

}
