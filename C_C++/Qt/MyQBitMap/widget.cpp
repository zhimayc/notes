#include "widget.h"
#include "ui_widget.h"
#include <QPainter>
#include <QBitmap>
#include <QImage>
#include <QPicture>

Widget::Widget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::Widget)
{
    ui->setupUi(this);

    //绘图设备 ， 100*100
    QPixmap pixmap(100,100);
    QPainter p(&pixmap);

    //填充背景色
    p.fillRect(0,0,100,100,QBrush(Qt::green));

    p.drawPixmap(0,0,80,80,QPixmap("../Image/2.jpg"));



    //保存图片
    pixmap.save("../pixmap.jpg");



    //QImage 绘图设备

    QImage image(200,200,QImage::Format_ARGB32);
    QPainter p1(&image);

    //绘图
    p1.drawImage(0,0,QImage("../Image/1.png"));

    //对绘图设备前50个像素点进行操作
    for(int i = 0; i < 50;i ++){

        for(int j = 0; j < 50;j ++){
            image.setPixel(QPoint(i,j),qRgb(0,255,0));
        }

    }



    image.save("../qimage.jpg");



    //QPicture绘图设备

    QPicture pic;
    QPainter p2;

    p2.begin(&pic);

    p2.drawPixmap(0,0,80,80,QPixmap("../Image/3.jpg"));
    p2.drawLine(50,50,150,50);

    p2.end();

    pic.save("../pic.png");




    //QImage 和 QPixmap相互转换

    QPainter p3(this);
    QPixmap pixmap2;
    pixmap2.load("../Image/5.jpg");

    //QPixmap => QImage
    QImage tempImage = pixmap2.toImage();
    p3.drawImage(0,0,tempImage);


    QImage image2;
    image2.load("../Image/2.jpg");

    //QImage => QPixmap
    QPixmap tempPixmap = QPixmap::fromImage(image2);
    p3.drawPixmap(100,0,tempPixmap);




}

Widget::~Widget()
{
    delete ui;
}


/*
 *  绘图设备
 *  QPixmap：针对屏幕进行优化了，和平台相关，不能对图片进行修改
 *  QImage: 和平台无关，可以对图片进行修改，在线程中绘图
 *  QPicture: 保存绘图的状态（二进制文件）
*/

void Widget::paintEvent(QPaintEvent *){
    QPainter p(this);

    //QPixmap
    p.drawPixmap(0,0,QPixmap("../Image/headimg.gif"));

    //QBitmap
    p.drawPixmap(200,0,QBitmap("../Image/headimg.gif"));

    //QPicture
    QPicture pic;
    pic.load("../pic.png");
    p.drawPicture(100,50,pic);

}
