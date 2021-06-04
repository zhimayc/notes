## ElasticSearch

6.x版本与7.x版本api区别大

- 7.x版本使用流行Restfull API
- 7.x版本废弃了type（相当于数据库表）概念

## 一、windows版安装

#### 1、下载windows版elasticsearch7.6.1

下载地址：https://www.elastic.co/cn/downloads/past-releases#elasticsearch

![image-20210412153855376](image/image-20210412153855376.png)

#### 2、解压即可使用

![image-20210412154008455](image/image-20210412154008455.png)

#### 3、目录结构

![image-20210412154139236](image/image-20210412154139236.png)

- config文件夹中有几个重要配置文件：

  log4j2.properties    日志配置文件

  jvm.options     java虚拟机相关配置

  elasticsearch.yml      主要配置，如集群配置端口配置

- modules   功能模块

- plugins    插件，ik分词器也是插件



#### 4、跨域配置

> 不配置跨域，则elasticsearch-head无法连接
>
> 在elasticsearch.yml中添加如下配置
>
> ![image-20210412161339315](image/image-20210412161339315.png)

#### 5、启动：双击.bat文件

![image-20210412154934211](image/image-20210412154934211.png)

#### 6、启动之后在浏览器中输入localhost:9200

![image-20210412155326079](image/image-20210412155326079.png)

> name  是主机名
>
> cluster_name  集群名，一台主机也是集群，默认集群名就是elasticsearch

#### 6、安装可视化插件 head

> 这是一个数据展示工具，查询操作可以使用Kibana

==确保本地拥有nodejs环境==

下载地址：https://github.com/mobz/elasticsearch-head

将 elasticsearch-head-master.zip 压缩包下载后解压

进入目录运行命令 ==cnmp install==

![image-20210412160650839](image/image-20210412160650839.png)

再使用命令启动 ==npm run start==

![image-20210412160752746](image/image-20210412160752746.png)

浏览器中输入localhost:9100，并连接elasticsearch服务

![image-20210412161559827](image/image-20210412161559827.png)

#### 7、安装Kibana

相当于可视化客户端，用于做查询操作

==需要本地nodejs环境==

下载地址：https://www.elastic.co/cn/downloads/past-releases#kibana

![image-20210412172300860](image/image-20210412172300860.png)

解压缩（非常耗时），通过bin目录中的命令启动

## 二、基本使用和概念

es就相当于数据库，一个索引就相当于一个数据库，下面就创建了一个chenlong索引（数据库）

![image-20210412162120331](image/image-20210412162120331.png)



#### 1、底层数据存储方式

> elasticsearch是面向文档，一切都是json，比如数据库也是一个json，每一个文档都是一个json文档，可以通过fastjson自动转换
>
> 索引：一个文档集合

![image-20210412194910246](image/image-20210412194910246.png)

#### 2、分片机制

#### 3、倒排索引

elasticsearch使用的是一种称为倒排索引的结构，采用Lucene倒排索引作为底层。这种结构适用于快速全文搜索，文档通过分词器进行分词操作，将一个文档分割成一系列词语，再通过词语建立倒排索引。

举例：1、假如有两个文档，内容分别如下

>  文档1：I like red apple.
>
> 文档2：I have an apple tree.

通过英语分词器可以得到如下词条（重复的只保留一个）

> I  like red apple have an tree

通过这些词条可以生成如下倒排索引，列出每个词条在哪个文档有出现

| term  | 文档1 | 文档2 |
| ----- | ----- | ----- |
| I     | √     | √     |
| like  | √     | ×     |
| red   | √     | ×     |
| apple | √     | √     |
| have  | ×     | √     |
| an    | ×     | √     |
| tree  | ×     | √     |

现在假如搜索 I like ，可以发现词条 I 在两个文档中都有，词条like只有文档1有，因为两个文档都有相关词条出现，但是文档1的匹配度更高，所以会按照匹配度输出搜索结果，这就是倒排索引的原理。中文搜索原理类似，只不过要使用中文分词器



