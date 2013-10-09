package protocol

import (
	"bytes"
	"strings"
)

type Lang map[uint8]string

func EncadeLang(lang Lang) []byte {
	var buff bytes.Buffer
	for n, _ := range lang {
		buff.WriteByte(byte(n))
		buff.WriteString(wipeEOT(lang[n]))
		buff.WriteRune(rune(0x17))
	}
	return buff.Bytes()
}

func DecadeLang(b []byte) (lang Lang) {
	var buff bytes.Buffer
	buff.Write(b)
	lang = make(Lang)
	for buff.Len() > 1 {
		key, _ := buff.ReadByte()
		text, _ := buff.ReadString(0x17)
		lang[uint8(key)] = wipeEOT(text)
	}
	return
}

func wipeEOT(s string) string {
	return strings.Replace(s, string(0x17), "", -1)
}
