#!/bin/bash

exec air --build.cmd 'go build -o ./tmp/megafactory.exe main.go' --build.bin 'tmp\megafactory.exe' --build.log './tmp/logs.txt' -build.stop_on_error true -build.include_ext go,html,js,css,svg -proxy.enabled true -proxy.proxy_port 8081 -proxy.app_port 8080
