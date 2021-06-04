#include "mainwindow.h"
#include "ui_mainwindow.h"
#include <QDebug>
#include <QCompleter>
#include <QStringList>

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);
}

MainWindow::~MainWindow()
{
    delete ui;
}
