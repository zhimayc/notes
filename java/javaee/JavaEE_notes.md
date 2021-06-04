# JavaEE实战开发笔记

-----

### 分布式基础概念

##### 常见的负载均衡算法

> * **轮询：**为第一个请求选择健康池中的第一个后端服务器，然后按顺序往后依次选择，直到最后一个，然后循环
> * **最小连接：**优先选择连接数最少，也就是压力最小的后端服务器，在会话较长的情况下可以考虑采取这种方式
> * **散列：**根据请求源的IP的散列(hash)来选择要转发的服务器。这种方式可以一定程度上保证特定用户能连接到相同的服务器。如果你的应用需要处理状态而要求用户能连接到和之前相同的服务器，可以考虑采取这种方式。

### Nacos基本概念

> 命名空间：用来做配置隔离
>
> ​	默认：public (保留空间)；默认新增的所有配置都在public空间
>
> ​		1、开发，测试，生产：利用命名空间来做环境隔离，需要在bootstrap.properties中，配置需要使用哪个命名空间的配置
>
> ​		2、每一个微服务之间互相隔离配置，每一个微服务都创建自己的命名空间，只加载自己命名空间下的配置
>
> 配置集：所有的配置的集合
>
> 配置集ID：类似文件名  Data ID
>
> 配置分组：默认所有的配置集都属于：DEFAULT_GROUP
>
> 同时加载多个配置集
>
> 1）、微服务任何配置信息，任何配置文件都可以放在配置中心中
>
> 2）、只需要在bootstrap.properties中说明加载配置中心中哪些配置文件即可
>
> 3）、@Value、@ConfigurationProperties，以前SpringBoot任何方式从配置文件中获取值都能使用，配置中心有的配置，优先使用配置中心的值
>
> 配置示例：(bootstrap.properties)
>
> ```properties
> spring.cloud.nacos.config.server-addr=127.0.0.1:8848
> spring.application.name=gulimall-coupon
> spring.cloud.nacos.config.namespace=9c2ff426-a493-4882-be5b-999566e70b2d
> spring.cloud.nacos.config.group=dev
> spring.cloud.nacos.config.extension-configs[0].data-id=datasource.yml
> spring.cloud.nacos.config.extension-configs[0].group=dev
> spring.cloud.nacos.config.extension-configs[0].refresh=true
> 
> spring.cloud.nacos.config.extension-configs[1].data-id=mybatis.yml
> spring.cloud.nacos.config.extension-configs[1].group=dev
> spring.cloud.nacos.config.extension-configs[1].refresh=true
> 
> spring.cloud.nacos.config.extension-configs[2].data-id=other.yml
> spring.cloud.nacos.config.extension-configs[2].group=dev
> spring.cloud.nacos.config.extension-configs[2].refresh=true
> ```
>
> 

