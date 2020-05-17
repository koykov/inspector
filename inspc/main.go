package main

import (
	"bytes"
	"flag"
	"log"
	"os"

	"github.com/koykov/inspector"
)

var (
	fPkg     = flag.String("pkg", "", "Package path. May be relative to GOPATH.")
	fOut     = flag.String("out", "", `Output dir. pkg + "_ins" by default.`)
	pkg, out string
	absPkg   string
)

func init() {
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		log.Fatal("No GOPATH variable found")
	}

	flag.Parse()
	pkg = *fPkg
	out = *fOut

	if len(pkg) == 0 {
		log.Fatal("pkg option is required")
	}

	ps := string(os.PathSeparator)
	absPkg = os.Getenv("GOPATH") + ps + "src" + ps + pkg
	_, err := os.Stat(absPkg)
	if os.IsNotExist(err) {
		log.Fatal("pkg doesn't exists: ", pkg)
	}

	if len(out) == 0 {
		out = pkg + "_ins"
	}
}

func main() {
	buf := &bytes.Buffer{}
	lg := log.New(os.Stdout, "", log.LstdFlags)
	c := inspector.NewCompiler(pkg, out, buf, lg)
	err := c.Compile()
	if err != nil {
		log.Fatal("compile failed with error: ", err)
	}
	log.Println("Total inspectors compiled:", c.GetTotal())
	log.Println("Complete")
}
