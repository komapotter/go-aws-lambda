APP := handler

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -o ${APP}

.PHONY: clean
clean:
	@rm -fr ${APP}.zip ${APP}

.PHONY: deps
deps:
	@dep ensure -v

.PHONY: zip
zip:
	@zip ${APP}.zip ${APP}

.PHONY: dist
dist: clean build zip
