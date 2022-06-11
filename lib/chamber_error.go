package lib

import "fmt"

// Defined Errors
var (
	UnsupportedHostError = ChamberError{StatusCode: 1, Message: "Anteroom Support GitHub"}
)

type ChamberError struct {
	StatusCode int
	Message    string
}

func (e *ChamberError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.StatusCode, e.Message)
}
