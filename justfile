release:
    go build -o autoinstall.exe -ldflags="-w -s" -trimpath cmd/main.go