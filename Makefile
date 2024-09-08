download:
	go mod download

install.tools: download
	@echo Installing tools from tools.go
	@grep _ tools.go | awk -F'"' '{print $$2}' | xargs -tI % go install %

install.assets:
	which pnpm || npm i -g pnpm
	cd assets && pnpm install && cd ..
	go mod download

install: install.tools install.assets

guard-%:
	@ test -n "${$*}" || (echo "FATAL: Environment variable $* is not set!"; exit 1)

db.test.prepare: guard-TEST_DATABASE_NAME guard-TEST_DATABASE_URL
	@ createdb ${TEST_DATABASE_NAME} 2>/dev/null || true
	@ env GOOSE_DBSTRING="${TEST_DATABASE_URL}" goose up

test: db.test.prepare
	go test -v ./...
