# mess-of-linux-configurations

Just bunch of different configs and scripts, which I use for my desktop. 

Repo consists of a few folders according with go compiler project structure.

cmd contains source code of simple cli tools with using pkg libs

scripts contains shell scripts to build all cmd tools and to link all system and user configs to configs from repo.

Configs directory contains my configs. root contains configs outside user home directory and home contains configs inside user home directory.

Repo has a Makefile, there are some possible targets:
- build - build my go help tools, put them to bin directory
- config - link all configs to host system
- sync - link all according configs to repo configs
- install - copy binaries and scripts to /usr/bin/
