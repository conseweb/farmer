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
	"sync"

	"github.com/go-xorm/xorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	log = logging.MustGetLogger("indexer")
	orm *xorm.Engine
)

type Device struct {
	ID      string `xorm:"pk" json:"id"`
	Address string `xorm:"notnull index" json:"address"`
}

type Indexer struct {
	cli   *http.Client
	devID string
	addr  string
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

func (idx *Indexer) ListLocalAll() ([]*FileInfo, error) {
	type walkArgs struct {
		fp   string
		info os.FileInfo
		e    error
	}
	files := []*FileInfo{}

	buf := make(chan *FileInfo, 100)
	wksbuf := make(chan *FileInfo, 100)
	done := make(chan struct{})

	MaxRunc := 4
	wg := new(sync.WaitGroup)

	for i := 0; i < MaxRunc; i++ {
		go func() {
			for {
				select {
				case file := <-wksbuf:
					f, err := os.Open(file.Path)
					if err != nil {
						log.Error(err)
						wg.Done()
						return
					}

					hash := sha256.New()
					_, err = io.Copy(hash, f)
					if err != nil {
						log.Error(err)
						wg.Done()
						f.Close()
						return
					}

					file.Hash = fmt.Sprintf("%x", hash.Sum(nil))
					f.Close()
					buf <- file

				case <-done:
					return
				}
			}
		}()
	}

	go func() {
		for {
			select {
			case fi := <-buf:
				wg.Done()
				files = append(files, fi)
			case <-done:
				return
			}
		}
	}()

	filepath.Walk("/", func(fp string, info os.FileInfo, e error) error {
		if e != nil {
			log.Error(e)
			return e
		}
		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		wksbuf <- &FileInfo{
			DeviceID: idx.devID,
			Path:     fp,
			Size:     info.Size(),
		}
		return nil
	})

	wg.Wait()
	close(done)

	return files, nil
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

	remoteFiles := []*FileInfo{}
	err = json.NewDecoder(resp.Body).Decode(remoteFiles)
	if err != nil {
		return err
	}

	localFiles, err := idx.ListLocalAll()
	if err != nil {
		return err
	}

	ret, err := idx.mergeFiles(remoteFiles, localFiles)
	if err != nil {
		return err
	}

	err = idx.updateRemote(ret)
	if err != nil {
		return err
	}

	return nil
}

func (idx *Indexer) mergeFiles(remotes, locals []*FileInfo) (*MergeResult, error) {
	ret := &MergeResult{
		Add:    []*FileInfo{},
		Update: []*FileInfo{},
		Del:    []*FileInfo{},
	}

	for i := 0; i < len(locals); i++ {
		local := locals[i]
		local.DeviceID = idx.devID
		exists := false
		var j int
	loop_remote:
		for ; j < len(remotes); j++ {
			remote := remotes[j]
			if remote.Hash == local.Hash && remote.Path == local.Path {
				exists = true
				break loop_remote
			} else if (remote.Hash != local.Hash && remote.Path == local.Path) ||
				(remote.Hash == local.Hash && remote.Path != local.Path) {
				exists = true
				local.ID = remote.ID
				ret.Update = append(ret.Update, local)
				break loop_remote
			}
		}
		if exists {
			if j+1 < len(remotes) {
				remotes = append(remotes[:j], remotes[j+1:]...)
			} else {
				remotes = remotes[:j]
			}
		} else {
			ret.Add = append(ret.Add, local)
		}
	}
	for _, rem := range remotes {
		rem.DeviceID = idx.devID
		ret.Del = append(ret.Del, rem)
	}
	return ret, nil
}

func (idx *Indexer) updateRemote(upd *MergeResult) error {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(upd)
	if err != nil {
		return err
	}

	u, _ := url.Parse(idx.addr)
	u.Path = "/devices/files"

	resp, err := idx.cli.Post(u.String(), "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("%s", resp.Status)
	}

	return nil
}
