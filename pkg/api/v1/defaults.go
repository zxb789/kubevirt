package v1

import (
	"github.com/pborman/uuid"
	"k8s.io/apimachinery/pkg/types"
)

var _true = t(true)
var _false = t(false)

func SetDefaults_TimerAttrs(obj *TimerAttrs) {
	if obj.Enabled == nil {
		obj.Enabled = _true
	}
}

func SetDefaults_RTCTimerAttrs(obj *RTCTimerAttrs) {
	if obj.Enabled == nil {
		obj.Enabled = _true
	}
}

func SetDefaults_FeatureState(obj *FeatureState) {
	if obj.Enabled == nil {
		obj.Enabled = _true
	}
}

func SetDefaults_FeatureVendorID(obj *FeatureVendorID) {
	if obj.Enabled == nil {
		obj.Enabled = _true
	}
}

func SetDefaults_DiskDevice(obj *DiskDevice) {
	if obj.Disk == nil &&
		obj.CDRom == nil &&
		obj.Floppy == nil &&
		obj.LUN == nil {
		obj.Disk = &DiskTarget{}
	}
}

func SetDefaults_InterfaceDevice(obj *InterfaceDevice) {
	if obj.E1000 == nil &&
		obj.VIRTIO == nil &&
		obj.RTL8139 == nil {
		obj.RTL8139 = &RTL8139Interface{}
	}
}

func SetDefaults_Watchdog(obj *Watchdog) {
	if obj.I6300ESB == nil {
		obj.I6300ESB = &I6300ESBWatchdog{}
	}
}

func SetDefaults_CDRomTarget(obj *CDRomTarget) {
	if obj.ReadOnly == nil {
		obj.ReadOnly = _true
	}
	if obj.Tray == "" {
		obj.Tray = TrayStateClosed
	}
}

func SetDefaults_FloppyTarget(obj *FloppyTarget) {
	if obj.Tray == "" {
		obj.Tray = TrayStateClosed
	}
}

func SetDefaults_FeatureSpinlocks(obj *FeatureSpinlocks) {
	if obj.Enabled == nil {
		obj.Enabled = _true
	}
	if obj.Enabled == _true && obj.Retries == nil {
		obj.Retries = ui32(4096)
	}
}

func SetDefaults_I6300ESBWatchdog(obj *I6300ESBWatchdog) {
	if obj.Action == "" {
		obj.Action = WatchdogActionReset
	}
}

func SetDefaults_Firmware(obj *Firmware) {
	if obj.UID == "" {
		obj.UID = types.UID(uuid.NewRandom().String())
	}
}

func SetDefaults_VirtualMachine(obj *VirtualMachine) {
	if obj.Spec.Domain == nil {
		// FIXME we need proper validation instead of this
		obj.Spec.Domain = NewMinimalDomainSpec()
	}
	if obj.Spec.Domain.Firmware == nil {
		obj.Spec.Domain.Firmware = &Firmware{}
	}

	if obj.Spec.Domain.Features == nil {
		obj.Spec.Domain.Features = &Features{}
	}
}

func t(v bool) *bool {
	return &v
}

func ui32(v uint32) *uint32 {
	return &v
}
