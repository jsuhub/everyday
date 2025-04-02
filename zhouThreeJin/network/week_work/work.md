## ssh
握手：  
![](./ssh/1.png)
![](./ssh/2.png)  
1. 建立在TCP三次握手的基础上  
2. 客户端和服务端协商版本  
3. 加密算法协商  
4. 密钥交换  
5. 服务器身份验证  
6. 客户端身份验证    
7. 安全通信  

挥手：  
1.客户端发送 SSH_MSG_DISCONNECT 表明要关闭回话了 
2.TCP 四次挥手  









## http  
三次握手建立连接：  
![](./http/1.jpeg)  
第一步 发送请求 HTTP 头部中FLAG为2 的SYN的信号报文表示要建立连接 Seq=0   
第二步 响应方返回响应报文 FLAG为12 的SYN+ ACK报文   Seq=0 Ack=请求报文的Seq+1    
第三部 请求方发送 响应报文 FLAG为10 的ACK 报文 Seq=1 Ack=1  
TCP窗口升级可以互相通信  

交流采用：请求-响应模式
![](./http/3.jpeg)  
双向箭头表示  


四次挥手：  
![](./http/2.jpeg)  
第一次 客户端给服务端发送 标志为11 的FIN 结束标志和ACK对上一个通信的确认  Seq=101
第二次 服务端给客户端发送 标志为10 的ACK 确认标志为，此时服务器即将关闭这条通道 Ack=102 Seq=123    
第三次 服务端给客户端  标志为11 的FIN 结束标志和 ACK 表示此事服务器已经快要关闭了 Ack=102 Seq=123  
第四次 客户端给服务器 标志为10 的Ack 确认报文 表示客户端已经知道服务器已经关闭seq=102 Ack=124 服务器接收到这条消息后就会关闭这条通道  

## https 
HTTP+TLS 构成HTTPS  
网站证书在CA机构去申请或者使用  
```
openssl req -x509 -newkey rsa:2048 -keyout server.key -out server.crt -days 365 -nodes \
-subj "/C=US/ST=CA/L=San Francisco/O=MyCompany/CN=localhost" \
-addext "subjectAltName = DNS:localhost,IP:127.0.0.1"
# 定义了证书可用于哪些主机名（DNS 或 IP 地址）。
```   
在当前目录下生成：server.key (私钥)、server.crt (证书)    
其中服务器中包含了 证书和私钥 ，而客户端储存的只有证书（一开始就已经被保存在客户端中了）

五次握手：  
![](./https/img/1.png) 
1~3次 建立TCP连接 
第一次 客户端向服务器发送SYN  
第二次 服务器向客户端回应ACK 和自己的SYN 
第三次 客户端向服务器响应 ACK   建立TCP连接成功  

第四次 客户端发送TCP升级—— Client hello（里面包含了TLS的版本等信息）  
第五次 服务器回应请求—— Server hello （表示服务器响应、里面就包含了服务器证书等信息）  
第六次 服务器向客户端发送密钥交换—— key Exchaneg (客户端生成对称密钥)
第七次 客户端对服务器的请求作出反应并接下来的通信都使用TLS—— Change Cipher Spec（ 切换到加密模式、验证出）
第八次 客户端回应握手完成进行HTTPS—— Finished 


挥手阶段：  
![](./https/img/2.png)  
第一次 客户端向服务器发送 ：Close Notify  
第二次 服务器回应客户端： Close Notify   
TCP的四次挥手 图中最后一次RST表示非正常关闭  