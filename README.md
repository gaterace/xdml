# eXtended Data Modeling Language (xdml)

Copyright 2017-2019 Demian Harvill

## Overview

This provides and implements a domain-specific language to describe a data dictionary of fields
and composite objects (structures, classes, enumerations, services and methods). The basic goals are 

* Cross-platform
* Computer language independent
* RDBMS independent
* Document each field and composite object once (Don't Repeat Yourself or DRY)
* Encourage consistent naming across a system
* Automate many of the software engineering tasks by automatically generating artifacts


As this grew out of a desire to automate many of the day-to-day tasks at my day job and on
special projects, it has some opinionated aspects in its current incarnation:

* Used Antlr4 to describe the language
* Written in Go
* Use of Google Protocol Buffers (protobuf) for serialization (and platform/language independence)
* Grpc for services
* Source DML code is compiled to a protobuf-based intermediate form
* Artifacts are generated from intermediate, compiled form
* Initial focus was on Java and MySql-compatible systems, though C# and SqlServer support is partially implemented
* Artifacts aree descibed with Go text/template instances and are thus extensible

The current artifacts that can be generated are:

* protobuf files
* gRPC services
* database table and index definitions
* stored procedures
* application binding code

## Building

Not wanting to commit a large vendor directory, please use golang/dep  

```
dep ensure 
```

to build the vendor directory from Gopkg.toml

## dmlgenerate

The command line tool in cmd/dmlgenerate/dmgenerate is the main utility.  Run without any parameters to see help.

## Examples

There are sample .dml source files in the examples/ directory

BrilligDomain.dml describes the domain objects for an application

BrilligService.dml  includes BrilligDomain adn describes a service with methods for an application.

## The XDML Domain Specific Language

The Antlr4 grammar for the language is defined in antlr/DML.g4

Each source file is a text file with a filename in the form <BasePackageName>.dml
Comments start with a double slash (//) and are ignored.
Comments that start with a triple slash (///) are significant and used to document whatever item follows.
The first non-blank, non-comment line should give the Java package name, as in:

```
package org.gaterace.brillig.domain;
```

In this version, the C# package and Go package are derived from this.
After this, other source xdml packages can be imported by their <BasePackageName>, ie:
import “BrilligDomain”;

Typically, the next section is a fieldset section in the form:

```
/// fieldset documentation
fieldset <fieldSetName> [extends <parentFieldSetName>]{
    // field definitions and documentation go here
}
```

Where the extends clause is optional, and implies that this field set is a continuation of the <parentFieldSetName>.  The field definitions are of the form:

/// field documentation, can use multiple lines
<field_type> <field_name> [<modifier>] [= <field_number>];

Where <modifier> and = <field_number> are optional. The builtin field types (and associated modifiers) are:

*	double
*	float
*	int32
*	int64
*	uint32
*	uint64
*	bool
*	string(<max_string_length>)
*	bytes(<max_byte_array length>) 
*	chararray(<fixed_string_length>)
*	bytearray(<fixed_byte_array_length>)
*	datetime
*	guid
*	decimal

The field type can also be a Class name or Enum name as described later. If the = <field_number> is not given, then the implied field_number is one greater than the last field number. This can be used in the generation of protobuf definitions. By convention, field names are lower case with underscores; they will be converted to camel case or other patterns as needed.

An Enum can be specified as

```
/// enum documentation
enum <enumName> {
    /// documentation for enum value
    <enum_value_name> = <enum_int>;

    // additional enum value definitions and documentation go here
}
```

By convention, <enum_value_name> is upper case with underscores, as in:

```
 MAKER_EVENT_UNKNOWN = 0;
 ```

A class or struct is defined as:

```
/// class documentation
class <className> fieldset(<fieldSetName>) [table (<tableName)}{
    // class field definitions go here, field documentation is automatic
    <class_field_modifier> <field_name>; 

    // if this class references a SQL table, index definitions go here
}
```

Where the field_name is described in the fieldSet (or one that it is derived from). By convention, the <className> is camel case (as in UserRole) , with the first character upper case.  The valid <class_field_modifiers> and semantics are:

*	required : class field is required (NOT NULL for SQL)
*	optional: class field is optional (NULL for SQL)
*	repeated: class field occurs multiple times, possibly zero times
*	auto: certain fields (created, modified, deleted, is_deleted, version) have predefined semantics
*	id: class field is used to identify the class instance
*	idgen:  class field is used to identify the class instance and will be generated automatically for SQL
*	virtual: class field is not really a part of the class but will be supplied virtually
*	map: class field is a one-to-many map to instances of another class
*	list: class field is an embedded list of instances of another class

If the table(tableName>) option is used, SQL indices can be defined.  Some examples are:

```
index primary (group_id);
index unique (account_id, group_name);
index nonunique (account_name);
```

If the table(tableName>) option is not used, then this class or struct will not be backed by a SQL table and just used for protobuf messages. 

To define a service, use:

```
/// service documentation
service <serviceName> fieldset(<fieldSetName>) {
    // method definitions and documentation go here   
}
```

By convention, <serviceName>, like <className> is camel case with first letter uppercase, as in BrilligSvc.

A method definition looks like:

```
/// method documentation
method <method_name> [pattern <proc_type> <qualified_class_name]{    
    <paramModifier> <field_name>;
    // repeated as necessary     
}
returns {
    <paramModifier> <field_name>;
    // repeated as necessary
}
```


Again, by convention, <method_name> is lowercase with underscores (like <fieldname>), as in create_account. If the optional pattern clause is given, this will be used to generate SQL stored procedures. The valid <proc_type> values are:

*	select
*	insert
*	update
*	delete

And the <qualified_class_name> is the <dmlPackageName>.<className> if the class is defined in a separate included dml package, or simply <className> if defined in this package.

The valid <paramModifier> values are a subset of the <classFieldModifier> values seen before:

*	required
*	optional
*	repeated


