#include "serverwidget.h"
#include "ui_serverwidget.h"
#include <QFileDialog>
#include <QDebug>
#include <QFileInfo>

ServerWidget::ServerWidget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::ServerWidget)
{
    ui->setupUi(this);

    //监听套接字
    tcpServer = new QTcpServer(this);

    //监听
    tcpServer->listen(QHostAddress::Any,8888);
    setWindowTitle("服务器端口为:8888");

    ui->buttonFile->setEnabled(false);
    ui->buttonSend->setEnabled(false);

    //如果客户端成功和服务器连接
    //tcpServer会自动触发newConnection()
    connect(tcpServer,&QTcpServer::newConnection,
       [=](){
            //取出建立好连接的套接字
            tcpSocket = tcpServer->nextPendingConnection();
            //获取对方的ip和端口
            QString ip = tcpSocket->peerAddress().toString();
            quint16 port = tcpSocket->peerPort();
            QString str = QString("[%1:%2] 成功连接！").arg(ip).arg(port);
            ui->textEdit->setText(str);  //显示到编辑区

            //成功连接后，才能选择文件
            ui->buttonFile->setEnabled(true);
            ui->buttonSend->setEnabled(false);
       }
    );

    connect(&timer,&QTimer::timeout,
    [=](){
        //关闭定时器
        timer.stop();
        //发送文件
        sendData();
    }
    );

}

ServerWidget::~ServerWidget()
{
    delete ui;
}

//选择文件按钮
void ServerWidget::on_buttonFile_clicked()
{
    QString filePath = QFileDialog::getOpenFileName(this,"open","../");

    if(false == filePath.isEmpty()){    //如果选择文件路径有效

        fileName.clear();
        fileSize = 0;

        //获取文件信息
        QFileInfo info(filePath);
        fileName = info.fileName(); //获取文件名
        fileSize = info.size();     //获取文件大小

        //只读方式打开文件
        //指定文件的名字
        file.setFileName(filePath);

        //打开文件
        bool isOk = file.open(QIODevice::ReadOnly);
        if(false == isOk){
            qDebug() << "只读方式打开文件失败！";
        }

        //提示打开文件的路径
        ui->textEdit->append(filePath);

        ui->buttonFile->setEnabled(false);
        ui->buttonSend->setEnabled(true);

    }else{
        qDebug() << "选择文件路径出错！";
    }
}

void ServerWidget::on_buttonSend_clicked()
{
    //先发送文件头信息  文件名##文件大小
    QString head = QString("%1##%2").arg(fileName).arg(fileSize);

    //发送头部信息
    qint64 len = tcpSocket->write(head.toUtf8());
    if(len > 0){    //头部信息发送成功
        //发送真正的文件信息
        //防止TCP黏包文件
        //需要通过定时器延时
        timer.start(20);
    }
    else{

        qDebug() << "头部信息发送失败！";
        file.close();
        ui->buttonFile->setEnabled(true);
        ui->buttonSend->setEnabled(false);

    }

}

void ServerWidget::sendData(){
    qint64 len = 0;
    sendSize = 0;
    do{
        //每次发送数据的大小

        char buf[4*1024] = {0};
        len = 0;
        len = file.read(buf,sizeof(buf));

        //发送数据
        len = tcpSocket->write(buf,len);

        //发送的数据需要累加
        sendSize += len;

    }while(len > 0);

    //是否发送文件完毕
    if(sendSize == fileSize){

         qDebug() << "文件发送完毕......";
        ui->textEdit->append("文件发送完毕");
        file.close();

        //把客户端断开
        tcpSocket->disconnectFromHost();
        tcpSocket->close();
    }

}
