## 数据类型命令以及落地运用
命令不区分大小写  
key值区分大小写  

help @类型查看帮助文档  

## STRING 类型 
单值——单value   
set 指令   
```
set key value [NX | XX ][GET][EX seconds|PX milliseconds|EXAT unix-time-seconds|PXAT unix-time-milliseconds|KEEPTTL]    
EX seconds:以秒为单位设置过期时间 
set k1 v1 ex 10 //10秒后过期 get k1 过了10秒就返回-2   
PX milliseconds：以毫秒为单位设置过期时间  
set k1 v1 px 1000 //1000毫秒后过期 get k1 过了1000毫秒就返回-2   
EXAT unix-time-seconds：以秒为单位UNIX时间戳设置过期时间  
EXAT k1 v1 1717334400 // 表示设置k1、v1在2024-06-01 00:00:00过期 
PXAT unix-time-milliseconds：以毫秒为单位UNIX时间戳设置过期时间  
PXAT k1 v1 1717334400000 // 表示设置k1、v1在2024-06-01 00:00:00过期 
KEEPTTL：保留设置前过期时间，因为使用set k1 v1xx 会覆盖之前设置的时间变成-1 即永久，所以想要保存这个时间就要使用KEEPTTL    set k1 v1xx keepttl   
NX：key不存在时设置值  
set k1 v1 nx //存在返回nil不存在返回v1 
XX：key存在时设置值  
set k1 v1xx xx //存在返回v1xx覆盖不存在返回nil
GET：获取值 若不存在返回nil 
set k1 v1 get  //返回k1的值v1xx 并且把v1付给k1
```

`mset/mget` 同时设置/获取多个键值  
```
mset k1 v1 k2 v2 k3 v3  //这是一个完整的事务 如果一个失败就都失败  
mget k1 k2 k3  
```

`getrange/setrange` 获取指定区间范围内的值  
```
getrange k1 0 2 //获取k1的值v1xx的0-2位 返回v1x 如果使用-1 就是截取全部    
setrange k1 0 v2 //将k1的值v1xx的0位开始替换为v2 返回v2xx   
```

`incr/decr` 自增/自减一定要是数字才能进行  
```
incr k1 4 //将k1的值自增4 返回增加后的数量
decr k1 //将k1的值自减1 返回减少后的值   
```

`strlen` 获取字符串长度  
```
strlen k1 //获取k1的值v1xx的长度 返回4  
```

`append` 追加字符串  
```
append k1 v2 //将v2追加到k1的值v1xx后面 返回v1xxv2  
```

分布式锁  
```
setnx lock uuid //设置lock的值uuid 如果lock不存在则设置成功 返回1 如果lock存在则设置失败 返回0这样就成功建立了一个锁 
del lock //删除锁  

setex loock 10 uuid //类似set loock uuid ex 10 设置10秒后过期   

```

`getset` 获取并设置值  
```
getset k1 v2 //获取k1的值v1xx 并设置k1的值为v2 类似 set k1 v2 get       
```






## LIST 类型


