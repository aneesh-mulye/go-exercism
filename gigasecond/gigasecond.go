// A package used to add gigaseconds to
package gigasecond

// import path for the time package from the standard library
import "time"

// Adds a gigasecond to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1_000_000_000)
}
