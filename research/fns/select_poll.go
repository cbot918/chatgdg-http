t := &syscall.Timeval{ /* timeout for the call*/ }
if _, err := syscall.Select(maxFD+1, fds, nil, nil, t); err != nil {
	return nil, err
}

for _, fd := range fds {
	if fdIsSet(fdset, fd) {
		// Consume data
	}
}