>  倒排索引可以完全过滤掉无关数据，无需搜索全文和全部数据，效率高
>
> 一个索引会分成若干个分片，存储在集群中，每个分片其实就是一个倒排索引

#### 4、IK分词器

下载地址：https://github.com/medcl/elasticsearch-analysis-ik

下载解压到elasticsearch的plugins目录

重启elasticsearch可以看到分词器插件加载，可以通过bin目录中的elasticsearch-plugin命令查看加载进来的插件

> elasticsearch-plugin list

IK分词器有两种算法，一种的==ik_max_word== 最细粒度划分，另一种是==ik_smart== 最少切分，最细粒度划分就是将一个文本尽可能划分出所有可能的词条，最少切分就是将一个文本尽可能少的切分但要保证切分出来的都是词条

在Kibana上进行测试：==ik_smart== 

```
get _analyze
{
  "analyzer":"ik_smart",
  "text":"中国共产党"
}
```

得到如下切分结果：

```
{
  "tokens" : [
    {
      "token" : "中国共产党",
      "start_offset" : 0,
      "end_offset" : 5,
      "type" : "CN_WORD",
      "position" : 0
    }
  ]
}
```

测试：==ik_max_word== 

```
get _analyze
{
  "analyzer":"ik_max_word",
  "text":"中国共产党"
}
```

得到如下切分结果：

```
{
  "tokens" : [
    {
      "token" : "中国共产党",
      "start_offset" : 0,
      "end_offset" : 5,
      "type" : "CN_WORD",
      "position" : 0
    },
    {
      "token" : "中国",
      "start_offset" : 0,
      "end_offset" : 2,
      "type" : "CN_WORD",
      "position" : 1
    },
    {
      "token" : "国共",
      "start_offset" : 1,
      "end_offset" : 3,
      "type" : "CN_WORD",
      "position" : 2
    },
    {
      "token" : "共产党",
      "start_offset" : 2,
      "end_offset" : 5,
      "type" : "CN_WORD",
      "position" : 3
    },
    {
      "token" : "共产",
      "start_offset" : 2,
      "end_offset" : 4,
      "type" : "CN_WORD",
      "position" : 4
    },
    {
      "token" : "党",
      "start_offset" : 4,
      "end_offset" : 5,
      "type" : "CN_CHAR",
      "position" : 5
    }
  ]
}
```

注意点：如果字段类型是keyword类型，那么分词器不会对他起作用，而text类型可以

![image-20210413143128316](image/image-20210413143128316.png)



#### 5、IK分词器字典

IK分词器之所以能将一句中文拆分成多个词条，那是因为IK分词器中的字典，字典中包含的词条IK分词器才能理解。

假如有一个词条 “苹果”， 但是字典中没有，那么IK分词器就无法理解，会继续将其拆分成可以理解的两个词条 "苹" 和 "果"

因此对于一些IK分词器字典中没有的词条，需要我们加入到该字典中，步骤如下：

在ik分词器的config目录下创建一个文件==kuang.dic==添加词条“苹果”，注意每个词条独占一行

![image-20210412215018619](image/image-20210412215018619.png)

然后在IKAnalyzer.cfg.xml文件中配置刚才创建的字典文件

![image-20210412215206630](image/image-20210412215206630.png)

重启elasticsearch服务即可

#### 6、elasticsearch  Restfull API

> get： 请求类型
>
> _analyze： 资源名
>
> json部分： 请求体中的数据

```
get _analyze
{
  "analyzer":"ik_smart",
  "text":"中国共产党"
}
```

基本命令：

![image-20210412215639955](image/image-20210412215639955.png)



## 三、api使用

#### 1、索引操作

###### 创建索引chen_index

```java
PUT /chen_index
```

###### 配置索引mapping

