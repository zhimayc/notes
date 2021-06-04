#include "widget.h"
#include "ui_widget.h"
#include <QHostAddress>

Widget::Widget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::Widget)
{
    ui->setupUi(this);

    //分配空间，指定父对象
    udpSocket = new QUdpSocket(this);

    //绑定
//    udpSocket->bind(8888);
    udpSocket->bind(QHostAddress::AnyIPv4,8888);
    setWindowTitle("服务器端口为8888");

    //加入某个组播
    //组播地址是D类地址
    udpSocket->joinMulticastGroup(QHostAddress("224.0.0.2"));

    //退出组播

//    udpSocket->leaveMulticastGroup(QHostAddress("224.0.0.2"));

    //当对方成功发送数据过来，自动触发readyRead信号
    connect(udpSocket,&QUdpSocket::readyRead,this,&Widget::dealMsg);
}


void Widget::dealMsg(){

    //读取对方发送的内容
    char buf[1024] = {0};
    QHostAddress clientAddr;    //对方IP地址
    quint16 port;    //对方端口
    qint64 len = udpSocket->readDatagram(buf,sizeof(buf),&clientAddr,&port);

    if(len > 0){

        //格式化
        QString str = QString("[%1:%2] %3").arg(clientAddr.toString()).arg(port).arg(buf);

        //给编辑器设置内容
        ui->textEdit->setText(str);

    }
}

Widget::~Widget()
{
    delete ui;
}

void Widget::on_ButtonSend_clicked()
{
    //先获取对方的ip和端口
    QString ip = ui->lineEditIP->text();
    qint16 port = ui->lineEditPort->text().toInt();

    //获取编辑区内容
    QString str = ui->textEdit->toPlainText();

    //给指定的ip发送数据
    udpSocket->writeDatagram(str.toUtf8(),QHostAddress(ip),port);
}
