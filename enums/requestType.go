package enums

type RequestType string

const (
	FULL    RequestType = "full"
	MIN     RequestType = "min"
	ADMIN   RequestType = "admin"
	UNKNOWN RequestType = "unknown"
)

var AllowedRequestType = []string{
	string(FULL),
	string(MIN),
	string(ADMIN),
	string(UNKNOWN),
}
