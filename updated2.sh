#!/bin/bash

cd ./mung
tar -zxvf ./mung.tar.gz

# 给目标文件加权
chmod +x ./mung   

kill -9 "$(pgrep -f mung)" && nohup ./mung >/dev/null 2>&1 &
