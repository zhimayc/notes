#ifndef MYWIDGET_H
#define MYWIDGET_H

#include <QWidget>

namespace Ui {
class MyWidget;
}

class MyWidget : public QWidget
{
    Q_OBJECT

public:
    explicit MyWidget(QWidget *parent = nullptr);
    ~MyWidget();

protected:
    //键盘按下事件
    void keyPressEvent(QKeyEvent *);
    //计时器事件
    void timerEvent(QTimerEvent *);
    //鼠标点击事件
    void mousePressEvent(QMouseEvent *);
    //窗口关闭事件
    void closeEvent(QCloseEvent *e);
    //事件分发函数
    bool event(QEvent *);
    //事件过滤器
    bool eventFilter(QObject * obj,QEvent * e);

private:
    Ui::MyWidget *ui;
    int timerId;
};

#endif // MYWIDGET_H
