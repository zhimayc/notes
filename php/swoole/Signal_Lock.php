<?php

// 触发函数 异步执行 达到10停止
// swoole_process::signal(SIGALRM,function($signo){
//     static $i = 0;
//     echo "$i  signo : $signo \n";
//     $i++;
//     if($i > 10){
//         swoole_process::alarm(-1);  // 清除定时器
//     }
// });

// 定时信号
// swoole_process::alarm(1000*1000);    // 1秒



// 锁机制

// 创建锁对象
$lock = new swoole_lock(SWOOLE_MUTEX);  // 互斥锁
echo "创建互斥锁 \n";
$lock->lock(); // 开始锁定  主进程
if(pcntl_fork() > 0){
    sleep(1);
    $lock->unlock();  // 解锁
}
else {
    echo "子进程 等待锁 \n";
    $lock->lock();  // 上锁
    echo "子进程 获取锁 \n";
    $lock->unlock();    // 释放锁
    exit("子进程退出");
}

echo "主进程 释放锁 \n";
unset($lock);
sleep(1);
echo "子进程退出";
