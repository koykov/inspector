package inspector

// Generic generators.
//go:generate inspc -package github.com/koykov/inspector/testobj
// //go:generate inspc -directory testobj --import github.com/koykov/inspector/testobj
// //go:generate inspc -file testobj/testobj.go --destination testobj_ins --import github.com/koykov/inspector/testobj --no-clean
// //go:generate inspc -file testobj/testobj1.go --destination testobj_ins --import github.com/koykov/inspector/testobj --no-clean

// Debug generators.
// Writes as result parsed data in XML format. Useful to guarantee similarity of loader and AST parsers.
//go:generate inspc -package github.com/koykov/inspector/testobj -xml testdata
// //go:generate inspc -directory testobj --import github.com/koykov/inspector/testobj -xml testdata
// //go:generate inspc -file testobj/testobj.go --destination testobj_ins --import github.com/koykov/inspector/testobj --no-clean -xml testdata
// //go:generate inspc -file testobj/testobj1.go --destination testobj_ins --import github.com/koykov/inspector/testobj --no-clean -xml testdata
