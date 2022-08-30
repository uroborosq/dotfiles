#!/bin/bash
while true; do
    echo "$(cat /sys/module/k10temp/drivers/pci:k10temp/0000:00:18.3/hwmon/hwmon4/temp1_input | cut -c1-2)°C $(cat /sys/module/amdgpu/drivers/pci:amdgpu/module/drivers/pci:amdgpu/0000:03:00.0/hwmon/hwmon6/temp1_input  | cut -c1-2)°C $(cat /sys/module/nvme/drivers/pci:nvme/0000:02:00.0/nvme/nvme0/hwmon3/temp2_input  | cut -c1-2)°C $(cat /sys/module/nvme/drivers/pci:nvme/0000:02:00.0/nvme/nvme0/hwmon3/temp3_input  | cut -c1-2)°C"
    sleep 2
done