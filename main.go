package main

import (
	"encoding/json"
	"fmt"
	"log"

	lua "github.com/yuin/gopher-lua"
)

func Double(l *lua.LState) int {
	lv := l.ToInt(1)
	l.Push(lua.LNumber(lv * 2))
	return 1
}
func main() {
	l := lua.NewState()
	defer l.Close()
	l.SetGlobal("double", l.NewFunction(Double))

	//lua call go
	if err := l.DoString(`print(double(5))`); err != nil {
		log.Println(err)
		return
	}

	//go call lua
	if err := l.DoFile("go_call_lua.lua"); err != nil {
		log.Println(err)
		return
	}

	tmp := struct {
		Id int `json:"id"`
	}{
		Id: 99,
	}
	b, err := json.Marshal(tmp)
	if err != nil {
		log.Println(err)
		return
	}
	a := make([]lua.LValue, 0)
	a = append(a, lua.LString(b))

	err = l.CallByParam(lua.P{
		Fn:      l.GetGlobal("process"),
		NRet:    1,
		Protect: true,
	}, a...)
	if err != nil {
		log.Println(err)
	}
	ret := l.Get(-1)
	l.Pop(1)
	fmt.Println(ret, ret.Type())

}
