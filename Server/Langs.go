package main

import (
	"../protocol"
	"sync"
)

const (
	HEJ        uint8 = 1
	INIT       uint8 = 2
	ADD        uint8 = 3
	LANG       uint8 = 4
	ENTACC     uint8 = 5
	ENTCODE    uint8 = 6
	ERR_ENTACC uint8 = 7
	ERR_CODE   uint8 = 8
	SUCLOGIN   uint8 = 9
	BALANCE    uint8 = 10

	//terminal colloring 
	Reset      = "\x1b[0m"
	Bright     = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underscore = "\x1b[4m"
	Blink      = "\x1b[5m"
	Reverse    = "\x1b[7m"
	Hidden     = "\x1b[8m"

	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"

	BgBlack   = "\x1b[40m"
	BgRed     = "\x1b[41m"
	BgGreen   = "\x1b[42m"
	BgYellow  = "\x1b[43m"
	BgBlue    = "\x1b[44m"
	BgMagenta = "\x1b[45m"
	BgCyan    = "\x1b[46m"
	BgWhite   = "\x1b[47m"
)

type Langs struct {
	langs   map[string]protocol.Lang
	blocker sync.Mutex
}

func (l *Langs) Lang(langcode string) (lang protocol.Lang) {
	l.blocker.Lock()
	defer l.blocker.Unlock()
	lang = l.langs[langcode]
	return
}

func (l *Langs) addLang(langcode string, lang protocol.Lang) {
	l.blocker.Lock()
	defer l.blocker.Unlock()
	l.langs[langcode] = lang

}

func MakeDefoultLangs() (l map[string]protocol.Lang) {
	l = make(map[string]protocol.Lang)
	l["sv"] = protocol.Lang{ADD: BgWhite + FgBlack + Underscore + "VÄLKOMMEN!" + Reset + FgBlue + BgWhite + "\n\nSMS lån 3000kr på 198 98 får 200kr/mån\nfler ärbjudanden på " + Underscore + "www.bank.se" + Reset, HEJ: "hej vällkomen", INIT: "\x1b[31m1\x1b[0m) låga in\n\x1b[31m2\x1b[0m) byta språk", LANG: "\tSelect lang:\n\t" + FgRed + "1" + Reset + " English\n\t" + FgRed + "1" + Reset + ") Svenska\n\t" + FgRed + "3" + Reset + ") 的例子", ENTACC: "Ange kontonumer:", ENTCODE: "Ange kod:", ERR_ENTACC: "Accountdosent Exist!"}
	l["en"] = protocol.Lang{ADD: "Welcome by stuff", HEJ: "welcome", INIT: "1) to login\n2) to change language", LANG: "1 en\n2 sv\n3 ch", ENTACC: "Enter account number:", ENTCODE: "Enter code:", ERR_ENTACC: FgRed + "Accountdosent Exist!" + Reset, ERR_CODE: FgRed + "Wrong code attempt %u/3!" + Reset, SUCLOGIN: FgGreen + "Sucsusful login!" + Reset + " Welcome %s"}
	l["en"][BALANCE] = "Your balance is %f"
	l["ch"] = protocol.Lang{ADD: BgWhite + FgRed + "我們一定會離開評價!" + Reset + BgWhite + FgBlack + "但也得到的Joomla管理和可靠的流量列在此頁面。除了一個簡單的例子，我們這是文字添加到購物車填寫表格時，一些他的人從下面的選擇。搜索編輯或想在近期的奧運惱人的後果，或在一系列規模疼痛有沒有貨運選擇這個文本和誰只是" + Reset, INIT: FgGreen + "\t1" + Reset + ") 但也\n" + FgGreen + "\t2" + Reset + ") 得到J 們行使", LANG: FgGreen + "\t1" + Reset + ") 但也\n" + FgGreen + "\t2" + Reset + ") 得到J 們行使\n\t" + FgGreen + "3" + Reset + ") 遵循一條", ENTACC: "Enter account number:", ENTCODE: "Enter code:", ERR_ENTACC: "Accountdosent Exist!"}
	return
}
