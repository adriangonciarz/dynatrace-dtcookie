package customservices

// Visibility has no documentation
type Visibility string

// Visibilities offers the known enum values
var Visibilities = struct {
	Internal         Visibility
	PackageProtected Visibility
	Private          Visibility
	Protected        Visibility
	Public           Visibility
}{
	"INTERNAL",
	"PACKAGE_PROTECTED",
	"PRIVATE",
	"PROTECTED",
	"PUBLIC",
}

func (me Visibility) Ref() *Visibility {
	return &me
}
