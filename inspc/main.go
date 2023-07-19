package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/koykov/inspector"
)

var (
	// Command line arguments.
	fPkg  = flag.String("pkg", "", "Package path. Should be relative to $GOPATH/src.")
	fDir  = flag.String("dir", "", "Path to directory contains Go files.")
	fFile = flag.String("file", "", "Path to Go file.")
	fImp  = flag.String("import", "", "Package import path to use together with -dir/-file options.")
	fDst  = flag.String("dst", "", `Destination dir. pkg + "_ins" by default.`)
	fBl   = flag.String("bl", "", "Path to blacklist file.")
	fXML  = flag.String("xml", "", "Path to generate XML output.")
	// Dereferenced arguments.
	pkg, dir, file, imp, src, dst string
	xml                           bool

	absPkg string
	target inspector.Target
	bl     = map[string]bool{}
)

func init() {
	// Check $GOPATH variable.
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		log.Fatal("No GOPATH variable found")
	}

	flag.Parse()
	pkg = *fPkg
	dir = *fDir
	file = *fFile
	imp = *fImp
	dst = *fDst
	if xml = len(*fXML) > 0; xml {
		dst = *fXML
	}

	if (len(dir) > 0 || len(file) > 0) && len(imp) == 0 {
		log.Fatal("Param -imp is required.")
	}

	switch {
	case len(pkg) > 0:
		target = inspector.TargetPackage
		src = pkg
		ps := string(os.PathSeparator)
		// Get absolute path to the input package and check it existence.
		absPkg = os.Getenv("GOPATH") + ps + "src" + ps + pkg
		_, err := os.Stat(absPkg)
		if os.IsNotExist(err) {
			log.Fatal("pkg doesn't exists: ", pkg)
		}
		if len(dst) == 0 {
			dst = pkg + "_ins"
		}
	case len(dir) > 0:
		target = inspector.TargetDirectory
		src = dir
		_, err := os.Stat(absPkg)
		if os.IsNotExist(err) {
			log.Fatal("dir doesn't exists: ", pkg)
		}
		if len(dst) == 0 {
			dst = dir + "_ins"
		}
	case len(file) > 0:
		target = inspector.TargetFile
		src = file
		if len(dst) == 0 {
			base := filepath.Base(file)
			path_ := file[:len(file)-len(base)]
			name := base[:len(base)-len(filepath.Ext(base))]
			dst = path_ + string(os.PathSeparator) + name + "_ins.go"
		}
		// todo implement me
	default:
		log.Fatal("No pkg, dir or file option provided.")
	}

	// Check and read blacklist file.
	if len(*fBl) > 0 {
		contents, err := os.ReadFile(*fBl)
		if err != nil {
			log.Fatal(err)
		}
		// Each blacklisted type should be on separate line in file.
		lines := bytes.Split(contents, []byte("\n"))
		for _, line := range lines {
			bl[string(line)] = true
		}
	}
}

func main() {
	// Prepare output buffer and logger.
	buf := &bytes.Buffer{}
	lg := log.New(os.Stdout, "", log.LstdFlags)

	// Initiate the compiler.
	c := inspector.NewCompiler(target, src, dst, imp, bl, buf, lg)
	// Parse and write compiled output to the destination directory.
	var (
		err error
		msg string
	)
	switch {
	case xml:
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
