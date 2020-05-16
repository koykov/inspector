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
	fPkg           = flag.String("pkg", "", "Package path. May be relative to GOPATH.")
	fOut           = flag.String("out", "", `Output dir. pkg + "_ins" by default.`)
	pkg, out       string
	absPkg, absOut string
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
	absOut = absPkg + "_ins"
	_, err = os.Stat(absOut)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(absOut, 0755); err != nil {
			log.Fatal("Couldn't prepare output dir: ", err)
		}
	}
}

func main() {
	buf := &bytes.Buffer{}
	lg := log.New(os.Stdout, "", log.LstdFlags)
	c := inspector.NewCompiler(pkg, buf, lg)
	err := c.Compile()
	if err != nil {
		log.Fatal("compile failed with error: ", err)
	}
	log.Println("Total inspectors compiled:", c.GetTotal())

	outFile := absOut + string(os.PathSeparator) + "ins001.go"
	_, err = os.Stat(absOut)
	if err != nil && !os.IsNotExist(err) {
		if err := os.Remove(outFile); err != nil {
			log.Fatal("Couldn't prepare out file: ", err)
		}
	}
	err = ioutil.WriteFile(outFile, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal("Couldn't write to file: ", err)
	}
	log.Println("Complete")
}