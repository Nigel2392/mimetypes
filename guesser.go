package mimetypes

type GuesserFunc func(filename string, data []byte) string

type Guesser struct {
	Funcs []GuesserFunc
}

func New() *Guesser {
	return &Guesser{}
}

func (g *Guesser) With(guesser ...GuesserFunc) {
	if g.Funcs == nil {
		g.Funcs = make([]GuesserFunc, 0)
	}
	g.Funcs = append(g.Funcs, guesser...)
}

// Guess returns the mime type of the given data
//
// Otherwise return "application/octet-stream"
func (g *Guesser) Guess(filename string, data []byte) string {
	var mime string
	for _, guesser := range g.Funcs {
		if mime = guesser(filename, data); mime != NO_MIME {
			return mime
		}
	}

	return "application/octet-stream"
}
