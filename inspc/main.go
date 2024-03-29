package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/koykov/inspector"
	"github.com/koykov/multiflag"
)

var conf inspector.Config

func init() {
	// Check $GOPATH variable.
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		log.Fatal("No GOPATH variable found")
	}

	var bl string
	multiflag.StringsVar(&conf.Package, []string{"p", "pkg", "package"}, "", "Path to Go `package`. Should be relative to $GOPATH/src.")
	multiflag.StringsVar(&conf.Directory, []string{"d", "dir", "directory"}, "", "Path to `directory` contains Go files.")
	multiflag.StringsVar(&conf.File, []string{"f", "file"}, "", "Path to single Go `file`.")
	multiflag.StringsVar(&conf.Import, []string{"i", "imp", "import"}, "", "Package import `path` to use together with -directory/-file options.")
	multiflag.StringsVar(&conf.Destination, []string{"dst", "destination"}, "", "Destination directory `path`.")
	multiflag.StringsVar(&bl, []string{"b", "bl", "black-list"}, "", "Path to blacklist `file`.")
	multiflag.BoolsVar(&conf.NoClean, []string{"nc", "no-clean"}, false, "Deny to cleanup destination directory.")
	multiflag.BoolsVar(&conf.NoSplit, []string{"ns", "no-split"}, false, "Deny to split output to separate files (reserved).")
	multiflag.BoolsVar(&conf.Force, []string{"force"}, false, "Continue even if error occurred.")
	multiflag.StringsVar(&conf.XML, []string{"x", "xml"}, "", "Debug XML data destination `path`.")
	multiflag.Parse()

	rootDir := func(path string) (string, bool, error) {
		if path == "." || path == "./" {
			wd, err := os.Getwd()
			if err != nil {
				return "", false, err
			}
			return wd, true, nil
		}
		return path, false, nil
	}

	root, ok, err := rootDir(conf.Destination)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		conf.Destination = root
		conf.NoClean = true
	}

	switch {
	case len(conf.Package) > 0:
		conf.Target = inspector.TargetPackage
		ps := string(os.PathSeparator)
		// Get absolute path to the input package and check it existence.
		absPkg := os.Getenv("GOPATH") + ps + "src" + ps + conf.Package
		if _, err = os.Stat(absPkg); os.IsNotExist(err) {
			log.Fatal("pkg doesn't exists: ", conf.Package)
		}
		if len(conf.Destination) == 0 {
			conf.Destination = conf.Package + "_ins"
		}
	case len(conf.Directory) > 0:
		conf.Target = inspector.TargetDirectory
		if _, err = os.Stat(conf.Directory); os.IsNotExist(err) {
			log.Fatal("dir doesn't exists: ", conf.Directory)
		}

		dirAbs, err := filepath.Abs(conf.Directory)
		if err != nil {
			log.Fatal(err)
		}
		dir := filepath.Dir(dirAbs)
		conf.InPlace = dir == conf.Destination

		if len(conf.Destination) == 0 {
			conf.Destination = conf.Directory + "_ins"
		}
	case len(conf.File) > 0:
		conf.Target = inspector.TargetFile

		fileAbs, err := filepath.Abs(conf.File)
		if err != nil {
			log.Fatal(err)
		}
		fileDir := filepath.Dir(fileAbs)
		conf.InPlace = fileDir == conf.Destination

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

	if (len(conf.Directory) > 0 || len(conf.File) > 0) && len(conf.Import) == 0 && !conf.InPlace {
		log.Fatal("Param -imp is required.")
	}

	// Check and read blacklist file.
	if len(bl) > 0 {
		conf.BlackList = make(map[string]struct{})
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
