package detectionrules

type Technology string

var Technologies = struct {
	Php    Technology
	Java   Technology
	Dotnet Technology
	Go     Technology
	Nodejs Technology
}{
	Technology("PHP"),
	Technology("Java"),
	Technology("DotNet"),
	Technology("Go"),
	Technology("Nodejs"),
}

type Matcher string

var Matchers = struct {
	BeginsWith Matcher
	Contains   Matcher
}{
	Matcher("BEGINS_WITH"),
	Matcher("CONTAINS"),
}

type Base string

var Bases = struct {
	Fqcn     Base
	FileName Base
	Package  Base
}{
	Base("FQCN"),
	Base("FILE_NAME"),
	Base("PACKAGE"),
}
