build:
	$(info Building...)
	(cd go/exports ; gopherjs build -o ../../js/dist/asu.js --minify)
	yarn --cwd js build
	cp js/dist/asu.d.ts js/dist/browser/asu.d.ts

test-js: build
	$(info Running JS tests...)
	yarn --cwd js test

test-go:
	$(info Running Go tests...)
	(cd go ; go test ./...)

test: test-go test-js

cover:
	mkdir -p go/coverage
	(cd go ; go test -coverprofile=coverage/cover.out)
	(cd go ; go tool cover -html=coverage/cover.out -o=coverage/cover.html)
	wslview go/coverage/cover.html
