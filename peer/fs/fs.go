package fs

import (
	"fmt"

	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

const (
	fsFuncName = "fs"
)

var (
	logger = logging.MustGetLogger("fsCmd")

	fsCmd = &cobra.Command{
		Use:   fsFuncName,
		Short: fmt.Sprintf("%s specific commands.", fsFuncName),
		Long:  fmt.Sprintf("%s specific commands.", fsFuncName),
	}

	srvAddr string
)

// Cmd returns the cobra command for Node
func Cmd() *cobra.Command {
	flags := fsCmd.PersistentFlags()

	flags.StringVarP(&srvAddr, "addr", "a", "http://localhost:9375", "server address")

	fsCmd.AddCommand(queryCmd())
	fsCmd.AddCommand(setKvCmd())
	fsCmd.AddCommand(syncFsCmd())
	fsCmd.AddCommand(initFsCmd())

	return fsCmd
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

func syncFsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "sync files from remote.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

func initFsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "init file system.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
