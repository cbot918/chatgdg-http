package main

import (
	"bufio"
	"net"
)

type Client struct {
	ID   string
	conn net.Conn
	rw   *bufio.ReadWriter
	f    *Frame
	buf  []byte
}

func NewClient(id string, conn net.Conn, rw *bufio.ReadWriter) *Client {
	return &Client{
		ID:   id,
		conn: conn,
		rw:   rw,
		f:    NewFrame(),
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
