export GOPROXY=https://goproxy.io,direct
export GO111MODULE=on

all: develop

# 开发环境
develop:
	@echo "开发环境"
	@go build -o cmd/gogeekbang -v .
	@cp conf/app.ini cmd/app.ini