> elasticsearch中的文档相当于java中的对象，都包含字段和字段的数据类型，在elasticsearch中使用mapping设置字段和数据类型
>
> 如果不设置mapping，elasticsearch在插入数据时会自动判断数据类型，这样数据在插入时如果没有严格按照格式插入，那么后期在对接java时会有问题，可能出现数据类型转换异常，所以应该提前配置好文档的字段信息
>
> mapping是处理数据的方式和规则方面做一些限制，如某个字段的数据类型、默认值、分析器、是否被索引等等，这些都是映射里面可以设置的，其它就是处理es里面数据的一些使用规则设置也叫做映射，按着最优规则处理数据对性能提高很大，因此才需要建立映射，并且需要思考如何建立映射才能对性能更好。
>
> 注意：mapping建立之后不可修改已存在的字段，可添加字段

```
PUT /chen_index
{
  "mappings": {
    "properties": {
      "name":{
        "type": "text"
      },
      "age":{
        "type": "long"
      },
      "birthday":{
        "type": "date"
      }
    }
  }
}
```

###### 配置索引setting

> 可以通过settings修改索引配置，副本数可以修改，分片数不能修改

```
PUT /chen_index
{
  "settings": {
		"number_of_shards": 3,
		"number_of_replicas": 2
	}
}
```



#### 2、文档操作

###### 添加一个文档

> 在kuangshen索引中添加第一个文档，1就是该文档的_id属性值

```
PUT /kuangshen/_doc/1
{
  "name":"狂神说",
  "age":12,
  "desc":"一顿操作猛如虎，一看是个250",
  "tags":["技术宅","暖男","直男"]
}
```

![image-20210413120409061](image/image-20210413120409061.png)

多添加几个数据

![image-20210413121252881](image/image-20210413121252881.png)



###### 更新文档数据

> 更新操作可以使用==PUT== 、==POST==以及 ==POST _update==方式，这三种方式有区别：
>
> 如果全部字段都有修改，那么三种方式都可以，如果只修改部分字段，那么PUT和POST方式未修改的字段会置空，也就是说只有POST _update方式可以动态修改，想改哪个字段就改哪个字段
>
> 写法也不同，POST _update方式需要doc包裹，而其他两种不用，如下：
>
> ```
> PUT /kuangshen/_doc/1
> {
>   "name":"chen1"
> }
> 
> POST /kuangshen/_doc/1
> {
>   "name":"chen2"
> }
> 
> POST /kuangshen/_doc/1/_update
> {
>   "doc":{
>     "name":"chen3"
>   }
> }
> ```

## 四、查询

#### 1、简单查询

```
GET /kuangshen/_doc/_search         //_search查询默认写法
{
  "query":{            //查询的匹配条件
    "match":{
      "name":"狂神"
     }
   },
   "_source":["name","desc"]    //查询结果中显示的字段，不写则默认返回所有字段
}
```

###### 

#### 2、复合查询关键字

> 复合查询需要使用==bool==包裹，可以将bool想象成括号，将包裹在内的条件当成整体
>
> bool中包含有以下四种判断条件：
>
> - ==must==：必须匹配，表示and，==参与算分==
>
> - ==should==：或，表示or，==参与算分==
>
> - ==must_not==：必须不匹配，表示not，==不参与算分==
>
> - ==filter==：必须匹配，==不参与算分==
>
> 这个条件中可以加入若干个 match 匹配或者 term 匹配

```
GET /jd_index/_doc/_search
{
  "query": {
    "bool": {
      
      "must": [],
      "should": [],
      "must_not": [],
      "filter": []
      
    }
  }
}
```

###### 单独使用must

> ==shop字段==值同时包含 ==电== 和 ==人民==的条件成立

```
GET /jd_index/_doc/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {"shop": "电"}
        },
        {
          "match": {"shop": "人民"}
        }
      ]
    }
  }
}
```

###### 单独使用should

> ==shop字段==包含 ==电== 或者 ==人民== 则条件成立

```
GET /jd_index/_doc/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "match": {"shop": "电"}
        },
        {
          "match": {"shop": "人民"}
        }
      ]
    }
  }
}
```

###### 单独使用must_not

> ==shop字段==既不包含 ==电== 又不包含 ==人民== 则条件成立

```
GET /jd_index/_doc/_search
{
  "query": {
    "bool": {
      "must_not": [
        {
          "match": {"shop": "电"}
        },
        {
          "match": {"shop": "人民"}
        }
      ]
    }
  }
}
```

