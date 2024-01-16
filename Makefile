pwd := $(shell pwd)
project := $(notdir $(pwd))

OS := $(shell go env GOOS)


.PHONY:build linux win all clean upx install-lint lint callvis-install callvis proxy pprof 

# 这个要是可以读取你的docker info就好了不是吗 ? 其实是可以写的就是说 
DockerImageVersion := $(shell docker image ls | rg $(project) | rg -o '\d+\.\d+\.\d+')
# 然后正则表达式一下 就可以获取到版本号了

# DockerImageVersion = $(if $(1),$(1),0.0.0)

# 1.2.14
VersionNumber := $(shell echo $(DockerImageVersion) | rg -o '\d+' )
# 1.2
MajorVersionNumber := $(shell echo $(DockerImageVersion) | rg -o '\d+\.\d+' )
# 14
LastVersionNumber := $(shell echo $(DockerImageVersion) | rev | rg  -o '^\d+' | rev)
NewLastVersionNumber :=  $(shell echo $$(($(LastVersionNumber)+1)))
ImageVersion := $(MajorVersionNumber).$(NewLastVersionNumber)








# VersionNumber = $(eval $(LastVersionNumber) + 1)



# DockerImageVersion := $($(MajorVersion).$(VersionNumber))

# 如果ImageVersion 是空字符串



imageInfo :
	# echo $(LastVersionNumber)

run: 
	@if [ $(OS) = "windows" ]; then \
		./$(project).exe;\
	else \
		./$(project) ;\
	fi


build:
	@if [ $(OS) = "windows" ]; then \
		go build -ldflags "-s -w" -o $(project).exe; \
		echo "build $(project).exe"; \
	else \
		go build -ldflags "-s -w" -o $(project); \
		echo "build $(project)"; \
	fi



release: linux docker-build win



linux:
	@set  GOOS=linux
	@go build -ldflags "-s -w" -o $(project)

win:
	@set GOOS=win
	@go build -ldflags "-s -w" -o $(project).exe
	@upx -9 $(project).exe

all:linux win

withoutwindow:
	@go build -ldflags "-s -w -H=windowsgui" 


static:
	CGO_ENABLED=0 
	GOOS=linux 
	go build -a -installsuffix cgo -o $(project) .

clean:
	-rm -f *.log
	-rm -f $(project)
	-rm -f *.exe
	-rm -f *.pprof
	-rm -f *.txt
	-rm -f main

rm: clean

# tools-chain for golang

fmt: 	
	@go fmt ./...

	
upx:build
	upx -9 $(project).exe

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	golangci-lint run

callvis-install:
	go install github.com/ofabry/go-callvis@master

callvis:
	go-callvis main.go

proxy:
	go env -w  GOPROXY=https://goproxy.io,direct

pprof:run
	go tool pprof -http=:8080 *.pprof


# Argument Check 
project:
	@echo $(project)


docker-clean-container:
	docker container prune -f
docker-remove: docker-clean-container
	docker image rm ${project}:${DockerImageVersion}

# Docker 
docker-build: docker-remove
	docker build -t ${project}:${ImageVersion} . 

# 要是可以查询就好了不是吗?


# docker-clean:
# 	docker image prune -f 
# 	docker container  prune -f
