package inspector

type Config struct {
	// Target points to source for parsing (package, directory or file).
	Target Target
	// Package indicates source directory relatives to GOPATH (or modules) directory.
	Package string
	// Directory indicates source directory contains go files.
	Directory string
	// File indicates path to Go file.
	File string
	// Destination directory path.
	Destination string
	// Package to import inspector code.
	// Required for directory or file targets.
	Import string
	// Registry of blacklisted types.
	BlackList map[string]struct{}
	// Flag denies destination directory cleaning.
	NoClean bool
	// Flag denies compiling inspectors to separate files.
	NoSplit bool
	// Path to build debug data in XML format.
	XML string
	// Output buffer.
	Buf ByteStringWriter
	// Logger of debug messages.
	Logger Logger
}

func (c *Config) Copy() *Config {
	cpy := *c
	return &cpy
}
