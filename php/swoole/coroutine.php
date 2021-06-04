<?php

// 执行DNS查询
// swoole_async_dns_lookup("www.baidu.com", function($host, $ip){  // 1.9.24版本之后不再需要显式调用
//     echo "{$host} : {$ip}\n";
// });


// 异步文件IO
// swoole_async_readfile(__DIR__."/1.txt", function($filename, $content) {
//     echo "$filename: $content";
// });

// use Swoole\Coroutine as co;
// $filename = __DIR__ . "/1.txt";
// co::create(function () use ($filename)
// {
//     $r =  co::readFile($filename);
//     var_dump($r);
// });


// V4.3.0中移除了异步模块。改用 Coroutine 协程模块。




// 创建协程


// 设置全局通道
$c = new chan(2);



function test(){
    co::sleep(1);
    global $c;
    $data = $c->pop();
    echo "test data = $data \n";
    
    $host = co::gethostbyname('www.baidu.com');
    var_dump($host);

    echo "test \n";
}

function test2(){
    co::sleep(1.5);
    global $c;
    $data = $c->pop();
    echo "test2 data = $data \n";
    echo "test2 \n";
}

function test3(){
    co::sleep(2);
    echo "test3 \n";
}

go(function(){
    co::sleep(0.5);
    echo "hello \n";
});


$c->push("hello");
$c->push("world");


go("test");

go("test2");


echo "主线程 \n";
// go([$object, "method"]);