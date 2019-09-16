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
	"fmt"
	"os"

	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gaterace/xdml/pkg/dml"
	"github.com/golang/protobuf/proto"
)

var PackageDmlMap map[string]*DmlHelper

var SearchPaths []string

var BuiltinTypes = [...]string{"double", "float", "int32", "int64", "uint32", "uint64", "bool", "string", "bytes", "datetime", "guid", "decimal", "bytearray", "chararray"}

var BuiltinMap map[string]bool

type DmlHelper struct {
	AstRoot     *dml.DmlTree
	BaseName    string
	FileName    string
	Dirty       bool
	EnumMap     map[string]*EnumHelper
	ClassMap    map[string]*ClassHelper
	FieldsetMap map[string]*FieldsetHelper
	ServiceMap  map[string]*ServiceHelper
}

type EnumHelper struct {
	Enum     *dml.DmlEnum
	ValueMap map[string]*dml.DmlEnumValue
}

type ClassHelper struct {
	Class    *dml.DmlClass
	FieldMap map[string]*dml.DmlClassField
}

type FieldsetHelper struct {
	Fieldset *dml.DmlFieldSet
	Extends  *FieldsetHelper
	FieldMap map[string]*dml.DmlField
}

type ServiceHelper struct {
	Service   *dml.DmlService
	MethodMap map[string]*dml.DmlServiceMethod
}

func init() {
	PackageDmlMap = make(map[string]*DmlHelper)
	SearchPaths = make([]string, 0, 10)
	BuiltinMap = make(map[string]bool)
	for _, builtin := range BuiltinTypes {
		BuiltinMap[builtin] = true
	}

}

func NewDmlHelper() *DmlHelper {
	helper := DmlHelper{}
	helper.EnumMap = make(map[string]*EnumHelper)
	helper.ClassMap = make(map[string]*ClassHelper)
	helper.FieldsetMap = make(map[string]*FieldsetHelper)
	helper.ServiceMap = make(map[string]*ServiceHelper)
	return &helper
}

func NewEnumHelper(proxy *dml.DmlEnum) *EnumHelper {
	helper := EnumHelper{}
	helper.Enum = proxy
	helper.ValueMap = make(map[string]*dml.DmlEnumValue)
	return &helper
}

func NewClassHelper(proxy *dml.DmlClass) *ClassHelper {
	helper := ClassHelper{}
	helper.Class = proxy
	helper.FieldMap = make(map[string]*dml.DmlClassField)
	return &helper
}

func NewFieldsetHelper(proxy *dml.DmlFieldSet) *FieldsetHelper {
	helper := FieldsetHelper{}
	helper.Fieldset = proxy
	helper.FieldMap = make(map[string]*dml.DmlField)
	return &helper
}

func NewServiceHelper(proxy *dml.DmlService) *ServiceHelper {
	helper := ServiceHelper{}
	helper.Service = proxy
	helper.MethodMap = make(map[string]*dml.DmlServiceMethod)
	return &helper
}

func (s *DmlHelper) Compile(infile string) error {
	astRoot, err := CompileAst(infile)
	if err != nil {
		return err
	}
	s.AstRoot = astRoot
	baseName := filepath.Base(infile)
	if strings.HasSuffix(baseName, ".dml") {
		baseName = baseName[0 : len(baseName)-4]
	}
	s.BaseName = baseName
	s.FileName = infile
	s.Dirty = true

	PackageDmlMap[s.BaseName] = s

	err = s.Build()
	if err != nil {
		return err
	}

	err = s.Validate()

	return err
}

func (s *DmlHelper) Load(infile string) error {
	if strings.HasSuffix(infile, ".dml") {
		infile = infile + "c"
	}
	dmlcStat, err := os.Stat(infile)
	if err != nil {
		return err
	}

	if strings.HasSuffix(infile, ".dmlc") {
		sourceFile := infile[0 : len(infile)-1]
		// fmt.Printf("infile: %s\n", infile)
		// fmt.Printf("sourceFile: %s\n", sourceFile)
		dmlStat, err := os.Stat(sourceFile)
		if err == nil {
			if dmlStat.ModTime().After(dmlcStat.ModTime()) {
				// need to recompile
				// fmt.Println("need to recompile")
				compileHelper := NewDmlHelper()
				err = compileHelper.Compile(sourceFile)
				if err != nil {
					return err
				}
				err = compileHelper.Save()
				if err != nil {
					return err
				}
			}

		}
	}

	ast := &dml.DmlTree{}

	data, err := ioutil.ReadFile(infile)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(data, ast)
	if err != nil {
		return err
	}

	baseName := filepath.Base(infile)
	if strings.HasSuffix(baseName, ".dmlc") {
		baseName = baseName[0 : len(baseName)-5]
	}

	s.AstRoot = ast
	s.BaseName = baseName
	s.FileName = infile
	s.Dirty = false

	PackageDmlMap[s.BaseName] = s

	err = s.Build()

	return err
}

