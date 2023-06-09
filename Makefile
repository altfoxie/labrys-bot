build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux-gnu" CXX="zig c++ -target x86_64-linux-gnu" go build -trimpath -ldflags "-s -w" --tags "fts5"