test:
	GOARCH=arm64 go build main.go
	ssh nas killall ./main
	scp main nas:main
	ssh nas ./main