func (s *DmlHelper) Build() error {
	if s.AstRoot == nil {
		return fmt.Errorf("no ast root for %s", s.BaseName)
	}

	buildSuccesful := true

	// make sure we have all the packages inPackageDmlMap
	for _, packageName := range s.AstRoot.GetImportedPackages() {
		_, ok := PackageDmlMap[packageName]
		// fmt.Printf("looking for: %s\n", packageName)
		if !ok {
			packagePath, err := FindDmlPackage(packageName)
			if err != nil {
				return err
			}

			// fmt.Printf("package path: %s\n", packagePath)

			importHelper := NewDmlHelper()

			err = importHelper.Load(packagePath)
			if err != nil {
				return err
			}

		}
	}

	// build the enums
	for _, enumChild := range s.AstRoot.GetEnumList() {
		enumName := enumChild.GetEnumName()
		_, ok := s.EnumMap[enumName]
		if ok {
			fmt.Printf("duplicate enum name: %s\n", enumName)
			buildSuccesful = false
		} else {
			enumHelper := NewEnumHelper(enumChild)
			s.EnumMap[enumName] = enumHelper
			for _, value := range enumChild.GetValues() {
				valueName := value.GetEnumValueName()
				_, found := enumHelper.ValueMap[valueName]
				if found {
					fmt.Printf("duplicate enum value %s\n", valueName)
					buildSuccesful = false
				} else {
					enumHelper.ValueMap[valueName] = value
				}

			}
		}
	}

	// build the classes
	for _, classChild := range s.AstRoot.GetClassList() {
		className := classChild.ClassName
		_, ok := s.ClassMap[className]
		if ok {
			fmt.Printf("duplicate class name: %s\n", className)
			buildSuccesful = false
		} else {
			classHelper := NewClassHelper(classChild)
			s.ClassMap[className] = classHelper
			for _, field := range classChild.GetFields() {
				fieldName := field.GetFieldName()
				_, found := classHelper.FieldMap[fieldName]
				if found {
					fmt.Printf("duplicate class field %s\n", fieldName)
					buildSuccesful = false
				} else {
					classHelper.FieldMap[fieldName] = field
				}
			}
		}
	}

	// build the fieldsets
	for _, setChild := range s.AstRoot.GetFieldsetList() {
		setName := setChild.GetFieldsetName()
		_, ok := s.FieldsetMap[setName]
		if ok {
			fmt.Printf("duplicate fieldset name: %s\n", setName)
			buildSuccesful = false
		} else {
			setHelper := NewFieldsetHelper(setChild)
			s.FieldsetMap[setName] = setHelper
			for _, field := range setChild.GetFields() {
				fieldName := field.GetFieldName()
				_, found := setHelper.FieldMap[fieldName]
				if found {
					fmt.Printf("duplicate fieldset field %s\n", fieldName)
					buildSuccesful = false
				} else {
					setHelper.FieldMap[fieldName] = field
				}
			}

			extendsFieldset := setChild.GetExtendsFieldsetName()
			if extendsFieldset != "" {
				// need to find the base fieldset
				foundExtends := false
				extendsHelper, ok := s.FieldsetMap[extendsFieldset]
				if ok {
					foundExtends = true
					setHelper.Extends = extendsHelper
				} else {
					// look in imports
					for _, importName := range s.AstRoot.GetImportedPackages() {

						importHelper, ok := PackageDmlMap[importName]
						if ok {
							extendsHelper, found := importHelper.FieldsetMap[extendsFieldset]
							if found {
								setHelper.Extends = extendsHelper
								foundExtends = true
								break
							}
						}
					}
				}
				if !foundExtends {
					fmt.Printf("dcould not find extends fieldset: %s\n", extendsFieldset)
					buildSuccesful = false
				}
			}
		}
	}

	// build the services
	for _, serviceChild := range s.AstRoot.GetServiceList() {
		serviceName := serviceChild.GetServiceName()
		_, ok := s.ServiceMap[serviceName]
		if ok {
			fmt.Printf("duplicate service name: %s\n", serviceName)
			buildSuccesful = false
		} else {
			serviceHelper := NewServiceHelper(serviceChild)
			s.ServiceMap[serviceName] = serviceHelper
			for _, method := range serviceChild.GetMethods() {
				methodName := method.GetMethodName()
				_, found := serviceHelper.MethodMap[methodName]
				if found {
					fmt.Printf("duplicate service method %s\n", methodName)
					buildSuccesful = false
				} else {
					serviceHelper.MethodMap[methodName] = method
				}
			}
		}
	}

	if buildSuccesful {
		return nil
	}

	return fmt.Errorf("build of %s encountered errors", s.BaseName)
}

