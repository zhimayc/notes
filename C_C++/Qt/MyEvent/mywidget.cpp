#include "mywidget.h"
#include "ui_mywidget.h"
#include <QDebug>
#include <QKeyEvent>
#include <QCloseEvent>
#include <QMessageBox>
#include <QEvent>

MyWidget::MyWidget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::MyWidget)
{
    ui->setupUi(this);
    timerId = this->startTimer(1000);     //毫秒为单位  每隔1s触发一次定时器
    connect(ui->pushButton,&MyButton::clicked,[=](){
        qDebug() << "按钮被按下";
    });

    //给label_2控件安装过滤器
    ui->label_2->installEventFilter(this);


    ui->label_2->setMouseTracking(this);

}

MyWidget::~MyWidget()
{
    delete ui;
}

void MyWidget::keyPressEvent(QKeyEvent *e){
    qDebug() << (char)e->key();
}

void MyWidget::timerEvent(QTimerEvent *e){
    static int sec = 0;
    ui->label->setText(QString("<center><h1>timer out: %1</h1></center>").arg(sec++));
    if(sec == 30){
        this->killTimer(this->timerId);
    }
}

void MyWidget::mousePressEvent(QMouseEvent *e){
    qDebug() << "+++++++++++++++++++++++++++++";
}

void MyWidget::closeEvent(QCloseEvent *e){
    int ret = QMessageBox::question(this,"question","是否确认关闭窗口");
    if(ret == QMessageBox::Yes){
        //关闭窗口
        //处理关闭窗口事件，接收事件，事件就不会再往下传递
        e->accept();

    }else{
        //不关闭窗口
        //忽略事件，事件继续给父组件传递
        e->ignore();
    }
}

bool MyWidget::event(QEvent *e){
    //事件分发
//    switch(e->type()){
//        case QEvent::Close:
//            closeEvent(e);
//            break;
//        case QEvent::MouseMove:
//            mouseMoveEvent(e);
//            break;
//    }

    //如果传入的事件已被识别并且处理，则需要返回true,否则返回false,如果返回值是true，那么Qt会认为这个事件已经处理完毕，不会将这个事件
    //发送给其他对象，而是会继续处理事件队列中的下一个事件。

    if(e->type() == QEvent::Timer){
        //定时器事件处理
        QTimerEvent *env = static_cast<QTimerEvent *>(e);
        timerEvent(env);
        return true;
    }
    else if(e->type() == QEvent::KeyPress){

        QKeyEvent *env = static_cast<QKeyEvent *>(e);
        if(env->key() == Qt::Key_B){
            return QWidget::event(e);
        }
        return true;

    }
    else{
        return QWidget::event(e);
    }
}

//事件过滤器

bool MyWidget::eventFilter(QObject * obj,QEvent * e){
    if(obj == ui->label_2){
        //判断事件类型
        if(e->type() == QEvent::MouseMove){
            QMouseEvent * env = static_cast<QMouseEvent *>(e);
            ui->label_2->setText(QString("Mouse Move: (%1,%2)").arg(env->x()).arg(env->y()));

            return true;
        }
        else{
             return QWidget::eventFilter(obj,e);
        }
    }
    else{
        return QWidget::eventFilter(obj,e);
    }
}
