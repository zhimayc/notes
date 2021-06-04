<?php

// 创建进程
// 进程对应的执行函数
// function doProcess(swoole_process $worker){
//     echo "PID,$worker->pid \n";
//     sleep(3);
// }

// // 创建子进程
// for($i = 0;$i < 3;$i++){
//     $process = new swoole_process("doProcess");
//     $pid = $process->start();
// }

// // 回收结束运行的子进程
// swoole_process::wait();



// 进程事件
// $workers = [];   // 进程数组
// $worker_num = 3;    // 创建进程的数据量


// // 创建启动进程
// for($i = 0;$i < $worker_num;$i++){
//     $process = new swoole_process("doProcess");    // 创建单独的新进程
//     $pid = $process->start();   // 启动进程，并获取进程ID
//     $workers[$pid] = $process;  // 存入进程数组
// }

// function doProcess(swoole_process $process){
//     $process->write("PID：$process->pid 。");  // 子进程写入信息
//     echo "写入信息：$process->pid $process->callback";
// }

// // 添加进程事件 向每一个子进程添加需要执行的动作
// foreach($workers as $process){
//     swoole_event_add($process->pipe,function($pipe) use($process){
//         $data = $process->read();
//         echo "接收到： $data \n";
//     });
// }




//  进程队列通信
$workers = [];  // 进程仓库
$worker_num = 2;  // 最大进程数

// 批量创建进程 完成
for($i = 0;$i < $worker_num;$i++){
    $process = new swoole_process('doProcess',false,false); // 创建子进程完成
    $process->useQueue();   // 开启队列，类似于全局函数
    $pid = $process->start();
    $workers[$pid] = $process;
}

// 进程执行函数
function doProcess(swoole_process $process){
    $recv = $process->pop();
    echo "从主进程获取到数据：$recv \n";
    sleep(5);
    $process->exit(0);
}

// 主进程 向子进程添加数据
foreach($workers as $pid => $process){
    $process->push("hello 子进程 $pid \n");
}

// 等待子进程结束 回收资源
for($i = 0;$i < $worker_num;$i++){
    $ret = swoole_process::wait();  // 等待执行完成
    $pid = $ret['pid'];
    unset($workers[$pid]);
    echo "子进程退出 $pid \n";
}