###### 单独使用filter

> 用于设置过滤条件

```
GET /jd_index/_doc/_search
{
  "query": {
    "bool": {
      "filter": [
        {
          "match": {"shop": "电"}
        },
        {
          "match": {"shop": "人民"}
        }
      ]
    }
  }
}
```



#### 3、各种匹配条件

###### match 匹配

> 既能处理全文字段，又能处理精确字段
>
> match在执行查询前，它将用正确的分析器去分析查询字符串，输出分词器分析的词条结果，然后分别对每个词条进行term查询，在倒排索引中查找出包含该词条的文档，计算评分得出结果



###### term 匹配

> 精确匹配，英文区分大小写，不对查询字符串进行分词器解析，直接查询倒排索引
>
> 注意：如果term中的查询字符串是英文一定要小写!!!，因为当初的文档经过分词器分析出词条生成倒排索引时，会将英文词条转成小写，而term查询会区分英文大小写，如果查询字符串包含英文大写则无法在倒排索引中找到词条，从而没有搜索结果。相反的，如果索引的某个字段设置为==not_analyzed不被分析==或者设置为==keyword==类型，那么生成索引时不会用分词器去分析该字段，那么该字段的值原本是大写或小写，生成索引时也是大写或小写。
>
> term查询是直接通过倒排索引指定的词条进行精确的查找
>
> match查询会使用分词器解析，先分析文档，然后通过分析的文档进行查询
>
> 注意：text和keyword的区别，keyword类型的字段不会被分词器解析，假如有一个keyword类型的字段name，值是 "陈龙2" ，通过term查询  name=陈龙  的文档是查不到的，因为分词器不会解析keyword类型的字段值，因此倒排索引中也没有 陈龙 这个词条，只有 陈龙2这个词条



###### range 匹配

> 范围匹配：
>
> gt：大于
>
> gte：大于等于
>
> lt：小于
>
> lte：小于等于

```
GET /jd_index/_doc/_search
{
    "query":{
        "range":{
            "price":{
                "gte":20,
                "lt":30
            }
        }
    }
}
```



###### exists 和 missing

> 判断文档中是否包含某个字段或者没有某个字段



###### prefix 前缀匹配

###### wildcard 通配符匹配

###### regexp 正则匹配

###### fuzzy 模糊匹配



###### 指定返回的字段

```
GET /jd_index/_doc/_search
{
  "query":{
    "bool": {
      "must": [
        { "match": {"shop": "电"} }
      ]
    }
  },
  "_source":["title","price"]
}
```

###### 指定排序字段

>  下面是按照age降序，有了排序规则后返回的结果就没有分值了，都是null

```
GET /jd_index2/_doc/_search
{
  "query":{
    "bool": {
      "must": [
        { "match": {"shop": "电"} }
      ]
    }
  },
  "sort":[
    {
      "price":{ "order":"desc" }
    }
  ]
}
```

###### 浅分页

> 从0开始取2条

```
"form": 0,
"size": 2
```

###### 深分页

> 上面的form和size是浅分页，随着from的值不断变大效率越来越慢。除了效率上的问题，还有一个无法解决的问题是，es 目前支持最大的 skip 值是 max_result_window ，默认为 10000 。也就是当 from + size > max_result_window 时，es 将返回错误，可以暂时通过调大max_result_window 的值来解决，但依旧无法规避效率问题
>
> 深分页就是解决深度分页时的效率问题



###### highlight 高亮查询

> 下面是对desc字段进行高亮，注意高亮字段必须是查询条件里的字段

```
GET /kuangshen/_doc/_search
{
  "query":{
    "match":{
      "desc":"一顿"
    }
  },

  "highlight":{
    "fields":{
      "desc":{}
    }
  }
}
```

查询结果中就多了一个highlight部分

![image-20210413170948261](image/image-20210413170948261.png)

也可以自定义highlight部分的标签

```
GET /kuangshen/_doc/_search
{
  "query":{
    "match":{
      "desc":"一顿"
    }
  },

  "highlight":{
    "pre_tags":"<p class='key' style='color:red'>",
    "post_tags":"</p>",
    "fields":{
      "desc":{}
    }
  }
}
```

