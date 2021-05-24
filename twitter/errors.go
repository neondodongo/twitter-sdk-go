package twitter

import "fmt"

var (
	// ErrTwitterClient base error for all errors thrown by the Twitter API client
	ErrTwitterClient error = fmt.Errorf("[twitter api client]")
)
