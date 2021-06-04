#include "mainwidget.h"
#include <QPushButton>
#include <QDebug>

MainWidget::MainWidget(QWidget *parent)
    : QWidget(parent)
{
    QPushButton b;
    b1.setParent(this);
    b1.setText(QString("确定"));
    b1.move(100,100);
    b2 = new QPushButton(this);
    b2->setText("取消");

    connect(&b1,&QPushButton::pressed,this,&MainWidget::close);
    connect(b2,&QPushButton::released,this,&MainWidget::mySlot);
    connect(b2,&QPushButton::released,&b1,&QPushButton::hide);

    setWindowTitle("父窗口");
    b3.setParent(this);
    b3.setText("切换到子窗口");
    b3.move(100,100);
    resize(400,300);

    connect(&b3,&QPushButton::released,this,&MainWidget::changeWin);

    //处理子窗口的信号  Qt5写法
//    void (SubWidget::*funSignal)() = &SubWidget::mySignal;
//    connect(&w,funSignal,this,&MainWidget::dealSub);

//    void (SubWidget::*testSignal)(int,QString) = &SubWidget::mySignal;
//    connect(&w,testSignal,this,&MainWidget::dealSlot);


    //Qt4写法
    connect(&w,SIGNAL(mySignal()),this,SLOT(dealSub()));
    connect(&w,SIGNAL(mySignal(int,QString)),this,SLOT(dealSlot(int,QString)));

    QPushButton *b4 = new QPushButton(this);
    b4->setText("Lambda表达式");
    b4->move(150,150);
    int x = 10,y = 100;
    connect(b4,&QPushButton::released,
            [=]() mutable{
                b4->setText("3242342");
                qDebug() << "11111111111";
                qDebug() << x << y;
                x = 1000;
            }


            );


}

void MainWidget::mySlot(){
    b2->setText(QString("123"));
}

void MainWidget::changeWin(){
    w.show();
    this->hide();
}

void MainWidget::dealSub(){
    w.hide();
    this->show();
}

void MainWidget::dealSlot(int a,QString str){
    qDebug() << a << str.toUtf8().data();
}

MainWidget::~MainWidget()
{

}
