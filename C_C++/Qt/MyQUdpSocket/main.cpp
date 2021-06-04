#include "widget.h"
#include "clientwidget.h"
#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    Widget w;
    w.show();

    ClientWidget w2;
    w2.show();

    return a.exec();
}
