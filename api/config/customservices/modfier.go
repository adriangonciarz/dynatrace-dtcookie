package customservices

// Modifier has no documentation
type Modifier string

// Modifiers offers the known enum values
var Modifiers = struct {
	Abstract Modifier
	Extern   Modifier
	Final    Modifier
	Native   Modifier
	Static   Modifier
}{
	"ABSTRACT",
	"EXTERN",
	"FINAL",
	"NATIVE",
	"STATIC",
}