替换成了自定义标签

![image-20210413171249740](image/image-20210413171249740.png)







## 五、springboot集成

#### 1、导入依赖

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-data-elasticsearch</artifactId>
</dependency>
```



#### 2、修改elasticsearch版本和线上一致

![image-20210413215955124](image/image-20210413215955124.png)

#### 3、创建elasticsearch配置类

```java
import org.apache.http.HttpHost;
import org.elasticsearch.client.RestClient;
import org.elasticsearch.client.RestHighLevelClient;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class ElasticsearchConfig {

    @Bean
    public RestHighLevelClient restHighLevelClient(){
        RestHighLevelClient restHighLevelClient = new RestHighLevelClient(
                //有多少台elasticsearch就配置多少个
                RestClient.builder(new HttpHost("192.168.100.199",9200,"http"))
        );
        return restHighLevelClient;
    }
}
```

#### 4、索引操作api

```java
public class TestController {

    @Autowired
    @Qualifier("restHighLevelClient")
    private RestHighLevelClient client;

    /**
     * 创建索引
     * @return
     */
    @RequestMapping("create_index")
    Object test1() throws IOException {

        //创建索引请求
        CreateIndexRequest request = new CreateIndexRequest("chen_index");
        //客户端执行请求   请求参数默认
        CreateIndexResponse createIndexResponse = client.indices().create(request, RequestOptions.DEFAULT);
        //打印结果
        System.out.println(createIndexResponse.toString());
        return null;
    }
    
    /**
     * 判断索引是否存在
     */
    @RequestMapping("exists_index")
    Object test2() throws IOException {
        GetIndexRequest request = new GetIndexRequest("chen_index");
        boolean exists = client.indices().exists(request, RequestOptions.DEFAULT);
        System.out.println(exists);
        return null;
    }

    /**
     * 删除索引
     */
    @RequestMapping("rm_index")
    Object test3() throws IOException {
        DeleteIndexRequest request = new DeleteIndexRequest("chen_index");
        AcknowledgedResponse delete = client.indices().delete(request, RequestOptions.DEFAULT);
        System.out.println(delete.isAcknowledged());
        return null;
    }
}
```

#### 5、文档操作api

###### 创建、更新、删除

```java
/**
     * 创建文档
     */
    @RequestMapping("create_doc")
    Object test4() throws IOException {
        User user = new User("狂神说", 21);
        //创建请求
        IndexRequest request = new IndexRequest("chen_index");
        // put /chen_index/_doc/1
        request.id("1");
        request.timeout(TimeValue.timeValueSeconds(1));
        request.timeout("1s");

        request.source(JSON.toJSONString(user), XContentType.JSON);

        IndexResponse index = client.index(request, RequestOptions.DEFAULT);

        System.out.println(index.toString());
        System.out.println(index.status());

        return null;
    }

    /**
     * 根据id判断某个文档是否存在
     */
    @RequestMapping("exists_doc")
    Object test5() throws IOException {
        GetRequest request = new GetRequest("chen_index", "1");
        //不获取_source的上下文
//        request.fetchSourceContext(new FetchSourceContext(false));
//        request.storedFields("name");

        boolean exists = client.exists(request, RequestOptions.DEFAULT);
        System.out.println(exists);
        return null;
    }

    /**
     * 根据id获取文档信息
     */
    @RequestMapping("get_doc")
    Object test6() throws IOException {
        GetRequest request = new GetRequest("chen_index", "1");
        GetResponse resp = client.get(request, RequestOptions.DEFAULT);

        //打印文档内容
        System.out.println(resp.getSourceAsString());
        System.out.println(resp.getSource());  //完整信息
        return null;
    }

    /**
     * 根据id更新文档
     */
    @RequestMapping("update_doc")
    Object test7() throws IOException {
        UpdateRequest request = new UpdateRequest("chen_index", "1");
        request.timeout("1s");

        User user = new User();
        user.setName("哼");
        request.doc(JSON.toJSONString(user), XContentType.JSON);

        UpdateResponse update = client.update(request, RequestOptions.DEFAULT);
        System.out.println(update.status());
        return null;
    }


    /**
     * 根据id删除文档
     */
    @RequestMapping("delete_doc")
    Object test8() throws IOException {
        DeleteRequest request = new DeleteRequest("chen_index", "1");
        request.timeout("1s");

        DeleteResponse update = client.delete(request, RequestOptions.DEFAULT);
        System.out.println(update.status());
        return null;
    }