func GetClassInPackage(packageName string, className string) (*ClassHelper, string) {

	for _, packageHelper := range PackageDmlMap {
		if packageHelper.AstRoot.GetPackageName() == packageName {
			classHelper, ok := packageHelper.ClassMap[className]
			if ok {
				// fmt.Printf("found %s in %s\n", className, baseName)
				return classHelper, packageHelper.BaseName
			}
		}

	}

	return nil, ""
}

func GetEnumInPackage(packageName string, enumName string) *EnumHelper {

	for _, packageHelper := range PackageDmlMap {
		if packageHelper.AstRoot.GetPackageName() == packageName {
			enumHelper, ok := packageHelper.EnumMap[enumName]
			if ok {
				// fmt.Printf("found %s in %s\n", className, baseName)
				return enumHelper
			}
		}

	}

	return nil
}

func (s *DmlHelper) GetPackageForFieldType(fieldType string) string {
	if fieldType == "" {
		return ""
	}

	_, foundBuiltin := BuiltinMap[fieldType]
	if foundBuiltin {
		return "builtin"
	}

	ourPackage := s.AstRoot.GetPackageName()

	_, foundEnum := s.EnumMap[fieldType]
	if foundEnum {
		return ourPackage
	}
	_, foundClass := s.ClassMap[fieldType]
	if foundClass {
		return ourPackage
	}

	// not in our package, look at imports
	for _, importName := range s.AstRoot.GetImportedPackages() {
		helper, ok := PackageDmlMap[importName]
		if ok {
			packageName := helper.GetPackageForFieldType(fieldType)
			if packageName != "" {
				return packageName
			}
		}
	}

	return ""
}

func (s *DmlHelper) GetField(fieldsetName string, fieldName string) *dml.DmlField {
	setHelper, ok := s.FieldsetMap[fieldsetName]
	if !ok {
		return nil
	}

	for setHelper != nil {
		field, ok := setHelper.FieldMap[fieldName]
		if ok {
			return field
		}

		setHelper = setHelper.Extends
	}

	return nil
}

func (s *DmlHelper) ValidParamName(fieldSetName string, paramName string) bool {
	if (fieldSetName == "") || (paramName == "") {
		return false
	}

	setHelper, ok := s.FieldsetMap[fieldSetName]
	if !ok {
		return false
	}

	for setHelper != nil {
		_, ok = setHelper.FieldMap[paramName]
		if ok {
			return true
		} else {
			setHelper = setHelper.Extends
		}
	}

	return false
}

