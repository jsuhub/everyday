 # 基础内容分类  
从TCP/IP协议栈为依托，由上至下、从应用层到基础设施介绍协议。  
1.应用层：  
    HTTP/1.1  
    Websocket  
    HTTP/2.0  
2.应用层的安全基础设施  
    LTS/SSL  
3.传输层  
    TCP 
4.网络层及数据链路层  
    IP层和以太网  

# HTTP协议  

## 网络页面形成基本
流程：    
![](./img/png1.jpeg)  
其中当在浏览器上面的网址搜索界面点击搜索的时候如果出现了一些记录就相当于图片中右边蓝色的“浏览器引擎”去搜索数据存储，然后找到发现一些存储在浏览器中的数据。  
详细操作：  
![](./img/png2.jpeg) 

## 定义  
一种无状态的、应用层的、以请求/应答方式运行的协议，它使用可扩展的语义和自描述消息格式、与基于网络的超文本信息系统灵活的互动   

基于ABNF（元语言）语法的HTTP格式：  
1." "字符：用来分隔定义中的各个元素  
2."/"选择：表示多个规则都是可供选择的规则  
3."%c##-###":表示从字符 ##到 ###  
4."()"序列组合：将规则组合起来，视为单个元素  
5."m*n"不定量重复：  
*表示零个或更多元素  
1*表示1个或更多元素  
2*4表示两个或者至多4个元素  
6.[]可选序列：包体是可选的，可有可无的  
操作符为：  
![](./img/png3.jpeg)  


根据上面的规则来分析格式  
![](./img/img4.jpeg)  
HTTP-message：表示http消息  
start-line：起始行  
request-line/status-line:表示请求行或者响应行构成起始行  
request-line=method SP request-target SP HTTP-version CRLF:请求行有 方法 空格 请求路径 空格 HTTP版本 换行
status-line=HTTP-version  SP status-code SP reason-phrase CRLF ：响应行由 HTTP版本 空格 响应码（三位数字） 空格 字符串形式描述的原因  换行  

### 详细分析请求头  
method——常见的方法有：  
GET:主要的获取信息方法  
HEAD:类似GET方法，但服务器不发送BODY,用以获取HEAD元数据，幂等方法  
POST:常用于提交HTML FROM表单、新增资源等  
PUT:更新资源、带条件时是幂等方法  
DELETE:删除资源、幂等方法  
CONNECT:建立tunnel隧道  
OPTION:显示服务器对访问资源支持的方法，幂等方法————跨域    
TRACE:回显服务器手到的请求，用于定位问题。  
Linux中使用： `curl static.taohui.tech -X OPTIONS `可以查看允许使用的方法  
用于文档管理的  
PROPFIND:从Web资源中检索以XML格式存储的属性——查看目录    
PROPPATCH:在单个原子性动作中更改和删除资源的多个属性  
MKCOL:创建集合或者目录  
COPY:将资源从一个URI复制到另一个URI   
MOVE:将资源从一个URI移动到另一个URI  
LOCK:锁定一个资源  
UNLOCK:接触资源的锁定  

### 详细分析响应行  
![](./img/png12.jpeg)  
响应码的规范:  
**1XX**:表示请求已经被服务器接收到了，需要进一步处理才能完成更进一步的操作    
100 Continue:上传大文件前使用  
101 Switch Protocols:协议升级使用  
102 Proccessing:表示服务器已经收到了请求但是这个响应需要很长的时间处理，放置客户端超市。  

**2XX**:成功处理请求  
200 OK :成功返回响应  
201 Created: 有新资源在服务器端杯成功创建  
202 Accepted:服务器接收并开始处理请求，请求并没有处理完成。异步、需要很长时间处理的任务  
203 Non-Authoritative Information:当代理服务器修改了origin server 的原始响应包体时。  
204 NO Content:成功执行了请求且不携带响应包体，并暗示客户端无需更新当前的页面视图。  
205 Reset Contern:成功执行了请求切不携带响应包体，同时指明客户端需要更新当前页面视图。  
206 Partial Content:使用range协议时返回部分响应内容时的响应码,多线程下载  
207 Multi-Status:在WEBDAV协议汇
208 Already Reported:为了解决WEb协议    

**3XX**:重定向  
300 Multiple Choices:允许客户端主动的显示  
301 Moved Permanently:表示资源永久的重定向  
302 Found:表示资源临时的重定向。  
303 See Other:重定向到其他资源  
304 Not Modified:客户端可复用的缓存  
307 Temporary Redirect:明确重定向后请求方法必须与原请求方法相同  
308 Permanent Redirect:类似301 

