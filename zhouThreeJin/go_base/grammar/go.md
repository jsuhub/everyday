##执行流程  
1.先编译成为可执行文件再运行（.go文件——go bulid/编译———>可执行文件——运行——>结果）  
2.直接对文件生成结果（.go文件——go run/编译运行———>结果）  

这两者的区别：  
1.如果先编译生成可执行文件，那么我们将该可执行文件拷贝到没有go语言开发环境中仍然可以运行。（编译之后会将它所需要的库倒入这个文件中）  
2.如果是直接利用go run生成的文件在另一台机器运行任然需要go的环境（因为直接使用go run进行的，所有并没有go的编译文件没有导入相应的库）    
3.在编译过后的可执行文件将包含程序执行的库文件，所以会变大很多。  
 
