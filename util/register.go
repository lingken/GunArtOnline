package util

type RegisterList struct {
	record map[string]bool
}

func NewRegisterList() *RegisterList {
	reg := RegisterList{
		record: make(map[string]bool),
	}
	return &reg
}

func (reg *RegisterList) Register(key string) {
	reg.record[key] = true
}

func (reg *RegisterList) Unregister(key string) {
	delete(reg.record, key)
}

func (reg *RegisterList) IsRegister(key string) bool {
	if _, ok := reg.record[key]; ok {
		return true
	}
	return false
}
