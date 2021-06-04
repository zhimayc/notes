# Hibernate学习笔记

### 基本介绍

##### 什么是Hibernate

> * 一个Java领域的持久化框架
> * 一个ORM框架

##### 对象的持久化

> * 狭义的理解，”持久化“仅仅指把对象永久保存到数据库中
> * 广义的理解，“持久化”包括和数据库相关的各种操作：
>   * 保存：把对象永久保存到数据库中
>   * 更新：更新数据库中对象（记录）的状态
>   * 删除：从数据库中删除一个对象
>   * 查询：根据特定的查询条件，把符合查询条件的一个或多个对象从数据库加载到内存中
>   * 加载：根据特定的OID，把一个对象从数据库加载到内存中。
>
> 为了在系统中能够找到所需对象，需要为每一个对象分配一个唯一的标志号。在关系数据库中称之为主键，而在对象术语中，则叫做对象标志( Object identifier-OID )

##### ORM

> ORM( Object / Relation Mapping ) ：对象/关系映射
>
> * ORM主要解决对象-关系的映射
>
>   | 面向对象概念 | 面向关系概念   |
>   | ------------ | -------------- |
>   | 类           | 表             |
>   | 对象         | 表的行（记录） |
>   | 属性         | 表的列（字段） |
>
> * ORM的思想：将关系数据库中表中的记录映射成为对象，以对象的形式展现，程序员可以把对数据库的操作转化为对对象的操作
>
> * ORM采用元数据来描述对象-关系映射细节，元数据通常采用XML格式，并且存放在专门的对象-关系映射文件中

##### 流行的ORM框架

> * Hibernate：
>   * 非常优秀，成熟的ORM框架
>   * 完成对象的持久化操作
>   * Hibernate允许开发者采用面向对象的方式来操作关系数据库
>   * 消除那些针对特定数据库厂商的SQL代码
> * MyBatis：
>   * 相比Hibernate更灵活，运行速度快
>   * 开发速度慢，不支持纯粹的面向对象操作，需熟悉sql语句，并且熟练使用sql语句优化功能
> * TopLink
> * OJB

##### 入门案例

