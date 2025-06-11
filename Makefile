BINARY_NAME=MarkDown.app
APP_NAME=MarkDown
VERSION=1.0.0

build-darwin:
	rm -rf ${BINARY_NAME}
	fyne package -os darwin

tidy:
	go mod tidy