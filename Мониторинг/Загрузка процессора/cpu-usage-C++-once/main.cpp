#include <iostream>
#include <fstream>
#include <string>
#include <csignal>
#include <ranges>


int main() {

    std::ifstream in("/proc/stat");
    std::string str;
        in.seekg(std::ios_base::beg);
        std::getline(in, str);

        str = str.substr(5, str.size() - 5);
        int first[10];
        int second[10];

        auto size = str.size();
        int *pointer = first;
        std::string buf;
        for (int i = 0; i < size; ++i) {
            if (str[i] != ' ') {
                buf.push_back(str[i]);
            } else {
                *pointer = std::stoi(buf);
                buf.clear();
                pointer++;
            }
        }
        *pointer = std::stoi(buf);

        int total_first = 0;
        int work_first = 0;

        for (int i = 0; i < 7; i++) {
            total_first += first[i];
        }

        for (int i = 0; i < 3; i++) {
            work_first += first[i];
        }

        sleep(1);

        in.seekg(std::ios_base::beg);

        std::getline(in, str);

        str = str.substr(5, str.size() - 5);

        size = str.size();
        pointer = second;
        buf.clear();
        for (int i = 0; i < size; ++i) {
            if (str[i] != ' ') {
                buf.push_back(str[i]);
            } else {
                *pointer = std::stoi(buf);
                buf.clear();
                pointer++;
            }
        }
        *pointer = std::stoi(buf);

        int total_second = 0;
        int work_second = 0;

        for (int i = 0; i < 7; i++) {
            total_second += second[i];
        }

        for (int i = 0; i < 3; i++) {
            work_second += second[i];
        }

        printf("%.1f%\n", (100 * (work_second - work_first) / (double) (total_second - total_first)));

}
