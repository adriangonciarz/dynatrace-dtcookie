package detectionrules

type Base string

var Bases = struct {
	FileName Base
	Fqcn     Base
	Package  Base
}{
	Base("FILE_NAME"),
	Base("FQCN"),
	Base("PACKAGE"),
}

type Matcher string

var Matchers = struct {
	BeginsWith Matcher
	Contains   Matcher
}{
	Matcher("BEGINS_WITH"),
	Matcher("CONTAINS"),
}

type Technology string

var Technologies = struct {
	Dotnet Technology
	Go     Technology
	Java   Technology
	Nodejs Technology
	Php    Technology
}{
	Technology("DotNet"),
	Technology("Go"),
	Technology("Java"),
	Technology("Nodejs"),
	Technology("PHP"),
}