func (s *DmlHelper) Validate() error {
	if s.AstRoot == nil {
		return fmt.Errorf("no ast root for %s", s.BaseName)
	}

	validateSuccesful := true

	// validate fieldsets
	for _, setChild := range s.AstRoot.GetFieldsetList() {
		idMap := make(map[int32]int32)
		for _, setField := range setChild.GetFields() {
			fieldName := setField.GetFieldName()
			fieldVal := setField.GetFieldValue()

			_, found := idMap[fieldVal]
			if found {
				fmt.Printf("duplicate fieldset field value, %s\n", fieldName)
				validateSuccesful = false
			} else {
				idMap[fieldVal] = fieldVal
			}

			packageName := s.GetPackageForFieldType(setField.GetFieldType())
			if packageName == "" {
				fmt.Printf("cannot determine package for fieldset field, %s\n", fieldName)
				validateSuccesful = false
			} else {
				setField.FieldTypePackage = packageName
			}

		}
	}

	// validate classes
	for _, classChild := range s.AstRoot.GetClassList() {
		className := classChild.GetClassName()
		fieldSetName := classChild.GetFieldsetName()
		setHelper, ok := s.FieldsetMap[fieldSetName]
		if ok {
			for _, fieldChild := range classChild.GetFields() {
				fieldName := fieldChild.GetFieldName()
				_, found := setHelper.FieldMap[fieldName]
				if !found {
					fmt.Printf("cannot find fieldset field for class field, %s\n", fieldName)
					validateSuccesful = false
				}
			}

		} else {
			fmt.Printf("cannot find fieldset for class, %s\n", className)
			validateSuccesful = false
		}
	}

	// validate services
	for _, serviceChild := range s.AstRoot.GetServiceList() {
		serviceName := serviceChild.GetServiceName()
		fieldSetName := serviceChild.GetFieldsetName()
		_, ok := s.FieldsetMap[fieldSetName]
		if ok {
			for _, methodChild := range serviceChild.GetMethods() {
				for _, inputChild := range methodChild.GetInputFields() {
					if !s.ValidParamName(fieldSetName, inputChild.GetParameterField()) {
						fmt.Printf("cannot find parameter map for input field, %s\n", inputChild.GetParameterField())
						validateSuccesful = false
					}
				}
				for _, outputChild := range methodChild.GetOutputFields() {
					if !s.ValidParamName(fieldSetName, outputChild.GetParameterField()) {
						fmt.Printf("cannot find parameter map for output field, %s\n", outputChild.GetParameterField())
						validateSuccesful = false
					}
				}
			}

		} else {
			fmt.Printf("cannot find fieldset for service, %s\n", serviceName)
			validateSuccesful = false
		}

	}

	if validateSuccesful {
		return nil
	}

	return fmt.Errorf("validate  of %s encountered errors", s.BaseName)
}

func (s *DmlHelper) Save() error {
	if s.AstRoot == nil {
		// throw error
	}

	if s.Dirty {
		if strings.HasSuffix(s.FileName, ".dml") {
			outfile := s.FileName + "c"
			out, err := proto.Marshal(s.AstRoot)
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(outfile, out, 0644); err != nil {
				return err
			}

		}

		s.Dirty = false
	}

	return nil
}

func FindDmlPackage(importName string) (string, error) {
	for _, pathpart := range SearchPaths {
		fullPath := pathpart + string(os.PathSeparator) + importName + ".dmlc"
		// fmt.Printf("fullPath: %s\n", fullPath)
		_, err := os.Stat(fullPath)
		if err == nil {
			return fullPath, nil
		}

	}
	return "", fmt.Errorf("import not found: %s", importName)
}

func AddSearchPath(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("not a directory: %s", path)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	//  fmt.Printf("absPath: %s\n", absPath)

	SearchPaths = append(SearchPaths, absPath)
	return nil
}

func Compile(infile string) error {

	AddSearchPath(".")
	// fmt.Printf("len SearchPaths: %d\n", len(SearchPaths))
	if len(SearchPaths) == 0 {

	}

	helper := NewDmlHelper()
	absPath, err := filepath.Abs(infile)
	if err != nil {
		return err
	}

	err = helper.Compile(absPath)

	if err != nil {
		return err
	}

	// fmt.Printf("helper.BaseName: %s\n", helper.BaseName)

	// err = helper.Build()

	if err != nil {
		return err
	}

	//for key, val := range PackageDmlMap {
	// fmt.Printf("base: %s, file: %s, package: %s\n", key, val.FileName, val.AstRoot.GetPackageName())
	// }

	err = helper.Save()

	return err
}

func Load(infile string) (*DmlHelper, error) {
	AddSearchPath(".")

	helper := NewDmlHelper()
	absPath, err := filepath.Abs(infile)
	if err != nil {
		return nil, err
	}

	err = helper.Load(absPath)

	if err != nil {
		return nil, err
	}

	return helper, nil
}
