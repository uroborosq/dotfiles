#include <iostream>
#include <fstream>
#include <csignal>


int main() {
    std::ifstream is;
    char before_dot[2];
    char after_dot;
    while (true)
    {
    
    
        is.open("/sys/module/k10temp/drivers/pci:k10temp/0000:00:18.3/hwmon/hwmon4/temp1_input");
        is.read((char *) before_dot, 2);
        is.read(&after_dot, 1);
        printf("%s.%c°C ", before_dot, after_dot);
        is.close();
        is.open("/sys/module/amdgpu/drivers/pci:amdgpu/module/drivers/pci:amdgpu/0000:03:00.0/hwmon/hwmon6/temp1_input");
        is.read((char *) before_dot, 2);
        is.read(&after_dot, 1);
        is.close();
        printf("%s.%c°C ", before_dot, after_dot);
        is.open("/sys/module/nvme/drivers/pci:nvme/0000:02:00.0/nvme/nvme0/hwmon3/temp3_input");
        is.read((char *) before_dot, 2);
        is.read(&after_dot, 1);
        printf("%s.%c°C\n", before_dot, after_dot);
        is.close();
        sleep(2);
    }
    return 0;
}
