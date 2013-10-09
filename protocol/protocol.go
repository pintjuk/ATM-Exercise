package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"
)

const (
	ETB        byte = 0x17
	ET         byte = 0x18
	SEND_UINT  byte = 0x01
	SEND_LOGON byte = 0x02
)

func MakeUpdate(add string, lang Lang) (result []byte) {
	var buff bytes.Buffer
	buff.WriteString(add)
	buff.WriteByte(ETB)
	buff.Write(EncadeLang(lang))
	return buff.Bytes()
}

func DecodeUpdate(data []byte) (add string, lang Lang) {
	var buff bytes.Buffer
	buff.Write(data)
	add, _ = buff.ReadString(ETB)
	add = strings.Replace(add, string(ETB), "", -1)
	lang = DecadeLang(buff.Bytes())
	return
}

func MakeSendUintMSG(data uint64) (msg []byte) {
	buf := bytes.NewBuffer(make([]byte, 0, 10))
	buf.WriteByte(SEND_UINT)
	binary.Write(buf, binary.BigEndian, data)
	msg = buf.Bytes()
	return
}
func IsUintSendMSG(msg []byte) bool {
	return msg[0] == SEND_UINT
}

func DecodeUintMSG(msg []byte) (result uint64, err error) {
	if IsUintSendMSG(msg) {
		return 0, errors.New("Not send Uint message")
	}
	result, n := binary.Uvarint(msg[1:9])
	if n < 0 {
		err = errors.New("Unable to read data")
	}
	return
}
func MakeSendLoginCMD()

func DecodeSendDataCMD()

func MakePrintCMD()
