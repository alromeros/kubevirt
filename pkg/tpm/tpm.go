package tpm

import v1 "kubevirt.io/api/core/v1"

func HasDevice(vmiSpec *v1.VirtualMachineInstanceSpec) bool {
	return vmiSpec.Domain.Devices.TPM != nil &&
		(vmiSpec.Domain.Devices.TPM.Enabled == nil || *vmiSpec.Domain.Devices.TPM.Enabled)
}

func HasPersistentDevice(vmiSpec *v1.VirtualMachineInstanceSpec) bool {
	if !HasDevice(vmiSpec) {
		return false
	}
	persistent := vmiSpec.Domain.Devices.TPM.Persistent
	// When the declarative virtualMachineState API is used, providing a state PVC
	// implies the TPM state should be kept, so the device is persistent unless it is
	// explicitly opted out with persistent: false. See VEP #312.
	if vmiSpec.VirtualMachineState != nil {
		return persistent == nil || *persistent
	}
	return persistent != nil && *persistent
}
