package dutf

import (
	"strings"
	"unicode/utf8"
)

func AppendRune(p []byte, last rune, r rune) (dutf []byte, _last rune) {
	if r <= 0x7F {
		p = utf8.AppendRune(p, r)
	} else {
		xorResult := last ^ r
		if r <= 0x3FFF {
			p = append(p,
				byte(0x80|(xorResult&0x7F)),
				byte((xorResult>>7)&0x7F))
		} else {
			p = append(p,
				byte(0x80|(xorResult&0x7F)),
				byte(0x80|((xorResult>>7)&0x7F)),
				byte((xorResult>>14)&0x7F))
		}
		last = r
	}
	return p, last
}

func EncodeString(s string) (dutf []byte) {
	last := rune(0)
	buffer := make([]byte, 0, len(s))
	for _, c := range s {
		buffer, last = AppendRune(buffer, last, c)
	}
	return buffer
}

func DecodeRune(p []byte, last rune) (r rune, siz int) {
	if p[0] <= 0x7F {
		return rune(p[0]), 1
	}
	xorDiff := rune(0)
	i := 0
	for {
		if i >= len(p) {
			return utf8.RuneError, -1
		}
		xorDiff |= (rune(p[i]) & 0x7F) << (7 * i)
		if p[i] <= 0x7F {
			return last ^ xorDiff, i + 1
		}
		i++
	}
}

func DecodeString(p []byte) (result string, undecodedBytes []byte) {
	var buffer strings.Builder
	last := rune(0)
	for len(p) > 0 {
		r, siz := DecodeRune(p, last)
		if r == utf8.RuneError {
			return buffer.String(), p
		}
		if r > 0x7F {
			last = r
		}
		buffer.WriteRune(r)
		p = p[siz:]
	}
	return buffer.String(), nil
}
