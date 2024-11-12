package aireqlog

import (
	"fmt"
	"net"
	"os"
	"os/user"
	"runtime"
)

func getSystemIdentity() (SystemIdentity, error) {
	var out SystemIdentity
	out.Hostname, _ = os.Hostname()
	out.OperatingSystem = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	out.MACAddress = getMACAddress()

	usr, err := user.Current()
	if err != nil {
		return out, fmt.Errorf("getting user: %w", err)
	}
	out.Username = usr.Username

	return out, nil
}

func getMACAddress() string {
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		if len(iface.HardwareAddr) > 0 {
			return iface.HardwareAddr.String()
		}
	}
	return ""
}
