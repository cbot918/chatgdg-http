1. ws one million connection
   1.1 open file limit: ulimit NOFILE
   1.2 reduce goroutine: epoll
   1.3 reduce buffer allocation: gobwas
   1.4 Conntrack table: concurrent in the OS

2. ws to millions connection
   1.1 get rid of reader goroutine: [netpoll]("github.com/mailru/easygo/netpoll")
   1.2 reuse goroutine: [gopool](https://github.com/gobwas/ws-examples/tree/master/src/gopool)
   1.3 zero copy upgrade
   1.4 library [gobwas](https://github.com/gobwas/ws)

3. My WebSocket Library

   1. Research

      1. With net package: Notice `\r\n`
      2. With http package: What is hijack

   2. Impl
      1. http upgrade
      2. first echo message
      3. wrap ui
      4. client class
      5. handle message

4. chatapp implementation

5. monitor my chatapp

6. chat app load test