**4XX**:客户端出现错误  
400 Bad Request:服务器认为客户端出现了错误，但不能明确判断那种错误  
401 Unauthorized: 用户认证信息缺失  
407 Proxy Authentication Required :对需要经有代理的请求，认证信息为通过代理服务器的验证  
403 Forbidden: 服务器理解请求的含义，但没有权限执行此请求  
404 Not Found: 服务器没有找到相对应的资源  
410 Gone :服务器没有找到对应的资源，且明确的知道该位置永久性找不到资源————对404补充    
405 Method Not Allowed:服务器不支持请求行中的method方法  
406 Not Acceptable:对客户端指定的资源表述不存在————语言不兼容   
408 Request Timeout:服务器接收请求超时  
409 Conflict:资源冲突  
411 Length Required:请求中未携带Content-Length头部  
412 Precondition Failed:条件类请求不满足时候返回  
413 Payload Too Large/Request Entity Too Large:请求的包体超出服务器能处理的最大限度  
414 URI Too Long:请求的URI超出服务器的最大长度  
415 Unsupported Media Type:上传的文件类型不支持  
416 Range Not Satisfiable: 无法提供Range请求中指定的那段包体  
417 Expectation Failed: 对于Expect 请求头部期待的情况无法满足  
421 Misdirected Request: 服务器认为该请求不该发给它  
426 Upgrade Required:服务器拒绝基于当前HTTP协议提供服务，必须要Upgrade告知客户端升级  
428 Precondition Required:用户请求中缺少条件类头部  
429 Too Many Requests:客户端发送请求的速率太快  
431 Request Header Fields Too Large:请求的头部大小超过限额  
451 Unavailable For Legal Reasons :由于法律原因资源不可访问  


**5XX**：表示服务端出现错误  
500 Internal Server Error :服务器内部错误  
501 Not Implemented: 服务器不支持实现请求所需要的功能  
502 Bad Gateway:代理服务器无法获取到合法响应  
503 Service Unavailable :服务器资源尚未准备好处理当前请求  
504 GateWay Timeout:代理服务器无法及时的从上游获得响应  
505 HTTP version Not Supported: 请求的HTTP协议版本不支持  
507 Insufficient Storage:服务器没有足够的空间处理请求————存在安全错误    
508 Loop Detected:访问资源是检测到循环  
511 Network Authentication Required:代理服务器发现客户端需要进行身份验证才能获取网络访问权限  




 
*（header-field CRLF）:表示0个或多个以CRFL为结尾的http头部 
header-field = field-name":"OWS field-value OWS 表示由一个头部名称 中间分隔符号为：和头部的值构成      
OWS = *(SP / HTAB ) : 表示由零个或者多个空格或者横向制表符构成    
field-name = token————表示头部字段的名字，它必须是一个token，即由一系列非控制、非空白字符组成，不能包含特殊字符如冒号、空格等。    
field-value = *(field-content/obs-fold) ————这是头部字段的实际内容，它可以包含各种字符    
**header-filed**是一种类似键值对的形式  

CRLF:空行    
[message-body]:表示后面可以有消息体，但他与头部必须中间隔了一个CRLF。  
message-body: 表示由二进制的包传递的  

## 利用telnet 工具查看  
`telnet www.taohui.pub 80`
![](./img/png5.jpeg)
对比上面发现发现  
**请求构成**：  
`GET SP request-target SP HTTP/1.1 CRLF ` 符合请求头的格式  
`Host : www.xxx.com ` 表示的就是header-field : field-value     

**响应构成**：
上面一大串就是status-line  
分隔开的单独的CRLF(空行，表示从这里开始这一行就没有了)  
最后的那一段就是message-body  

## 网络为什么要分层  
OSI 概念模型  
![](./img/png6.jpeg)  
应用层：解决业务问题--七层设备nginx等    
表示层：把网络中的消息转化成为应用层的东西（TLS/LLS加密）  
会话层：建立关系连接/握手————概念层    
传输层：解决进程与进程的关系，TCP协议等--四层设备  
网络层：确保在广域网中从一个ip发送到另一个ip上--三层设备  
数据链路层：利用mac地址连接到局域网的交换机的--二层设备  
物理层：物理介质   

分层的好处是：每一层只需要关注自己这一层的东西（封装）不需要关心其他层做了什么。这样技术迭代的时候就不用担心兼容问题。  
坏处：因为分层之间需要数据处理，就需要更长的时间影响了效率。  

每一层的详细处理：  

![](./img/png7.jpeg)  

## 架构风格  
### 1.数据流风格  
优点：简单性、可进化性、可扩展性、可配置性、可重用性  
![](./img/png8.jpeg)  
管道与过滤器：每个Filter都有输入端和输出，只能从输入端读取数据，处理后再从输出端输出数据  
统一接口的管道与过滤器：在PF上增加了统一接口的约束，所有的Filter过滤器必须具有同样的接口  

### 2.复制风格 
优点：用户可察觉的性能、可伸缩性、网络效率、可靠性  
![](./img/png9.jpeg)  
复制仓库（RR）：多个进程提供相同的服务，通过反向代理对外提供集中服务  
缓存：RR的变体，通过复制请求的结果，为后续请求服用  


### 3.分层风格
优点：简单性、可进化性、可伸缩性  
客户端服务器（CS）：由客户端触发请求，服务端监听到请求后响应，客户端一直等待收到响应后，会话结束  
分层系统（LS）:每一层为其上的层服务，并使用在其下的层所提供的服务，如：TCP/IP  
分层客户端服务器（LCS）：LS+CS,如正向代理/反向代理，从空间上分为外部层与内部层  
无状态、客户端服务器（CSS）:基于CS、服务器不允许有session state会话状态  
缓存、无状态、客户端服务器（C$SS）:提升性能  

