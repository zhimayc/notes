#include "mainwindow.h"
#include <QPushButton>
#include <QMenuBar>
#include <QMenu>
#include <QDebug>
#include <QToolBar>
#include <QStatusBar>
#include <QLabel>
#include <QTextEdit>
#include <QDockWidget>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
{
    //菜单栏
    QMenuBar *mBar = menuBar();

    //添加菜单
    QMenu *pFile = mBar->addMenu("文件");

    //添加菜单项
    QAction *pNew = pFile->addAction("新建");
    pFile->addSeparator();
    QAction *pOpen = pFile->addAction("打开");

    //工具栏、菜单项的快捷方式
    QToolBar *toolBar = addToolBar("toolBar");

    //工具栏添加快捷键
    toolBar->addAction(pNew);

    QPushButton *b = new QPushButton(this);

    //添加小控件
    toolBar->addWidget(b);
    connect(b,&QPushButton::clicked,[=](){
         b->setText("123");
    });

    //状态栏
    QStatusBar *sb = statusBar();
    QLabel *label = new QLabel(this);
    label->setText("Normal text file");

    //从左往右添加
    sb->addWidget(label);

    //从右往左添加
    sb->addPermanentWidget(new QLabel("2",this));
    sb->addPermanentWidget(new QLabel("3",this));

    //核心控件
    QTextEdit *textEdit = new QTextEdit(this);
    setCentralWidget(textEdit);

    //浮动窗口
    QDockWidget *dock = new QDockWidget(this);
    addDockWidget(Qt::RightDockWidgetArea,dock);

    QTextEdit *textEdit1 = new QTextEdit(this);
    dock->setWidget(textEdit1);




}

MainWindow::~MainWindow()
{

}
