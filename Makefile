test:
	go test ./lang/*

coverage:
	go test -v -coverprofile cover.out ./lang/*

coverage-html: coverage
	go tool cover -html=cover.out