### 4.移动代码风格
优点：可扩展性、网络效率  
虚拟机(VM)：分离指令与实现  
远程求值(REV):基于CS的VM，将代码发送至服务器执行  
按需代码（COD）:服务器在响应发回处理代码，在客户端执行  


### 5.点对点风格
优点：可进化性、可重用性、可拓展性、可配置性  
EBI:基于事件集成系统  

##  使用Chrome抓包  
![](./img/png10.jpeg)   
控制器：控制面板的外观与功能。  
过滤器：过滤请求列表中显示的资源  
概览：显示HTTP请求、响应的时间轴  
请求列表：默认时间排序，可选择显示列  
概要：请求总数、总数据量、总花费时间等  

### 浏览器的加载过程  
1.解析HTML结构  
2.加载外部脚本和样式表文件  
3.解析并执行脚本代码//部分脚本会阻塞页面的加载  
4.DOM树构建完成//DOMContentLoaded事件  
5.加载图片等外部文件  
6.页面加载完毕//load事件  

请求时间详细分布：  
Queueing:浏览器在排队  
Stalled:请求可能会因Queueing中描述的任何原因停止  
DNS Lookup:浏览器正在解析请求的IP地址  
Proxy Negotiation:浏览器正在与代理服务器协商请求  
Request sent :正在发送请求  
ServiceWorker Preparation:浏览器正在启动Service Worker  
Request to ServiceWork:正在将请求发送到Service Workder  
Waiting(TTFB)：将浏览器正在等待响应的第一个字节  
Content Download:浏览器正在接受响应  
Receiving Push:浏览器正在通过HTTP/2服务器推送接受此响应数据  
Reading Push:浏览器正在读取之前的本地数据  


## URI
定义：  
Uniform Resource Identifier 统一资源标识符  
组成：  
![](./img/png11.jpeg)  
schme, user information, host, port（HTTP默认43）, path, query,fragment（分段）      
用元语表示URI：  
URI= schme ":" hire-pary["?"query]["#"fragment]  
解释：  
schme=ALPHA*(ALPHA/DIGIT/"+"/"-"/".")————方案/协议后面跟：  
如：https,http,ftp,mailto,file,telnet等  
ALPHA表示字母  
DIGIT表示数字  

hire-pary="//" authority path-abempty / path-absolute /path-rootless /path-empty————（层次部分）：这个术语指的是URI中紧跟在方案标识符之后的部分，包含了用户信息、主机、端口以及路径  
 authority=[ userinfo"@"]host[":" port]————用户信息@域名:端口    
  userinfo=*(unreserved /pct-encoded /sub-delims/":")————这部分是可选的，用来存储登录服务器所需的用户名和密码  
  host=IP-literal/ipv4address/reg-name  
  port=*DIGIT(浏览器默认端口为80)  