```

###### 批量操作

```java
/**
 * 批量插入，其他批量操作也可以这样
 */
@RequestMapping("batch_create_doc")
Object test9() throws IOException {
    BulkRequest bulkRequest = new BulkRequest();
    bulkRequest.timeout("10s");

    ArrayList<User> users = new ArrayList<>();
    users.add(new User("chen1",21));
    users.add(new User("chen2",22));
    users.add(new User("chen3",23));
    users.add(new User("chen4",24));

    for(int i=0 ; i<users.size(); i++){
        IndexRequest request = new IndexRequest("chen_index");
        request.id(""+(i+1));
        request.source(JSON.toJSONString(users.get(i)),XContentType.JSON);
        bulkRequest.add(request);
    }

    BulkResponse bulk = client.bulk(bulkRequest,RequestOptions.DEFAULT);
    System.out.println(bulk.hasFailures());  //false表示成功

    return null;
}
```

###### 查询

```java
/**
 * 查询
 */
Object test10() throws IOException {

    SearchRequest searchRequest = new SearchRequest("chen_index");

    //构建搜索条件
    SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();
    //查询条件使用工具来实现
    TermQueryBuilder termQueryBuilder = QueryBuilders.termQuery("name", "chen");

    searchSourceBuilder.query(termQueryBuilder);
    searchSourceBuilder.timeout(new TimeValue(60, TimeUnit.SECONDS));

    searchRequest.source(searchSourceBuilder);

    SearchResponse search = client.search(searchRequest, RequestOptions.DEFAULT);

    System.out.println(search.getHits());
    System.out.println(JSON.toJSONString(search.getHits()));

    return null;
}
```

## 六、实战

#### 1、java爬虫

html爬取可以使用jsoup包，音乐影视爬取可以使用tika包

导入依赖

```xml
<dependency>
    <groupId>org.jsoup</groupId>
    <artifactId>jsoup</artifactId>
    <version>1.10.2</version>
</dependency>
```

创建工具类

```java
import com.example.elasticsearch.pojo.Product;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;
import org.springframework.util.StringUtils;

import java.net.URL;
import java.net.URLEncoder;
import java.util.ArrayList;
import java.util.List;

public class HttpParseUtils {


//    public static void main(String[] args) throws Exception {
//
//        List<Product> list = getProductWithKey("心理学");
//        for (Product p: list) {
//            System.out.println(p);
//        }
//
//    }

    public static List<Product> getProductWithKey(String key) throws Exception {
        if (StringUtils.isEmpty(key)){
            throw new Exception();
        }

        String url = "https://search.jd.com/Search?keyword=" + URLEncoder.encode(key, "UTF-8");
        System.out.println(url);
        URL url1 = new URL(url);

        //解析网页  document就是js中的Document页面对象，所有在js中可以使用的方法这里都可以使用
        Document document = Jsoup.parse(url1, 30000);
        Element j_goodsList = document.getElementById("J_goodsList");
        Elements lis = j_goodsList.getElementsByTag("li");

        ArrayList<Product> products = new ArrayList<>();
        for (Element li : lis) {

            String img = li.getElementsByTag("img").eq(0).attr("data-lazy-img");  //source-data-lazy-img
            String title = li.getElementsByClass("p-name").eq(0).text();
            String price = li.getElementsByClass("p-price").eq(0).text();
            String shop = li.getElementsByClass("p-shopnum").eq(0).text();

            Product product = new Product();
            product.setImg(img);
            product.setTitle(title);
            product.setPrice(price);
            product.setShop(shop);

            products.add(product);
        }

        return products;
    }

}
```

爬取数据批量存入es

```java
/**
 * 爬取数据存入es
 */
