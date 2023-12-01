package main

import (
	"bufio"
	"chatgdg-http/examples/5_message/ws"
	"net"
)

type Client struct {
	ID   string
	conn net.Conn
	rw   *bufio.ReadWriter
	f    *ws.Frame
	buf  []byte
}

func NewClient(id string, conn net.Conn, rw *bufio.ReadWriter) *Client {
	return &Client{
		ID:   id,
		conn: conn,
		rw:   rw,
		f:    ws.NewFrame(),
		buf:  make([]byte, 4096),
	}
}

func (c *Client) ReadString() (msg []byte, err error) {
	_, err = c.rw.Read(c.buf)
	if err != nil {
		return []byte(""), err
	}
	return c.f.DecodeFrame(c.buf), nil
}

func (c *Client) WriteString(m []byte) (err error) {
	outMessage := c.f.EncodeFrame(m)
	_, err = c.rw.Write(outMessage)
	if err != nil {
		return err
	}
	err = c.rw.Flush()
	if err != nil {
		return err
	}
	return err
}

func (c *Client) WriteMessage(cm *ChanMsg) (err error) {

	outMessage := c.f.EncodeFrame(EnChanMsg(cm))

	_, err = c.rw.Write(outMessage)
	if err != nil {
		return err
	}
	err = c.rw.Flush()
	if err != nil {
		return err
	}
	return err
}

func (c *Client) ReadMessage() (*ChanMsg, error) {
	_, err := c.rw.Read(c.buf)
	if err != nil {
		return nil, err
	}
	bytes := c.f.DecodeFrame(c.buf)

	return DeChanMsg(string(bytes)), nil
}
