# MyBatis学习笔记

### MyBatis基本介绍

##### 入门案例

> mybatis-config.xml
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE configuration
> PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
> "http://mybatis.org/dtd/mybatis-3-config.dtd">
> <configuration>
> <environments default="development">
> <environment id="development">
>    <transactionManager type="JDBC"/>
>    <dataSource type="POOLED">
>        <property name="driver" value="com.mysql.jdbc.Driver"/>
>        <property name="url" value="jdbc:mysql://192.168.100.199/test"/>
>        <property name="username" value="root"/>
>        <property name="password" value="root"/>
>    </dataSource>
> </environment>
> </environments>
> <!--将我们写好的sql映射文件一定要注册到全局配置文件中-->
> <mappers>
> <mapper resource="EmployeeMapper.xml"/>
> </mappers>
> </configuration>
> ```
>
> EmployeeMapper.xml
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE mapper
> PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
> "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
> <mapper namespace="com.example.mybatis.EmployeeMapper">
> <!--
> namespace：名称空间（随便起）
> id：唯一标识
> resultType：返回值类型
> #{id}：从传递过来的参数中取出id值【占位符】
> -->
> <select id="selectEmp" resultType="com.example.mybatis.bean.Employee">
> select id,last_name as lastName,email,gender from tbl_employee where id = #{id}
> </select>
> </mapper>
> ```
>
> ```java
> package com.example.mybatis.test;
> import com.example.mybatis.bean.Employee;
> import org.apache.ibatis.io.Resources;
> import org.apache.ibatis.session.SqlSession;
> import org.apache.ibatis.session.SqlSessionFactory;
> import org.apache.ibatis.session.SqlSessionFactoryBuilder;
> import org.junit.Test;
> import java.io.IOException;
> import java.io.InputStream;
> 
> public class MyBatisTest {
> /**
>      *  1.根据xml配置文件（全局配置文件）创建一个SqlSessionFactory对象，有数据源一些运行环境信息
>      *  2.sql映射文件，配置了每一个sql，以及sql的封装规则
>      *  3.将sql的映射文件注册在全局配置文件中
>      *  4.写代码：
>      *      1）根据全局配置文件得到SqlSessionFactory
>      *      2）使用SqlSessionFactory，获取到SqlSession对象，使用它来执行增删改查
>      *         一个sqlSession就是代表和数据库的一次会话，用完关闭
>      *      3）使用sql的唯一标识来告诉MyBatis执行哪个sql，sql语句都是保存在sql映射文件中的
>      *
>      */
>     @Test
>     public void test() throws IOException {
>         String resource = "mybatis-config.xml";
>         InputStream inputStream = Resources.getResourceAsStream(resource);
>         SqlSessionFactory sqlSessionFactory = new SqlSessionFactoryBuilder().build(inputStream);
>         // 获取sqlSession实例，能直接执行已经映射的sql语句
>         SqlSession openSession = sqlSessionFactory.openSession();
>         // selectOne()第一个参数：sql的唯一标识（namespace + id）; 第二个参数：执行sql要用的参数
>         try {
>             Employee employee = openSession.selectOne("com.example.mybatis.EmployeeMapper.selectEmp", 1);
>             System.out.println(employee);
>         } catch (Exception e) {
>             e.printStackTrace();
>         } finally {
>             openSession.close();
>         }
>     }
> }
> ```
>
> ```java
> package com.example.mybatis.bean;
> 
> public class Employee {
>     private Integer id;
>     private String lastName;
>     private String email;
>     private String gender;
> 
>     public Integer getId() {
>         return id;
>     }
> 
>     public void setId(Integer id) {
>         this.id = id;
>     }
> 
>     public String getLastName() {
>         return lastName;
>     }
> 
>     public void setLastName(String lastName) {
>         this.lastName = lastName;
>     }
> 
>     public String getEmail() {
>         return email;
>     }
> 
>     public void setEmail(String email) {
>         this.email = email;
>     }
> 
>     public String getGender() {
>         return gender;
>     }
> 
>     public void setGender(String gender) {
>         this.gender = gender;
>     }
> 
>     @Override
>     public String toString() {
>         return "Employee{" +
>                 "id=" + id +
>                 ", lastName='" + lastName + '\'' +
>                 ", email='" + email + '\'' +
>                 ", gender='" + gender + '\'' +
>                 '}';
>     }
> }
> 
> ```
>
> 第二种实现方式（通过绑定接口实现）：
>
> ```java
> package com.example.mybatis.dao;
> import com.example.mybatis.bean.Employee;
> public interface EmployeeMapper {
>     public Employee getEmpById(Integer id);
> }
> 
> ```
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE mapper
>         PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
>         "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
> <mapper namespace="com.example.mybatis.dao.EmployeeMapper">
>     <!--
>         namespace：名称空间指定为接口的全类名
>         id：写接口里面对应的方法名
>         resultType：返回值类型
>         #{id}：从传递过来的参数中取出id值【占位符】
>     -->
>     <select id="getEmpById" resultType="com.example.mybatis.bean.Employee">
>       select id,last_name as lastName,email,gender from tbl_employee where id = #{id}
>     </select>
> </mapper>
> ```
>
> ```java
> public SqlSessionFactory getSqlSessionFactory() throws IOException {
>     String resource = "mybatis-config.xml";
>     InputStream inputStream = Resources.getResourceAsStream(resource);
>     return new SqlSessionFactoryBuilder().build(inputStream);
> }
> 
> @Test
> public void test1() throws IOException {
>     SqlSessionFactory sqlSessionFactory = getSqlSessionFactory();
>     SqlSession openSession = sqlSessionFactory.openSession();
> 
>     try {
>         // ；获取接口实现类对象
>         // 会为接口自动的创建一个代理对象，代理对象去执行增删改查
>         EmployeeMapper mapper = openSession.getMapper(EmployeeMapper.class);
>         Employee employee = mapper.getEmpById(1);
>         System.out.println(employee);
>     } finally {
>         openSession.close();
>     }
> }
> ```
>
> *小结*
>
> 接口式编程：
>
> 原生：		Dao ====> DaoImpl
>
> mybatis：  Mapper ====> xxMapper.xml
>
> sqlSession代表和数据库的一次会话，用完必须关闭；
>
> SqlSession 和 Connection 一样，都是非线程安全的，每次使用都应该去获取新的对象
>
> mapper接口没有实现类，按时mybatis会为这个接口生成一个代理对象（将接口和xml绑定）
>
> 两个重要的配置文件：
>
> ​	mybatis的全局配置文件：包含数据库连接池信息，事务管理器信息，系统运行环境信息……
>
> ​	sql映射文件：保存了每一个sql语句的映射信息，将sql抽取出来

### 基础配置

##### 全局配置文件

> **引入外部配置文件**
>
> ```properties
> jdbc.driver=com.mysql.jdbc.Driver
> jdbc.url=jdbc:mysql://192.168.100.199/test
> jdbc.username=root
> jdbc.password=root
> ```
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE configuration
> PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
> "http://mybatis.org/dtd/mybatis-3-config.dtd">
> <configuration>
> <!--
> mybatis可以使用properties来引入外部properties配置文件的内容
> resource：引入类路径下的资源
> url：引入网络路径或者磁盘路径下的资源
> -->
> <properties resource="dbconfig.properties"></properties>
> <environments default="development">
> <environment id="development">
> <transactionManager type="JDBC"/>
> <dataSource type="POOLED">
> <property name="driver" value="${jdbc.driver}"/>
> <property name="url" value="${jdbc.url}"/>
> <property name="username" value="${jdbc.username}"/>
> <property name="password" value="${jdbc.password}"/>
> </dataSource>
> </environment>
> </environments>
> <!--将我们写好的sql映射文件一定要注册到全局配置文件中-->
> <mappers>
> <mapper resource="EmployeeMapper.xml"/>
> </mappers>
> </configuration>
> ```
>
> **驼峰命名**
>
> ```xml
> <!--
> settings包含很多重要的设置项
> setting：用来设置每一个设置项
> name：设置项名
> value：设置项取值
> -->
> <settings>
> <!--是否开启驼峰命名自动映射，即从经典数据库列名 A_COLUMN 映射到经典 Java 属性名 aColumn。-->
> <setting name="mapUnderscoreToCamelCase" value="true"/>
> </settings>
> ```
>
> **类型别名（typeAliases）**
> 类型别名可为 Java 类型设置一个缩写名字。 它仅用于 XML 配置，意在降低冗余的全限定类名书写。
> typeAlias：为某个java类型取别名
> type：指定要起别名的类型全类名，默认别名是类名小写
> alias：指定新的别名
>
> *单个类别名*
>
> mybatis-config.xml
>
> ```xml
> <typeAliases>
> <typeAlias type="com.example.mybatis.bean.Employee" alias="emp"/>
> </typeAliases>
> ```
>
> EmployeeMapper.xml
>
> ```xml
> <mapper namespace="com.example.mybatis.dao.EmployeeMapper">
> <!--
> namespace：名称空间指定为接口的全类名
> id：写接口里面对应的方法名
> resultType：返回值类型
> #{id}：从传递过来的参数中取出id值【占位符】
> -->
> <select id="getEmpById" resultType="emp">
> select * from tbl_employee where id = #{id}
> </select>
> </mapper>
> ```
>
> *批量起别名*
>
> ```xml
> <typeAliases>
> <!--
> 批量起别名
> package：为某个包下的所有类批量起别名
> name：指定包名（为当前包以及下面所有的包的每一个类都起一个默认别名）
> 每一个在包 com.example.mybatis.bean 中的 Java Bean，在没有注解的情况下，会使用 Bean 的首字母小写的非限定类名来作为它的别名。
> 比如 com.example.mybatis.bean.Author 的别名为 author；若有注解，则别名为其注解值。
> @Alias("author")
> public class Author {
> ...
> }
> -->
> <package name="com.example.mybatis.bean" />
> </typeAliases>
> ```
>
> **environments运行环境**
>
> ```xml
> <!--environments：mybatis可以配置多种环境，default指定使用某种环境
> environment：配置一个具体的环境信息；必须有两个标签,id代表当前环境的唯一标识
> transactionManager：事务管理器
>  type：事务管理器类型；JDBC(JdbcTransactionFactory)   |   MANAGED(ManagedTransactionFactory)
>        自定义事务管理器；实现TransactionFactory接口 type指定为全类名
> dataSource：数据源
>  type：数据源类型；UNPOOLED(UnpooledDataSourceFactory) | POOLED(PooledDataSourceFactory)  |  JNDI(JndiDataSourceFactory)
>  自定义数据源：实现DataSourceFactory接口，type是全类名
> -->
> <environments default="development">
> <environment id="test">
> <transactionManager type="JDBC"/>
> <dataSource type="POOLED">
>    <property name="driver" value="${jdbc.driver}"/>
>    <property name="url" value="${jdbc.url}"/>
>    <property name="username" value="${jdbc.username}"/>
>    <property name="password" value="${jdbc.password}"/>
> </dataSource>
> </environment>
> <environment id="development">
> <transactionManager type="JDBC"/>
> <dataSource type="POOLED">
>    <property name="driver" value="${jdbc.driver}"/>
>    <property name="url" value="${jdbc.url}"/>
>    <property name="username" value="${jdbc.username}"/>
>    <property name="password" value="${jdbc.password}"/>
> </dataSource>
> </environment>
> </environments>
> ```
>
> **配置多数据库支持**
>
> mybatis-config.xml
>
> ```xml
> <!--databaseIdProvider支持多数据库厂商
> DB_VENDOR：VendorDatabaseIdProvider，作用就是得到数据库厂商标识（驱动），mybatis就能根据数据库厂商标识来执行不同的sql
> MySQL，Oracle，SQL Server，
> -->
> <databaseIdProvider type="DB_VENDOR">
> <!--为不同数据库厂商取别名-->
> <property name="MySQL" value="mysql"/>
> <property name="Oracle" value="oracle"/>
> <property name="SQL Server" value="sqlserver"/>
> </databaseIdProvider>
> ```
>
> EmployeeMapper.xml
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE mapper
> PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
> "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
> <mapper namespace="com.example.mybatis.dao.EmployeeMapper">
> <!--这里databaseId的值填上面property属性的value值-->
> <select id="getEmpById" resultType="employee" databaseId="mysql">
> select * from tbl_employee where id = #{id}
> </select>
> </mapper>
> ```
>
> **映射器（mappers）**
>
> mybatis-config.xml
>
> ```xml
> <!--将我们写好的sql映射文件一定要注册到全局配置文件中-->
> <!--mappers：将sql映射注册到全局配置中
> resource：引用类路径下的sql映射文件
> url：引用网络路径或者磁盘路径下的sql映射文件
> class：引用（注册）接口
> 1.必须有sql映射文件，映射文件名必须和接口同名，并且放在与接口同一目录下
> 2.没有sql映射文件，所有的sql都是利用注解写在接口上
>     推荐：比较重要的，复杂的Dao接口我们来写sql映射文件
>     不重要Dao接口为了开发快速可以使用注解
> -->
> <mappers>
> <mapper resource="EmployeeMapper.xml"/>
> <mapper class="com.example.mybatis.dao.EmployeeMapper"/>
> <!--批量注册，接口名和映射文件名必须相同，并且在同一目录下，或者可以使用注解-->
> <package name="com.example.mybatis.dao"/>
> </mappers>
> ```
>
> 注解实现sql映射
>
> ```java
> package com.example.mybatis.dao;
> import com.example.mybatis.bean.Employee;
> import org.apache.ibatis.annotations.Select;
> 
> public interface EmployeeMapper {
> @Select("select * from tbl_employee where id = #{id}")
> public Employee getEmpById(Integer id);
> 
> @Select("select * from tbl_employee where last_name = #{lastName}")
> public Employee getEmpByLastName(String lastName);
> }
> ```
>
> **增删改查**
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE mapper
>   PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
>   "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
> <mapper namespace="com.example.mybatis.dao.EmployeeMapper">
> <!--
>   namespace：名称空间指定为接口的全类名
>   id：写接口里面对应的方法名
>   resultType：返回值类型
>   #{id}：从传递过来的参数中取出id值【占位符】
> -->
> <select id="getEmpById" resultType="com.example.mybatis.bean.Employee">
> select * from tbl_employee where id = #{id}
> </select>
> 
> <!--parameterType可以省略-->
> <insert id="addEmp" parameterType="com.example.mybatis.bean.Employee">
>   insert into tbl_employee(last_name,email,gender) values (#{lastName},#{email},#{gender})
> </insert>
> 
> <update id="updateEmp">
>   update tbl_employee set last_name=#{lastName},email=#{email},gender=#{gender} where id=#{id}
> </update>
> 
> <delete id="deleteEmpById">
>   delete from tbl_employee where id=#{id}
> </delete>
> </mapper>
> ```
>
> ```java
> package com.example.mybatis.dao;
> import com.example.mybatis.bean.Employee;
> public interface EmployeeMapper {
> public Employee getEmpById(Integer id);
> 
> public Long addEmp(Employee employee);
> 
> public Boolean updateEmp(Employee employee);
> 
> public void deleteEmpById(Integer id);
> }
> ```
>
> ```java
> public SqlSessionFactory getSqlSessionFactory() throws IOException {
> String resource = "mybatis-config.xml";
> InputStream inputStream = Resources.getResourceAsStream(resource);
> return new SqlSessionFactoryBuilder().build(inputStream);
> }
> 
> /**
>      *  测试增删改查
>      *  mybatis允许增删改直接定义以下返回值
>      *      Integer，Long，Boolean
>      *  我们需要手动提交数据  sqlSessionFactory.openSession();  =====> 手动提交
>      *	sqlSessionFactory.openSession(true); =====> 自动提交
>      */
> @Test
> public void test2() throws IOException {
>     SqlSessionFactory sqlSessionFactory = getSqlSessionFactory();
>     SqlSession openSession = sqlSessionFactory.openSession();
> 
>     try {
>         EmployeeMapper mapper = openSession.getMapper(EmployeeMapper.class);
>         // 查询
>         Employee employee = mapper.getEmpById(1);
>         System.out.println(employee);
>         // 添加
>         Long rows = mapper.addEmp(new Employee(null, "jerry", "jerry@gmail.com", "1"));
>         System.out.println(rows);
>         // 修改
>         Boolean flag = mapper.updateEmp(new Employee(3, "mike", "mike@icloud.com", "1"));
>         System.out.println(flag);
>         // 删除
>         mapper.deleteEmpById(2);
>         // 手动提交数据
>         openSession.commit();
>     } finally {
>         openSession.close();
>     }
> }
> ```

