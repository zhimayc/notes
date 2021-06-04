#include "widget.h"
#include "ui_widget.h"
#include <QFile>
#include <QFileDialog>
#include <QFileInfo>
#include <QDateTime>
#include <QDebug>

Widget::Widget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::Widget)
{
    ui->setupUi(this);
}

Widget::~Widget()
{
    delete ui;
}


//读文件
void Widget::on_buttonRead_clicked()
{
    QString path = QFileDialog::getOpenFileName(this,"open","../","TXT(*.txt)");
    if(path.isEmpty() == false){
        //文件对象
        QFile file(path);

        //打开文件，只读方式
        bool isOk = file.open(QIODevice::ReadOnly);

        if(isOk == true){

#if 0
            //读文件，默认只识别utf8编码
            QByteArray array = file.readAll();
            //显示到编辑区
            ui->textEdit->setText(QString(array));
#endif
            QByteArray array;
            while(file.atEnd() == false){

                //读一行
                array += file.readLine();

            }

            ui->textEdit->setText(QString(array));
        }

        //关闭文件
        file.close();

        //获取文件信息
        QFileInfo info(path);
        qDebug() << "文件名字：" << info.fileName();
        qDebug() << "文件后缀名：" << info.suffix();
        qDebug() << "文件大小：" << info.size();
        qDebug() << "文件创建时间：" << info.created().toString("yyyy-MM-dd hh:mm:ss");
    }
}


//写文件
void Widget::on_pushButton_2_clicked()
{
    QString path = QFileDialog::getSaveFileName(this,"save","../","TXT(*.txt)");
    if(path.isEmpty() == false){
        //创建文件对象
        QFile file;
        file.setFileName(path);

        //打开文件，只写方式
        bool isOk = file.open(QIODevice::WriteOnly);
        if(isOk == true){

            //获取编辑区内容
            QString str = ui->textEdit->toPlainText();

            //写文件
//            file.write(str.toUtf8());
            file.write(str.toStdString().data());  //c++标准 char * 方式
        }

        file.close();
    }
}
