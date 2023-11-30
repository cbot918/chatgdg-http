package main

type Chat struct {
	UserCount int
	Clients   []Client
}

func NewChat() *Chat {
	return &Chat{
		UserCount: 0,
		Clients:   []Client{},
	}
}
