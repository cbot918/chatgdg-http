package main

import (
	"strings"
)

type ChanMsg struct {
	Chan string
	Msg  Msg
}

type Msg struct {
	Data string
}

func EnChanMsg(m *ChanMsg) []byte {

	str := m.Chan + ";" + m.Msg.Data

	return []byte(str)
}

func DeChanMsg(input string) *ChanMsg {
	return &ChanMsg{
		Chan: strings.Split(input, ";")[0],
		Msg: Msg{
			Data: strings.Split(input, ";")[1],
		},
	}
}