### mapper映射文件

##### insert获取自增主键

> ```xml
> <!--parameterType 参数类型，可以省略
> useGeneratedKeys ：使用自增主键，获取主键值策略
> keyProperty：指定对应的主键属性，也就是mybatis获取到主键值以后，将这个值封装到JavaBean指定的属性
> -->
> <insert id="addEmp" parameterType="com.example.mybatis.bean.Employee" useGeneratedKeys="true" keyProperty="id">
> insert into tbl_employee(last_name,email,gender) values (#{lastName},#{email},#{gender})
> </insert>
> ```
>
> ```java
> // 添加
> Employee employee1 = new Employee(null, "jerry", "jerry@gmail.com", "1");
> Long rows = mapper.addEmp(employee1);
> System.out.println(employee1.getId());
> ```

##### 获取非自增主键值

> ```xml
> <insert id="addEmp" parameterType="com.example.mybatis.bean.Employee" databaseId="oracle">
> <!--
> keyProperty：查出的主键值封装给javaBean的哪个属性
> order="BEFORE"：当前sql在插入sql之前运行
>       AFTER：当前sql在插入sql之后运行
> resultType：查出的数据的返回类型
> -->
> <selectKey keyProperty="id" order="BEFORE" resultType="Integer">
>   <!--编写查询主键的sql语句-->
>   select EMPLOYEES_SEQ.nextval from dual
> </selectKey>
> insert into employees(EMPLOYEE_ID,LAST_NAME,EMAIL) values (#{id},#{lastName},#{email})
> </insert>
> ```

