# NIO学习笔记

### Java NIO简介

##### 基本介绍

> Java NIO ( New IO ) 是从Java1.4版本开始引入的一个新的IO API，可以替代标准的Java IO API。NIO与原来的IO有同样的作用和目的，但是使用的方式完全不同，NIO支持面向缓冲区的、基于通道的IO操作。NIO将以更加高效的方式进行文件的读写操作。

##### Java NIO与IO的主要区别

| IO                         | NIO                            |
| -------------------------- | ------------------------------ |
| 面向流 ( Stream Oriented ) | 面向缓冲区 ( Buffer Oriented ) |
| 阻塞IO ( Blocking IO )     | 非阻塞IO ( Non Blocking IO )   |
| （无）                     | 选择器 ( Selectors )           |

##### 通道和缓冲区

> Java NIO 系统的核心在于：通道(Channel)和缓冲区(Buffer)。通道表示打开到IO设备（例如：文件、套接字）的连接。若需要使用NIO系统，需要获取用于连接IO设备的通道以及用于容纳数据的缓冲区。然后操作缓冲区，对数据进行处理。
>
> ==简而言之，Channel负责传输，Buffer负责存储==
>
> **缓冲区**
>
> * 缓冲区（Buffer）：一个用于特定基本数据类型的容器。由java.nio包定义的，所有缓冲区都是Buffer抽象类的子类。
> * Java NIO 中的Buffer主要用于与NIO通道进行交互，数据是从通道读入缓冲区，从缓冲区写入通道中的。
>
> **缓冲区常用方法**
>
> ```java
> package com.nio;
> 
> import org.junit.Test;
> 
> import java.nio.ByteBuffer;
> 
> /**
>  *  一、缓冲区（Buffer）：在Java NIO中负责数据的存取。缓冲区就是数组，用于存储不同数据类型的数据
>  *  根据数据类型不同(boolean除外)，提供了相应类型的缓冲区
>  *  ByteBuffer,CharBuffer,ShortBuffer,IntBuffer,LongBuffer,FloatBuffer,DoubleBuffer
>  *  上述缓冲区的管理方式几乎一致，通过allocate获取缓冲区
>  *
>  *  二、缓冲区存取数据的两个核心方法
>  *  put() : 存入数据到缓冲区中
>  *  get() : 获取缓冲区中的数据
>  *
>  *  三：缓冲区中的四个核心属性
>  *  capacity : 容量，表示缓冲区中最大存储数据的容量，一旦声明不能改变
>  *  limit : 界限，表示缓冲区中可以操作数据的大小。（limit后数据不能进行读写）
>  *  position : 位置，表示缓冲区中正在操作数据的位置
>  *  position <= limit <= capacity
>  *  mark : 标记，表示记录当前position的位置，可以通过reset() 恢复到mark的位置
>  *  0 <= mark <= position <= limit <= capacity
>  */
> public class TestBuffer {
> 
>     @Test
>     public void test1(){
>         String str = "abcde";
>         // 1.分配一个指定大小的缓冲区
>         ByteBuffer buf = ByteBuffer.allocate(1024);
>         System.out.println("-------------------");
>         System.out.println(buf.position());
>         System.out.println(buf.limit());
>         System.out.println(buf.capacity());
>         // 2.利用put存入数据到缓冲区中
>         buf.put(str.getBytes());
>         System.out.println("--------put()-----------");
>         System.out.println(buf.position());
>         System.out.println(buf.limit());
>         System.out.println(buf.capacity());
>         // 3.flip()切换成读取数据模式
>         buf.flip();
>         System.out.println("--------flip()-----------");
>         System.out.println(buf.position());
>         System.out.println(buf.limit());
>         System.out.println(buf.capacity());
>         // 4.利用get() 读取缓冲区中的数据
>         byte[] dst = new byte[buf.limit()];
>         buf.get(dst);
>         System.out.println(new String(dst,0,dst.length));
>         System.out.println("--------get()-----------");
>         System.out.println(buf.position());
>         System.out.println(buf.limit());
>         System.out.println(buf.capacity());
>         // 5.可重复读，即把指针重新指向数据开头
>         buf.rewind();
>         System.out.println("--------rewind()-----------");
>         System.out.println(buf.position());
>         System.out.println(buf.limit());
>         System.out.println(buf.capacity());
>         // 6.clear() : 清空缓冲区，但是缓冲区中的数据依然存在，但是处于“被遗忘”状态，指针指向了缓冲区开头
>         buf.clear();
>         System.out.println("--------clear()-----------");
>         System.out.println(buf.position());
>         System.out.println(buf.limit());
>         System.out.println(buf.capacity());
>         System.out.println((char)buf.get());
>     }
> 
>     @Test
>     public void test2(){
>         String str = "abcde";
>         ByteBuffer buf = ByteBuffer.allocate(1024);
>         buf.put(str.getBytes());
>         buf.flip();
>         byte[] dst = new byte[buf.limit()];
>         buf.get(dst,0,2);
>         System.out.println(new String(dst,0,2));
>         System.out.println(buf.position());
>         // mark() : 标记
>         buf.mark();
>         buf.get(dst,2,2);
>         System.out.println(new String(dst,2,2));
>         System.out.println(buf.position());
>         // reset() : 恢复到mark的位置
>         buf.reset();
>         System.out.println(buf.position());
>         // 判断缓冲区中是否还有剩余数据
>         if(buf.hasRemaining()){
>             // 获取缓冲区中可以操作的数量
>             System.out.println(buf.remaining());
>         }
>     }
> }
> ```
>
> 

