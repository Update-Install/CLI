go install

go build -ldflags "-s -w" -o ./dist/ui-cli_linux_amd64_uncompressed

upx -f --best --lzma ./dist/ui-cli_linux_amd64_uncompressed -o ./dist/ui-cli_linux_amd64
