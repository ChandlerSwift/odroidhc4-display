test:
	GOARCH=arm64 go build main.go
	ssh nas killall ./main || true
	scp main nas:main
	ssh nas ./main
