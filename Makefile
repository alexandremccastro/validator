
test:
	- go test .

test-cov:
	- go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
