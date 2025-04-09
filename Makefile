.PHONY: test
test:
	cd heapmap && go test -v -count=1 ./...
	cd priorityquery -v && go test -v -count=1 ./...

race:
	go test -v -race -count=100 ./...

.PHONY: coverHeapMap
coverHeapMap:
	cd heapmap && go test -short -count=1 -coverpkg=./... -race -coverprofile=coverage.out ./...
	cd heapmap && go tool cover -html=coverage.out
	rm ./heapmap/coverage.out

.PHONY: coverHeapSlice
coverHeapSlice:
	cd heapslice && go test -short -count=1 -coverpkg=./... -race -coverprofile=coverage.out ./...
	cd heapslice && go tool cover -html=coverage.out
	rm ./heapslice/coverage.out

.PHONY: precommitInstall
precommitInstall:
	pre-commit install

.PHONY: precommitRunAll
precommitRunAll:
	pre-commit run --all-files

.PHONY: linterAll
linterAll:
	golangci-lint run --fix