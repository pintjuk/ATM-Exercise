package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

const (
	ETB        byte = 0x17
	ET         byte = 0x18
	SEND_UINT  byte = 0x02
	SEND_INT   byte = 0x06
	SEND_FLOAT byte = 0x0E
	SEND_LOGIN byte = 0x1E
	SEND_ASCII byte = 0x3E
	SEND_OTHER byte = 0x7E

	PRINT_OPCODE byte = 0x01
	CLEAR_FlAG   byte = 0x80
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
	if !IsUintSendMSG(msg) {
		return 0, errors.New("Not send Uint message")
	}
	buf := bytes.NewBuffer(make([]byte, 0, 10))
	buf.Write(msg[1:])
	fmt.Println()
	binary.Read(buf, binary.BigEndian, &result)
	return
}

func MakeSendLoginCMD()

func DecodeSendDataCMD()

func MakePrintCMD(stringRef uint8, livedata []byte, flags ...byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 10))
	var opcode byte = 0x01
	for _, e := range flags {
		opcode = opcode | e
	}
	buf.WriteByte(opcode)
	buf.WriteByte(byte(stringRef))
	buf.Write(livedata)
	return buf.Bytes()
}
func mergFlags(flags []byte) (opcode byte) {
	opcode = PRINT_OPCODE
	for _, e := range flags {
		opcode = opcode | e
	}
	return
}

func MakePrintUintCMD(stringRef uint8, livedata uint64, flags ...byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 8))
	binary.Write(buf, binary.BigEndian, livedata)

	return MakePrintCMD(stringRef, buf.Bytes(), mergFlags(flags))
}
func MakePrintIntCMD(stringRef uint8, livedata int64, flags ...byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 8))
	binary.Write(buf, binary.BigEndian, livedata)

	return MakePrintCMD(stringRef, buf.Bytes(), mergFlags(flags))
}

func MakePrintFloatCMD(stringRef uint8, livedata float64, flags ...byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 8))
	binary.Write(buf, binary.BigEndian, livedata)

	return MakePrintCMD(stringRef, buf.Bytes(), mergFlags(flags))
}

func MakePrintStringCMD(stringRef uint8, livedata string, flags ...byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 8))
	buf.WriteString(livedata)
	buf.WriteByte(0x00)
	buf.Truncate(8)

	return MakePrintCMD(stringRef, buf.Bytes(), mergFlags(flags))
}

func DecodePrintCMD(data []byte, lang Lang) (usertext string, clear bool, dataRequest byte, err error) {
	if data[0]&PRINT_OPCODE != PRINT_OPCODE {
		err = errors.New("this is not a print command")
	}
	dataRequest = data[0] & SEND_OTHER
	clear = data[0]&CLEAR_FlAG == CLEAR_FlAG

	var dataAsUint uint64
	var dataAsInt int64
	var dataAsFlout float64

	buf := bytes.NewBuffer(make([]byte, 0, 10))
	read := func(d interface{}) {
		buf.Reset()
		buf.Write(data[2:10])
		binary.Read(buf, binary.BigEndian, d)
	}
	read(&dataAsUint)
	read(&dataAsInt)
	read(&dataAsFlout)

	buf.Reset()
	buf.Write(data[2:10])
	buf.WriteByte(0x00)
	dataAsString, _ := buf.ReadString(0x00)

	usertext = strings.Replace(lang[data[1]], "%u", fmt.Sprintf("%d", dataAsUint), -1)
	usertext = strings.Replace(usertext, "%i", fmt.Sprintf("%d", dataAsInt), -1)
	usertext = strings.Replace(usertext, "%f", fmt.Sprintf("%f", dataAsFlout), -1)
	usertext = strings.Replace(usertext, "%s", fmt.Sprintf("%s", dataAsString), -1)
	return
}
