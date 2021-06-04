#include "mainwindow.h"
#include "ui_mainwindow.h"
#include <QSqlDatabase>
#include <QSqlError>
#include <QSqlQuery>
#include <QMessageBox>
#include <QSqlQueryModel>
#include <QDebug>
#include "domxml.h"

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);

    //开始  车辆管理页面
    on_actionCar_triggered();

    //打开数据库
    connectDB();

    //初始化数据
    initData();

    DomXML::createXML("../demo.xml");

//    QStringList list;
//    list << "二汽神龙" << "毕加索" << "39" << "1" << "39";
//    DomXML::appendXML("../demo.xml",list);

//    QStringList fList;
//    QStringList bList;
//    QStringList pList;
//    QStringList nList;
//    QStringList tList;

//    DomXML::readXML("../demo.xml",fList,bList,pList,nList,tList);
//    for(int i = 0;i < fList.size();i++){

//        QString str = QString("%1:%2 卖出了%3,单价：%4,总价：%5").arg(fList.at(i)).arg(bList.at(i)).arg(nList.at(i)).arg(pList.at(i)).arg(tList.at(i));

//        qDebug() << str.toUtf8().data();
//    }


}

MainWindow::~MainWindow()
{
    delete ui;
}

//车辆管理菜单
void MainWindow::on_actionCar_triggered()
{
    //切换到车辆管理页面
    ui->stackedWidget->setCurrentWidget(ui->car);
    ui->label->setText("车辆管理");
}

//销售统计菜单
void MainWindow::on_actionCalc_triggered()
{
    //切换到销售统计页面
    ui->stackedWidget->setCurrentWidget(ui->calc);
    ui->label->setText("销售统计");
}

//连接数据库
void MainWindow::connectDB(){

    //添加数据库
    QSqlDatabase db = QSqlDatabase::addDatabase("QMYSQL");
    db.setHostName("192.168.0.199");
    db.setUserName("root");
    db.setPassword("root");
    db.setDatabaseName("car");

    //打开数据库
    if(!db.open()){
        QMessageBox::warning(this,"error",db.lastError().text());
        return;
    }
}

//初始化数据
void MainWindow::initData(){

    QSqlQueryModel *queryModel = new QSqlQueryModel(this);  //新建模型
    queryModel->setQuery("select name from factory");   //sql语句
    ui->comboBoxFactory->setModel(queryModel);

    ui->labelLast->setText("0");    //剩余数量
    ui->lineEditTotal->setEnabled(false);   //金额
}


//厂家下拉框槽函数
void MainWindow::on_comboBoxFactory_currentIndexChanged(const QString &arg1)
{
    if(arg1 == "请选择厂家"){
        //内容清空
        ui->comboBoxBrand->clear();    //品牌下拉框清空
        ui->lineEditPrice->clear();    //报价清空
        ui->labelLast->setText("0");   //剩余数量
        ui->lineEditTotal->clear();    //金额清空
        ui->spinBox->setValue(0);      //数量选择框
        ui->spinBox->setEnabled(false);
        ui->buttonSure->setEnabled(false);
    }
    else{

        ui->comboBoxBrand->clear();
        QSqlQuery query;
        QString sql = QString("select name from brand where factory = '%1'").arg(arg1);

        //执行sql语句
        query.exec(sql);

        //获取内容
        while(query.next()){
            QString name = query.value("name").toString();
            ui->comboBoxBrand->addItem(name);
        }

        ui->spinBox->setEnabled(true);
    }
}

void MainWindow::on_comboBoxBrand_currentIndexChanged(const QString &arg1)
{
    QSqlQuery query;
    QString sql = QString("select price,last from brand where factory = '%1' and name = '%2'").arg(ui->comboBoxFactory->currentText()).arg(arg1);

    //执行sql语句
    query.exec(sql);

    //获取内容
    while(query.next()){
        //报价
        int price = query.value("price").toInt();
        //剩余数量
        int last = query.value("last").toInt();
        //更新到对应的控件
        ui->lineEditPrice->setText(QString::number(price));
        ui->labelLast->setText(QString::number(last));
    }
}


//数量选择框槽函数
void MainWindow::on_spinBox_valueChanged(int arg1)
{

    if(0 == arg1){
        ui->buttonSure->setEnabled(false);
    }
    else{
        ui->buttonSure->setEnabled(true);
    }

    QSqlQuery query;
    QString sql = QString("select sum,last from brand where factory = '%1' and name = '%2'").arg(ui->comboBoxFactory->currentText()).arg(ui->comboBoxBrand->currentText());

    //执行sql语句
    query.exec(sql);

    //获取内容
    int sum,last;
    while(query.next()){
        //报价
        sum = query.value("sum").toInt();
        //剩余数量
        last = query.value("last").toInt();
    }

    if(arg1 > last){   //如果选择的数目大于剩余数目，则中断程序
        ui->spinBox->setValue(last);
        return;
    }

    //更新剩余数量
    int tempNum = last - arg1;
    ui->labelLast->setText(QString::number(tempNum));


    //金额
    int price = ui->lineEditPrice->text().toInt();  //报价
    int total = price * arg1;
    ui->lineEditTotal->setText(QString::number(total));
}

void MainWindow::on_buttonCancel_clicked()
{
    ui->comboBoxFactory->setCurrentIndex(0);
    ui->labelLast->setText("0");
}

void MainWindow::on_buttonSure_clicked()
{
    //获取信息
    //1)销售数据
    int num = ui->spinBox->value();
    //2) 剩余
    int last = ui->labelLast->text().toInt();

    //获取数据库的销量
    QSqlQuery query;
    QString sql = QString("select sell from brand where factory = '%1' and name = '%2'").arg(ui->comboBoxFactory->currentText()).arg(ui->comboBoxBrand->currentText());

    //执行sql语句
    query.exec(sql);

    //获取内容
    int sell;
    while(query.next()){
        //报价
        sell = query.value("sell").toInt();
    }

    //更新数据库，剩余数量，销售总量
    sell += num;
    sql = QString("update brand set sell = %1,last = %2 where factory = '%3' and name = '%4'")
            .arg(sell)
            .arg(last)
            .arg(ui->comboBoxFactory->currentText())
            .arg(ui->comboBoxBrand->currentText());

    query.exec(sql);

    //把确认后的数据更新到xml中
    QStringList list;
    list << ui->comboBoxFactory->currentText()
         << ui->comboBoxBrand->currentText()
         << ui->lineEditPrice->text()
         << QString::number(num)
         << ui->lineEditTotal->text();

    DomXML::appendXML("../demo.xml",list);

    QStringList fList;
    QStringList bList;
    QStringList pList;
    QStringList nList;
    QStringList tList;

    DomXML::readXML("../demo.xml",fList,bList,pList,nList,tList);
    for(int i = 0;i < fList.size();i++){

        QString str = QString("%1:%2 卖出了%3,单价：%4,总价：%5").arg(fList.at(i)).arg(bList.at(i)).arg(nList.at(i)).arg(pList.at(i)).arg(tList.at(i));
        ui->textEdit->append(str);

        qDebug() << str.toUtf8().data();
    }

    ui->buttonSure->setEnabled(false);
    on_buttonCancel_clicked();
}
