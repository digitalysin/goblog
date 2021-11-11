package ftp

import (
	"io"
	"time"

	gftp "github.com/jlaffaye/ftp"
)

type (
	EntryType string

	Ftp interface {
		Get(path string, dst io.Writer) error
		Put(path string, src io.Reader) error
		List(path string) ([]Entry, error)
		Close() error
	}

	Entry struct {
		Name   string
		Target string
		Type   EntryType
		Size   uint64
		Time   time.Time
	}

	auth struct {
		user, password string
	}

	impl struct {
		ftp  *gftp.ServerConn
		auth *auth
	}
)

const (
	Folder  = EntryType("FOLDER")
	File    = EntryType("FILE")
	SymLink = EntryType("SYMBOLIC_LINK")
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

func (i *impl) List(path string) ([]Entry, error) {
	var (
		entries, err = i.ftp.List(path)
		result       = make([]Entry, 0)
	)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		var (
			res = Entry{
				Name:   entry.Name,
				Size:   entry.Size,
				Time:   entry.Time,
				Target: entry.Target,
			}
		)

		if entry.Type == gftp.EntryTypeFile {
			res.Type = File
		}

		if entry.Type == gftp.EntryTypeFolder {
			res.Type = Folder
		}

		if entry.Type == gftp.EntryTypeLink {
			res.Type = SymLink
		}

		result = append(result, res)
	}

	return result, nil
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
