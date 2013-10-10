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
		buff.WriteByte(ETB)
	}
	buff.WriteByte(ET)
	return buff.Bytes()
}

func DecadeLang(b []byte) (lang Lang) {
	var buff bytes.Buffer
	buff.Write(b)
	temp, _ := buff.ReadBytes(ET)
	buff.Reset()
	buff.Write(temp)
	lang = make(Lang)
	for buff.Len() > 1 {
		key, _ := buff.ReadByte()
		text, _ := buff.ReadString(ETB)
		lang[uint8(key)] = wipeEOT(text)
	}
	return
}

func wipeEOT(s string) string {
	return strings.Replace(strings.Replace(s, string(ETB), "", -1), string(ET), "", -1)
}
