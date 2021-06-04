# Swagger学习笔记

### 与SpringBoot集成

> 1. **导入依赖**
>
> ```java
> <dependency>
>     <groupId>io.springfox</groupId>
>     <artifactId>springfox-swagger2</artifactId>
>     <version>2.9.2</version>
> </dependency>
> <dependency>
>     <groupId>io.springfox</groupId>
>     <artifactId>springfox-swagger-ui</artifactId>
>     <version>2.9.2</version>
> </dependency>
> ```
>
> 2. **配置**
>
> ```java
> package com.example.demo.config;
> import org.springframework.context.annotation.Bean;
> import org.springframework.context.annotation.Configuration;
> import springfox.documentation.builders.PathSelectors;
> import springfox.documentation.builders.RequestHandlerSelectors;
> import springfox.documentation.service.ApiInfo;
> import springfox.documentation.service.Contact;
> import springfox.documentation.spi.DocumentationType;
> import springfox.documentation.spring.web.plugins.Docket;
> import springfox.documentation.swagger2.annotations.EnableSwagger2;
> import java.util.ArrayList;
> 
> @Configuration
> @EnableSwagger2
> public class SwaggerConfig {
>     @Bean
>     public Docket docket1(){
>         // 配置多个分组的时候，可以配置多个docket
>         return new Docket(DocumentationType.SWAGGER_2).groupName("A分组");
>     }
> 
>     // 配置了Swagger的Docket的bean实例
>     @Bean
>     public Docket docket(){
> 
>         return new Docket(DocumentationType.SWAGGER_2)
>                 .apiInfo(apiInfo())
>                 .groupName("轻狂书生分组")  // 配置API分组
>                 .enable(true)  // 配置是否启用swagger
>                 .select()
>                 // RequestHandlerSelectors，配置要扫描接口的方式
>                 // basePackage:指定要扫描的包
>                 // any() 扫描全部
>                 // none() 不扫描
>                 // withClassAnnotation() 根据类注解扫描  RequestHandlerSelectors.withClassAnnotation(Controller.class)
>                 // withMethodAnnotation() 根据方法注解扫描  RequestHandlerSelectors.withMethodAnnotation(GetMapping.class)
>                 .apis(RequestHandlerSelectors.basePackage("com.example.demo.controller"))
>                 // 过滤什么路径(接口的url请求路径)
>                 .paths(PathSelectors.ant("/**"))
>                 .build();
>     }
> 
>     // 配置Swagger信息=apiInfo
>     public ApiInfo apiInfo(){
>         return new ApiInfo(
>                 "轻狂书生的API文档",
>                 "API文档描述",
>                 "v1.0" ,
>                 "http://www.baidu.com",
>                 new Contact("charles", "http://github.com/charles", "charles@icloud.com"),
>                 "Apache2.0",
>                 "http://www.apache.org",
>                 new ArrayList());
>     }
> }
> ```
>
> 3. **注解**
>
> ```java
> @ApiModel("雇员实体类")	// 给实体类添加注释
> public class Employee {
>     @ApiModelProperty("雇员ID")	// 给实体类的属性添加注释
>     private Integer id;
>     private String lastName;
>     private Integer gender;
>     private String email;
>     private Integer dId;
> }
> ```
>
> **常用注解**
>
> * @Api 标识一个java类型是文档类，用controller类的类名上
>
> * @ApiModel 表示一个实体类/模型文档，用在类名上；
>
> * @ApiModelProperty 作用在属性上，添加属性描述；
>
> * @ApiOperation 作用在接口类的方法上，控制方法的相关描述；
>
> * @ApiImplicitParam 作用在接口方法上，描述单个参数信息，只能作用在方法上；
>
> * @ApiImplicitParams 作用在接口方法上，@ApiImplicitParam参数组；
>
> * @ApiParam 作用在接口方法上，描述单个参数信息，属性基本与@ApiImplicitParam一样，但可以作用在方法、参数、属性上；
>
> 下面分别对每个注解的常用参数作讲解。
>
> #### @Api： 
>
> value：字符串，对controller类的作用描述，代替原来的description（已过时），一般用此属性；
>
> tags：字符串数组，标签组，同样可以描述controller的作用；
>
> #### @ApiModel  
>
> value：字符串，模型的简短别名，使得在文档的导航中便于识别；
>
> description：字符串，模型的附加描述；
>
> #### @ApiOperation 
>
> value：字符串，方法的功能描述；
>
> tags：字符串数组，标签组，同样可以描述方法的作用；
>
> response：ClassType，显示指出返回的对象类型；在响应示例中会显示出改对象的字段以及示例、描述；
>
> code：响应代码，默认200，一般不改；
>
> #### @ApiModelProperty 
>
> value：字符串，字段描述；
>
> required：boolean；指定参数是否必须，默认false；
>
> example：字符串，参数值的示例
>
> #### @ApiImplicitParam 
>
> name：字符串，参数名；
>
> value：字符串，参数描述；
>
> defaultValue：字符串，参数默认值；
>
> required：boolean，标识是否必须传值，默认false；
>
> dataType：字符串，参数类型，可以是某个类名，也可以是基本数据类型的引用类名，如Integer；
>
> example：字符串，参数值示例；
>
> #### @ApiImplicitParams
>
> value：@ApiImplicitParam类型数组，当方法有多个@ApiImplicitParam参数时，需要放到@ApiImplicitParams注解中
>
> #### @ApiParam 
>
> name：字符串，参数名；
>
> value：字符串，参数描述；
>
> defaultValue：字符串，设置默认值；
>
> required：boolean，是否必须，默认false；
>
> example：字符串，参数值示例；