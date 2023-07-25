package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/koykov/inspector"
)

var conf inspector.Config

func init() {
	// Check $GOPATH variable.
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		log.Fatal("No GOPATH variable found")
	}

	var bl string
	flag.StringVar(&conf.Package, "pkg", "", "Package path. Should be relative to $GOPATH/src.")
	flag.StringVar(&conf.Directory, "dir", "", "Path to directory contains Go files.")
	flag.StringVar(&conf.File, "file", "", "Path to single Go file.")
	flag.StringVar(&conf.Import, "import", "", "Package import path to use together with -dir/-file options.")
	flag.StringVar(&conf.Destination, "dst", "", "Destination directory path.")
	flag.StringVar(&bl, "bl", "", "Path to blacklist file.")
	flag.BoolVar(&conf.NoClean, "no-clean", false, "Deny to cleanup destination directory.")
	flag.BoolVar(&conf.NoSplit, "no-split", false, "Deny to split output to separate files.")
	flag.StringVar(&conf.XML, "xml", "", "Path to generate XML output.")

	flag.Parse()

	if (len(conf.Directory) > 0 || len(conf.File) > 0) && len(conf.Import) == 0 {
		log.Fatal("Param -imp is required.")
	}

	switch {
	case len(conf.Package) > 0:
		conf.Target = inspector.TargetPackage
		ps := string(os.PathSeparator)
		// Get absolute path to the input package and check it existence.
		absPkg := os.Getenv("GOPATH") + ps + "src" + ps + conf.Package
		_, err := os.Stat(absPkg)
		if os.IsNotExist(err) {
			log.Fatal("pkg doesn't exists: ", conf.Package)
		}
		if len(conf.Destination) == 0 {
			conf.Destination = conf.Package + "_ins"
		}
	case len(conf.Directory) > 0:
		conf.Target = inspector.TargetDirectory
		_, err := os.Stat(conf.Directory)
		if os.IsNotExist(err) {
			log.Fatal("dir doesn't exists: ", conf.Directory)
		}
		if len(conf.Destination) == 0 {
			conf.Destination = conf.Directory + "_ins"
		}
	case len(conf.File) > 0:
		conf.Target = inspector.TargetFile
		if len(conf.Destination) == 0 {
			base := filepath.Base(conf.File)
			path_ := conf.File[:len(conf.File)-len(base)]
			path_ = strings.Trim(path_, "./")
			name := base[:len(base)-len(filepath.Ext(base))]
			conf.Destination = path_ + string(os.PathSeparator) + name + "_ins"
		}
	default:
		log.Fatal("No pkg, dir or file option provided.")
	}

	// Check and read blacklist file.
	if len(bl) > 0 {
		contents, err := os.ReadFile(bl)
		if err != nil {
			log.Fatal(err)
		}
		// Each blacklisted type should be on separate line in file.
		lines := bytes.Split(contents, []byte("\n"))
		for _, line := range lines {
			conf.BlackList[string(line)] = struct{}{}
		}
	}
}

func main() {
	// Prepare output buffer and logger.
	conf.Buf = &bytes.Buffer{}
	conf.Logger = log.New(os.Stdout, "", log.LstdFlags)

	// Initiate the compiler.
	c, err := inspector.NewCompiler(&conf)
	if err != nil {
		log.Fatal(err)
	}
	// Parse and write compiled output to the destination directory.
	var msg string
	switch {
	case len(conf.XML) > 0:
		msg = "Total files generated:"
		err = c.WriteXML()
	default:
		msg = "Total inspectors compiled:"
		err = c.Compile()
	}
	if err != nil {
		log.Fatal("compile failed with error: ", err)
	}

	// Brief report.
	log.Println(msg, c.GetTotal())
	log.Println("Complete")
}
