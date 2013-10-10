package main

import (
	"../protocol"
	"sync"
)

const (
	HEJ   uint8 = 1
	LOGIN uint8 = 2
	ADD   uint8 = 3
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

func MakeDefoultLangs() (l *Langs) {
	l = new(Langs)
	l.langs = make(map[string]protocol.Lang)
	l.langs["en"] = protocol.Lang{ADD: "Välkommenochköpsaker", HEJ: "hej vällkomen", LOGIN: "skriv 1 för att låga in, skriv 2 för att låga in"}
	l.langs["sv"] = protocol.Lang{ADD: "Welcome by stuff", HEJ: "welcome", LOGIN: "Pres 1 to login, pres 2 to change language"}
	l.langs["ch"] = protocol.Lang{ADD: "我們一定會離開評價，但也得到的Joomla管理和可靠的流量列在此頁面。除了一個簡單的例子，我們這是文字添加到購物車填寫表格時，一些他的人從下面的選擇。搜索編輯或想在近期的奧運惱人的後果，或在一系列規模疼痛有沒有貨運選擇這個文本和誰只是", HEJ: "它遵循一條評", LOGIN: "我們一定會離開評價 1，但也 2 得到J 們行使"}
	return
}
