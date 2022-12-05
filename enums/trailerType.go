package enums

type Trailer string

const (
	TEASER  Language = "Teaser"
	TRAILER Language = "Trailer"
)

var TrailerTypes = []string{
	string(TEASER),
	string(TRAILER),
}