> ```java
> // News.java
> package com.atguigu.demo;
> import java.sql.Blob;
> import java.util.Date;
> 
> public class News {
> 	
> 	private Integer id;
> 	private String title;
> 	private String author;
> 	
> 	private String desc;
> 	
> 	//使用 title + "," + content 可以来描述当前的 News 记录. 
> 	//即 title + "," + content 可以作为 News 的 desc 属性值
> 	
> 	private String content;
> 	
> 	private Blob picture;
> 	
> 	public Blob getPicture() {
> 		return picture;
> 	}
> 
> 	public void setPicture(Blob picture) {
> 		this.picture = picture;
> 	}
> 
> 	public String getContent() {
> 		return content;
> 	}
> 
> 	public void setContent(String content) {
> 		this.content = content;
> 	}
> 	
> 	public String getDesc() {
> 		return desc;
> 	}
> 
> 	public void setDesc(String desc) {
> 		this.desc = desc;
> 	}
> 
> 
> 
> 	private Date date;
> 
> 	public Integer getId() { //property
> 		return id;
> 	}
> 
> 	public void setId(Integer id) {
> 		this.id = id;
> 	}
> 
> 	public String getTitle() {
> 		return title;
> 	}
> 
> 	public void setTitle(String title) {
> 		this.title = title;
> 	}
> 
> 	public String getAuthor() {
> 		return author;
> 	}
> 
> 	public void setAuthor(String author) {
> 		this.author = author;
> 	}
> 
> 	public Date getDate() {
> 		return date;
> 	}
> 
> 	public void setDate(Date date) {
> 		this.date = date;
> 	}
> 
> 	public News(String title, String author, Date date) {
> 		super();
> 		this.title = title;
> 		this.author = author;
> 		this.date = date;
> 	}
> 	
> 	public News() {
> 		// TODO Auto-generated constructor stub
> 	}
> 
> 	@Override
> 	public String toString() {
> 		return "News [id=" + id + ", title=" + title + ", author=" + author
> 				+ ", date=" + date + "]";
> 	}
> }
> ```
>
> ```java
> // HibernateTest.java
> package com.atguigu.demo;
> import org.hibernate.Session;
> import org.hibernate.SessionFactory;
> import org.hibernate.Transaction;
> import org.hibernate.cfg.Configuration;
> import org.hibernate.service.ServiceRegistry;
> import org.hibernate.service.ServiceRegistryBuilder;
> import org.junit.Test;
> import java.util.Date;
> 
> public class HibernateTest {
> 
>     @Test
>     public void test(){
> 
>         // 1.创建一个SessionFactory对象
>         // 1) 创建Configuration对象：对应hibernate的基本配置信息和对象关系映射
>         SessionFactory sessionFactory = null;
>         Configuration configuration = new Configuration().configure();
> 
>         /// 4.0之前这样创建 ，后面已标记为废弃
>         // sessionFactory = configuration.buildSessionFactory();
>         // 2) 创建一个ServiceRegistry对象 : hibernate4.X 新添加的对象
>         // hibernate的任何配置和服务都需要在该对象中注册后才能有效
>         ServiceRegistry serviceRegistry = new ServiceRegistryBuilder().applySettings(configuration.getProperties()).buildServiceRegistry();
>         // 3)
>         sessionFactory = configuration.buildSessionFactory(serviceRegistry);
> 
>         // 2.创建一个Session对象
>         Session session = sessionFactory.openSession();
>         // 3.开启事务
>         Transaction transaction = session.beginTransaction();
>         // 4.执行保存操作
>         News news = new News("Java","ATGUIGU",new Date(new java.util.Date().getTime()));
>         session.save(news);
>         // 5.提交事务
>         transaction.commit();
>         // 6.关闭session
>         session.close();
>         // 7.关闭SessionFactory
>         sessionFactory.close();
>     }
> }
> ```
>
> ```xml
> hibernate.cfg.xml
> <?xml version="1.0" encoding="UTF-8"?>
> <!DOCTYPE hibernate-configuration PUBLIC
>         "-//Hibernate/Hibernate Configuration DTD 3.0//EN"
>         "http://hibernate.sourceforge.net/hibernate-configuration-3.0.dtd">
> <hibernate-configuration>
>     <session-factory>
> 
>         <!-- 配置连接数据库的基本信息 -->
>         <property name="connection.username">root</property>
>         <property name="connection.password">root</property>
>         <property name="connection.driver_class">com.mysql.jdbc.Driver</property>
>         <property name="connection.url">jdbc:mysql://192.168.0.199:3306/test</property>
> 
>         <!-- 配置 hibernate 的基本信息 -->
>         <!-- hibernate 所使用的数据库方言 -->
>         <property name="dialect">org.hibernate.dialect.MySQL5InnoDBDialect</property>
> 
>         <!-- 执行操作时是否在控制台打印 SQL -->
>         <property name="show_sql">true</property>
> 
>         <!-- 是否对 SQL 进行格式化 -->
>         <property name="format_sql">true</property>
> 
>         <!-- 指定自动生成数据表的策略 -->
>         <property name="hbm2ddl.auto">update</property>
> 
>         <!-- 指定关联的 .hbm.xml 文件 -->
>         <mapping resource="com/atguigu/demo/News.hbm.xml"/>
>     </session-factory>
> </hibernate-configuration>
> ```
>
> ```xml
> News.hbm.xml
> <?xml version="1.0"?>
> <!DOCTYPE hibernate-mapping PUBLIC "-//Hibernate/Hibernate Mapping DTD 3.0//EN"
> "http://hibernate.sourceforge.net/hibernate-mapping-3.0.dtd">
> <hibernate-mapping>
>     <class name="com.atguigu.demo.News" table="NEWS">
>     	
>         <id name="id" type="java.lang.Integer">
>             <column name="ID" />
>             <!-- 指定主键的生成方式, native: 使用数据库本地方式 -->
>             <generator class="native" />
>         </id>
>         <property name="title" type="java.lang.String">
>             <column name="TITLE" />
>         </property>
>         <property name="author" type="java.lang.String">
>             <column name="AUTHOR"/>
>         </property>
>         <property name="date" type="java.sql.Date">
>             <column name="DATE" />
>         </property>
>     </class>
> </hibernate-mapping>
> ```
>
> **创建持久化Java类的条件**
>
> * 提供一个无参的构造器：使Hibernate可以使用Constructor.newInstance() 来实例化持久化类
> * 提供一个标识属性( identifier property )：通常映射为数据库表的主键字段，如果没有该属性，一些功能将不起作用，如：Session.saveOrUpdate() 
> * 为类的持久化类字段声明访问方法( get/set )：Hibernate对JavaBean风格的属性属性持久化。
> * 使用非final类：在运行时生成代理是Hibernate的一个重要的功能。如果持久化类没有实现任何接口，Hibernate使用CGLIB生成代理。如果使用的是final类，则无法生成CGLIB代理
> * 重写equals和hashCode方法：如果需要把持久化类的实例放到Set中（当需要进行关联映射时），则应该重写这两个方法
> * Hibernate不要求持久化类继承任何父类或实现接口，这可以保证代码不被污染，这就是Hibernate被称为低侵入式设计的原因
>
> **Hibernate配置文件的两个配置项**
>
> * hbm2ddl.auto：该属性可帮助程序员实现正向工程，即由java代码生成数据库脚本，进而生成具体的表结构。取值create | update | create_drop | validate
>   * create：会根据.hbm.xml文件来生成数据表，但是每次运行都会删除上一次的表，重新生成表，哪怕二次没有任何改变
>   * create-drop：会根据.hbm.xml文件生成表，但是SessionFactory一关闭，表就自动删除
>   * update：==最常用的属性值==，也会根据.hbm.xml文件生成表，但若.hbm.xml文件和数据库中对应的数据表的表结构不同，Hibernate将更新数据表结构，但不会删除已有的行和列
>   * validate：会和数据库中的表进行比较，若.hbm.xml文件中的列在数据表中不存在，则抛出异常
> * format_sql：是否将SQL转化为格式良好的SQL，取值true | false

