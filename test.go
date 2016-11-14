package testrediss

import (
	"net"
	"net/url"
	"strconv"
)

func RedissURL(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	h, p, err := net.SplitHostPort(u.Host)
	if err != nil {
		return "", err
	}
	port, err := strconv.Atoi(p)
	if err != nil {
		return "", err
	}
	port++

	u.Host = net.JoinHostPort(h, strconv.Itoa(port))
	return u.String(), nil
}
