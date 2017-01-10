package farmer

import (
	"github.com/hyperledger/fabric/farmer/api"
	"github.com/hyperledger/fabric/farmer/daemon"
)

func StartFarmer() {
	d, err := daemon.LoadDaemon()
	if err != nil {
		panic(err)
		return
	}
	d.GetLogger().Debug("Start farmer server.")

	go func() {
		if err := d.StartPeer(); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := api.Serve(d); err != nil {
			panic(err)
		}
	}()
	d.WaitExit()
}

func StopFarmer() error {

	return nil
}