##### 参数处理

> **单个参数**
>
> mybatis不会做特殊处理，#{参数名}  ：取出参数值
>
> **多个参数**
>
> mybatis会做特殊处理，多个参数会被封装成一个map
>
> ​		key：param1,param2...param10，或者参数的索引也可以
>
> ​		value：传入的参数值
>
> #{} 就是从map中获取指定的key值
>
> ==命名参数：==明确指定封装参数时map的key
>
> ​	  key：使用@Param注解指定的值
>
> ​	  value：参数值
>
> #{指定的key}取出对应的参数值
>
> ==POJO：==
>
> 如果多个参数正好是我们业务逻辑的数据模型，我们就可以直接传入pojo
>
> #{属性名} ：取出传入的pojo的属性值
>
> ==Map：==
>
> 如果多个参数不是业务模型中的数据，没有对应的pojo，为了方便，我们也可以传入map
>
> #{key} ：取出map中对应的值
>
> ==TO：==
>
> 如果多个参数不是业务模型中的数据，但是经常使用，推荐来编写一个TO（Transfer Object）数据传输对象
>
> ```text
> Page{
> int index;
> int size;
> }
> ```
>
> 
>
> *示例1*
>
> 接口方法定义：
>
> ```java
> public Employee getEmpByIdAndLastName(Integer id,String lastName);
> ```
>
> sql映射文件：
>
> ```xml
> <select id="getEmpByIdAndLastName" resultType="com.example.mybatis.bean.Employee">
> select * from tbl_employee where id = #{param1} and last_name = #{param2}
> </select>
> ```
>
> *示例2*
>
> 接口方法定义：
>
> ```java
> public Employee getEmpByIdAndLastName(@Param("id") Integer id, @Param("lastName") String lastName);
> ```
>
> sql映射文件：
>
> ```xml
> <select id="getEmpByIdAndLastName" resultType="com.example.mybatis.bean.Employee">
> select * from tbl_employee where id = #{id} and last_name = #{lastName}
> </select>
> ```
>
> *示例3*
>
> 接口方法定义：
>
> ```java
> public Long addEmp(Employee employee);
> ```
>
> sql映射文件：
>
> ```xml
> <insert id="addEmp" parameterType="com.example.mybatis.bean.Employee">
> insert into tbl_employee(last_name,email,gender) values (#{lastName},#{email},#{gender})
> </insert>
> ```
>
> *示例4*
>
> 接口方法定义：
>
> ```java
> public Employee getEmpByMap(Map<String,Object> map);
> ```
>
> sql映射文件：
>
> ```xml
> <select id="getEmpByMap" resultType="com.example.mybatis.bean.Employee">
> select * from tbl_employee where id = #{id} and last_name = #{lastName}
> </select>
> ```
>
> 测试：
>
> ```java
> EmployeeMapper mapper = openSession.getMapper(EmployeeMapper.class);
> Map<String,Object> map = new HashMap<>();
> map.put("id", 1);
> map.put("lastName", "tom");
> Employee employee = mapper.getEmpByMap(map);
> System.out.println(employee);
> ```
>
> **参数值的获取**
>
> #{} ：可以获取map的值或者pojo对象属性的值
>
> ${} ：可以获取map的值或者pojo对象属性的值
>
> 区别：#{} ：是以预编译的形式，将参数设置到sql语句中；PreparedStatement；sql语句中是以占位符？的形式替代参数Preparing: select * from tbl_employee where id = ? and last_name = ?
>
> ​			${}：是直接把值拼装到sql语句中，Preparing: select * from tbl_employee where id = 1 and last_name = ?
>
> 小结：大多数情况下，我们取参数的值都应该去使用#{}，原生JDBC不支持使用占位符的地方我们就可以使用${}取值，比如分表，排序，按照年份拆分表
>
> `select * from ${year}_salary where xxx		select * from tbl_employee order by  ${f_name} ${order}`
>
> #{}：更丰富的用法
>
> ​		规定参数的一些规则：
>
> ​		javaType、jdbcType、mode(存储过程)、numericScale、resultMap、typeHandler、jdbcTypeName、expression
>
> ​		jdbcType通常需要在某种特定的条件下被设置：
>
> ​			在我们的数据为null的时候，有些数据库可能不能识别mybatis对null的默认处理，比如Oracle（报错）
>
> ​			insert into employees(EMPLOYEE_ID,LAST_NAME,EMAIL) values (#{id},#{lastName},#{email,jdbcType=NULL})  默认jdbcType=OTHER
>
> ​		由于全局配置中：jdbcTypeForNull=OTHER；oracle不支持
>
> 		1. #{email,jdbcType=OTHER}
> 		2. jdbcTypeForNull=NULL

### 查询

##### select返回List

> 接口方法定义：
>
> ```java
> public List<Employee> getEmpsByLastName(String lastName);
> ```
>
> sql映射文件：
>
> ```xml
> <!--resultType：如果返回的是一个集合，要写集合中元素的类型-->
> <select id="getEmpsByLastName" resultType="com.example.mybatis.bean.Employee">
> select * from tbl_employee where last_name like #{lastName}
> </select>
> ```
>
> 测试文件：
>
> ```java
> EmployeeMapper mapper = openSession.getMapper(EmployeeMapper.class);
> List<Employee> employees = mapper.getEmpsByLastName("%e%");
> for (Employee employee : employees) {
> System.out.println(employee);
> }
> ```

##### select返回map

> 接口方法定义：
>
> ```java
> // 返回一条记录的map；key就是列名，值就是对应的值
> public Map<String,Object> getEmpByIdReturnMap(Integer id);
> ```
>
> sql映射文件：
>
> ```xml
> <select id="getEmpByIdReturnMap" resultType="map">
> select * from tbl_employee where id=#{id};
> </select>
> ```
>
> 测试文件：
>
> ```java
> EmployeeMapper mapper = openSession.getMapper(EmployeeMapper.class);
> Map<String, Object> map = mapper.getEmpByIdReturnMap(1);
> System.out.println(map);
> ```
>
> **多条记录封装成一个map**
>
> 接口方法定义：
>
> ```java
> // 多条记录封装成一个map：Map<Integer,Employee>；key是这条记录的主键，值是记录封装后的javaBean
> // 告诉mybatis封装这个map的时候使用哪个属性作为map的key
> @MapKey("id")
> public Map<Integer,Employee> getEmpByLastNameLikeReturnMap(String lastName);
> ```
>
> sql映射文件：
>
> ```xml
> <select id="getEmpByLastNameLikeReturnMap" resultType="com.example.mybatis.bean.Employee">
> select * from tbl_employee where last_name like #{lastName};
> </select>
> ```
>
> 测试文件：
>
> ```java
> EmployeeMapper mapper = openSession.getMapper(EmployeeMapper.class);
> Map<Integer, Employee> maps = mapper.getEmpByLastNameLikeReturnMap("%e%");
> System.out.println(maps);
> ```

##### resultMap

> **自定义结果映射规则**
>
> *示例1*
>
> 接口方法定义：
>
> ```java
> package com.example.mybatis.dao;
> import com.example.mybatis.bean.Employee;
> public interface EmployeeMapperPlus {
> public Employee getEmpById(Integer id);
> }
> ```
>
> sql映射文件：
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE mapper
>   PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
>   "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
> <mapper namespace="com.example.mybatis.dao.EmployeeMapperPlus">
> 
> <!--自定义某个javaBean封装规则
>   type：自定义规则的java类型
>   id：唯一id，方便引用
> -->
> <resultMap id="myEmp" type="com.example.mybatis.bean.Employee">
>   <!--指定主键列的封装规则
>       id定义主键底层会有优化
>       column：指定哪一列对应
>       property：指定对应的javaBean属性
>   -->
>   <id column="id" property="id" />
>   <!--定义普通列封装规则-->
>   <result column="last_name" property="lastName"/>
>   <!--其他不指定的列会自动封装，但是推荐写resultMap就把全部的列都写上-->
>   <result column="email" property="email"/>
>   <result column="gender" property="gender"/>
> </resultMap>
> <!--resultMap：自定义结果集映射-->
> <select id="getEmpById" resultMap="myEmp">
>   select * from tbl_employee where id=#{id}
> </select>
> </mapper>
> ```
>
> 将sql映射文件注册到全局配置文件中
>
> ```xml
> <mappers>
> <mapper resource="EmployeeMapper.xml"/>
> <mapper resource="EmployeeMapperPlus.xml"/>
> <!--批量注册，接口名和映射文件名必须相同，并且在同一目录下，或者可以使用注解-->
> </mappers>
> ```
>
> **关联查询**
>
> *级联属性封装结果集*
>
> 类定义：
>
> ```java
> package com.example.mybatis.bean;
> public class Employee {
> private Integer id;
> private String lastName;
> private String email;
> private String gender;
> private Department dept;   
> 	......
> }
> ```
>
> ```java
> package com.example.mybatis.bean;
> public class Department {
> private Integer id;
> private String departmentName;
> 	......
> }
> ```
>
> sql映射文件：
>
> ```xml
> <!--联合查询：级联属性封装结果集-->
> <resultMap id="MyComplexEmp" type="com.example.mybatis.bean.Employee">
> <id column="id" property="id"/>
> <result column="last_name" property="lastName"/>
> <result column="gender" property="gender"/>
> <result column="did" property="dept.id"/>
> <result column="dept_name" property="dept.departmentName"/>
> </resultMap>
> <select id="getEmpAndDept" resultMap="MyComplexEmp">
> select e.id id,e.last_name last_name,e.gender gender,e.dept_id dept_id,d.id did,d.dept_name dept_name
> from tbl_employee e,tbl_dept d where e.dept_id = d.id and e.id = 1
> </select>
> ```
>
> 测试文件：
>
> ```java
> EmployeeMapperPlus mapper = openSession.getMapper(EmployeeMapperPlus.class);
> Employee employee = mapper.getEmpAndDept(1);
> System.out.println(employee);
> ```
>
> 表结构：
>
> tbl_employee表
>
> ![](./images/tbl_employee.jpg)
>
> tbl_dept表
>
> ![](./images/tbl_dept.jpg)
>
> **association定义关联对象封装规则**
>
> 接口方法定义
>
> ```java
> public Employee getEmpAndDept(Integer id);
> ```
>
> sql映射文件定义
>
> ```xml
> <resultMap id="MyComplexEmp2" type="com.example.mybatis.bean.Employee">
> <id column="id" property="id"/>
> <result column="last_name" property="lastName"/>
> <result column="gender" property="gender"/>
> <!--association可以指定联合的JavaBean
>       property="dept"：指定哪个属性是联合的对象
>       javaType：指定这个属性对象的类型【不能省略】
>   -->
> <association property="dept" javaType="com.example.mybatis.bean.Department">
>   <id column="did" property="id"/>
>   <result column="dept_name" property="departmentName"/>
> </association>
> </resultMap>
> <select id="getEmpAndDept" resultMap="MyComplexEmp2">
> select e.id id,e.last_name last_name,e.gender gender,e.dept_id dept_id,d.id did,d.dept_name dept_name
> from tbl_employee e,tbl_dept d where e.dept_id = d.id and e.id = 1
> </select>
> ```
>
> 其他文件跟上面的例子一样
>
> **使用association进行分步查询**
>
> 接口方法定义
>
> EmployeeMapperPlus.java
>
> ```java
> public Employee getEmpByIdStep(Integer id);
> ```
>
> DepartmentMapper.java
>
> ```java
> package com.example.mybatis.dao;
> import com.example.mybatis.bean.Department;
> public interface DepartmentMapper {
> public Department getDeptById(Integer id);
> }
> ```
>
> sql映射文件
>
> EmployeeMapperPlus.xml
>
> ```xml
> <!--使用association分步查询
>   1.先按照员工id查出员工信息
>   2.根据查询员工信息中的dept_id值，去部门表查出部门信息
>   3.部门设置到员工中
> -->
> <resultMap id="MyEmpByStep" type="com.example.mybatis.bean.Employee">
> <id column="id" property="id"/>
> <result column="last_name" property="lastName"/>
> <result column="email" property="email"/>
> <result column="gender" property="gender"/>
> <!--association定义关联对象的封装规则
>       select：表明当前属性是调用select指定的方法查出的结果
>       column：指定将哪一列的值传给这个方法
>       流程：使用select指定的方法（传入column指定的这列参数的值）查出对象，并封装给property指定的属性
>   -->
> <association property="dept" select="com.example.mybatis.dao.DepartmentMapper.getDeptById" column="dept_id"></association>
> </resultMap>
> <select id="getEmpByIdStep" resultMap="MyEmpByStep">
> select * from tbl_employee where id=#{id}
> </select>
> ```
>
> DepartmentMapper.xml
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE mapper
>   PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
>   "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
> <mapper namespace="com.example.mybatis.dao.DepartmentMapper">
> <select id="getDeptById" resultType="com.example.mybatis.bean.Department">
>   select id,dept_name departmentName from tbl_dept where id=#{id}
> </select>
> </mapper>
> ```
>
> **延迟加载**
>
> 在分步查询的基础之上，在全局配置上添加以下配置，关联属性会延迟加载
>
> * lazyLoadingEnabled：延迟加载的全局开关。当开启时，所有关联对象都会延迟加载。 特定关联关系中可通过设置 `fetchType` 属性来覆盖该项的开关状态。
> * aggressiveLazyLoading：开启时，任一方法的调用都会加载该对象的所有延迟加载属性。 否则，每个延迟加载属性会按需加载（参考 lazyLoadTriggerMethods)。
>
> **collection定义关联集合封装规则**
>
> 类定义
>
> ```java
> package com.example.mybatis.bean;
> import java.util.List;
> public class Department {
>     private Integer id;
>     private String departmentName;
>     private List<Employee> emps;
>     ......
> }
> ```
>
> ```java
> package com.example.mybatis.bean;
> public class Employee {
>     private Integer id;
>     private String lastName;
>     private String email;
>     private String gender;
>     private Department dept;
>     ......
> }
> ```
>
> 接口方法定义
>
> ```java
> public Department getDeptByIdPlus(Integer id);
> ```
>
> sql映射文件
>
> ```xml
> <resultMap id="MyDept" type="com.example.mybatis.bean.Department">
>     <id column="did" property="id"/>
>     <result column="dept_name" property="departmentName"/>
>     <!--collection嵌套结果集的方式，定义关联的集合类型元素的封装规则
>             collection：定义关联的集合类型的属性的封装
>                 ofType：指定集合里面元素的类型
>         -->
>     <collection property="emps" ofType="com.example.mybatis.bean.Employee">
>         <!--定义这个集合中元素的封装规则-->
>         <id column="eid" property="id"/>
>         <result column="last_name" property="lastName"/>
>         <result column="email" property="email"/>
>         <result column="gender" property="gender" />
>     </collection>
> </resultMap>
> <select id="getDeptByIdPlus" resultMap="MyDept">
>     select d.id did,d.dept_name dept_name,e.id eid,e.last_name last_name,e.email email,e.gender gender
>     from tbl_dept d left join tbl_employee e on d.id=e.dept_id where d.id=#{id}
> </select>
> ```
>
> 测试文件
>
> ```java
> DepartmentMapper mapper = openSession.getMapper(DepartmentMapper.class);
> Department dept = mapper.getDeptByIdPlus(1);
> System.out.println(dept);
> ```
>
> **collection分步查询**
>
> 类文件定义和上面的上面的collection定义关联结果一样
>
> 接口方法定义
>
> EmployeeMapperPlus.java
>
> ```java
> public List<Employee> getEmpsByDeptId(Integer deptId);
> ```
>
> DepartmentMapper.java
>
> ```java
> public Department getDeptByIdStep(Integer id);
> ```
>
> sql映射文件
>
> DepartmentMapper.xml
>
> ```xml
> <resultMap id="MyDeptStep" type="com.example.mybatis.bean.Department">
>     <id column="id" property="id"/>
>     <result column="dept_name" property="departmentName"/>
>     <collection property="emps" select="com.example.mybatis.dao.EmployeeMapperPlus.getEmpsByDeptId" column="id" 			fetchType="lazy"></collection>
> </resultMap>
> <select id="getDeptByIdStep" resultMap="MyDeptStep">
>     select * from tbl_dept where id=#{id}
> </select>
> ```
>
> EmployeeMapperPlus.xml
>
> ```xml
> <select id="getEmpsByDeptId" resultType="com.example.mybatis.bean.Employee">
>     select * from tbl_employee where dept_id=#{deptId}
> </select>
> ```
>
> 测试文件
>
> ``` java
> DepartmentMapper mapper = openSession.getMapper(DepartmentMapper.class);
> Department dept = mapper.getDeptByIdStep(1);
> System.out.println(dept);
> ```
>
> 扩展：多列的值传递过去，将多列的值封装map传递
>
> column="{key1=column1,key2=column2}"
>
> fetchType="lazy" ：表示使用延迟加载
>
> **discriminator鉴别器**
>
> 鉴别器：mybatis可以使用discriminator判断某列的值，然后根据某列的值改变封装行为
>
> sql映射文件
>
> ```xml
> <!--封装Employee:
>         如果查出的是女生：就把部门信息查询出来，否则不查询
>         如果是男生，就把last_name这一列的值赋值给email-->
> <resultMap id="MyEmpDis" type="com.example.mybatis.bean.Employee">
>     <id column="id" property="id"/>
>     <result column="last_name" property="lastName"/>
>     <result column="email" property="email"/>
>     <result column="gender" property="gender"/>
>     <!--column：指定要判断的列名
>             javaType：列值对应的java类型
>         -->
>     <discriminator javaType="string" column="gender">
>         <!--女生 resultType：指定封装的结果类型，不能省略，或者可以用resultMap-->
>         <case value="0" resultType="com.example.mybatis.bean.Employee">
>             <association property="dept" select="com.example.mybatis.dao.DepartmentMapper.getDeptById" column="dept_id">			</association>
>         </case>
>         <!--男生-->
>         <case value="1" resultType="com.example.mybatis.bean.Employee">
>             <id column="id" property="id"/>
>             <result column="last_name" property="lastName"/>
>             <result column="last_name" property="email"/>
>             <result column="gender" property="gender"/>
>         </case>
>     </discriminator>
> </resultMap>
> 
> <select id="getEmpByIdStep" resultMap="MyEmpDis">
>     select * from tbl_employee where id=#{id}
> </select>
> ```

