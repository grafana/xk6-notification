package notification

import "fmt"

// ReportError is a handy function to show errors
func ReportError(err error, msg string) {
	// TODO: Use k6 logger (?)
	if err != nil {
		fmt.Println(err, msg)
	}
}
