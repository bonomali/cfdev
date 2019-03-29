package kvm

import (
	"code.cloudfoundry.org/cfdev/driver"
)

// consider moving to more native go implementation
// rather than shelling out
func (d *KVM) setupNetworking(tapDevice, bridge string) {
	d.SudoShell.Run("ip", "tuntap", "add", "dev", tapDevice, "mode", "tap")
	d.SudoShell.Run("ip", "link", "set", tapDevice, "master", bridge)
	d.SudoShell.Run("ip", "link", "set", "dev", bridge, "up")
	d.SudoShell.Run("ip", "link", "set", "dev", tapDevice, "up")
}

func (d *KVM) setupRoutes(ip string) {
	d.SudoShell.Run("ip", "route", "add", driver.ContainerSubnet, "via", ip)
}

func (d *KVM) teardownRoutes() {
	d.SudoShell.Run("ip", "route", "flush", driver.ContainerSubnet)
}

func (d *KVM) teardownNetworking(tapDevice string) {
	d.SudoShell.Run("ip", "link", "set", "dev", tapDevice, "down")
	d.SudoShell.Run("ip", "link", "del", "dev", tapDevice)
}