### 动态SQL

##### if判断

> 接口方法定义
>
> ```java
> package com.example.mybatis.dao;
> import com.example.mybatis.bean.Employee;
> import java.util.List;
> public interface EmployeeMapperDynamicSQL {
> public List<Employee> getEmpsByConditionIf(Employee employee);
> }
> ```
>
> sql映射文件
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE mapper
>   PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
>   "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
> <mapper namespace="com.example.mybatis.dao.EmployeeMapperDynamicSQL">
> <!--查询员工，要求携带了哪个字段查询条件就带上这个字段的值-->
> <select id="getEmpsByConditionIf" resultType="com.example.mybatis.bean.Employee">
>   select * from tbl_employee where
>   <!--test：判断表达式（OGNL）
>       从参数中取值进行判断
>       遇见特殊符号应该写转义字符
>   -->
>   <if test="id!=null">
>       id=#{id}
>   </if>
>   <if test="lastName!=null and lastName!='' and id==null">
>       last_name like #{lastName}
>   </if>
>   <if test="lastName!=null and lastName!='' and id!=null">
>       and last_name like #{lastName}
>   </if>
>   <if test="email!=null and email.trim()!=&quot;&quot;">
>       and email=#{email}
>   </if>
>   <!--OGNL会进行字符串与数字的转换判断-->
>   <if test="gender==0 or gender==1">
>       and gender=#{gender}
>   </if>
> </select>
> </mapper>
> ```
>
> 测试文件
>
> ```java
> public SqlSessionFactory getSqlSessionFactory() throws IOException {
>  String resource = "mybatis-config.xml";
>  InputStream inputStream = Resources.getResourceAsStream(resource);
>  return new SqlSessionFactoryBuilder().build(inputStream);
> }
> 
> @Test
> public void test5() throws IOException {
>  SqlSessionFactory sqlSessionFactory = getSqlSessionFactory();
>  SqlSession openSession = sqlSessionFactory.openSession();
> 
>  try {
>      EmployeeMapperDynamicSQL mapper = openSession.getMapper(EmployeeMapperDynamicSQL.class);
>      Employee employee = new Employee(null, "%e%", "jerry@gmail.com", null);
>      List<Employee> emps = mapper.getEmpsByConditionIf(employee);
>      System.out.println(emps);
>      // 手动提交数据
>      openSession.commit();
>  } finally {
>      openSession.close();
>  }
> }
> ```

##### where查询条件

> 查询的时候如果某些条件没带可能sql拼装会有问题，解决办法如下：
>
> 1. 给where后面加上1=1，以后的条件都and xxx.
> 2. mybatis使用where标签来将所有的查询条件包括在内，mybatis就会将where标签中拼装的sql，多出来的and或者or去掉
> 3. where标签只会去掉第一个多出来的and或者or
>
> ```xml
> <?xml version="1.0" encoding="UTF-8" ?>
> <!DOCTYPE mapper
>         PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
>         "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
> <mapper namespace="com.example.mybatis.dao.EmployeeMapperDynamicSQL">
> 
>     <!--查询员工，要求携带了哪个字段查询条件就带上这个字段的值-->
>     <select id="getEmpsByConditionIf" resultType="com.example.mybatis.bean.Employee">
>         select * from tbl_employee
>         <where>
>             <!--test：判断表达式（OGNL）
>                 从参数中取值进行判断
>                 遇见特殊符号应该写转义字符
>             -->
>             <if test="id!=null">
>                 id=#{id}
>             </if>
>             <if test="lastName!=null and lastName!=''">
>                 and last_name like #{lastName}
>             </if>
>             <if test="email!=null and email.trim()!=&quot;&quot;">
>                 and email=#{email}
>             </if>
>             <!--OGNL会进行字符串与数字的转换判断-->
>             <if test="gender==0 or gender==1">
>                 and gender=#{gender}
>             </if>
>         </where>
>     </select>
> </mapper>
> ```

##### trim自定义字符串截取

> 接口方法定义
>
> ```java
> public List<Employee> getEmpsByConditionTrim(Employee employee);
> ```
>
> sql映射文件
>
> ```xml
> <select id="getEmpsByConditionTrim" resultType="com.example.mybatis.bean.Employee">
>  select * from tbl_employee
>  <!--后面多出的and或者or where标签不能解决
>          prefix：前缀，trim标签体中是整个字符串拼串后的结果，prefix就是给拼串后的字符串加一个前缀
>          prefixOverrides：前缀覆盖，去掉整个字符串前面多余的字符
>          suffix：给整个拼串后的整个字符串加后缀
>          suffixOverrides：后缀覆盖，去掉整个字符串后面多余的字符
>      -->
>  <trim prefix="where" prefixOverrides="" suffix="" suffixOverrides="and">
>      <if test="id!=null">
>          id=#{id} and
>      </if>
>      <if test="lastName!=null and lastName!=''">
>          last_name like #{lastName} and
>      </if>
>      <if test="email!=null and email.trim()!=&quot;&quot;">
>          email=#{email} and
>      </if>
>      <!--OGNL会进行字符串与数字的转换判断-->
>      <if test="gender==0 or gender==1">
>          gender=#{gender}
>      </if>
>  </trim>
> </select>
> ```

##### choose分支选择

> 接口方法定义
>
> ```java
> public List<Employee> getEmpsByConditionChoose(Employee employee);
> ```
>
> sql映射文件
>
> ```xml
> <select id="getEmpsByConditionChoose" resultType="com.example.mybatis.bean.Employee">
>  select * from tbl_employee
>  <where>
>      <!--如果带了id，就用id查，如果带了lastName就用lastName查，只会进入其中一个-->
>      <choose>
>          <when test="id!=null">
>              id=#{id}
>          </when>
>          <when test="lastName!=null">
>              last_name like #{lastName}
>          </when>
>          <when test="email!=null">
>              email=#{email}
>          </when>
>          <otherwise>
>              gender=0
>          </otherwise>
>      </choose>
>  </where>
> </select>
> ```

##### 动态set与if结合

> 接口方法定义
>
> ```java
> public void updateEmp(Employee employee);
> ```
>
> sql映射文件
>
> ```xml
> <update id="updateEmp">
>  update tbl_employee
>  <set>
>      <if test="lastName!=null">
>          last_name=#{lastName},
>      </if>
>      <if test="email!=null">
>          email=#{email},
>      </if>
>      <if test="gender!=null">
>          gender=#{gender}
>      </if>
>  </set>
>  where id=#{id}
> </update>
> ```

##### foreach遍历集合

> 接口方法定义
>
> ```java
> public List<Employee> getEmpsByConditionForeach(@Param("ids")List<Integer> ids);
> ```
>
> sql映射文件
>
> ```xml
> <select id="getEmpsByConditionForeach" resultType="com.example.mybatis.bean.Employee">
>  select * from tbl_employee where id in
>  <!--
>          collection：指定要遍历的集合
>              list类型的参数会特殊处理封装在map中，map的key就叫list
>              item：将当前遍历出的元素赋值给指定的变量
>              separator：每个元素之间的分隔符
>              open：遍历出所有结果，拼接一个开始字符
>              close：遍历出所有结果，拼接一个结束字符
>              index：索引，遍历list的时候是索引，遍历map的时候是map的key
>              #{变量名}：就能取出变量的值，也就是当前遍历出的元素
>      -->
>  <foreach collection="ids" item="item_id" separator="," open="(" close=")" index="">
>      #{item_id}
>  </foreach>
> </select>
> ```
>
> 测试文件
>
> ```java
> EmployeeMapperDynamicSQL mapper = openSession.getMapper(EmployeeMapperDynamicSQL.class);
> List<Employee> list = mapper.getEmpsByConditionForeach(Arrays.asList(1, 2, 3, 4));
> System.out.println(list);
> ```
>
> **批量插入**
>
> 接口方法定义
>
> ```java
> public void addEmps(@Param("emps") List<Employee> emps);
> ```
>
> sql映射文件
>
> ```xml
> <insert id="addEmps">
>  insert into tbl_employee(last_name,email,gender,dept_id) values
>  <foreach collection="emps" item="emp" separator=",">
>      (#{emp.lastName},#{emp.email},#{emp.gender},#{emp.dept.id})
>  </foreach>
> </insert>
> ```
>
> 测试文件
>
> ```java
> EmployeeMapperDynamicSQL mapper = openSession.getMapper(EmployeeMapperDynamicSQL.class);
> List<Employee> emps = new ArrayList<>();
> emps.add(new Employee(null,"smith","smith@icloud.com","1",new Department(1)));
> emps.add(new Employee(null,"lucy","lucy@facebook.com","0",new Department(2)));
> mapper.addEmps(emps);
> ```

##### 内置参数

> mybatis默认还有两个内置参数：
>
> * _parameter：代表整个参数
>   * 单个参数：_parameter就是这个参数
>   * 多个参数：参数会被封装为一个map；_parameter就是代表这个map
> * _databaseId：如果配置了DatabaseIdProvider标签
>   * databaseId就是代表当前数据库的别名
>
> 接口方法定义
>
> ```java
> public List<Employee> getEmpsTestInnerParameter(Employee employee);
> ```
>
> sql映射文件
>
> ```xml
> <select id="getEmpsTestInnerParameter" resultType="com.example.mybatis.bean.Employee">
>     <if test="_databaseId=='mysql'">
>         select * from tbl_employee
>         <if test="_parameter!=null">
>             where last_name=#{_parameter.lastName}
>         </if>
>     </if>
>     <if test="_databaseId=='oracle'">
>         select * from employees
>     </if>
> </select>
> ```

##### bind

> 可以将OGNL表达式的值绑定到一个变量中，方便后来引用这个变量的值
>
> ```xml
> <select id="getEmpsTestInnerParameter" resultType="com.example.mybatis.bean.Employee">
>  <bind name="_lastName" value="'%'+lastName+'%'"></bind>
>  <if test="_databaseId=='mysql'">
>      select * from tbl_employee
>      <if test="_parameter!=null">
>          where last_name like #{_lastName}
>      </if>
>  </if>
>  <if test="_databaseId=='oracle'">
>      select * from employees
>  </if>
> </select>
> ```

##### 抽取可重用sql片段

> ```xml
> <!--抽取可重用的sql片段，方便后面引用-->
> <sql id="insertColumn">
>  <if test="_databaseId=='oracle'">employee_id,last_name,email </if>
>  <if test="_databaseId=='mysql'">last_name,email,gender,dept_id </if>
> </sql>
> <insert id="addEmps">
>  insert into tbl_employee(
>  <include refid="insertColumn"/>
>  )values
>  <foreach separator="," item="emp" collection="emps">(#{emp.lastName},#{emp.email},#{emp.gender},#{emp.dept.id}) </foreach>
> </insert>
> ```

### 缓存

##### 一级缓存

> 一级缓存即本地缓存：sqlSession级别的缓存，一级缓存是一直开启的，sqlSession级别的一个map
>
> 与数据库同一次会话期间查询到的数据会放在本地缓存中，以后如果需要获取相同的数据，直接从缓存中拿，没必要再去查询数据库
>
> 一级缓存失效情况（没有使用到当前一级缓存的情况，效果就是还需要再向数据库发出查询）
>
> 一级缓存失效的四种情况：
>
> 1. sqlSession不同
> 2. sqlSession相同，查询条件不同（当前缓存中还没有这个数据）
> 3. sqlSession相同，两次查询之间执行了增删改操作
> 4. sqlSession相同，手动清除了缓存openSession.clearCache()

##### 二级缓存

> 二级缓存即全局缓存：基于namespace级别的缓存，一个namespace对应一个二级缓存
>
> 工作机制：
>
> 1. 一个会话，查询一条数据，这个数据就会被放在当前会话的一级缓存中
> 2. 如果会话关闭，一级缓存中的数据会被保存到二级缓存中；新的会话查询信息，就可以参照二级缓存中内容
> 3. 不同namespace查出的数据会放在自己对应的缓存中(map)
> 4. 只有会话提交，或者关闭以后，一级缓存中的数据才会转移到二级缓存中
>
> **使用：**
>
> 1. 开启全局二级缓存配置：<setting name="cacheEnabled" value="true"/>
> 2. 去mapper.xml中配置使用二级缓存
> 3. 我们的POJO需要实现序列化接口(Serializable)
>
> 在全局配置文件中开启缓存配置
>
> ```java
> <settings>
>     <setting name="cacheEnabled" value="true"/>
> </settings>
> ```
>
> sql映射文件配置
>
> ```xml
> <!--
> eviction：缓存的回收策略：
>     • LRU – 最近最少使用的：移除最长时间不被使用的对象。
>     • FIFO – 先进先出：按对象进入缓存的顺序来移除它们。
>     • SOFT – 软引用：移除基于垃圾回收器状态和软引用规则的对象。
>     • WEAK – 弱引用：更积极地移除基于垃圾收集器状态和弱引用规则的对象。
>     • 默认的是 LRU。
> flushInterval：缓存刷新间隔，缓存多长时间清空一次，设置一个毫秒值
> readOnly：是否只读
>     true：只读，mybatis认为所有从缓存中获取数据的操作都是只读操作，不会修改数据
>           mybatis为了加快获取速度，直接就会将数据在缓存中的引用交给用户，不安全，速度快
>     false：非只读，mybatis觉得获取的数据可能会被修改
>            mybatis会利用序列化和反序列化的技术复制一份新的数据
>     size：缓存存放多少元素
>     type：指定自定义缓存的全类名，实现Cache接口即可
> -->
> <cache eviction="LRU" flushInterval="6000" readOnly="false" size="2048"></cache>
> ```
>
> 测试文件
>
> ```java
> @Test
> public void test6() throws IOException {
>     SqlSessionFactory sqlSessionFactory = getSqlSessionFactory();
>     SqlSession openSession = sqlSessionFactory.openSession();
>     SqlSession openSession1 = sqlSessionFactory.openSession();
>     try {
>         EmployeeMapper mapper = openSession.getMapper(EmployeeMapper.class);
>         EmployeeMapper mapper1 = openSession1.getMapper(EmployeeMapper.class);
>         Employee emp = mapper.getEmpById(1);
>         System.out.println("========");
>         System.out.println(emp);
>         openSession.close();
>         // 第二次查询是从二级缓存中拿到的数据，并没有发送新的sql【只有第一个会话关闭，才能从二级缓存中取到第一次会话的缓存】
>         Employee emp1 = mapper1.getEmpById(1);
>         System.out.println("=====11111=====");
>         System.out.println(emp1);
>         openSession1.close();
>     }finally {
>     }
> }
> ```

##### 缓存设置

> 1. cacheEnabled=true；false：关闭缓存（二级缓存关闭）一级缓存一直可用
>
> 2.  每个select标签都有useCache="true";false，控制的是二级缓存的关闭与开启
>
> 3. 每个增删改标签都有flushCache="true"  控制是否清除缓存（一级缓存和二级缓存都会被清空）
>
> 4. sqlSession.clearCache只是清除当前session的一级缓存
>
> 5. ```xml
>    <!--引用缓存：指定和哪个名称空间下的缓存一样-->
>    <cache-ref namespace="com.example.mybatis.dao.EmployeeMapper"></cache-ref>
>    ```

##### 第三方缓存整合

> 1. 导入第三方缓存包
> 2. 导入与第三方缓存的适配包，官方有
> 3. mapper.xml中使用自定义缓存
>
> ```xml
> <cache type="org.mybatis.caches.ehcache.EhcacheCache"></cache>
> ```
>
> 

### 插件

##### 插件原理

> 在四大对象创建的时候
>
> 1. 每个创建出来的对象不是直接返回的，而是interceptorChain.pluginAll(parameterHandler)；
> 2. 获取到所有的Interceptor（拦截器）（插件需要实现的接口）；调用interceptor.plugin(target)；返回target 包装后的对象
> 3. 插件机制：我们可以使用插件为目标对象创建一个代理对象；AOP（面向切面），我们的插件可以为四大对象创建出代理对象；代理对象就可以拦截到四大对象的每一个执行

##### 插件编写

> 1. 编写Interceptor的实现类
> 2. 使用@Intercepts注解完成插件签名
> 3. 将写好的插件注册到全局配置文件中
>
> ```java
> package com.example.mybatis.dao;
> import org.apache.ibatis.executor.statement.StatementHandler;
> import org.apache.ibatis.plugin.*;
> import java.util.Properties;
> /**
>  *  完成插件签名：告诉mybatis当前插件用来拦截哪个对象的哪个方法
>  */
> @Intercepts({@Signature(type = StatementHandler.class,method = "parameterize",args = java.sql.Statement.class)})
> public class MyFirstPlugin implements Interceptor {
>     // 拦截目标对象目标方法的执行
>     @Override
>     public Object intercept(Invocation invocation) throws Throwable {
> 
>         System.out.println("intercept 拦截 : " + invocation.getMethod());
>         // 执行目标方法
>         Object proceed = invocation.proceed();
> 
>         // 返回执行后的返回值
>         return proceed;
>     }
>     // 包装目标对象，为目标对象创建一个代理对象
>     @Override
>     public Object plugin(Object target) {
> 
>         System.out.println("MyFirstPlugin...plugin：mybatis将要包装的对象" + target);
>         // 我们可以借助Plugin的wrap方法来使用当前Interceptor包装我们的目标对象
>         Object wrap = Plugin.wrap(target, this);
>         // 返回为当前target创建的动态代理
>         return wrap;
>     }
>     // 将插件注册时的property属性设置进来
>     @Override
>     public void setProperties(Properties properties) {
>         System.out.println("插件配置的信息：" + properties);
>     }
> }
> ```
>
> 在全局配置类中注册
>
> ```xml
> <!--插件注册-->
> <plugins>
>     <plugin interceptor="com.example.mybatis.dao.MyFirstPlugin">
>         <!--这里传的property会传到setProperties方法的properties参数中-->
>         <property name="username" value="root" />
>         <property name="password" value="123456" />
>     </plugin>
> </plugins>
> ```
>
> **动态改变sql语句参数**
>
> ```java
> package com.example.mybatis.dao;
> import org.apache.ibatis.executor.statement.StatementHandler;
> import org.apache.ibatis.plugin.*;
> import org.apache.ibatis.reflection.MetaObject;
> import org.apache.ibatis.reflection.SystemMetaObject;
> import java.util.Properties;
> 
> /**
>  *  完成插件签名：告诉mybatis当前插件用来拦截哪个对象的哪个方法
>  */
> @Intercepts({@Signature(type = StatementHandler.class,method = "parameterize",args = java.sql.Statement.class)})
> public class MyFirstPlugin implements Interceptor {
> 
>     // 拦截目标对象目标方法的执行
>     @Override
>     public Object intercept(Invocation invocation) throws Throwable {
> 
>         // 动态改变sql运行的参数，从查询1号员工改为查询8号员工
>         Object target = invocation.getTarget();
>         // 拿到target的元数据
>         MetaObject metaObject = SystemMetaObject.forObject(target);
>         Object value = metaObject.getValue("parameterHandler.parameterObject");
>         System.out.println("sql语句用到的参数是：" + value);
> 
>         // 修改sql语句要用到的参数
>         metaObject.setValue("parameterHandler.parameterObject", 8);
>         // 执行目标方法
>         Object proceed = invocation.proceed();
> 
>         // 返回执行后的返回值
>         return proceed;
>     }
> 
>     // 包装目标对象，为目标对象创建一个代理对象
>     @Override
>     public Object plugin(Object target) {
> 
>         System.out.println("MyFirstPlugin...plugin：mybatis将要包装的对象" + target);
>         // 我们可以借助Plugin的wrap方法来使用当前Interceptor包装我们的目标对象
>         Object wrap = Plugin.wrap(target, this);
>         // 返回为当前target创建的动态代理
>         return wrap;
>     }
> 
>     // 将插件注册时的property属性设置进来
>     @Override
>     public void setProperties(Properties properties) {
>         System.out.println("插件配置的信息：" + properties);
>     }
> }
> ```

##### 批量操作

> ```java
> @Test
> public void testBatch() throws IOException{ 
>     SqlSessionFactory sqlSessionFactory = getSqlSessionFactory();
>     //可以执行批量操作的sqlSession
>     SqlSession openSession = sqlSessionFactory.openSession(ExecutorType.BATCH);
>     long start = System.currentTimeMillis();
>     try{
>         EmployeeMapper mapper = openSession.getMapper(EmployeeMapper.class);
>         for (int i = 0; i < 20; i++) {
>             mapper.addEmp(new Employee(UUID.randomUUID().toString().substring(0, 5), "b", "1","1","1"));
>         }
>         openSession.commit();
>         long end = System.currentTimeMillis();
>         //批量：（预编译sql一次==>设置参数===>10000次===>执行（1次））
>         //Parameters: 616c1(String), b(String), 1(String)==>4598
>         //非批量：（预编译sql=设置参数=执行）==》10000    10200
>         System.out.println("执行时长："+(end-start));
>     }finally{
>         openSession.close();
>     }
> }
> ```

### 类型处理

##### 枚举类型处理

> mybatis在处理枚举类型对象的时候默认保存的是枚举的名字
>
> 配置枚举类型处理器
>
> ```xml
> <typeHandlers>
>     <typeHandler handler="org.apache.ibatis.type.EnumOrdinalTypeHandler" javaType="com.example.mybatis.bean.EmpStatus"/>
> </typeHandlers>
> ```
>
> 这里表示EmpStatus类型的数据由EnumOrdinalTypeHandler来进行处理

##### 自定义类型处理器

> 实现类型处理器接口
>
> ```java
> package com.example.mybatis.typehandler;
> import com.example.mybatis.bean.EmpStatus;
> import org.apache.ibatis.type.JdbcType;
> import org.apache.ibatis.type.TypeHandler;
> import java.sql.CallableStatement;
> import java.sql.PreparedStatement;
> import java.sql.ResultSet;
> import java.sql.SQLException;
> 
> public class MyEnumEmpStatusTypeHandler implements TypeHandler<EmpStatus> {
> 
>     /**
>      * 定义当前数据如何保存到数据库中
>      *
>      */
>     @Override
>     public void setParameter(PreparedStatement preparedStatement, int i, EmpStatus empStatus, JdbcType jdbcType) throws SQLException {
>         preparedStatement.setString(i,empStatus.getCode().toString());
>     }
> 
>     @Override
>     public EmpStatus getResult(ResultSet resultSet, String s) throws SQLException {
>         // 需要根据从数据库中拿到的枚举的状态码，返回一个枚举对象
>         // 按列名拿数据
>         int code = resultSet.getInt(s);
>         return EmpStatus.getEmpStatusByCode(code);
>     }
> 
>     @Override
>     public EmpStatus getResult(ResultSet resultSet, int i) throws SQLException {
>         // 按列的索引拿数据
>         int code = resultSet.getInt(i);
>         return EmpStatus.getEmpStatusByCode(code);
>     }
> 
>     @Override
>     public EmpStatus getResult(CallableStatement callableStatement, int i) throws SQLException {
>         // 存储过程拿数据
>         int code = callableStatement.getInt(i);
>         return EmpStatus.getEmpStatusByCode(code);
>     }
> }
> ```
>
> 在全局配置文件中注册类型处理器
>
> ```xml
> <typeHandlers>
>     <!--配置自定义类型处理器-->
>     <typeHandler handler="com.example.mybatis.typehandler.MyEnumEmpStatusTypeHandler" 						javaType="com.example.mybatis.bean.EmpStatus"/>
> </typeHandlers>
> ```
>
> 自定义类型
>
> ```java
> package com.example.mybatis.bean;
> public enum EmpStatus {
>     LOGIN(100,"用户登录"),LOGOUT(200,"用户登出"),REMOVE(300,"用户不存在");
>     private Integer code;
>     private String msg;
> 
>     EmpStatus(Integer code, String msg) {
>         this.code = code;
>         this.msg = msg;
>     }
> 
>     public Integer getCode() {
>         return code;
>     }
> 
>     public void setCode(Integer code) {
>         this.code = code;
>     }
> 
>     public String getMsg() {
>         return msg;
>     }
> 
>     public void setMsg(String msg) {
>         this.msg = msg;
>     }
> 
>     // 按照状态码返回枚举对象
>     public static EmpStatus getEmpStatusByCode(Integer code){
>         switch (code){
>             case 100 :
>                 return LOGIN;
>             case 200 :
>                 return LOGOUT;
>             case 300 :
>                 return REMOVE;
>             default:
>                 return LOGOUT;
>         }
>     }
> }
> ```
>
> 除了配置全局类型处理器，也可以在处理某个字段的时候告诉mybatis使用什么类型处理器
>
> ```xml
> <insert id="addEmp" parameterType="com.example.mybatis.bean.Employee">
>     insert into tbl_employee(id,last_name,email,gender,empStatus) values
>     (#{id},#{lastName},#{email},#{gender},#{empStatus,typeHandler=XXXX})
> </insert>
> ```
>
> 查询设置，可以自定义resultMap
>
> ```xml
> <resultMap id="MyEmp" type="com.example.mybatis.bean.Employee">
>     <id column="id" property="id"/>
>     <result column="empStatus" property="empStatus" typeHandler="com.example.mybatis.typehandler.MyEnumEmpStatusTypeHandler"/>
> </resultMap>
> ```
>
> 

