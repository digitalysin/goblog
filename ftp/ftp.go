package ftp

import (
	"io"
	"time"

	gftp "github.com/jlaffaye/ftp"
)

type (
	Ftp interface {
		Get(path string, dst io.Writer) error
		Put(path string, src io.Reader) error
		Close() error
	}

	auth struct {
		user, password string
	}

	impl struct {
		ftp  *gftp.ServerConn
		auth *auth
	}
)

func (a *auth) IsEmptyAuth() bool {
	return len(a.user) == 0 && len(a.password) == 0
}

func (i *impl) Close() error {
	if !i.auth.IsEmptyAuth() {
		i.ftp.Logout()
	}
	return i.ftp.Quit()
}

func (i *impl) Get(path string, dst io.Writer) error {
	var (
		resp, err = i.ftp.Retr(path)
	)

	if err != nil {
		return err
	}

	if _, err := io.Copy(dst, resp); err != nil {
		return err
	}

	return resp.Close()
}

func (i *impl) Put(path string, src io.Reader) error {
	return i.ftp.Stor(path, src)
}

func New(address, user, password string, timeout time.Duration) (Ftp, error) {
	ftp, err := gftp.DialTimeout(address, timeout)

	if err != nil {
		return nil, err
	}

	i := &impl{ftp: ftp, auth: &auth{user, password}}

	if i.auth.IsEmptyAuth() {
		return i, nil
	}

	if err := i.ftp.Login(i.auth.user, i.auth.password); err != nil {
		i.ftp.Quit()
		return nil, err
	}

	return i, nil
}
