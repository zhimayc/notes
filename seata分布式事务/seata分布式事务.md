#### 分布式架构中的CAP定理

![image-20210602141345605](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20210602141345605.png)



#### 一致性

一致性分为强一致性和弱一致性，是根据节点间数据同步时间决定的，看是同步复制还是异步复制

强一致性：数据复制是同步的，保持节点间数据实时同步

弱一致性：数据复制是异步的，保持节点间数据最终同步

一般业务中都是选择AP，牺牲强一致性但支持最终一致性，保留可用性和分区容错性



#### 分布式事务协议

###### 1、二阶段协议 2PC

将事务的提交过程分为准备阶段和提交阶段两个阶段来进行处理

==准备阶段==：协调者向所有参与者发送事务内容，询问是否可以提交事务，各参与者执行操作，若执行成功则向协调者反馈同意，否则反馈中止

==提交阶段==：当协调者收到的反馈全部都为同意时，协调者向所有参与者发起正式提交请求，所有参与者完成事务提交并向协调者发送ack响应，协调者收到所有ack响应之后即代表事务完成



![image-20210602143821123](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20210602143821123.png)



![image-20210602144305918](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20210602144305918.png)



![image-20210602144836507](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20210602144836507.png)



二阶段协议的缺点：

性能问题：执行过程中所有参与节点都是事务阻塞型的，占有公共资源，因为

可靠性问题：参与者发生故障，协调者需要额外提供超时机制

一致性问题：协调者在发出commit消息之后宕机，而唯一接收到这条消息的参与者也宕机了，那么即使通过选举选出了新的协调者，这条事务的状态也是不确定的



2、三阶段提交协议 3PC

他是二阶段协议的升级版，在协调者和参与者中都引入了超时机制，在第一阶段和第二阶段之间插入一个准备阶段，保证了最后提交阶段之前各个参与节点状态的一致





#### 分布式事务解决方案

1、事务补偿 TCC

TCC方案是一种应用层面侵入业务的两阶段提交，可保持强一致性，但对资源侵占比较强

具体产品：==ByteTCC==

2、全局消息

3、基于可靠消息服务的分布式事务

4、最大努力通知



#### Seata

![image-20210602153420448](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20210602153420448.png)



什么是全局事务服务GTS



###### seata术语

![image-20210602153800868](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20210602153800868.png)



![image-20210602153929125](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20210602153929125.png)



如上图：RM相当于参与者，TM相当于需要分布式事务的地方，TC事务协调者



###### UNDO_LOG表

需要在每个业务数据库中创建，用于保存回滚操作数据，当全局提交时，UNDO_LOG记录直接删除，当全局回滚时，将现有数据撤销，还原至操作前状态



###### AT模式

![image-20210602154937513](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20210602154937513.png)