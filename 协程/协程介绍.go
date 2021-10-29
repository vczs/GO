package main
/**
进程和线程：
进程就是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位。
线程是进程的一个执行实例，是程序执行的最小单元。
一个进程可以创建和销毁多个线程，同一个进程中的多个线程可以并发执行。
一个程序至少有一个进程，一个进程至少有一个线程。

并发和并行：
并发：多个线程在同一个CPU上运行，在某个时间点上只有一个线程在运行，这就是并发。
并行：多个线程在多个CPU上运行，在某个时间点上有多个（CPU个数）线程在运行，这就是并行。
理解：并发是一个人同时吃三个馒头，而并行是三个人同时吃三个馒头。

GO主线程（可以称为线程、也可以理解为进程）和GO协程：
GO主线程：一个GO主线程上，可以起多个协程(轻轻松松起上万个协程)，协程是轻量级的线程。
GO协程特点：1.有独立的栈空间  2.共享程序堆空间   3.调度由用户控制  4.协程是轻量级的线程
区别：
1.如果主线程退出了，主线程起的协程即使没有执行完毕，也会退出；当协程执行完任务后就会结束，即使主线程还在运行。
2.主线程是一个物理线程，直接作用在CPU上。是重量级的，非常消耗资源。
3.协程是从线程开启的，是轻量级的线程，是逻辑态，对资源消耗很少。
4.GO的协程的机制，使GO主线程可以轻松开启上万个协程，运行非常稳定，这就是突显GO在并发上的优势。
 */
