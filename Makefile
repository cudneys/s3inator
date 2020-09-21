.DEFAULT_GOAL := build

clean:
	@rm -rvf dist

build: build_linux build_windows build_darwin

build_linux:
	@mkdir -p dist/linux
	GOOS=linux go build -o dist/linux/s3inator

build_windows:
	@mkdir -p dist/windows
	GOOS=windows go build -o dist/windows/s3inator.exe

build_darwin:
	@mkdir -p dist/osx
	GOOD=darwin go build -o dist/osx/s3inator
