package myssh

import (
	"net"
	"strconv"
)

const (
	myhost     = "0.0.0.0"
	myport     = 2222
	myuser     = "foo"
	mypassword = "bar"
)

func MyAddress() string {
	return net.JoinHostPort(myhost, strconv.Itoa(myport))
}
