#!/bin/bash
#run
if [ "$1" == "run" ];then
CGO_ENABLED=1 go build
./sqlite3 $*
fi

#linux
if [ "$1" == "linux" ];then
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=1 GOOS=linux GOARCH=386 go build
fi


#windows
if [ "$1" == "windows" ];then
CC=x86_64-w64-mingw32-gcc GOARCH=amd64 CGO_ENABLED=1 GOOS=windows go build  
CC=i686-w64-mingw32-gcc-win32 GOARCH=386 CGO_ENABLED=1 GOOS=windows go build  
fi
#darwin
if [ "$1" == "darwin" ];then
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=darwin GOARCH=adm64 go build 
CGO_ENABLED=1 GOOS=darwin GOARCH=386 go build 
fi 