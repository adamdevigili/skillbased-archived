.PHONY: build-frontend
build-frontend:
	cd frontend; \
	yarn install; \
	NEXT_TELEMETRY_DISABLED=1 yarn run export

.PHONY: build
build: build-frontend
	go build -o bin/