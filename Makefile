all: linux windows

linux:
	CGO_ENABLED=0 GOOS=linux go build -o sq

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o sq.exe

clean:
	rm sq sq.exe
