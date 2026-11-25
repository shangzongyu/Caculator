/*******************************************************************************
 * 类名: Ui
 * 功能: 实现计算器界面显示
 ******************************************************************************/

#ifndef UI_H
#define UI_H

#include <QFrame>
#include <QLineEdit>

#include "Button.h"
#include "TitleBar.h"
#include "AlgorithmStrategy.h"

class Ui : public QFrame
{
    Q_OBJECT
public:
    explicit Ui(QWidget *parent = 0);

private:
    void init();
    void layout();
    void initShortcuts();
    void initWindowControls();
    void selectOperator(EAlgorithmType operation);
    bool commitOperand();
    void startEntry();

private slots:
    void numButtonClicked();
    void clearButtonClicked();
    void dotButtonClicked();
    void delButtonClicked();
    void equalButtonClicked();

private:
    enum class State { Entering, AwaitingOperand, Result, Error };

    QLineEdit m_display;
    Button m_clearButton; // 清除按钮
    Button m_delButton;   // 删除按钮
    Button m_addButton;   // 加法按钮
    Button m_subButton;   // 减法按钮
    Button m_mulButton;   // 乘法按钮
    Button m_divButton;   // 除法按钮
    Button m_dotButton;   // 小数点按钮
    Button m_equalButton; // 等号按钮
    Button m_numButtons[10]; // 数字按钮集合

    TitleBar m_titleBar;

    State m_state = State::Entering;
    EAlgorithmType m_operation = EAlgorithmType_Unknown; // 算法类型设置为未知
    double m_accumulator = 0.0; // 左操作数设置默认值
};

#endif // UI_H