query=*(pchar/"//"?")————以"?"开头的可选项，通常用来向服务器传递参数，以键值对的形式出现    

fragment=*(pchar/"/"/"?")————以"#"开头的可选项，用来指向文档内部的某个部分  

相对URI:去除绝对路径只留下/html/rfc7231?test=1#page-7  

## TCP编程角度看HTTP请求处理  
![](./img/png13.jpeg) 

事务：对应着一个响应和一个请求  
短连接：每次客户端和服务端完成一次数据交互之后就立即断开连接。下次再有请求时重新建立连接。  
长连接：客户端和服务端建立连接后，即使完成了数据的交互也不立即关闭连接，而是保持这个连接，以便后续的数据交互。  
区分长连接还是短连接靠的就是Connection头部：  
 Keep-Alive:长连接  
 Close:短连接  
 如果存在代理服务器的话，不转发Connection头部，只与当前连接有关  
 存在的问题中有：陈旧的服务器可能无法识别Connection的长连接，就会导致自动退化为短连接，但是服务器和客户端仍然认为发送的是长连接——————Proxy-Connection就是用来解决这个问题的，它就会替代Connection表示当前使用的是代理连接    

 ### HTTP消息在路由中的传播  
 Host 头部————出现在URL  
 定义为： uri-host[":"port]  
HTTP/1.1请求中必须要传递Host头部否则返回400错误响应码  
因为HPPT/1.0的域名比较少就不会出现多个域名对应一个服务器的情况，但是现在域名变多了，所以1.1中就规定了请求路径必须是绝对路径的形式，而不能是相对路径  

### 处理方法  
1.建立TCP连接  
2.读取请求行、请求头部  
3.服务器寻找理由中对应的URI的路径处理相应请求  
4.访问相应资源并且生成响应  
5.返回给客户端并且生成一个日志  

### 代理服务器的头部  
HTTP头部中会有X-Forwarded-For传递IP———  X-Forwarded-For中包含的是之前经过的公网ip（所以这个里面的第一个Ip就是起始ip）,还有一些采用的是X-Real-IP(nginx)传递ip    

Max-Forwards头部————表示的限制Proxy代理服务器的最大转发次数  
Max-Forwards=1*DIGIT  

Via头部————进过了那些代理服务器，版本和名称  
via-1#(receuved-protocol(代理服务器版本号)| RWS received-by（代理服务器名称） [RWS comment（代理器的注释）])  

Cache-Control:notarnsform——————表示禁止代理服务器修改响应包体  


### 表示上下文的HTTP头部  
User-Agent————指明客户端的类型信息，服务器可以根据对资源的表述做抉择  
通用：  
User-Agent=porduct*(RWS(prodect/comment))  
product=token["/"product-version]  
RWS=1*(SP-空格/HTAB-换行)  
具体：  
user-agent:
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36  
解释：Mozilla表示浏览器兼容的版本 （）里面表示的就是注释，比如基于那个操作系统上  
 AppleWebKit 表示的是浏览器的渲染引擎   
 Chrome、Safari表示是浏览器    

 Referer————当前请求来自另外一个页面浏览器就会自动添加这个页面到这个头部/是来自上一个请求的   
 referer:
https://www.csdn.net/  
表示我这个请求是来自csdn首页的然后跳转到这个页面——可以用于防爆链    
不会添加Refer情况：    
当源页面采用的协议表示本地文件的"file"或者"data" URI  
当前请求页面采用的是http协议，而来源页面采用的是https协议  
用处：用于统计分析、缓存优化、防盗链等功能  

From————主要告诉服务器的如何通过邮箱联系到爬虫的负责人  
  
Server:指明服务器上所用软件的信息，用于帮助客户端定位问题或者统计数据  
通用：  
Server=product*(RWS(product/comment))  
product=token["/"product-version]  
RWS=1*(SP-空格/HTAB-换行)  
具体：  
server: elb
server:nginx  

Allow————告诉客户端，服务器上该URI对应的资源允许那些方法执行  
通用：  
Allow=#method  
具体：  
Allow:GET、HEAD、PUT  

Accept-Range————告诉客户端服务器上该资源是否允许range请求  
通用：  
Accept-Ranges= acceptable-ranges  
如：  
Accept-Ranges:bytes  接收range请求  
Accept-Ranges:none  不接收range请求    

### 内容协商与资源表述  
内容协商：每个URI指向的资源可以是任何事物，它有着多种不同的表述（语言、方式、压缩方式等）  
![](./img/png14.jpeg)  
方式:   
Proactive————主动式内容协商: 客户端表明想要的服务方式，服务端根据这些头部提供特定的representation表述。客户端————>服务器。    
Reactive————响应式内容协商: 服务器将多种表述方式都返回给客户端让客户端自己选择，服务器———>客户端。       

**要素**：
质量因子q————内容的质量、可接受类型的优先级  
媒体类型、编码————content-type:text/html;charset=utf-8  
内容编码————压缩算法——content-encoding:gzip    
表达语言————accept-language:zh-CN,zh;q=0.9,en;q=0.8  

### HTTP包体  
请求或响应都可以携带包体：  
HTTP-message = start-line*(header-field CRLF) CRLF [message-body]  
message-body = *OCTET(二进制字节流)    

#### 包体长度
发送HTTP消息时已经确认包体的全部长度————直接传输    
使用Content-Length 头部（HEAD，可见性高）明确指明了包体长度  
Content-Length = 1* DIGIT  （十进制）  
如果定义的Content-Length 少于实际的字节数量的话，那么浏览器中不能查看到全部内容，WireShark 可以看到请求是都有但是浏览器自动丢掉了。如果定义的是比实际多的话，会返回failed，因为不符合HTTP规范了  

不确定包体的全部长度——分块传输 
前提：
客户端支持Trailer头部————TE:trailers  
服务器告知chunk包体后会传输那些Trailer尾部中，比如包体的大小    


使用Transfer-Encoding头部指明使用Chunk传输方式————会自动忽略Content-Length头部    
通用：transfer-coding=**"chunked"**/"compress"/"deflate"/"gzip"/transfer-extension  
chunked-body=*chunk last-chunk trailer-part CRLF   

chunk = chunk-size[chunk-ext] CRLF chunk-data CRLF  
chunk-size(表示当前块的长度)=1*HEXDIG(16进制而不是10进制)————需要包括换行符号。  
chunk-data(实际的数据内容)=1*OCTET  

last-chunk（最后一块）=1*("0")[chunk-ext]CRLF  
trailer-part=*(header-field CRLF)————尾部部分紧跟在最后一个数据块之后，

它由零个或多个头部字段  
浏览器接收的是只有一个请求  
优点：基于长连接持续推送动态内容  
可是实现变压缩变发送————因为长度不知道  
先发送完包体再发送头部  



以下消息不含包体：  
HEAD方法请求对应相对的响应  
1XX、204、304  对应的响应  
CONNECT 方法对应的2XX的响应  


#### 包体内容  
MIME——媒体类型  
content:="Content-Type":"type"/"subtype*(";"parameter)"    
"Content-Type": 这是HTTP头部字段的名称  
type: 表示媒体类型的类别，如text, image, application等。  
subtype: 表示具体子类型的标识符，例如对于text/html中的html部分。  
举例：Content-Type: application/json; charset=utf-8  

Content-Disposition 是一个用于指示响应内容如何被处理的HTTP头部字段  
disposition-type = "inline" |"attachment"|disp-ext-type   
inline: 这是最常见的处置类型之一，意味着返回的内容应该在浏览器中直接显示（如果浏览器能够渲染的话）。例如，HTML页面或内联图像通常会使用这种类型。  
attachment: 表示返回的内容应该被视为下载附件，并提示用户保存到本地文件系统，而不是直接在浏览器中打开。
disp-ext-type: 这指的是扩展处置类型，允许定义非标准的处置类型。然而，在实际应用中很少使用，因为这可能不会被所有用户代理正确识别或处理    
`Content-Disposition: inline; filename="example.jpg"`   

#### 表单提交时的包体格式  

HTML：是一种结构化的标记语言，没有交互功能  
里面的FORM作用
FORM表单：提供了交互控制元件用来向服务器通过HTTP协议提交信息    
核心属性：  
action————提交时发起HTTP请求的URI    

method————提交时发起HTTP请求的http方法  

enctype————在POST方法下，对表单内容在请求包体中的编码方式   
application/x-www-form-urlencoded：这是默认类型，它会将表单数据转换成URL编码格式，键值对之间用&连接，空格变成+号，特殊字符转换为ASCII十六进制字节值  
multipart/form-data：这种编码方式不进行URL编码，而是使用多部分边界分隔每个字段的数据，特别适用于需要上传文件的情况。  
使用上面这种的时候就需要添加Content-type 头部指明这是一个多表述包体    
```

Content-type:multipart/form-data;   //认为每个输入框都是一个独立资源描述    
        boundary=---Webxxxx            //分隔  
    //boundary:=0*69<bchars>bcharsnospace  

multipart-body = 1*encapsulation close-delimiter //一个多部分消息体由一个或多个封装（encapsulation）组成，并以关闭定界符（close-delimiter）结束


//全部内容 
POST /upload HTTP/1.1
Host: example.com
Content-Type: multipart/form-data; boundary=boundary-string

--boundary-string
Content-Disposition（内容形式）: form-data; name="field1"

value   //表示field1的值是value
--boundary-string
Content-Disposition: form-data; name="field2"; filename="example.txt"
Content-Type: text/plain

(data)   //example.txt 里面的内容   
--boundary-string-- //最后加两个`--`表示结束了
```

## 多线程、断电续传、随机点播等场景的步骤  
### 步骤 
1.客户端明确任务：从哪里开始下载  
本地是否已有部分文件、文件已下载部分在服务器端发送改变？  
使用几个线程并发下载  
2.下载文件的指定部分内容  
3.下载完毕后拼装成统一的文件  

### HTTP Range规范  
允许服务器基于客户端的请求只发送响应包体的一部分给客户端，而客户端自动将多个片段的包体组合成完整的体积更大的包体————许客户端请求资源的特定部分（字节范围），而不是整个资源  

前提：Accept-Range头部表示支持Range请求  
#### Range用法
Accept-Ranges:byte 表示支持并且按照字节的形式传输类似于封装成贞的思想  
总长度为500  
直接表示：bytes=0-499    
分段表示：bytes=0-99,100-499 /(使用重合)0-200,100-499 客户端用来组装  
从后开始:bytes=-500,0-  
仅要第1个和最后一个字节：bytes=0-0,-1  
实际传输的时候需要在头部中Range: bytes=0-5  

#### Range条件请求  
如果客户端已经获得了Range响应的一部分——就是之前已经发送过了Range请求（服务器端会为这个响应生成ETag），并想在这部分响应未过期的情况下获取其他部分的响应。
`IF-Match:"ETag"`    
`Content-Range: bytes（单位） 0-999/1458934`表示显示传输资源的大小/表示总体大小如果不知道可以用*代替    
如果返回不符合就会返回416报错   
服务器不支持Range返回200    

### Cookie  
Cook是HTTP的状态管理机制  
保存在客户端的、由浏览器维护、表示应用状态的HTTP头部    
存放在内存或者磁盘中，服务器端生成Cookie在响应中通过Set-Cookie头部告知客户端，客户端得到Cookie后，后续请求都会自动将Cookie头部携带至请求中  

Set——Cookie头部一次只能传递1个name/value对，响应中可以含有多个头部  
set-cookie-header="Set-Cookie:" SP set-cookie-string  
"Set-Cookie:"：这是头部的名字，表明服务器希望客户端设置一个或多个cookie。
set-cookie-string：这是实际的cookie数据字符串，包含了cookie的各种属性。最基本的set-cookie-string至少包含一个键值对（即名称和值），但也可以包括额外的属性。   
可选属性有：  
Expires 指定了cookie的过期时间
Max-Age=3600 表示该cookie将在3600秒（即1小时）后过期，优先级高于expires    
Domain=example.com 定义了cookie有效的域名，默认是当前域名    
Path   表明指定路径下使用cookie  
Secure 标志表示此cookie只能通过HTTPS传输
HttpOnly 标志意味着这个cookie不能通过JavaScript的Document.cookie API访问，有助于减轻XSS攻击的风险。


Cookie用于客户端（浏览器）向服务器发送之前通过Set-Cookie头部存储的cookie信息。它仅包含键值对，没有额外的属性。 
Cookie-header="Cookie:"OWS cookie-string OWS   
"Cookie"这是Cookie头部的字面名称。在HTTP请求中，客户端通过包含这个头部来向服务器发送存储在客户端上的cookie信息  
cookie-string：这是实际携带的cookie数据，由一个或多个键值对组成，这些键值对通常是由服务器通过Set-Cookie响应头设置，并存储在客户端上。多个cookie用分号加空格（; ）分隔。

#### Cookie于Session的常见用法  
![](./img/png15.jpeg)  
步骤：  
1.首先客户端请求发送POST请求给服务器，服务器读取数据库用户名和密码对比。  
2.通过后就会把已经登陆的状态生成一个session(有实效)，把session的id和有效时间放入内存数据库中  
3.服务器再返回这个时间给客户端，按照Set-Cookie返回  
4.浏览器保存session之后就会保存在cookie  
5.浏览器之后发送请求就会携带session
6.服务器接收到了session之后就会和内存数据库去比对  
7.成功后执行业务  

第三方Cookie  
浏览器允许对于不安全域下的资源响应中的Set-Cookie保存，并在后续访问该域时自动使用Cookie，就类似在第一个浏览器引用第三方的资源的时候，可能会把第三方的Cookie给携带走   
默认发送机制：当浏览器向网站发起请求时，会自动附带该网站域下的所有未过期的Cookie（根据路径匹配原则）。这意味着，如果一个网页加载了来自第三方的内容（如图片、脚本或iframe），而这些内容触发了对第三方服务器的请求，那么该第三方域下对应的Cookie也会被自动发送  
所以引用第三方Cookie时就会携带走一些隐私  

#### 同源策略  
原因：  
同一个浏览器发出的请求，未必都是用户自愿发出的请求如加载的时候就会发送GET请求：网页可以悄悄地向受害者的银行或电子邮件服务等敏感站点发送请求。如果用户的浏览器已经登录到这些服务并且相应的会话Cookie没有过期，这些请求将会自动携带有效的会话信息，从而可能导致未经授权的操作被执行。   
广告网络和其他第三方服务提供商经常使用Cookie来跟踪用户行为。例如，当一个用户访问不同的网站时，如果这些网站都加载了同一个第三方广告服务，那么该服务可以通过在用户浏览器中设置的Cookie来识别用户，并收集关于其浏览习惯的信息，用于个性化广告展示等目的。   

![](./img/png16.jpeg)  

定义：限制了从同一个源加载的文档或脚本与另一个源的资源进行交互    
同源的定义：协议，主机，端口号都完全相同   
跨域问题：当一个网页尝试请求另一个不同源的资源时就会遇到跨域问题  

可用性：  
允许跨域写操作、允许带有SRC属性的可以跨域访问  

跨站请求伪造攻击（CSRF）  
![](./img/png17.jpeg)    
当用户访问银行的服务器登陆之后，又访问一个木马网站。  
网站利用同源策略中的允许跨域写操作，对银行页面进行操作就可以使用Cookie了  

解决办法：1.禁止来源其他网页的发起请求————Reference: 如果带的有值就不允许操作，但这个一些版本过低的页面无法使用  
2.服务端返回一个唯一且有实效性的token 这样就当木马网站发送请求的时候这个实效就失效了，且token无法伪造  

CORS————跨域的资源共享  
这是正规、合理的跨域访问  
站点A允许站点B的脚本访问其资源，必须HTTP响应中显式的告知浏览器：站点B是被允许的  
所以要做到浏览器要告知站点A这个请求来自站点B并且站点A的响应中，应明确哪些跨域请求是被允许的  
策略1:使用简单请求  
GET/POST/HEAD方法之一  
仅能使用CORS安全的头部：Accept、Accept-Language、Content-Language、Content-type————协商类的  
Content-Type的值:text/plain、multipart/form-data、applicatio/x-www-form-urlencoded  
![](./img/png18.jpeg)  
如同上图：客户端先给服务器发送一个请求携带了Origin告知服务器来自那个域，服务器返回一个允许通过的域给浏览器，从而浏览器放行。  
注意：服务端并不会阻拦相关的跨域，只有浏览器才能处理想对应的跨域请求  


策略2:使用复杂请求    
![](./img/png19.jpeg)  
首先发送一个预检请求头部有：  
Access-Control-Request-Method: POST ———— 表示实际请求将使用 POST 方法。   
Access-Control-Request-Headers ————表示接下来发送的实际请求中需要包含自定义的头部  

服务器响应预请求  
Access-Control-Allow-Method————允许跨域请求的方法  
Access-Control-Allow-Origin————服务器允许来自XXX的请求访问资源。   
Access-Control-Allow-Headers————服务器允许请求中包含XXX头。
Access-Control-Max-Age: 86400: 服务器设置预检请求的结果可以被缓存 86400 秒（一天）。  

客户端发送实际跨域请求   

服务端返回一个响应跨域请求，并且告知浏览器哪里的请求可以跨域访问  
  
### 条件请求  
定义：由客户端携带条件判断信息，而服务器预执行条件验证过程成功，再返回资源表述  

####验证器
验证器：根据客户端请求中携带的相关头部，以及服务器资源的信息，执行两端的资源验证  
强验证器：服务器上的资源表只要有变动，那么旧的验证头部访问一定会验证不通过  
弱验证器：服务器资源变动时，允许一定程度上任然可以验证通过   
验证器响应头部有：  

1.Etag 响应头部  
定义:entity-tag=[weak]opaque-tag   
weak:弱验证器以W/前缀来标记  
opaque-tag : 这是指实体标签的具体值，它对客户端来说应该是不透明的（即客户端不需要理解其意义）。服务器端决定如何生成这个值，通常是为了唯一地标识一个特定版本的资源。    
弱验证:ETag: W/"67ab43"  
强验证:ETag: "67ab43"   

2.Last-Modified 响应头部  
表示上次修改的时间   
Last-Modified=HTTP-date  
Date头部表示的是响应包体生成的时间  

#### 基于上述的判断的条件请求  
if-Match="*"/1#entity-tag  
这个请求头用于确保只有当目标资源的当前实体标签（ETag）匹配指定的值时，才执行请求。如果使用了通配符\*，则表示无论当前的ETag是什么，都应执行请求。通常用于PUT或其他更新操作，以防
止覆盖他人的更改。  

if-None-Match="*"/1#entity-tag  
该请求头用于检查服务器上的资源是否不同于客户端已有的副本。若指定的ETag与服务器
端资源的ETag不匹配（即资源已更新），则执行请求（如GET）。如
果匹配，则返回304未修改状态码。  

if-Modified-Since=HTTP-date  
此请求头告知服务器仅当所请求的资源自指定日期以来已被修改时，才返回资源的内容；否则返回304未修改状态码。  

if-Unmodified-Since=HTTP-date  
如果资源自指定日期后没有被修改过，则允许执行请求。   

if-Range=entity-tag/HTTP-date  
这个请求头结合了范围请求（Range）一起使用。它指示服务器仅当资源的ETag或
最后修改时间与请求中指定的一致时，才处理范围请求。如果不匹配，服务器应当返回整个资源而不是部分内容。  

#### 应用场景
1.缓存更新   
a.首次缓存  
![](./img/png20.jpeg)  
当客户端一开始没有任何缓存的时候，客户端会发送一个GET请求给服务器，服务器会返回给浏览器一个最后修改时间和Etag标记，和相应的数据    

b.基于过期缓存发送条件请求  
![](./img/png21.jpeg)  
出现两种情况  
情况一：服务器并没有对数据进行修改。客户端向服务器发
送请求（包含了条件请求）检验是否修改如果回复304表示没有修改那么浏览器可以继续使用缓存大大减少了流量的浪费  
情况二：服务器对数据进行修改。那么客户端会回复200并且给他发送新的数据给缓存和Etag和Last-Modified  

2.增量更新  
当服务器支持Range服务时，连接意外中断时以及接收到部分数据  
![](./img/png22.jpeg)  
为了防止两次下载间资源已发生了变更  
如果没有发送变化：  
![](./img/png23.jpeg)  

如果发送了变化：  
![](./img/png24.jpeg) 
服务器返回412然后浏览器重新发送GET请求，他就会使用新的Etag    
若要使If-Range工作，需要确保所使用的条件（无论是ETag还是HTTP-date）与服务器端的状态相匹配。如果条件不符合，服务器将默认返回完整的资源而不是部分内容——就不需要再发送Get请求了而是合并了。  

3.更新丢失问题  
获取资源，再把本地修改后的资源提交  
![](./img/png25.jpeg)  
相当于两个客户端同时访问同一片共享资源的时候可能会出现一个客户端把另一个客户端的内容给覆盖了。  
解决办法：乐观锁————基于条件请求判断是否基于最新的标签  
乐观锁还能解决首次上传的问题：  
![](./img/png26.jpeg)  
首次上传的时候默认为* 修改之后标签改变不为*就会返回412  


#### Nginx的处理图  
![](./img/png27.jpeg)  


## 缓存
HTTP缓存：为了当前请求复用前请求的响应  
目的：减少时延；降低带宽消耗  
![](./img/png28.jpeg)  
表示先从浏览器发送请求给服务器，服务器响应请求并且返回客户端。浏览器Cache储存响应中的数据  
1.如果缓存没有过期的话，下次的请求不会直接发送给服务器而显查找本地的缓存————就不会发送请求给Web服务器。    
查看是否使用缓存可以查看浏览器中的Network中的size有From cache 
2.如果缓存过期了，客户端先从浏览器Cache中检查缓存过期，Local Browser Cache会先返回一个标签，然后给服务器返回标签。服务器对比标签发现里面的内容没有改变依旧可以使用就会返回一个304 然后重新发一个这个缓存使用时间的标签，并不会传递全部的缓存。  

### 私有缓存与共享缓存  
私有缓存：仅提供一个用户使用的缓存，通常只存在浏览器这样的客户端上  
共享缓存：可以提供多个用户的缓存，存在于网络中负责转发消息的代理服务器————相当于发出请求的时候访问的不是服务器，而是访问的是代理服务器的缓存。  

![](./img/png29.jpeg)  
首先客户端给代理服务器发送Get请求，代理服务器查找自己Cache是否有对应的数据（并且检查是否过期），如果没有就给服务器发送Get请求要求给代理服务器要缓存。然后服务端响应之后会有max-age表示在代理服务器存放的时间，并且给分支给客户端。  
如果代理服务器有了缓存之后客户端下次访问的代理服务器缓存就是直接返回给客户端。  
如果代理服务器缓存过期了，那么下次客户端发送请求给代理服务器，代理服务器会发现自己缓存过期再次向服务器发送请求更新自己的数据。  

缓存实现的机制  
![](./img/png30.jpeg)  
相当于缓存按照关键字能够快速查找：使用字典（哈希表）作为主索引可以提供对存储在双向链表中的数据项的快速访问能力。HTTP请求信息（如schema、path、host）作为键，允许系统迅速定位到相应的缓存项。  
顺序管理：双向链表的存在使得数据项之间能够保持一定的顺序，这通常用于实现LRU（Least Recently Used，最近最少使用）等缓存淘汰策略。新访问的数据项可以通过移动到链表头部来标记为最近使用过，而长时间未被访问的数据项则会逐渐移向链表尾部，便于淘汰。  

判断缓存是否过期：  
response_is_fresh=(freshness_lifetime>current_age)  
freshness_lifetime:按优先级，取以下响应头部的值  
s-maxage>max-age>Expires>预估过期时间（默认的缓存）————(DownloadTime-LastModified)*10%  
current_age定义：在HTTP缓存机制中表示一个缓存响应的当前年龄，即自该响应资源被最初生成或最后验证以来经过的时间长度。它是决定缓存内容是否仍然新鲜（即是否可以继续使用而不需重新验证或获取）的一个关键因素————是一个概念并不会在http头部出现    
current_age=corrected_initial_age+resident_time
corrected_initial_age:最初始的发出可以使用的时间  
resident_time:是指从缓存接收到响应开始直到当前时刻为止的时间长度（时延）。    
![](./img/png31.jpeg)  
先有源服务器发送为date1表示该响应是在 date1 时刻生成的。，然后代理服务器1接收到这个响应的时间此时的Age表示传输花费的时间。发给代理服务器2now-date1 的表示在代理服务器停留的时间  所以综上Currnet_age表示的是传输途中消耗的时间  
因为需要缓存会被不停的调用所以Currnet_age会不断被更新从而就出现currnet_age更新，它是接收了一次请求就会更新  
在network中的age表示从产生到这个浏览器的时间  

### Cache-Control 头部
Cache-Control=1#chache-directive  
cache-directive = token [ "=" ( token quoted-string ) ]  
token=1*DIGIT    

#### 请求中的值
max-age ————客户端希望接收缓存时间不超过指定秒数的响应（即使缓存中存在更旧的资源）。  
max-stale————允许接收已过期的缓存响应（默认允许任何过期时间，或指定最大可接受的过期时间）。  
min-fresh————要求响应至少在接下来的指定秒数内保持新鲜（即未过期）。  
no-cache————强制缓存服务器在使用缓存前必须向原始服务器验证资源的有效性。  
no-store————禁止任何缓存存储请求或响应的内容（用于敏感数据）。  
no-transform————禁止代理缓存对资源进行转换（如压缩、图片格式转换等）。  
only-if-cached————客户端仅接受缓存的响应，若缓存中无有效响应，则返回 504 Gateway Timeout。    

#### 响应中的值  
no-cache: 要求缓存在提供存储的副本之前必须先提交给源服务器进行验证（通过发送条件请求），以确保副本仍然有效。这并不意味着不允许缓存，而是强调了重新验证的重要性。   
no-store————禁止任何缓存存储请求或响应的内容（用于敏感数据）。  
max-age: 指定资源被视为新鲜的最大时间长度。在这个时间段内，缓存可以直接使用存储的响应而无需重新验证。  
must-revalidate: 一旦资源过期，在再次使用之前必须向源服务器验证其有效性。它强制缓存遵循新鲜度检查规则，即使这样做可能会导致网络访问失败。  
proxy-revalidate: 与 must-revalidate 类似，但是仅适用于共享缓存（如代理服务器）。它不影响私有缓存的行为。    
   
### 什么样的HTTP响应会被缓存  
请求方法可以被缓存理解（不只GET方法）  
响应码可以被缓存理解（404、205等）     
响应于请求中没有指明no-store  
响应中至少应含有以下1个或者多个头部：Expires、max-age、s-maxage、public当响应中没有明确指示过期间的头部，如果有响应码也可以被缓存。  
如果缓存存在代理服务器上  
不含有private、Authorization  

使用缓存的条件：  
URI是匹配的：URI作为主要的缓存关键字，当一个URI同时对应多分缓存时，选择日期最近的缓存  
缓存中的响应允许当前请求的方法使用缓存  

