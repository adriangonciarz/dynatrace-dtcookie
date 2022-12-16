package usersettings

type Language string

var Languages = struct {
	Auto Language
	En   Language
	Ja   Language
}{
	Language("Auto"),
	Language("En"),
	Language("Ja"),
}

type Theme string

var Themes = struct {
	Auto  Theme
	Dark  Theme
	Light Theme
}{
	Theme("Auto"),
	Theme("Dark"),
	Theme("Light"),
}
