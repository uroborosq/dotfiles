#/usr/bin/python3
import time
tctl = open('/sys/module/k10temp/drivers/pci:k10temp/0000:00:18.3/hwmon/hwmon4/temp1_input', 'r')
edge = open('/sys/module/amdgpu/drivers/pci:amdgpu/module/drivers/pci:amdgpu/0000:03:00.0/hwmon/hwmon6/temp1_input', 'r')
sensor_1 = open('/sys/module/nvme/drivers/pci:nvme/0000:02:00.0/nvme/nvme0/hwmon3/temp2_input', 'r')
sensor_2 = open('/sys/module/nvme/drivers/pci:nvme/0000:02:00.0/nvme/nvme0/hwmon3/temp3_input', 'r')

#while True:
print(f"{tctl.readline()[:2]}°C {edge.readline()[:2]}°C {sensor_1.readline()[:2]}°C {sensor_2.readline()[:2]}°C")
tctl.seek(0)
edge.seek(0)
sensor_1.seek(0)
sensor_2.seek(0)
#    time.sleep(2)