### shell脚本
概念：shell 是命令解释器，用于解释用户对系统的操作，把用户所执行的命令翻译给内核，再根据结果反馈给内核。  
shell的类型分配在/etc/shells文件中，CentOS7默认使用bash  

Linux 的启动过程：  
BIOS 基本的输入输出功能在主板半，选择引导的介质  
MBR 硬盘的主引导记录，是否可引导，包含了磁盘分区表等内容    
BootLoader(grub) 启动和引导内核的主要工具(选择内核以及版本号)    
kernel 内核
systemd/init 系统初始化进程，管理硬件，启动服务，管理服务————启动一号进程  
shell 用户与内核之间的桥梁，用户通过shell与内核进行交互，shell脚本就是用户与内核之间的桥梁  

```
dd if=/dev/sda of=mbr.bin bs=446 count=1 //将硬盘的MBR备份到mbr.bin文件中
hexdump -C mbr.bin //查看mbr.bin文件的十六进制表示-c表示显示为字符，因为没有安装文件管理系统。在512字节的最后内容为55aa，表示可引导。  

boot/grub2/grub.cfg  //grub的配置文件，用于配置内核的启动参数  
grub2-editenv list  //表示查看引导的内核的版本号    

/etc/re.d //CentOS6启动的是/usr/sbin/init引导程序,init会在/etc/rc.d下面发现大量的脚本。
/etc/systemd/system //CentOS7启动的是/etc/etc/systemd/system就会有很多启动级别，就是1号进程，然后根据启动级别/usr/lib/systemd/syetem目录下读取各种service服务。  
```


### shell格式  
为了组合命令和多次执行，使用脚本文件保存需要执行的命令。  

```
cd /var/ ;ls   //这就是执行两条命令，进入var目录并且查看。  
```



