#!/bin/bash

# GET MACHINE_UUID with VBoxManage vms list

MACHINE_UUID="22b58b18-c9b9-4246-afb2-fb11a23cf8fc"

VBoxManage modifyvm $MACHINE_UUID --cpuidset 00000001 000106e5 00100800 0098e3fd bfebfbff
VBoxManage setextradata $MACHINE_UUID "VBoxInternal/Devices/efi/0/Config/DmiSystemProduct" "iMac11,3"
VBoxManage setextradata $MACHINE_UUID "VBoxInternal/Devices/efi/0/Config/DmiSystemVersion" "1.0"
VBoxManage setextradata $MACHINE_UUID "VBoxInternal/Devices/efi/0/Config/DmiBoardProduct" "Iloveapple"
VBoxManage setextradata $MACHINE_UUID "VBoxInternal/Devices/smc/0/Config/DeviceKey" "ourhardworkbythesewordsguardedpleasedontsteal(c)AppleComputerInc"
VBoxManage setextradata $MACHINE_UUID "VBoxInternal/Devices/smc/0/Config/GetKeyFromRealSMC" 1
VBoxManage setextradata $MACHINE_UUID "VBoxInternal2/EfiGopMode 4"
