package mimetypes

import (
	"bytes"
	"path"
	"unicode"
)

// Check for signatures of known mimetypes in the data
func LocalDatabaseGuesser(filename string, data []byte) string {
	for _, mt := range mimeTypes {
		for _, sign := range mt.Signs {
			if len(data) < sign.Offset+len(sign.Bytes) {
				continue
			}
			if string(data[sign.Offset:sign.Offset+len(sign.Bytes)]) == string(sign.Bytes) {
				return mt.Mime
			}
		}
	}
	return NO_MIME
}

// PlaintextGuesser returns a guesser function that will check if the data is ascii or unicode
//
// If it is ascii, it will return "text/plain; charset=utf-8"
//
// If it is unicode, it will return "text/plain; charset=utf-16"
func PlaintextGuesser(filename string, data []byte) string {
	var (
		isAscii   = true
		isMod2    = len(data)%2 == 0
		isUnicode = isMod2
		b         byte
		b1        byte
	)
	for i := 0; i < len(data) && isAscii; i++ {
		b = data[i]
		// break if it is not ascii and not unicode, or if it is
		if isAscii = isAscii && b <= 126; !isAscii && !isUnicode {
			break
		}

		// If it is not a modulo of 2, then it is not unicode
		if isMod2 && isUnicode {
			b1 = data[i+1]
			if isUnicode = isUnicode && unicode.IsPrint(rune(b)<<8|rune(b1)); !isUnicode {
				break
			}
		}
	}
	if isAscii {
		return "text/plain; charset=utf-8"
	}
	if isUnicode {
		return "text/plain; charset=utf-16"
	}
	return NO_MIME
}

//
//	// RegexGuesser returns a guesser function that will check if the data matches the given regex
//	//
//	// If it does, it will return the given mimetype
//	func RegexGuesser(regex string, mimetype string) func([]byte) string {
//		var re = regexp.MustCompile(regex)
//		return func(data []byte) string {
//			if re.Match(data) {
//				return mimetype
//			}
//			return NO_MIME
//		}
//	}

// ContainsGuesser returns a guesser function that will check if the data contains any of the given byte slices
//
// If it does, it will return the given mimetype
func ContainsGuesser(mimeType string, contains [][]byte) func(filename string, data []byte) string {
	return func(filename string, data []byte) string {
		for _, c := range contains {
			if bytes.Contains(data, c) {
				return mimeType
			}
		}
		return NO_MIME
	}
}

var DefaultMimetypeMap = map[string]string{
	"html": "text/html",
	"htm":  "text/html",
	"txt":  "text/plain",
	"css":  "text/css",
	"csv":  "text/csv",
	"js":   "application/javascript",
	"json": "application/json",
	"xml":  "application/xml",
	"gif":  "image/gif",
	"jpeg": "image/jpeg",
	"jpg":  "image/jpeg",
	"png":  "image/png",
	"svg":  "image/svg+xml",
	"ico":  "image/x-icon",
	"webp": "image/webp",
	"mp4":  "video/mp4",
	"webm": "video/webm",
	"pdf":  "application/pdf",
	"zip":  "application/zip",
	"gz":   "application/gzip",
	"tar":  "application/x-tar",
	"7z":   "application/x-7z-compressed",
	"rar":  "application/x-rar-compressed",
	"bz2":  "application/x-bzip2",
	"mp3":  "audio/mpeg",
	"wav":  "audio/wav",
	"ogg":  "audio/ogg",
	"flac": "audio/flac",
	"mpg":  "video/mpeg",
	"mpeg": "video/mpeg",
	"mov":  "video/quicktime",
	"avi":  "video/x-msvideo",
	"wmv":  "video/x-ms-wmv",
}

// ExtensionGuesser returns a guesser function that will check if the filename ends with the given extension.
// If it does, it will return the given mimetype.
//
// This function takes a map of extensions to mimetypes
func ExtensionGuesser(m map[string]string) func(filename string, data []byte) string {
	return func(filename string, data []byte) string {
		if mime, ok := m[path.Ext(filename)]; ok {
			return mime
		}
		return NO_MIME
	}
}
