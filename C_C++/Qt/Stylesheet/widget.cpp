#include "widget.h"
#include "ui_widget.h"

Widget::Widget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::Widget)
{
    ui->setupUi(this);
    ui->label->setStyleSheet("QLabel{color:rgb(0,255,255);background-color:blue;background-image:url(:/new/prefix1/headimg.gif)}");
    ui->pushButton->setStyleSheet("QPushButton{border:2px outset green;border-radius:15px;}QPushButton:pressed{color:blue;border:2px outset blue;}");
}

Widget::~Widget()
{
    delete ui;
}
