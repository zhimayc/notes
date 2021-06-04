#include "mybutton.h"
#include <QMouseEvent>
#include <QDebug>

MyButton::MyButton(QWidget *parent) : QPushButton(parent)
{

}

void MyButton::mousePressEvent(QMouseEvent *e){
    if(e->button() == Qt::LeftButton){
        qDebug() << "按下的是左键";
        e->ignore();
    }else{
       QPushButton::mousePressEvent(e);
    }
}