@RequestMapping("data")
Object test11(@RequestParam String key) throws Exception {

    List<Product> list = HttpParseUtils.getProductWithKey(key);

    BulkRequest bulkRequest = new BulkRequest();
    bulkRequest.timeout("10s");

    for(int i=0 ; i<list.size(); i++){
        IndexRequest request = new IndexRequest("jd_index");
        request.id(""+(i+1));
        request.source(JSON.toJSONString(list.get(i)),XContentType.JSON);
        bulkRequest.add(request);
    }

    BulkResponse bulk = client.bulk(bulkRequest,RequestOptions.DEFAULT);
    System.out.println(bulk.hasFailures());  //false表示成功

    return null;
}
```

es查询数据返回前端

```java
/**
 * 数据查询
 */
@RequestMapping("getdata")
Object test12(@RequestParam String key) throws IOException {
    SearchRequest searchRequest = new SearchRequest("jd_index");

    //构建搜索条件
    SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();
    //查询条件使用工具来实现
    TermQueryBuilder termQueryBuilder = QueryBuilders.termQuery("title", key);

    searchSourceBuilder.query(termQueryBuilder);
    searchSourceBuilder.timeout(new TimeValue(60, TimeUnit.SECONDS));

    searchRequest.source(searchSourceBuilder);

    SearchResponse response = client.search(searchRequest, RequestOptions.DEFAULT);

    ArrayList<Object> objects = new ArrayList<>();
    //解析结果
    for(SearchHit hit : response.getHits().getHits()){
        objects.add(hit.getSourceAsMap());
    }

    return objects;
}
```



## 七、更新索引

#### 1、零停机重新索引

> 零停机重新索引主要用于索引的全量更新和索引数据转移，通过使用索引别名实现
>
> 对于线上的es索引库通常需要建立一个索引别名，查询业务使用别名进行，当需要对原索引重新建立索引，为了不影响线上业务，需要先建立好新索引，然后将别名指定到新索引即可



> 常见的需要重建索引的情况：修改字段类型、修改分片数



###### 创建一个索引 my_index

```
PUT /my_index
```

###### 配置映射

```
PUT /my_index/_mapping
{
	"properties": {
        "name": {
            "type": "text"
        },
        "age": {
        	"type": "long"
        }
    }
}
```

###### 添加一个文档

```
PUT /my_index/_doc/1
{
	"name": "李峰",
	"age": 22
}
```

###### 设置一个别名 prod_index 指向 my_index

```
PUT /my_index/_alias/prod_index
```

###### 检查别名是否设置成功

```
GET /my_index/_alias/*
```

###### 使用 prod_index 也可以查找数据

```
GET /prod_index/_doc/1
```

###### 重建索引

```
PUT /my_index2
```

###### 创建映射

```
PUT /my_index2/_mapping
{
	"properties": {
        "name": {
            "type": "keyword"
        },
        "age": {
        	"type": "long"
        },
        "height": {
        	"type": "int"
        }
    }
}
```

###### 将数据从my_index中索引到my_index2

> ES提供了_reindex这个API。相对于我们重新导入数据肯定会快不少，实测速度大概是bulk导入数据的5-10倍。

```
POST /_reindex
{
	"source": {
		"index": "my_index"
	},
	"dest": {
		"index": "my_index2"
	}
}
```

###### 查看有无数据

```
GET /my_index2/_doc/1
```



###### 修改索引别名指向

> 接下来就是将别名`prod_index`指到`my_index2`上，因为一个别名可以指向多个索引，所以我们还需要删除`my_index`的别名。这个操作需要原子化，所以我们使用`_aliases API`来操作。

```bash
POST /_aliases
{
    "actions": [
        {
            "remove": {
                "index": "my_index",
                "alias": "prod_index"
            }
        },
        {
            "add": {
                "index": "my_index2",
                "alias": "prod_index"
            }
        }
    ]
}
```

#### 2、增量更新索引

