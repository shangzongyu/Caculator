#include <QApplication>
#include <QMessageBox>
#include <QDir>
#include <QFile>
#include <QDebug>
#include "Ui.h"

int main(int argc, char *argv[])
{
    QApplication app(argc, argv);
    QApplication::addLibraryPath("./plugins");

    // 中文文本使用 UTF-8，Qt 6 会按 UTF-8 读取字符串，无需设置 GB2312 来解决中文乱码。
    QFile file(":my.qss");
    if (file.open(QFile::ReadOnly))
    {
        qApp->setStyleSheet(file.readAll());
        file.close();
    }
    else
    {
        QMessageBox::warning(NULL, "警告", "加载  <font color='red'>my.qss</font> 失败");
        exit(-1);
    }

    Ui ui;
    ui.show();

    return app.exec();
}
