GO_FILE=sql-migrator.go
QUERIES_DIR=queries
# Target to run the Go program
run:
	@echo "Running Go file: $(GO_FILE)"
	go run $(GO_FILE) $(QUERIES_DIR) 

# Target to build the Go program
build:
	@echo "Building Go file: $(GO_FILE)"
	go build -o main $(GO_FILE)

# Target to clean up built files
clean:
	@echo "Cleaning up..."
	rm -f main