### Session接口

##### session概述

> * Session接口是Hibernate向应用程序提供的操纵数据库的最主要的接口，它提供了基本的保存，更新，删除和加载Java对象的方法
> * Session具有一个缓存，位于缓存中的对象称为持久化对象，它和数据库中的相关记录对应。Session能够在某些时间点，按照缓存中对象的变化来执行相关的SQL语句，来同步更新数据库，这一过程被称为刷新缓存( flush )
> * 站在持久化的角度，Hibernate把对象分为4种状态：持久化状态，临时状态，游离状态，删除状态。Session的特定方法能使对象从一个状态转换到另一个状态

##### Session缓存

> * 在Session接口的实现中包含一系列的Java集合，这些Java集合构成了Session缓存，只要Session实例没有结束生命周期，且没有清理缓存，则存放在它缓存中的对象也不会结束生命周期
> * Session缓存可减少Hibernate应用程序访问数据库的频率

##### flush缓存

> * flush：Session按照缓存中对象属性的变化来同步更新数据库
> * 默认情况下Session在以下时间点刷新缓存：
>   * 显示调用Session的flush方法
>   * 当应用程序调用Transaction的commit方法时，该方法先刷新缓存，然后再向数据库提交事务
>   * 当应用程序执行一些查询( HQL，Criteria )操作时，如果缓存中持久化对象的属性已经发生了变化，会先flush缓存，以保证查询结果能够反映持久化对象的最新状态
> * flush缓存的例外情况：如果对象使用native生成器生成OID，那么当调用Session的save() 方法保存对象时，会立即执行向数据库插入该实体的insert语句
> * commit() 和 flush() 方法的区别：flush执行一系列sql语句，但不提交事务；commit 方法先调用flush() 方法，然后提交事务。提交事务意味着对数据库的操作永久保存下来。
>
> **实例代码**
>
> ```java
> package com.atguigu.demo;
> import org.hibernate.Session;
> import org.hibernate.SessionFactory;
> import org.hibernate.Transaction;
> import org.hibernate.cfg.Configuration;
> import org.hibernate.service.ServiceRegistry;
> import org.hibernate.service.ServiceRegistryBuilder;
> import org.junit.After;
> import org.junit.Before;
> import org.junit.Test;
> import java.util.Date;
> 
> public class HibernateTest {
>     private SessionFactory sessionFactory;
>     private Session session;
>     private Transaction transaction;
> 
>     @Before
>     public void init(){
>         Configuration configuration = new Configuration().configure();
>         ServiceRegistry serviceRegistry = new ServiceRegistryBuilder().applySettings(configuration.getProperties()).buildServiceRegistry();
>         sessionFactory = configuration.buildSessionFactory(serviceRegistry);
>         session = sessionFactory.openSession();
>         transaction = session.beginTransaction();
>     }
> 
>     @After
>     public void destroy(){
>         transaction.commit();
>         session.close();
>         sessionFactory.close();
>     }
> 
>     @Test
>     public void test(){
>         News news = (News) session.get(News.class, 1);
>         System.out.println(news);
>         // 第二次查询不会发送sql语句，会直接从Hibernate一级缓存中取数据
>         News news1 = (News) session.get(News.class, 1);
>         System.out.println(news1);
>     }
> 
>     /**
>      *  flush：使数据表中的记录和Session缓存中的对象的状态保持一致。为了保持一致，则可能会发送对应的SQL语句
>      *  1. 在Transaction的commit() 方法中：先调用session的flush方法，再提交事务
>      *  2. flush() 方法可能会发送SQL语句，但不会提交事务
>      *  3. 注意：在未提交事务或显示的调用session.flush方法之前，也有可能会进行flush操作
>      *     1）执行HQL或QBC查询，会先进行flush操作，以得到数据表的最新的记录
>      *     2) 若记录的ID是由底层数据库使用自增的方式生成的，则在调用save()方法后，就会立即发送insert语句，因为
>      *     save方法后，必须保证对象的ID是存在的
>      */
>     @Test
>     public void testSessionFlush(){
>         News news = (News) session.get(News.class, 1);
>         news.setAuthor("oracle");
>         News news2 = (News) session.createCriteria(News.class).uniqueResult();
>         System.out.println(news2);
>     }
> 
>     @Test
>     public void testSessionFlush2(){
>         News news = new News("Java", "Sun",new Date());
>         session.save(news);
>     }
> 
>     /**
>      *  refresh()：会强制发送SELECT语句，以使Session缓存中的对象状态和数据表中对应的记录保持一致
>      *  这里要体现效果，必须要修改mysql事务的隔离级别，在hibernate.cfg.xml修改配置
>      *   <!--设置Hibernate的事务隔离级别-->
>      *   <property name="connection.isolation">2</property>
>      */
>     @Test
>     public void testRefresh(){
>         News news = (News) session.get(News.class, 1);
>         System.out.println(news);
>         session.refresh(news);
>         System.out.println(news);
>     }
> 
>     /**
>      *  clear(): 清理缓存
>      */
>     @Test
>     public void testClear(){
>         News news = (News) session.get(News.class,1);
>         session.clear();
>         News news1 = (News) session.get(News.class,1);
>     }
> }
> ```
>
> 