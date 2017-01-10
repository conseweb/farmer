package fs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/hyperledger/fabric/farmer/account"
	"github.com/hyperledger/fabric/storage/indexer"
)

func query(name string) error {
	cli := http.DefaultClient

	u, err := url.Parse(srvAddr)
	if err != nil {
		return err
	}

	u.Path = "/api/namesrv"

	q := u.Query()
	q.Set("name", name)
	u.RawQuery = q.Encode()

	logger.Infof("url: %s", u.String())

	resp, err := cli.Get(u.String())
	if err != nil {
		return err
	}

	bs, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		logger.Infof("status: %+v", resp.Status)
	}
	fmt.Println(string(bs))

	return nil
}

func setName(name, val string) error {
	cli := http.DefaultClient

	u, err := url.Parse(srvAddr)
	if err != nil {
		return err
	}
	u.Path = "/api/namesrv/new"

	var kv struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	kv.Key = name
	kv.Value = val
	logger.Infof("url: %s", u.String())

	buf := &bytes.Buffer{}
	err = json.NewEncoder(buf).Encode(kv)
	if err != nil {
		return err
	}

	resp, err := cli.Post(u.String(), "application/json", buf)
	if err != nil {
		return err
	}

	bs, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		logger.Infof("status: %+v", resp.Status)
	}
	fmt.Println(string(bs))

	return nil
}

func syncFiles() error {

	return nil
}

func initFileSystem() error {
	acc, err := account.LoadFromFile()
	if err != nil {
		return err
	}

	if len(acc.Devices) > 0 {
		return fmt.Errorf("required device.")
	}

	devID := acc.Devices[0].DeviceID

	idx, err := indexer.NewIndexer(srvAddr, devID)
	if err != nil {
		return err
	}

	return idx.SendIndexer()
}
