package main

type account struct {
	codes   []uint8
	balance float64
	name    string
}

func (a *account) login(code uint8) bool {
	if (*a).codes[1] == code {
		(*a).codes = (*a).codes[1:]
		return true
	}
	return false
}

func makeDumyCodes() (res []uint8) {
	res = make([]uint8, 0)
	for i := 0; i < 99; i++ {
		res = append(res, uint8(i))
	}
	return
}

func makeDumyAccounts() (res map[uint64]account) {
	res = make(map[uint64]account)
	res[1337] = account{makeDumyCodes(), 24, "Kalle"}
	res[4711] = account{makeDumyCodes(), 0, "Olle"}
	res[5011] = account{makeDumyCodes(), 3.14, "Jossan"}
	res[9001] = account{makeDumyCodes(), -100000, "Fredde"}
	return
}
