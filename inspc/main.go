package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/koykov/inspector"
)

var (
	// Command line arguments.
	fPkg = flag.String("pkg", "", "Package path. Should be relative to $GOPATH/src.")
	fDst = flag.String("dst", "", `Destination dir. pkg + "_ins" by default.`)
	fBl  = flag.String("bl", "", "Path to blacklist file.")
	// Dereferenced arguments.
	pkg, dst string
	absPkg   string
	bl       = map[string]bool{}
)

func init() {
	// Check $GOPATH variable.
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		log.Fatal("No GOPATH variable found")
	}

	flag.Parse()
	pkg = *fPkg
	dst = *fDst

	if len(pkg) == 0 {
		log.Fatal("pkg option is required")
	}

	ps := string(os.PathSeparator)
	// Get absolute path to the input package and check it existence.
	absPkg = os.Getenv("GOPATH") + ps + "src" + ps + pkg
	_, err := os.Stat(absPkg)
	if os.IsNotExist(err) {
		log.Fatal("pkg doesn't exists: ", pkg)
	}

	// Prepare destination dir.
	if len(dst) == 0 {
		dst = pkg + "_ins"
	}

	// Check and read blacklist file.
	if len(*fBl) > 0 {
		contents, err := ioutil.ReadFile(*fBl)
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
	c := inspector.NewCompiler(pkg, dst, bl, buf, lg)
	// Parse and write compiled output to the destination directory.
	err := c.Compile()
	if err != nil {
		log.Fatal("compile failed with error: ", err)
	}

	// Brief report.
	log.Println("Total inspectors compiled:", c.GetTotal())
	log.Println("Complete")
}
