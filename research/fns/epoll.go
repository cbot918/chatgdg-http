package fns

epfd, _ := unix.EpollCreate1(0)

_ := unix.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd,
&unix.EpollEvent{Events: unix.PILLIN | unix.POLLHUP, Fd: fd})

events := make([]unix.EpollEvent, 100)

n, _ := unix.EpollWait(e.fd, events, 100)
// blocking

for i := 0; i < n; i++ {
	// consume data
}