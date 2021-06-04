#include "clientwidget.h"
#include "ui_clientwidget.h"
#include <QDebug>
#include <QMessageBox>
#include <QHostAddress>
ClientWidget::ClientWidget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::ClientWidget)
{
    ui->setupUi(this);

    setWindowTitle("客户端");

    tcpSocket = new QTcpSocket(this);

    isStart = true;

    connect(tcpSocket,&QTcpSocket::readyRead,
    [=](){
        //取出接收的数据
        QByteArray buf = tcpSocket->readAll();
        if(true == isStart){
            //接收头
            isStart = false;
            //解析头部信息 QString buf = "hello##1024"
            fileName = QString(buf).section("##",0,0);
            fileSize = QString(buf).section("##",1,1).toInt();
            recvSize = 0;

            //打开文件
            file.setFileName(fileName);
            bool isOk = file.open(QIODevice::WriteOnly);
            if(false == isOk){
                qDebug() << "WriteOnly error!";
            }

        }
        else{//文件信息
            qint64 len = file.write(buf);
            recvSize += len;
            if(recvSize == fileSize){
                file.close();
                QMessageBox::information(this,"完成","文件接收完成！");
                tcpSocket->disconnectFromHost();
                tcpSocket->close();

            }
        }
    }
    );
}

ClientWidget::~ClientWidget()
{
    delete ui;
}

void ClientWidget::on_buttonConnect_clicked()
{
    //获取服务器的IP和端口
    QString ip = ui->lineEditIP->text();
    quint16 port = ui->lineEditPort->text().toInt();
    tcpSocket->connectToHost(QHostAddress(ip),port);
}
