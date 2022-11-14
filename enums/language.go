package enums

type Language string

const (
	FRENCH  Language = "fr"
	ENGLISH Language = "en"
	DUTCH   Language = "nl"
	ITALIAN Language = "it"
)

var AllowedLanguages = []string{
	string(ENGLISH),
	string(FRENCH),
	string(ITALIAN),
	string(DUTCH),
}
