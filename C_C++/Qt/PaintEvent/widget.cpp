#include "widget.h"
#include "ui_widget.h"
#include <QPainter>
#include <QPen>
#include <QBrush>

Widget::Widget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::Widget)
{
    ui->setupUi(this);

    x = 0;
}

Widget::~Widget()
{
    delete ui;
}

void Widget::paintEvent(QPaintEvent *){
//    QPainter p(this);
    QPainter p;  //创建画家对象
    p.begin(this);  //指定当前窗口为绘图设备

    //绘图操作

    //画背景图
//    p.drawPixmap(0,0,width(),height(),QPixmap("../Image/2.jpg"));
      p.drawPixmap(rect(),QPixmap("../Image/4.jpg"));

      //定义画笔
      QPen pen;
      pen.setWidth(5);

      //设置颜色
//      pen.setColor(Qt::red);
      pen.setColor(QColor(14,9,234));   //设置RGB颜色
      pen.setStyle(Qt::DashLine);      //设置风格


      //设置画笔
      p.setPen(pen);

      //画直线
      p.drawLine(50,50,150,50);
      p.drawLine(50,50,50,150);

      //创建画刷对象
      QBrush brush;
      brush.setColor(Qt::green);
      brush.setStyle(Qt::Dense1Pattern);

      //设置画刷
      p.setBrush(brush);

      //画矩形
      p.drawRect(150,150,100,50);

      //画圆形（椭圆）
      p.drawEllipse(QPoint(150,150),50,25);

      //画笑脸
      p.drawPixmap(x,180,80,80,QPixmap("../Image/5.jpg"));

    p.end();
}

void Widget::on_pushButton_clicked()
{
    x += 20;
    if(x > width()){
        x = 0;
    }

    //刷新窗口，让窗口重绘
    update();       //间接调用paintEvent()
}
