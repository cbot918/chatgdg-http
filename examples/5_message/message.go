package main

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
