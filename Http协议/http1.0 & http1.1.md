# http1.0 & http1.1

> 区别一：1.0版本每次请求响应都会建立新的连接，浪费资源，1.1版本可以复用tcp连接



> http协议是以ASCII码传输的

### 请求消息数据格式

> 请求消息数据格式包括如下几部分：
>
> ==请求行==：主要标明方法、请求URL和http的版本号
>
> ![img](image/None.gif)**GET /index.html HTTP/1.1**
>
> ==请求头==：主要是一些请求头参数
>
> ​    **Accept: \*/\***
> ![img](image/None.gif)**Accept-Language: zh-cn**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**Accept-Encoding: gzip, deflate**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**If-Modified-Since: Wed, 17 Oct 2007 02:15:55 GMT**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**If-None-Match: W/"158-1192587355000"**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**User-Agent: Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**Host: 192.168.2.162:8080**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**Connection: Keep-Alive**
>
> ==请求体==：一些附加数据
>
> 注意：GET请求不包含请求体、POST包含请求体

### 响应消息数据格式

> 响应消息数据格式包括如下几部分：
>
> ==响应行==：主要标明协议版本和状态码
>
> ​    **HTTP/1.0 200 OK**
>
> ==响应头==：响应头的参数
>
> ​    **ETag: W/"158-1192590101000"**
> ![img](image/None.gif)**Last-Modified: Wed, 17 Oct 2007 03:01:41 GMT**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**Content-Type: text/html**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**Content-Length: 158**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**Date: Wed, 17 Oct 2007 03:01:59 GMT**
> ![img](http://www.blogjava.net/Images/OutliningIndicators/None.gif)**Server: Apache-Coyote/1.1**
>
> ==响应体==：返回的数据

### POST请求中常见的请求体消息格式

> POST 提交的数据必须放在消息主体（entity-body）中，但协议并没有规定数据必须使用什么编码方式。实际上，开发者完全可以自己决定消息主体的格式，只要最后发送的 HTTP 请求满足上面的请求格式就可以
>
> 但是，数据发送出去，还要服务端解析成功才有意义。一般服务端语言如 php、python 等，以及它们的 framework，都内置了自动解析常见数据格式的功能。服务端通常是根据请求头（headers）中的 Content-Type 字段来获知请求中的消息主体是用何种方式编码，再对主体进行解析。所以说到 POST 提交数据方案，包含了 Content-Type 和消息主体编码方式两部分



##### application/x-www-form-urlencoded

> 浏览器的原生 form 表单，如果不设置 enctype 属性，那么最终就会以 application/x-www-form-urlencoded 方式提交数据，如下，省略了部分请求头字段
>
> ```
> POST http://www.example.com HTTP/1.1
> Content-Type: application/x-www-form-urlencoded;charset=utf-8
> 
> title=test&sub%5B%5D=1&sub%5B%5D=2&sub%5B%5D=3
> ```
>
> 第一行为请求行、第二行为请求头、第三行为请求体，form表单数据就放在请求体里面。数据经过了URL转码，否则会报错



##### multipart/form-data

> 使用表单上传文件时，必须让 form标签 的 enctyped 属性等于这个值
>
> ```
> POST http://www.example.com HTTP/1.1      //请求行
> Content-Type:multipart/form-data; boundary=----WebKitFormBoundaryrGKCBY7qhFd3TrwA
> 
> ------WebKitFormBoundaryrGKCBY7qhFd3TrwA
> Content-Disposition: form-data; name="name"
>  
> chenlong
> ------WebKitFormBoundaryrGKCBY7qhFd3TrwA
> Content-Disposition: form-data; name="password"
>  
> 111111
> ------WebKitFormBoundaryrGKCBY7qhFd3TrwA
> Content-Disposition: form-data; name="headerimg"; filename="chrome.png"
> Content-Type: image/png
>  
> PNG ... content of chrome.png ...
> ------WebKitFormBoundaryrGKCBY7qhFd3TrwA--
> ```
>
> 定义的boundary用于分割不同的字段，为了避免与正文内容重复
>
> 消息主体里按照字段个数又分为多个结构类似的部分，每部分都是以 ––boundary 开始，紧接着内容描述信息，然后是回车，最后是字段具体内容（文本或二进制）。如果传输的是文件，还要包含文件名和文件类型信息。消息主体最后以 ––boundary–– 标示结束。



##### application/json

> 用来告诉服务端消息主体是序列化后的 JSON 字符串，除了低版本 IE 之外的各大浏览器都原生支持 JSON.stringify，服务端语言也都有处理 JSON 的函数，使用 JSON 不会遇上什么麻烦。
>
> ```
> POST http://www.example.com HTTP/1.1
> Content-Type: application/json;charset=utf-8
> 
> {"title":"test","sub":[1,2,3]}
> ```



##### text/xml

> ```
> POST http://www.example.com HTTP/1.1
> Content-Type: text/xml
> 
> <?xml version="1.0"?>
> <methodCall>
>     <methodName>examples.getStateName</methodName>
>     <params>
>         <param>
>             <value><i4>41</i4></value>
>         </param>
>     </params>
> </methodCall>
> ```



##### 上传文件

>  使用multipart/form-data

controller

```java
    @RequestMapping("upload")
    @ResponseBody
    Object upload(String username,@RequestParam(value = "header") MultipartFile file){

        if (file.isEmpty()) {
            System.out.println("文件为空空");
        }
        String fileName = file.getOriginalFilename();  // 文件名
        String suffixName = fileName.substring(fileName.lastIndexOf("."));  // 后缀名
        String filePath = "D://temp-rainy//"; // 上传后的路径
        fileName = UUID.randomUUID() + suffixName; // 新文件名
        File dest = new File(filePath + fileName);
        if (!dest.getParentFile().exists()) {
            dest.getParentFile().mkdirs();
        }
        try {
            file.transferTo(dest);
        } catch (IOException e) {
            e.printStackTrace();
        }
//        String filename = "/temp-rainy/" + fileName;
//        model.addAttribute("filename", filename);
//        return "file";

        return "上传成功";
    }
```

thymeleaf页面

```xml
<!DOCTYPE html>
<html lang="en" xmlns:th="http://www.thymeleaf.org">
<head>
    <meta charset="UTF-8">
    <title>index</title>
</head>
<body>

<form action="/testa/upload" method="post" enctype="multipart/form-data">
    用户名：<input type="text" name="username"><br>
    图片：<input type="file" name="header"><br>
    <input type="submit" value="登录">

</form>

</body>
</html>
```



GET方式提交的数据最多只能有1024字节，而POST则没有此限制