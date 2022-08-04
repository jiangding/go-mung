"# go-mung" 


绿豆科技小程序 +
bingo 后端


prometheus  + grafana  监控

部署的 sh文件

```html 
tar -zxvf .....

/bin/cp -rf ./* ../mung/


// 无输出的linux后台运行 执行./mung
kill -9 12345
nohup /root/go/src/mung/mung >/dev/null 2>&1 &

kill -9 "$(pgrep -f mung)" && nohup /root/go/src/mung/mung >/dev/null 2>&1 &



// prometheus 启动,监听端口 , 热加载
./prometheus --config.file=prometheus.yml --web.listen-address="0.0.0.0:3333" --web.enable-lifecycle

// grafana 启动，状态
systemctl start grafana-server.service 
systemctl restatus grafana-server.service 
systemctl status grafana-server.service 
```