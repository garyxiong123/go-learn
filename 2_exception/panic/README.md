reference the address

https://www.jianshu.com/p/f30da01eea97/

我们在调用recover的延迟函数中以最合理的方式响应该异常：

打印堆栈的异常调用信息和关键的业务信息，以便这些问题保留可见； 
将异常转换为错误，以便调用者让程序恢复到健康状态并继续安全运行。 

我们看一个简单的例子：

原理

如果遇到引用空指针、下标越界或显式调用panic函数等情况，则先触发panic函数的执行， 然后调用延迟函数。调用者继续传递panic，因此该过程一直在调用栈中重复发生：函数停止执行，
调用延迟执行函数等。如果一路在延迟函数中没有recover函数的调用，则会到达该携程的起点，该携程结束， 然后终止其他所有携程，包括主携程（类似于C语言中的主线程，该携程ID为1）