# 動機
強化自身對於反向代理以及封包傳送的觀念，架設一個反向代理的環境來進行實踐

# 目的
使用nginx作為reverse proxy server，反向代理一個websocket server，在請求的過程中並使用tcpdump去觀察封包的傳輸狀況
- 請求使用工具: chrome-websocket_client-extension
- 抓封包工具
> tcpdump -i any -s0 -nnvvvA
or 
> tcpdump -i any -w ./test.pcap; wireshark test.pcap

# 備註
1. 使用docker-compose.yml中的external參數去append一個既有的網路
2. 使用SNI去同時連線到多個virtualhost

# refer
- https://www.nginx.com/blog/websocket-nginx/
- http://fu7771.blogspot.com/2017/09/nginx-websocket.html

- nginx遇到499錯誤
https://www.jianshu.com/p/aa5a06fef39c
