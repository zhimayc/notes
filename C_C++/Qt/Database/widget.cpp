#include "widget.h"
#include "ui_widget.h"
#include <QSqlDatabase>
#include <QDebug>
#include <QMessageBox>
#include <QSqlError>
#include <QSqlQuery>
#include <QVariantList>

Widget::Widget(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::Widget)
{
    ui->setupUi(this);

    //打印Qt支持的数据库驱动
    qDebug() << QSqlDatabase::drivers();

    //添加MySQL数据库
    QSqlDatabase db = QSqlDatabase::addDatabase("QMYSQL");
    //连接数据库
    db.setHostName("192.168.1.199");  //数据库服务器ip
    db.setUserName("root");   //数据库用户名
    db.setPassword("root");    //数据库密码
    db.setDatabaseName("test");  //使用哪个数据库


    /* windows下使用mysql，需要把libmyql.dll 文件放入qt\5.11.2\mingw53_32\bin 类似目录下 */
    //打开数据库
    if(!db.open()){   //数据库打开失败
        QMessageBox::warning(this,"错误",db.lastError().text());
    }

    QSqlQuery query;
//    query.exec("create table student(id int primary key auto_increment,name varchar(255),age int,score int);");


    /*
     *      数据库连接 连接名用法
    QSqlDatabase db1 = QSqlDatabase::addDatabase("QMYSQL","db1name");
    //连接数据库
    db1.setHostName("192.168.0.199");  //数据库服务器ip
    db1.setUserName("root");   //数据库用户名
    db1.setPassword("root");    //数据库密码
    db1.setDatabaseName("test1");  //使用哪个数据库

    //打开数据库
    if(!db1.open()){   //数据库打开失败
        QMessageBox::warning(this,"错误",db1.lastError().text());
    }

    QSqlQuery query1(db1);
    query1.exec("create table student(id int primary key auto_increment,name varchar(255),age int,score int);");

    */


    //插入
//    QSqlQuery query;
//    query.exec("insert into student(name,age,score)values('tom',18,120);");

//    //批量插入
//    //odbc风格
//    //预处理语句
    //占位符 ?
//    query.prepare("insert into student(name,age,score)values(?,?,?)");
//    //给字段设置内容 list
//    QVariantList nameList,ageList,scoreList;
//    nameList << "jack" << "mike" << "charles" << "maria" << "lucy";
//    ageList << 15 << 15 << 16 << 18 << 20;
//    scoreList << 60 << 70 << 80 << 90 << 100;

//    //给字段绑定相应的值 按顺序绑定
//    query.addBindValue(nameList);
//    query.addBindValue(ageList);
//    query.addBindValue(scoreList);

//    //执行预处理命令
//    query.execBatch();



    //oracle 风格
    //占位符 : + 加自定义名字
//    query.prepare("insert into student(name,age,score)values(:name,:age,:score)");
//    //给字段设置内容 list
//    QVariantList nameList,ageList,scoreList;
//    nameList << "jane" << "jenny" << "charles" << "maria" << "lucy";
//    ageList << 18 << 20 << 16 << 17 << 18;
//    scoreList << 80 << 70 << 85 << 90 << 120;

//    //给字段绑定（不用按顺序绑定）
//    query.bindValue(":name",nameList);
//    query.bindValue(":age",ageList);
//    query.bindValue(":score",scoreList);

//    //执行预处理命令
//    query.execBatch();

    query.exec("select * from student where name='charles'");
    while(query.next()){    //一行一行遍历
        //取出当前行的内容
        qDebug() << query.value(0).toInt() << query.value(1).toString() << query.value("age").toInt() << query.value("score").toInt();
    }



}

Widget::~Widget()
{
    delete ui;
}

void Widget::on_buttonDel_clicked()
{
    //获取行编辑内容
    QString name = ui->lineEdit->text();

    QString sql = QString("delete from student where name = '%1'").arg(name);

    //开启一个事务
    QSqlDatabase::database().transaction();
    QSqlQuery query;
    query.exec(sql);

    qDebug() << "删除";


}

void Widget::on_buttonSure_clicked()
{
    //确定删除
    QSqlDatabase::database().commit();
     qDebug() << "确定删除";
}

void Widget::on_buttonCancel_clicked()
{
    //回滚，撤销
     QSqlDatabase::database().rollback();
     qDebug() << "撤销";
}
