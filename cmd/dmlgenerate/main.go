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

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gaterace/xdml/pkg/compiler"
	"github.com/gaterace/xdml/pkg/gen"
)

var resequence bool
var help bool
var outdir string
var dbname string
var dbengine string
var language string
var extemplate string

func init() {
	flag.BoolVar(&help, "h", false, "Show dmlgenerate help.")
	flag.BoolVar(&help, "help", false, "Show dmlgenerate help.")
	flag.StringVar(&outdir, "o", ".", "Output directory for generated files.")
	flag.BoolVar(&resequence, "resequence", false, "Resequence protogen files.")
	flag.StringVar(&dbname, "dbname", "", "Database name for generated files.")
	flag.StringVar(&dbengine, "dbengine", "mysql", "Database engine for generated files.")
	flag.StringVar(&language, "language", "java", "Computer language for generated code.")
	flag.StringVar(&extemplate, "extemplate", "", "Use file as external template for generated code.")

}

func showHelp() {
	fmt.Println("Usage: dmlgenerate [flags] <command> <infile>")
	fmt.Println("where <command> is:")
	fmt.Println("\tversion")
	fmt.Println("\tcompile")
	fmt.Println("\tprotogen")
	fmt.Println("\ttablegen")
	fmt.Println("\tprocgen")
	fmt.Println("\tservicegen")
	fmt.Println("and flags are:")
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
}

func main() {
	flag.Parse()
	var command string
	var infile string

	myArgs := flag.Args()

	if len(myArgs) > 0 {
		command = myArgs[0]
		if command == "version" {
			fmt.Println("xdml dmlgenerate - version: 0.9.2")
			fmt.Println("Copyright Demian Harvill 2018")
			return
		}
	}

	if len(myArgs) > 1 {

		infile = myArgs[1]
	} else {
		help = true
	}

	if help {
		showHelp()
		return
	}

	// normalize infile
	if strings.HasSuffix(infile, ".dmlc") {
		infile = infile[:len(infile)-1]
	} else if !strings.HasSuffix(infile, ".dml") {
		infile = infile + ".dml"
	}

	// normalize outdir
	absPath, err := filepath.Abs(outdir)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	stat, err := os.Stat(absPath)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if !stat.IsDir() {
		fmt.Printf("not a directory: %s", absPath)
		return
	}

	outdir = absPath

	var externalTemplate string

	if extemplate != "" {
		buf, err := ioutil.ReadFile(extemplate)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		externalTemplate = string(buf)
	}

	if command == "protogen" {

		astHelper, err := compiler.Load(infile)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		err = gen.ProtoGen(astHelper, outdir, resequence, externalTemplate)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	} else if command == "compile" {

		err := compiler.Compile(infile)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	} else if command == "tablegen" {
		astHelper, err := compiler.Load(infile)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		err = gen.TableGen(astHelper, outdir, dbname, dbengine, externalTemplate)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	} else if command == "procgen" {
		astHelper, err := compiler.Load(infile)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		err = gen.ProcGen(astHelper, outdir, dbname, dbengine, externalTemplate)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	} else if command == "servicegen" {
		astHelper, err := compiler.Load(infile)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		err = gen.ServiceGen(astHelper, outdir, dbname, dbengine, language, externalTemplate)

		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}

}
