os-archs ?= \
	windows:amd64 \
	windows:arm64
module = $(shell go list -m)

.PHONY: build
build:
	for os_arch in $(os-archs); do \
		os=$$(echo $$os_arch | cut -d: -f1); \
		arch=$$(echo $$os_arch | cut -d: -f2); \
		out="$(module)-$$os-$$arch"; \
		[ "$$os" = "windows" ] && out="$$out.exe"; \
		echo "Building $$out"; \
		CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch \
			go build -trimpath -ldflags "-s -w -H=windowsgui" -o ./dist/$$out; \
	done
