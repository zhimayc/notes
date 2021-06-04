Mysql学习笔记
===============================================


SQL基础
-----------------------------------------------

* 创建数据库  create database test1;
* 删除数据库  drop database dbname;
* 创建表      CREATE TABLE tablename (column_name_1 column_type_1 constraints，
  column_name_2  column_type_2  constraints ， ……column_name_n  column_type_n
  constraints）

  create table emp(ename varchar(10),hiredate date,sal decimal(10,2),deptno int(2));


* 创建新表并复制旧表结构
  create table 新表 like 旧表

* 将旧表中的数据复制到新表中
  * 结构相同的两张表
    insert 新表 select * from 旧表

  * 结构不同的两张表
    insert 新表(字段1,字段2,字段3) select 字段1,字段2,字段3 from 旧表

  * 复制整个表（包括结构和数据）
    create table 新表 select * from 旧表

  * 注：只是用select 语句的结果建了一个新表，所以新表不会有主键，索引

* 删除表      DROP TABLE tablename

* 修改表  ALTER TABLE tablename MODIFY [COLUMN] column_definition [FIRST | AFTER col_name]

  alter table emp modify ename varchar(20);

* 增加表字段
  ALTER TABLE tablename ADD [COLUMN] column_definition [FIRST | AFTER col_name]

  alter table emp add column age int(3);

* 删除表字段
  ALTER TABLE tablename DROP [COLUMN] col_name

  alter table emp drop column age;

* 字段改名
  ALTER TABLE tablename CHANGE [COLUMN] old_col_name column_definition
  [FIRST|AFTER col_name]

  alter table emp change age age1 int(4) ;

* 表改名
  ALTER TABLE tablename RENAME [TO] new_tablename

  alter table emp rename emp1;


SQL聚合操作
-----------------------------------------------
* 在公司中统计部门人数并统计总人数

  select deptno,count(1) from emp group by deptno with rollup;


* 统计人数大于1的部门

  select deptno,count(1) from emp group by deptno having count(1)>1;

* 统计公司所有员工的薪水总额、最高和最低薪水

  select sum(sal),max(sal),min(sal) from emp;

* 查询出所有雇员的名字和所在部门的名称

  select ename,deptname from emp ,dept where emp.deptno = dept.deptno;


* 查询emp中所有用户名和所在部门名称

  select ename,deptname from emp left join dept on emp.deptno=dept.deptno;

联合查询
---------------------------------------------
* UNION 和 UNION ALL 的主要区别是 UNION ALL 是把结果集直接合并在一起，而 UNION 是将
UNION ALL 后的结果进行一次 DISTINCT，去除重复记录后的结果。

  select deptno from emp union select deptno from dept;

* DCL 语句主要是 DBA 用来管理系统中的对象权限时所使用，一般的开发人员很少使用。

* 创建一个数据库用户test,密码为123,具有对mydb数据库所有表的INSERT/SELECT权限

  grant select,insert on mydb.* to 'test'@'localhost' identified by '123';


* 收回用户test的insert权限

  revoke insert on mydb.* from 'test'@'localhost';


Mysql数据类型
-----------------------------------------------

* 对于小数的表示，MySQL 分为两种方式：浮点数和定点数。浮点数包括 float（单精度）
和 double（双精度），而定点数则只有 decimal 一种表示。定点数在 MySQL 内部以字符串形
式存放，比浮点数更精确，适合用来表示货币等精度高的数据。
浮点数后面跟“(M,D)”的用法是非标准用法，如果要用于数据库的迁移，则最好不要这么使用。
float 和 double 在不指定精度时，默认会按照实际的精度（由实际的硬件和操作系统决定）
来显示，而 decimal 在不指定精度时，默认的整数位为 10，默认的小数位为 0。


* TIMESTAMP的插入和查询都受当地时区的影响，更能反应出实际的日期。而
DATETIME则只能反应出插入时当地的时区，其他时区的人查看数据必然会有误差
的。


* CHAR 和 VARCHAR 很类似，都用来保存 MySQL 中较短的字符串。二者的主要区别在于存储
方式的不同：CHAR 列的长度固定为创建表时声明的长度，长度可以为从 0～255 的任何值；
而 VARCHAR 列中的值为可变长字符串，长度可以指定为 0～255 （5.0.3以前）或者 65535 （5.0.3
以后）之间的值。在检索的时候，CHAR 列删除了尾部的空格，而 VARCHAR 则保留这些空格。


* ENUM 类型只允许从值集合中选取单个值，而不能一次取多个值。Set 和 ENUM 类型非常类似，
也是一个字符串对象，里面可以包含 0～64 个成员。根据成员的不同，存储上也有所不同。Set 和 ENUM 除了存储之外，最主要的区别在于Set类型一次可以选取多个成员，而 ENUM则只能选一个。Set中对于（ ' a,d,a ' ）这样包含重
复成员的集合将只取一次，写入后的结果为“a,d”。


MySQL常用函数
--------------------------------------------------------

* 字符串函数
  * CONCAT(S1,S2,…Sn)函数：把传入的参数连接成为一个字符串。任何字符串与 NULL 进行连接的结果都将是 NULL。

  * INSERT(str ,x,y,instr)函数：将字符串 str 从第 x 位置开始，y 个字符长的子串替换为字符串 instr

  * LOWER(str)和 UPPER(str)函数：把字符串转换成小写或大写。

  * LEFT(str,x)和 RIGHT(str,x)函数：分别返回字符串最左边的x个字符和最右边的x个字符。如果第二个参数是 NULL，那么将不返回任何字符串。

  * LPAD(str,n ,pad)和 RPAD(str,n ,pad)函数：用字符串 pad 对 str 最左边和最右边进行填充,直到长度为 n 个字符长度。

  * LTRIM(str)和 RTRIM(str)函数：去掉字符串 str 左侧和右侧空格。

  * REPEAT(str,x)函数：返回 str 重复 x 次的结果。

  * REPLACE(str,a,b)函数：用字符串 b 替换字符串 str 中所有出现的字符串 a。

  * TRIM(str)函数：去掉目标字符串的开头和结尾的空格。

  * SUBSTRING(str,x,y)函数：返回从字符串 str 中的第 x 位置起 y 个字符长度的字串。
  此函数经常用来对给定字符串进行字串的提取



* 数值函数

  * ABS(x)函数：返回 x 的绝对值。

  * CEIL(x)函数：返回大于 x 的最大整数。

  * FLOOR(x)函数：返回小于 x 的最大整数，和 CEIL 的用法刚好相反。

  * MOD(x，y)函数：返回 x/y 的模。

  * RAND()函数：返回 0 到 1 内的随机值

  * ROUND(x,y)函数：返回参数 x 的四舍五入的有 y 位小数的值。如果是整数，将会保留 y 位数量的 0；如果不写 y，则默认 y 为 0，即将 x 四舍五入后取整。适合于将所有数字保留同样小数位的情况。

  * TRUNCATE(x,y)函数：返回数字 x 截断为 y 位小数的结果。




日期和时间函数
------------------------------------------------------------

* CURDATE()函数：返回当前日期，只包含年月日。

* CURTIME()函数：返回当前时间，只包含时分秒。

* NOW()函数：返回当前的日期和时间，年月日时分秒全都包含。

* UNIX_TIMESTAMP(date)函数：返回日期 date 的 UNIX 时间戳。

* FROM_UNIXTIME （ unixtime ） 函 数 ： 返 回 UNIXTIME 时 间 戳 的 日 期 值 ， 和
UNIX_TIMESTAMP(date)互为逆操作。

* WEEK(DATE)和 YEAR(DATE)函数：前者返回所给的日期是一年中的第几周，后者返回所
给的日期是哪一年。

* DATE_FORMAT(date,fmt)函数：按字符串 fmt 格式化日期 date 值，此函数能够按指定的
格式显示日期

* DATEDIFF（date1，date2）函数：用来计算两个日期之间相差的天数。


流程函数
------------------------------------------------------

* 模拟了对职员薪水进行分类，这里首先创建并初始化一个职员薪水表

  create table salary (userid int,salary decimal(9,2));

  insert into salary values(1,1000),(2,2000), (3,3000),(4,4000),(5,5000), (1,null);


* IF(value,t,f)函数：

  select if(salary>2000,'high','low') from salary;

* IFNULL(value1,value2)函数：这个函数一般用来替换 NULL 值的，我们知道 NULL 值是不
能参与数值运算的，下面这个语句就是把 NULL 值用 0 来替换。

  select ifnull(salary,0) from salary;

* CASE WHEN [value1] THEN[result1]…ELSE[default]END 

  select case when salary<=2000 then 'low' else 'high' end from salary;


* CASE [expr] WHEN [value1] THEN[result1]…ELSE[default]END 函数：这里还可以分多种情况把职员的薪水分多个档次

  select case salary when 1000 then 'low' when 2000 then 'mid' else 'high' end from salary;



其他常用函数
----------------------------------------------------


* DATABASE()        返回当前数据库名
* VERSION()         返回当前数据库版本
* USER()            返回当前登录用户名
* INET_ATON(IP)     返回 IP 地址的数字表示
* INET_NTOA(num)    返回数字代表的 IP 地址
* PASSWORD(str)     返回字符串 str 的加密版本
* MD5()             返回字符串 str 的 MD5 值


MySQL数据库引擎
-------------------------------------------------------

* 可以通过“ALTER TABLE *** AUTO_INCREMENT = n;”语句强制设置自动增长列的初始值，
默认从 1 开始

* 可以使用 LAST_INSERT_ID()查询当前线程最后插入记录使用的值。


约束
--------------------------------------------------------

* MySQL 支持外键的存储引擎只有 InnoDB，在创建外键的时候，要求父表必须有对应的
索引，子表在创建外键的时候也会自动创建对应的索引

* 下面是样例数据库中的两个表，country 表是父表，country_id 为主键索引，city 表是子
表，country_id 字段对 country 表的 country_id 有外键。

```

CREATE TABLE country (
country_id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
country VARCHAR(50) NOT NULL,
last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (country_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE city (
city_id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
city VARCHAR(50) NOT NULL,
country_id SMALLINT UNSIGNED NOT NULL,
last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (city_id),
KEY idx_fk_country_id (country_id),
CONSTRAINT `fk_city_country` FOREIGN KEY (country_id) REFERENCES country (country_id) ON DELETE RESTRICT ON UPDATE CASCADE)ENGINE=InnoDB DEFAULT CHARSET=utf8;


```

* 为已创建的表添加外键约束
  alter table student add foreign key(class_id) references class(id) on delete restrict on update cascade;

* 删除主键约束
  alter table 表名 drop primary key;

* 删除外键约束
  alter table 表名 drop foreign key 外键名称


* 在导入多个表的数据时，如果需要忽略表之前的导入顺序，可以暂时关闭外键的检查；
同样，在执行 LOAD DATA 和 ALTER TABLE 操作的时候，可以通过暂时关闭外键约束来加快处
理的速度，关闭的命令是“SET FOREIGN_KEY_CHECKS = 0;”，执行完成之后，通过执行“SET
FOREIGN_KEY_CHECKS = 1;”语句改回原状态。


索引
--------------------------------------------------------

* 创建普通索引

  * create index 索引名 on 表名(字段名)

  * alter table 表名 add index 索引名(字段名)
* 唯一索引 必须唯一，可以有空值
  create unique index 索引名 on 表名(字段名)
  alter table 表名 add unique(字段名)
* 主键索引
  创建表时添加 或 alter table 表名 add primary key(字段名)
* 组合索引
  create index 索引名 on 表名(字段1,字段2)
* 删除索引
  drop index 索引名称 on 表名



* explain字段解析
  
  > ==id:选择标识符==
  >
  > 【 id相同时，执行顺序由上至下。】
  >
  > 【如果是子查询，id的序号会递增，id值越大优先级越高，越先被执行。】
  >
  > 【 id如果相同，可以认为是一组，从上往下顺序执行；在所有组中，id值越大，优先级越高，越先执行】
  >
  > 
  >
  > ==select_type:表示查询的类型。==
  >
  > 【SIMPLE(简单SELECT，不使用UNION或子查询等)】
  >
  > 【PRIMARY(子查询中最外层查询，查询中若包含任何复杂的子部分，最外层的select被标记为PRIMARY)】
  >
  > 【UNION(UNION中的第二个或后面的SELECT语句)】
  >
  > 【 DEPENDENT UNION(UNION中的第二个或后面的SELECT语句，取决于外面的查询)】
  >
  > 【UNION RESULT(UNION的结果，union语句中第二个select开始后面所有select)】
  >
  > 【 SUBQUERY(子查询中的第一个SELECT，结果不依赖于外部查询)】
  >
  > 【DEPENDENT SUBQUERY(子查询中的第一个SELECT，依赖于外部查询)】
  >
  > 【 DERIVED(派生表的SELECT, FROM子句的子查询)】
  >
  > 【UNCACHEABLE SUBQUERY(一个子查询的结果不能被缓存，必须重新评估外链接的第一行)】
  >
  > 
  >
  > ==table:输出结果集的表【显示这一步所访问数据库中表名称（显示这一行的数据是关于哪张表的），有时不是真实的表名字，可能是简称，例如上面的e，d，也可能是第几步执行的结果的简称】==
  >
  > partitions:匹配的分区
  >
  > 
  >
  > ==type:表示表的连接类型【常用的类型有：ALL、index、range、 ref、eq_ref、const、system、NULL（从左到右，性能从差到好）】==
  >
  > 【ALL：Full Table Scan， MySQL将遍历全表以找到匹配的行】
  >
  > 【index: Full Index Scan，index与ALL区别为index类型只遍历索引树】
  >
  > 【range:只检索给定范围的行，使用一个索引来选择行】
  >
  > 【ref: 表示上述表的连接匹配条件，即哪些列或常量被用于查找索引列上的值】
  >
  > 【eq_ref: 类似ref，区别就在使用的索引是唯一索引，对于每个索引键值，表中只有一条记录匹配，简单来说，就是多表连接中使用primary key或unique key作为关联条件】
  >
  > 【const、system: 当MySQL对查询某部分进行优化，并转换为一个常量时，使用这些类型访问。如将主键置于where列表中，MySQL就能将该查询转换为一个常量，system是const类型的特例，当查询的表只有一行的情况下，使用system】
  >
  > 【NULL: MySQL在优化过程中分解语句，执行时甚至不用访问表或索引，例如从一个索引列里选取最小值可以通过单独索引查找完成。】
  >
  > 
  >
  > ==possible_keys:表示查询时，可能使用的索引【指出MySQL能使用哪个索引在表中找到记录，查询涉及到的字段上若存在索引，则该索引将被列出，但不一定被查询使用（该查询可以利用的索引，如果没有任何索引显示 null）】==
  >
  > 
  >
  > ==key:表示实际使用的索引【key列显示MySQL实际决定使用的键（索引），必然包含在possible_keys中】==
  >
  > 
  >
  > ==key_len:索引字段的长度==
  >
  > 
  >
  > ==ref:列与索引的比较==
  >
  > 
  >
  > ==rows:扫描出的行数(估算的行数)==
  >
  > ==filtered:按表条件过滤的行百分比==
  >
  > 
  >
  > ==Extra:执行情况的描述和说明【该列包含MySQL解决查询的详细信息,有以下几种情况：】==
  >
  > 【Using where:不用读取表中所有信息，仅通过索引就可以获取所需数据，这发生在对表的全部的请求列都是同一个索引的部分的时候，表示mysql服务器将在存储引擎检索行后再进行过滤】
  >
  > 【Using temporary：表示MySQL需要使用临时表来存储结果集，常见于排序和分组查询，常见 group by ; order by】
  >
  > 【Using filesort：当Query中包含 order by 操作，而且无法利用索引完成的排序操作称为“文件排序”】
  >
  > 【Using join buffer：改值强调了在获取连接条件时没有使用索引，并且需要连接缓冲区来存储中间结果。如果出现了这个值，那应该注意，根据查询的具体情况可能需要添加索引来改进能。】
  >
  > 【Impossible where：这个值强调了where语句会导致没有符合条件的行（通过收集统计信息不可能存在结果）。】
  >
  > 【Select tables optimized away：这个值意味着仅通过使用索引，优化器可能仅从聚合函数结果中返回一行】
  >
  > 【No tables used：Query语句中使用from dual 或不含任何from子句】



* 避免索引失效的情况

1. 尽量用全值匹配（即索引列=value值，例如select * from emp where name='tom'）
2. 最佳左前缀法则（如果索引了多列，要遵守最左前缀法则。指的是查询从从索引的最左前列开始，并且*不跳过索引中的列*）
3. 不要在索引列上做任何的操作（计算、函数【自动or手动】类型转换），会导致索引失效而转向全表扫描
4. 存储引擎不能使用索引中范围条件右边的列(例如 ==select * from emp where name = "mike" and age > 15 and nickname="m"== 则age之后的索引列全部失效)
5. 尽量使用覆盖索引（只访问索引的查询【索引列和查询列一致】，减少select *,例如索引列是name,age 则查询语句为select name,age from emp where age > 10）
6. mysql在使用不等于（!=或<>）的时候无法使用索引会导致全表扫描
7. is null , is not null 也无法使用索引
8. like以通配符开头（'%abc...'）mysql索引失效会变成全表扫描（尽量加在右边'abc%'），如果一定要用'like %abc%'，则可以用覆盖索引（即查询的列跟索引的列完全一致。 例如 select name,age,email from emp 【索引是idx_nameAgeEmail】）
9. 字符串不加单引号索引失效（例如select * from emp where name=2000,这里name是varchar，name=2000会发生隐式的类型转换，将int型转换为varchar类型，需要写成select * from emp where name='2000'，其实符合第3条原则，不在索引类上做自动的类型转换）
10. 少用or，用它来连接时，索引会失效
11. order by 尽量在索引列上，否则会导致filesort



* 设计索引的原则
1. 搜索的索引列，不一定是所要选择的列。换句话说，最适合索引的列是出现在 WHERE
子句中的列，或连接子句中指定的列，而不是出现在 SELECT 关键字后的选择列表中的列。

2. 使用惟一索引。考虑某列中值的分布。索引的列的基数越大，索引的效果越好。例
如，存放出生日期的列具有不同值，很容易区分各行。而用来记录性别的列，只含有“ M”
和“F”，则对此列进行索引没有多大用处，因为不管搜索哪个值，都会得出大约一半的行。

3. 使用短索引。如果对字符串列进行索引，应该指定一个前缀长度，只要有可能就应
该这样做。例如，如果有一个 CHAR(200)列，如果在前 10 个或 20 个字符内，多数值是惟一
的，那么就不要对整个列进行索引。对前10个或20个字符进行索引能够节省大量索引空间，
也可能会使查询更快。较小的索引涉及的磁盘 IO 较少，较短的值比较起来更快。更为重要
的是，对于较短的键值，索引高速缓存中的块能容纳更多的键值，因此，MySQL 也可以在
内存中容纳更多的值。这样就增加了找到行而不用读取索引中较多块的可能性。

4. 利用最左前缀。在创建一个n列的索引时，实际是创建了MySQL可利用的n个索引。
多列索引可起几个索引的作用，因为可利用索引中最左边的列集来匹配行。这样的列集称为
最左前缀。

5. 不要过度索引。不要以为索引“越多越好”，什么东西都用索引是错误的。每个额
外的索引都要占用额外的磁盘空间，并降低写操作的性能。在修改表的内容时，索引必须进
行更新，有时可能需要重构，因此，索引越多，所花的时间越长。如果有一个索引很少利用
或从不使用，那么会不必要地减缓表的修改速度。此外，MySQL 在生成一个执行计划时，
要考虑各个索引，这也要花费时间。创建多余的索引给查询优化带来了更多的工作。索引太
多，也可能会使 MySQL 选择不到所要使用的最好索引。只保持所需的索引有利于查询优化。

6. 对于 InnoDB 存储引擎的表，记录默认会按照一定的顺序保存，如果有明确定义的主
键，则按照主键顺序保存。如果没有主键，但是有唯一索引，那么就是按照唯一索引的顺序
保存。如果既没有主键又没有唯一索引，那么表中会自动生成一个内部列，按照这个列的顺
序保存。按照主键或者内部列进行的访问是最快的，所以 InnoDB 表尽量自己指定主键，当
表中同时有几个列都是唯一的，都可以作为主键的时候，要选择最常作为访问条件的列作为
主键，提高查询的效率。另外，还需要注意，InnoDB 表的普通索引都会保存主键的键值，
所以主键要尽可能选择较短的数据类型，可以有效地减少索引的磁盘占用，提高索引的缓存
效果。


视图
------------------------------------------------------

* 创建视图
  create view 视图名[字段...] as 查询的SQL语句

  * create view myview as select id from mytable2;

  * create view myview2(city_id,myname) as select city_id,city from city limit 2;

* 更新视图
create or replace view 视图名 [字段] as 查询的SQL语句  (有就替换，没有就创建)
或alter view 视图名 [字段...] as 查询的SQL语句

* 删除视图

  drop view 视图名

* 查看数据库中那些是表 那些是视图

  show table status



存储过程和函数
---------------------------------------------------------


* 变量定义
	delimiter $$
	create procedure p_vartest()
	begin 
	declare a varchar(20) default 'abc';
	select a;
	end $$

* 调用：

	delimiter ;
	call p_vartest();


* 参数类型
	delimiter $$
	
	create procedure p_test_out(out v_out_int int)
	begin
	select v_out_int;
	set v_out_int=15;
	select v_out_int;
	end $$

* 调用：
	delimiter ;
	set @v_out_int=10;
	call p_test_out(@v_out_int);

* 输出：
	v_out_int 
	 NULL
	
	 v_out_int
	 15

 * 输入参数(in)，存储过程会复制一份输入参数的值，不会改变它本身的值，而输出参数(out)有这样一个特点，即便我们之前给这个变量赋了值，但它的初始值仍然为空值。输入输出参数(inout)
 兼并输入和输出这两个参数的值。


* 选择语句

	delimiter $$
	
	create procedure p_showage(in age int)
	begin
	if age>=18 then
	select '成年人';
	else 
	select '未成年人';
	end if;
	end $$

* 调用：
	delimiter ;
	set @age=19;
	call p_showage(@age);

* case分支
	create procedure p_addsalary(in v_empno int)
	begin
	declare adds int;
	case v_empno
	when 1001 then
	set adds=1500;
	when 1002 then
	set adds=2500;
	when 1003 then
	set adds=3200;
	else
	set adds=1000;
	end case;
	end $$


* while循环语句
	create procedure p_addnum()
	begin
	declare i int default 1;
	declare addresult int default 0;
	while i <= 100 do
	set addresult=addresult+i;
	set i=i+1;
	end while;
	select addresult;
	end $$

	create procedure p_insertEmp()
	begin 
	declare maxempno int default 0;
	declare i int default 1;
	while i<=100 do
	select max(emp_no) into maxempno from employees;  /*将最大值赋值给maxempno*/
	set maxempno=maxempno+1;
	insert into employees(emp_no,birth_date,first_name,last_name,gender,hire_date)values(
	maxempno,'1991-01-01','jmi','ui','M','2011-01-05');
	set i=i+1;
	end while;
	end $$


* repeat循环语句
	create procedure p_updateGender()
	begin
	declare imin int default 1;
	declare imax int default 1;
	select min(emp_no) into imin from employees;
	select max(emp_no) into imax from employees;
	repeat
	if imin%2=0 then
	update employees set gender='F' where emp_no = imin;
	end if;
	set imin=imin+1;
	until imin>imax
	end repeat;
	end $$


* loop循环语句
	create procedure p_updateHire()
	begin
	declare imin int default 1;
	declare imax int default 1;
	select min(emp_no) into imin from employees;
	select max(emp_no) into imax form employees;
	myloop:loop
	if imin%2=1 then
	update employees set hire_date='1991-10-12' where emp_no=imin;
	end if;
	set imin=imin+1;
	if imin>imax then
	leave myloop;
	end if;
	end loop;
	end $$


* 定义条件和处理（错误处理，当前面的代码执行出错，后面的代码也能执行）
declare continue handler for sqlstate '错误代码值' set 变量=变量值;

* 存储过程管理
show procedure status where db='数据库名';
select specific_name from mysql.proc; //查看存储过程的名字
select specific_name,body from mysql.proc where specific_name='存储过程名字';

drop procedure if exists 存储过程名字;//删除存储过程

alter procedure 存储过程名 选项 //修改存储过程



* 创建自定义函数
  1. 查看是否开启创建函数功能。
     show variables like '%fun%';
  2. 如果值为OFF，设置开启
     set global log_bin_trust_function_creators=1;
  3. 创建函数的语法
	    create function 函数名( 变量1，变量2.....)    //它有点类似于我们的存储过程，只是多了个变量1，变量2
	     returns 数据类型
	    begin       //函数执行体
	   ......执行的程序代码
	    return 数据;
	    end;


* 示例：
	delimiter $$
	create function fun_add(a int,b int)
	returns int
	begin 
	return a+b;
	end;
	$$
	
	delimiter ;
	select fun_add(3,4);
	
	
	show create function fun_add;     //查看刚才创建的函数
	drop function if exists fun_add;  //删除函数


触发器
----------------------------------------------------
* 创建触发器的语法

	CREATE TRIGGER 语法
	CREATE TRIGGER trigger_name trigger_time trigger_event
	ON tbl_name FOR EACH ROW trigger_stmt


* 触发程序是与表有关的命名数据库对象，当表上出现特定事件时，将激活该对象。触发程序与命名为tbl_name的表相关。tbl_name必须引用永久性表。不能将触发程序与TEMPORARY表或视图关联起来。

* trigger_time

触发程序的动作时间。它可以是BEFORE或AFTER，以指明触发程序是在激活它的语句之前或之后触发。

* trigger_event

>指明了激活触发程序的语句的类型。trigger_event可以是下述值之一：INSERT：将新行插入表时激>活触发程序，例如，通过INSERT、LOAD DATA和REPLACE语句。UPDATE：更改某一行时激活触发程序，例如，通过UPDATE语句。DELETE：从表中删除某一行时激活触发程序，例如，通过DELETE和REPLACE语句。请注意，trigger_event与以表操作方式激活触发程序的SQL语句并不很类似，这点很重要。例如，关于INSERT的BEFORE触发程序不仅能被INSERT语句激活，也能被LOAD DATA语句激活。可能会造成混淆的例子之一是INSERT INTO .. ON DUPLICATE UPDATE ...语法：BEFORE INSERT触发程序对于每一行将激活，后跟AFTER INSERT触发程序，或BEFORE UPDATE和AFTER UPDATE触发程序，具体情况取决于行上是否有重复键。对于具有相同触发程序动作时间和事件的给定表，不能有两个触发程序。例如，对于某一表，不能有两个BEFORE UPDATE触发程序。但可以有1个BEFORE UPDATE触发程序和1个BEFOREINSERT触发程序，或1个BEFORE UPDATE触发程序和1个AFTER UPDATE触发程序。

* trigger_stmt

>是当触发程序激活时执行的语句。如果你打算执行多个语句，可使用BEGIN ... END复合语句结构。这样，就能使用存储子程序中允许的相同语句。

* 特别说明

1. 对于insert而言，新插入的行用new来表示，行中的每一列的值用new.列名来表示；

2. 对于delete而言，对于delete而言：原本有一行,后来被删除，想引用被删除的这一行，用old来表示，old.列名可以引用被删除的行的值。

3. 对update而已，它实际上是有删除和新增这两个操作，对于update而言，对于update而言：被修改的行，修改前的数据，用old来表示，old.列名引用被修改之前行中的值；修改的后的数据，用new来表示，new.列名引用被修改之后行中的值。


* 示例：
	delimiter $$
	
	create trigger tr_insert after insert on emp for each row
	begin
	insert into dept(deptname)values(new.ename);
	end;
	$$
	delimiter $$
	create trigger tr_update after update on emp for each row
	begin
	update dept set deptname=new.ename where deptno=6;
	end;
	$$
	
	

show triggers;   //查看所有触发器
	select * from information_schema.triggers where trigger_name='tr_update';   //查看某个触发器
	drop trigger tr_deleteEmp;   //删除tr_deleteEmp这个触发器


锁
-------------------------------------------------------------

* MySQL锁概述
>相对其他数据库而言，MySQL的锁机制比较简单，其最显著的特点是不同的存储引擎支持不同的锁机制。比如：MyISAM和MEMORY存储引擎采用的是表级锁(table-level locking)。BDB存储引擎采用的是页面锁(page-level locking),但也支持表级锁。InnoDB存储引擎既支持行级锁(row-level locking)也支持表级锁，但默认情况下是采用行级锁。

> MySQL这3种锁的特性可大致归纳如下：
开销，加锁速度，死锁，粒度，并发性能。
1、表级锁:开销小，加锁快，不会出现死锁。锁定粒度大，发生锁冲突概率最高，并发度最低。
2、行级锁:开销大，加锁慢，会出现死锁。锁定粒度最小。发生锁冲突的概率最低，并发度也最高。
3、页面锁:开锁和加锁时间介于行锁和表锁之间，会出现死锁，锁定粒度介于表锁和行锁之间，
并发度一般。仅从锁的角度来说：表级锁更适合以查询为主，只有少量按索引条件更新数据的应用，
如web应用。行级锁更适合有大量按索引条件并发更新少量不同数据。同时又有并发查询的应用，如
一些在线事务处理系统。


>一个session 用lock table 给表f加了读锁，这个session可以查询锁定表中的记录，但更新或访问其他表都会提示错误，同时另一个session可以查询表中的记录，但更新就会出现锁等待。

>lock table 表名 read   //给表加共享读锁,只有当前session可以进行写操作，其他session只能进行读操作
>lock tables 表名1 read,表名2 write
>lock table 表名 write  //给表加独占写锁，只有当前session可以进行读写操作，其他session不能进行读，写操作。
>unlock tables   //给表解锁

* 并发插入
>MyISAM存储引擎有一个系统变量concurrent_insert，专门用以控制其并发插入的行为，其值分别可以为0、1或2。

>1）当concurrent_insert设置为0时，不允许并发插入。

>2）当concurrent_insert设置为1时，如果MyISAM表中没有空洞（即表的中间没有被删除的行），MyISAM允许在一个进程读表的同时，另一个进程从表尾插入记录。这也是MySQL的默认设置。

>3）当concurrent_insert设置为2时，无论MyISAM表中有没有空洞，都允许在表尾并发插入记录。

>set global concurrent_insert=2;
>lock table 表名 read local;

>一个进程请求某个MyISAM表的读锁，同时另一个进程也请求同一表的写锁？Mysql如何处理呢？
答案：写进程先获得锁，即使读请求先到锁等待队列，写请求后到，写请起后到，写锁也会插到读锁队列。Mysql认为写请求比读请求更重要。MyISAM不太适合于有大量更新操作和查询操作的原因，因为，大量的更新操作会造成查询操作很难获得读锁，从而可能永远阻塞。读的这个进程会一直等到这个写的进程完成之后，它才会得到这个请求。

* 解决方法

>执行 set low_priority_updates =1;使该连接发出的更新请求优先级降低。它的优先级降低那么实际上这个查询操作就更能获得我们这个锁。 其中insert，delete也可以通过此种方法指定。




事务
-------------------------------------------------------

* set autocommit=0 //取消自动提交


* 事务提交
	start transaction;
	update account set accmoney=accmoney-1000 where aid=1;
	update account set accmoney=accomoney+1000 where aid=2;
	commit;

* 事务回滚
	 start transaction;
	 insert into account value(3,'C',1000);
	 savepoint s1;  // 创建还原点
	 insert into account value(4,'D',6000);
	 savepoint s2;
	 insert into account value(5,'E',2000);
	 savepoint s3;
	 rollback to s2;  //回滚到还原点s2
	 commit;


* 补充
	commit and chain;   //表示提交事务之后重新开启了新事务
	rollback release;   //表示事务回滚之后断开和客户端的连接
	
	
	
	##### 事务的隔离级别
	
	事务基本要素（ACID）
	
	1、原子性（Atomicity）：
	事务开始后所有操作，要么全部做完，要么全部不做。事务执行过程中出错，会回滚到事务开始前的状态，所有的操作就像没有发生一样。也就是说事务是一个不可分割的整体，就像化学中学过的原子，是物质构成的基本单位。
	
	2、一致性（Consistency）：
	事务开始前和结束后，数据库的完整性约束没有被破坏 。比如A向B转账，不可能A扣了钱，B却没收到。
	3、隔离性（Isolation）：
	同一时间，只允许一个事务请求同一数据，不同的事务之间彼此没有任何干扰。比如A正在从一张银行卡中取钱，在A取钱的过程结束前，B不能向这张卡转账。
	4、持久性（Durability）：
	事务完成后，事务对数据库的所有更新将被保存到数据库，不能回滚。
	
	
	
	事务的并发问题
	
	1、脏读：
	事务A读取了事务B更新的数据，然后B回滚操作，那么A读取到的数据是脏数据
	
	2、不可重复读：
	事务 A 多次读取同一数据，事务 B 在事务A多次读取的过程中，对数据作了更新并提交，导致事务A多次读取同一数据时，结果不一致。
	
	3、幻读：
	系统管理员A将数据库中所有学生的成绩从具体分数改为ABCDE等级，但是系统管理员B就在这个时候插入了一条具体分数的记录，当系统管理员A改结束后发现还有一条记录没有改过来，就好像发生了幻觉一样，这就叫幻读
	
	**小结：**
	不可重复读的和幻读很容易混淆，不可重复读侧重于**修改**，幻读侧重于**新增或删除**（条数不同）。
	
	MySQL事务隔离级别
	
	| 事务隔离级别                   | 脏读 | 不可重复读 | 幻读 |
	| :----------------------------- | :--- | :--------- | :--- |
	| 读未提交（read-uncommitted）   | 是   | 是         | 是   |
	| 不可重复读（read-committed）） | 否   | 是         | 是   |
	| 可重复读（repeatable-read）    | 否   | 否         | 是   |
	| 串行化（serializable）         | 否   | 否         | 否   |
	
	**隔离级别小结：**
	（1）不可重复读解决脏读问题
	（2）可重复读解决脏读、不可重复读问题
	（3）串行化解决脏读、不可重复读和幻读问题
	（4）隔离级别越高，越能保证数据的完整性和一致性，但是对并发性能的影响也越大
	
	mysql默认的事务隔离级别为可重复读（repeatable-read）