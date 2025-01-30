BROKER_BINARY=brokerApp
USER_BINARY=userApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerApp
PRODUCT_COMMAND_BINARY=productCommandApp
PRODUCT_EVENTS_BINARY=productEventsApp
PRODUCT_QUERY_BINARY=productQueryApp

build_broker:
	@echo Building broker service...
	cd ./broker-service&& set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${BROKER_BINARY} -tags timetzdata
	@echo Done!

build_user:
	@echo Building user service...
	cd ./user-service&& set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${USER_BINARY}
	@echo Done!

build_auth:
	@echo Building user service...
	cd ./auth-service&& set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${AUTH_BINARY}
	@echo Done!

build_logger:
	@echo Building logger service...
	cd ./logger-service&& set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${LOGGER_BINARY} 
	@echo Done!

build_product_command:
	@echo Building product command service...
	cd ./product-command-service&& set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${PRODUCT_COMMAND_BINARY} 
	@echo Done!

build_product_events:
	@echo Building product events service...
	cd ./product-events-service&& set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${PRODUCT_EVENTS_BINARY} 
	@echo Done!

build_product_query:
	@echo Building product query service...
	cd ./product-query-service&& set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${PRODUCT_QUERY_BINARY} 
	@echo Done!

up_build: build_broker build_user build_auth build_logger build_product_command build_product_events build_product_query
	@echo Running all services...
	docker compose -f docker-compose.yaml -p myapp up --build -d 
	@echo Done!

up:
	@echo Running all services...
	docker compose -f docker-compose.yaml -p myapp up -d  
	@echo Done!

down:
	@echo Stopping and removing all services...
	docker compose -p myapp down
	@echo Done!