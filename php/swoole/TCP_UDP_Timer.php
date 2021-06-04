<?php

// use Swoole\Server;

// //创建服务器
// $host = '0.0.0.0';
// $port = 9501;
// $mode = SWOOLE_PROCESS;
// $sock_type = SWOOLE_SOCK_TCP;
// $serv = new Server($host,$port,$mode,$sock_type);
// /**
//  *  $host: 参数用来指定监听的ip地址，如127.0.0.1，或者外网地址，或者0.0.0.0监听全部地址
//  *  IPv4使用 127.0.0.1表示监听本机，0.0.0.0表示监听所有地址
//  *  IPv6使用::1表示监听本机，:: (相当于0:0:0:0:0:0:0:0) 表示监听所有地址
//  * 
//  *  $port监听的端口，如9501
//  *  如果$sock_type为UnixSocket Stream/Dgram，此参数将被忽略
//  *  监听小于1024端口需要root权限
//  *  如果此端口被占用server->start时会失败
//  *  $mode运行的模式
//  *  SWOOLE_PROCESS多进程模式（默认）
//  *  SWOOLE_BASE基本模式
//  *  $sock_type指定Socket的类型，支持TCP、UDP、TCP6、UDP6、UnixSocket Stream/Dgram 6种
//  *  使用$sock_type | SWOOLE_SSL可以启用SSL隧道加密。启用SSL后必须配置ssl_key_file和ssl_cert_file
//  * 
//  */

//  $serv->on('connect',function($serv,$fd){
  
//     echo '建立连接 \n';
//     var_dump($serv);
//     var_dump($fd);
//  });

//  $serv->on('receive',function($serv,$fd,$reactor_id,$data){
//     echo '接收到数据 \n';
//     var_dump($data);
//  });

//  $serv->on('close',function($serv,$fd){
//     echo '连接关闭';
//  });

//  $serv->start();


//  /**
//   * bool $swoole_server->on(string $event,mixed $callback);
//   * $event:
//   * connect : 当建立连接的时候 $serv:服务器信息  $fd: 客户端信息
//   * receive : 当接收到数据 $serv:服务器信息 $fd: 客户端信息 $reactor_id: 客户端id $data : 数据
//   * close : 关闭连接
//   *
//   */



// 建立http服务器
// use Swoole\Http\Server;

// $http = new Server("0.0.0.0", 9501);
// $http->on('request', function ($request, $response) {
//     $response->end("<h1>Hello Swoole. #".rand(1000, 9999)."</h1>");
// });
// $http->start();


// 循环执行定时器

// swoole_timer_tick(2000,function($timer_id){
//     echo "执行 $timer_id \n";
// });

// echo "主进程下面的语句 \n";



// 单次执行定时器 ，在指定的时间后执行函数

// swoole_timer_after(3000,function(){
//     echo "延时执行语句 \n";
// });

// echo "主进程下面的语句 \n";



// 异步TCP
$serv = new swoole_server("0.0.0.0",9501);

// 设置异步 进程工作数
$serv->set(array('task_worker_num' => 4));

// 投递异步任务
$serv->on('receive',function($serv,$fd,$reactor_id,$data){
    $task_id = $serv->task($data);      // 异步ID
    echo "异步ID：$task_id \n";
});

// 处理异步任务
$serv->on('task',function($serv,$task_id,$reactor_id,$data){
    echo "执行异步ID：$task_id \n";
    $serv->finish("$data -> Ok");
});

// 处理结果
$serv->on('finish',function($serv,$task_id,$data){
    echo "执行完成\n";
});

$serv->start();