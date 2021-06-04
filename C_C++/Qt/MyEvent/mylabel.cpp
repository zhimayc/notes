#include "mylabel.h"
#include <QMouseEvent>
#include <QDebug>

MyLabel::MyLabel(QWidget *parent) : QLabel(parent)
{
    //设置追踪鼠标
    this->setMouseTracking(true);

}

void MyLabel::mousePressEvent(QMouseEvent *ev){
    int i = ev->x();
    int j = ev->y();
    QString text = QString("<center><h1>Mouse press:(%1,%2)</h1></center>").arg(i).arg(j);
    this->setText(text);

    if(ev->button() == Qt::LeftButton){
        qDebug() << "left";
    }else if(ev->button() == Qt::MidButton){
        qDebug() << "Mid";
    }else if(ev->button() == Qt::RightButton){
        qDebug() << "Right";
    }
}

void MyLabel::mouseReleaseEvent(QMouseEvent *ev){
    QString text = QString("<center><h1>Mouse Release:(%1,%2)</h1></center>").arg(ev->x()).arg(ev->y());
    this->setText(text);
}

void MyLabel::mouseMoveEvent(QMouseEvent *ev){
    QString text = QString("<center><h1>Mouse Move:(%1,%2)</h1></center>").arg(ev->x()).arg(ev->y());
    this->setText(text);
}

void MyLabel::enterEvent(QEvent *e){
    qDebug() << "enter";
}

void MyLabel::leaveEvent(QEvent *e){
    qDebug() << "leave";
}
