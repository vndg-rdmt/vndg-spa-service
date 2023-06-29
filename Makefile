APP_NAME		= vndg-service
CMD_DIR			= ./cmd
DEST_DIRECTORY	= ./bin

clean:
	@rm -r $(DEST_DIRECTORY)/*

build:
	@go build -o $(DEST_DIRECTORY)/$(APP_NAME) $(CMD_DIR)/$(APP_NAME);
	
run: build
	@$(DEST_DIRECTORY)/$(APP_NAME)
