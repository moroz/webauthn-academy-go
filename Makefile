install:
	which modd || go install github.com/cortesi/modd/cmd/modd@latest
	which templ || go install github.com/a-h/templ/cmd/templ@latest
	which goose || go install github.com/pressly/goose/v3/cmd/goose@latest
	which dlv || go install github.com/go-delve/delve/cmd/dlv@latest
	which pnpm || npm i -g pnpm
	cd assets && pnpm install && cd ..
	go mod download

guard-%:
	@ test -n "${$*}" || (echo "FATAL: Environment variable $* is not set!"; exit 1)

db.test.prepare: guard-TEST_DATABASE_NAME guard-TEST_DATABASE_URL
	@ createdb ${TEST_DATABASE_NAME} 2>/dev/null || true
	@ env GOOSE_DBSTRING="${TEST_DATABASE_URL}" goose up

test: db.test.prepare
	go test -v ./...
