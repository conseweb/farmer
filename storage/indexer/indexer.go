package indexer

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	log = logging.MustGetLogger("indexer")
	orm *xorm.Engine
)

type FileInfo struct {
	ID       int64  `xorm:"pk autoincr 'id'" json:"id"`
	DeviceID string `xorm:"notnull index 'device_id'" json:"device_id"`
	Path     string `xorm:"notnull index 'path'" json:"path"`
	Hash     string `xorm:"notnull index 'hash'" json:"hash"`
	Size     int64  `xorm:"'size'" json:"hash"`

	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

type Device struct {
	ID      string `xorm:"pk" json:"id"`
	Address string `xorm:"notnull index" json:"address"`
}

type Indexer struct {
	cli   *http.Client
	devID string
	addr  string
}

type mergeResult struct {
	Add, Update, Del []*FileInfo
}

func NewIndexer(addr, devID string) (*Indexer, error) {
	_, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	if devID == "" {
		return nil, fmt.Errorf("required device ID.")
	}

	return &Indexer{
		cli:   http.DefaultClient,
		addr:  addr,
		devID: devID,
	}, nil
}

func (idx *Indexer) SendIndexer() error {
	buf := make(chan *FileInfo, 10)
	done := make(chan struct{})

	go func() {
		filepath.Walk("root", func(fp string, info os.FileInfo, e error) error {
			if e != nil {
				return e
			}

			hash := sha256.New()
			f, err := os.Open(fp)
			if err != nil {
				return err
			}
			defer f.Close()

			fi, err := f.Stat()
			if err != nil {
				return err
			}
			if fi.IsDir() {
				return nil
			}

			_, err = io.Copy(hash, f)
			if err != nil {
				return err
			}

			buf <- &FileInfo{
				DeviceID: idx.devID,
				Path:     fp,
				Hash:     fmt.Sprintf("%x", hash.Sum(nil)),
				Size:     info.Size(),
			}

			return nil
		})

		var nl struct{}
		done <- nl
	}()

	files := []*FileInfo{}
	for {
		select {
		case fi := <-buf:
			if len(files) < 10 {
				files = append(files, fi)
			} else {
				err := idx.send(files)
				if err != nil {
					return err
				}
				files = []*FileInfo{}
			}

		case <-done:
			if len(files) > 0 {
				err := idx.send(files)
				if err != nil {
					return err
				}
			}

			log.Debugf("Walk dir over.")
			break
		}
	}

	return nil
}

func (idx *Indexer) send(files []*FileInfo) error {
	u, _ := url.Parse(idx.addr)
	u.Path = fmt.Sprintf("/devices/%s/online", idx.devID)

	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(files)
	if err != nil {
		return err
	}
	resp, err := idx.cli.Post(u.String(), "application/json", body)
	if err != nil {
		return err
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Infof("Send over. %s", bs)
	return nil
}

func (idx *Indexer) Online() error {
	addr := viper.GetString("daemon.address")
	dev := &Device{
		Address: addr,
	}

	u, _ := url.Parse(idx.addr)
	u.Path = fmt.Sprintf("/devices/%s/online", idx.devID)

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(dev)
	if err != nil {
		return err
	}

	resp, err := idx.cli.Post(u.String(), "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 300 {
		log.Errorf("status: %s", resp.Status)
		return fmt.Errorf("online failed.")
	}

	return nil
}

func (idx *Indexer) SyncFiles() error {
	u, _ := url.Parse(idx.addr)
	u.Path = fmt.Sprintf("/devices/%s/", idx.devID)

	resp, err := idx.cli.Get(u.String())
	if err != nil {
		return err
	}

	if resp.StatusCode > 300 {
		bs, _ := ioutil.ReadAll(resp.Body)
		log.Errorf("Status: %v, body: %s", resp.Status, bs)
		return fmt.Errorf("%s", bs)
	}

	files := []*FileInfo{}
	err = json.NewDecoder(resp.Body).Decode(files)
	if err != nil {
		return err
	}

	return nil
}

func mergeFiles(remotes, locals []*FileInfo) (*mergeResult, error) {
	ret := &mergeResult{
		Add:    []*FileInfo{},
		Update: []*FileInfo{},
		Del:    []*FileInfo{},
	}

	for i := 0; i < len(locals)-1; i++ {
		local := locals[i]
		exists := false

		for j := i; j < len(remotes); j++ {
			remote := remotes[j]
			if remote.Hash == local.Hash && remote.Path == local.Path {
				exists = true
				break loop_remote
			} else if remote.Hash != local.Hash && remote.Path == local.Path {
				exists = true
				local.ID = remote.ID
				local.DeviceID = remote.DeviceID
				ret.Update = append(ret.Update, local)
				break loop_remote
			} else if remote.Hash == local.Hash && remote.Path != local.Path {
				exists = true
				local.ID = remote.ID
				local.DeviceID = remote.DeviceID
				ret.Update = append(ret.Update, local)
				break loop_remote
			}
		loop_remote:
		}
		if exists {
			ret.Add = append(ret.Add, local)
		}
	}
	return nil, nil
}
