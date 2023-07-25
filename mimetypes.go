package mimetypes

const NO_MIME string = ""

// Guess returns the mime type of the given data
//
// Otherwise return "application/octet-stream"
func Guess(data []byte) string {
	return defaultGuesser.Guess(data)
}

// With registers a new guesser function
func With(guesser ...GuesserFunc) {
	defaultGuesser.With(guesser...)
}

// DefaultGuesser returns the default guesser
//
// It has two guesserFuncs registered:
//
// 1. LocalDatabaseGuesser
// 2. PlaintextGuesser
var defaultGuesser = func() *Guesser {
	var g = New()
	g.With(
		LocalDatabaseGuesser,
		PlaintextGuesser,
	)
	return g
}()
