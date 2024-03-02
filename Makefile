CMD_FILES=$(shell ls cmd)
CONF_FILES=$(shell find config)

build:
	@go mod tidy
	@for i in $(shell ls cmd) ; do \
		cd cmd/$$i ; \
		echo "Building $$i" ; \
		go build -o ../../bin/uq-$$i main.go && echo 'Build successful' ; \
		cd ../.. ; \
	done


install: build
	@chmod 755 bin/*
	@chmod 755 scripts/*
	@cp -fp bin/* /usr/bin/
	@cp -fp scripts/* /usr/bin

uninstall:

clean:
	rm bin/*

sync:
	@for i in $(shell find configs/root); do \
		host_path=$$(echo "$$i" | cut -c13-) ; \
		if [[ -f $$i ]]; then \
			echo "Linking $$host_path to $$i" ; \
			ln -f $$host_path $$i ; \
		fi ; \
	done

	@for i in $(shell find configs/home); do \
		host_path="$$HOME/$$(echo "$$i" | cut -c14-)" ; \
		if [[ -f $$i ]]; then \
			echo "Linking $$i to $$host_path" ; \
			ln -f $$i $$host_path ; \
		fi ; \
	done


config: 
	@for i in $(shell find configs/root); do \
		path=$$(echo "$$i" | cut -c13-) ; \
		echo $$i $$path ; \
		if [[ -d $$i ]]; then \
			echo "Creating directory $$i" ; \
			mkdir -p $$path ; \
		elif [[ -f $$i ]]; then \
			echo "Linking $$i to $$path" ; \
			ln -f $$i $$path ; \
		fi \
	done
	@for i in $(shell find configs/home); do \
		path="$$HOME/$$(echo "$$i" | cut -c14-)" ; \
		echo $$i $$path ; \
		if [[ -d $$i ]]; then \
			echo "Creating directory $$i" ; \
			mkdir -p $$path ; \
		elif [[ -f $$i ]]; then \
			echo "Linking $$i to $$path" ; \
			ln -f $$i $$path ; \
		fi \
	done