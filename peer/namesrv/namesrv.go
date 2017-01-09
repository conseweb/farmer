package namesrv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

const (
	nameSrvFuncName = "namesrv"
)

var (
	logger = logging.MustGetLogger("nameSrvCmd")

	nameSrvCmd = &cobra.Command{
		Use:   nameSrvFuncName,
		Short: fmt.Sprintf("%s specific commands.", nameSrvFuncName),
		Long:  fmt.Sprintf("%s specific commands.", nameSrvFuncName),
	}

	srvAddr string
)

// Cmd returns the cobra command for Node
func Cmd() *cobra.Command {
	flags := nameSrvCmd.PersistentFlags()

	flags.StringVarP(&srvAddr, "addr", "a", "http://localhost:9375", "server address")

	nameSrvCmd.AddCommand(queryCmd())
	nameSrvCmd.AddCommand(setKvCmd())

	return nameSrvCmd
}

func queryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "query name.",
		RunE: func(cmd *cobra.Command, args []string) error {
			names := cmd.Flags().Args()
			if len(names) != 1 {
				return fmt.Errorf("invalid name")
			}
			return query(names[0])
		},
	}

	return cmd
}

func setKvCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "set name.",
		RunE: func(cmd *cobra.Command, args []string) error {
			names := cmd.Flags().Args()
			if len(names) != 2 {
				return fmt.Errorf("invalid name or value.")
			}
			return setName(names[0], names[1])
		},
	}

	return cmd
}

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
