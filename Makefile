HELM_PLUGIN_DIR ?= $(shell helm env | grep HELM_PLUGINS | cut -d\" -f2)/helm-ssm-sm
HELM_PLUGIN_NAME := helm-ssm-sm
DIST := $(shell pwd)/_dist



.PHONY: build
build:
	go build -o bin/${HELM_PLUGIN_NAME}


.PHONY: test
test:
	go test -v ./internal


.PHONY: dist
dist:
	mkdir -p ${DIST}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${HELM_PLUGIN_NAME}
	tar -zcvf $(DIST)/${HELM_PLUGIN_NAME}-linux.tgz ${HELM_PLUGIN_NAME} README.md LICENSE plugin.yaml
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${HELM_PLUGIN_NAME}
	tar -zcvf $(DIST)/${HELM_PLUGIN_NAME}-linux-arm.tgz ${HELM_PLUGIN_NAME} README.md LICENSE plugin.yaml
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${HELM_PLUGIN_NAME}
	tar -zcvf $(DIST)/${HELM_PLUGIN_NAME}-macos.tgz ${HELM_PLUGIN_NAME} README.md LICENSE plugin.yaml
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${HELM_PLUGIN_NAME}.exe
	tar -zcvf $(DIST)/${HELM_PLUGIN_NAME}-windows.tgz ${HELM_PLUGIN_NAME}.exe README.md LICENSE plugin.yaml

.PHONY: install
install: dist
	if [ ! -f .version ] ; then echo "dev" > .version ; fi
	mkdir -p $(HELM_PLUGIN_DIR)
	if [ "$$(uname)" = "Darwin" ]; then file="${HELM_PLUGIN_NAME}-macos"; \
 	elif [ "$$(uname)" = "Linux" ]; then file="${HELM_PLUGIN_NAME}-linux"; \
	else file="${HELM_PLUGIN_NAME}-windows"; \
	fi; \
	mkdir -p $(DIST)/$$file ; \
	tar -xf $(DIST)/$$file.tgz -C $(DIST)/$$file ; \
	cp -r $(DIST)/$$file/* $(HELM_PLUGIN_DIR) ;\
	rm -rf $(DIST)/$$file


