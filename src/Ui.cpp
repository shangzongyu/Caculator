#include "Ui.h"

#include "AlgorithmContext.h"

#include <QGridLayout>
#include <QShortcut>
#include <cmath>

namespace {
constexpr int maxInputLength = 19; // 操作数输入的最大长度为 19，计算结果不受此限制
constexpr int displayPrecision = 15;
}

/*******************************************************************************
 * 函数名称： Ui
 * 功能描述： Ui构造函数
 * 参数列表： parent - 父窗口指针
 *
 * 返回结果： 无
 ******************************************************************************/
Ui::Ui(QWidget *parent) :
    QFrame(parent)
{
    init();
    layout();
    initShortcuts();
}

/*******************************************************************************
 * 函数名称： init
 * 功能描述： 初始化
 * 参数列表： void
 *
 * 返回结果： void
 ******************************************************************************/
void Ui::init()
{
    setWindowFlags(Qt::FramelessWindowHint); // 去掉窗体的边框
    setWindowTitle(tr("计算器"));
    setWindowIcon(QIcon(":/image/1.png")); // 设置状态栏图片
    setFixedSize(324, 576); // 设置窗体大小

    m_titleBar.setFixedSize(324, 20); // 设置可拖拽区域大小
    m_titleBar.setObjectName("titleBar_Label"); // 设置可拖拽区域类名
    m_titleBar.setParent(this); // 设置可拖拽区域父窗体
    initWindowControls();

    // 初始化数字按钮
    QString strNum = "";
    for (int i = 0; i < 10; ++i)
    {
        strNum = ("Num" + QString::number(i));
        m_numButtons[i].setObjectName(strNum); // 设置数字按钮的类名
        m_numButtons[i].setFixedSize(60, 75);  // 设置数字按钮的大小
        m_numButtons[i].setName(QString::number(i));
        connect(&m_numButtons[i], SIGNAL(clicked()), this, SLOT(numButtonClicked())); // 设置数字按钮和相应的槽函数关联
    }

    m_display.setFixedSize(324, 90); // 设置结果显示窗体的大小(是一个QLineEdit)
    m_display.setObjectName("result_Lable"); // 设置结果显示窗体的类名
    m_display.setParent(this); // 设置结果显示窗体的父窗体
    m_display.setText("0"); // 结果显示窗体的默认内容是为0
    m_display.setAlignment(Qt::AlignVCenter | Qt::AlignRight); // 结果显示窗体的文本设置向右对齐,上下居中
    m_display.setReadOnly(true); // 结果显示窗体的文本设置为只读

    m_clearButton.setFixedSize(210,75); // 设置清除按钮大小
    m_clearButton.setObjectName("clear_btn"); // 设置清除按钮的类名
    connect(&m_clearButton, SIGNAL(clicked()), this, SLOT(clearButtonClicked())); // 设置清除按钮和相应的槽函数关联

    m_delButton.setFixedSize(90,75); // 设置删除按钮大小
    m_delButton.setObjectName("del_btn"); // 设置删除按钮的类名
    connect(&m_delButton, SIGNAL(clicked()), this, SLOT(delButtonClicked())); // 设置删除按钮和相应的槽函数关联

    m_addButton.setFixedSize(90,75); // 设置加法按钮大小
    m_addButton.setObjectName("add_btn");  // 设置加法按钮的类名
    connect(&m_addButton, &Button::clicked, this, [this] { selectOperator(EAlgorithmType_Add); }); // 设置加法按钮和相应的处理函数关联

    m_subButton.setFixedSize(90,75); // 设置减法按钮大小
    m_subButton.setObjectName("sub_btn"); // 设置减法按钮的类名
    connect(&m_subButton, &Button::clicked, this, [this] { selectOperator(EAlgorithmType_Sub); }); // 设置减法按钮和相应的处理函数关联


    m_mulButton.setFixedSize(90,75); // 设置乘法按钮大小
    m_mulButton.setObjectName("mul_btn"); // 设置乘法按钮的类名
    connect(&m_mulButton, &Button::clicked, this, [this] { selectOperator(EAlgorithmType_Mul); }); // 设置乘法按钮和相应的处理函数关联

    m_divButton.setFixedSize(90,75); // 设置除法按钮大小
    m_divButton.setObjectName("div_btn"); // 设置除法按钮的类名
    connect(&m_divButton, &Button::clicked, this, [this] { selectOperator(EAlgorithmType_Div); }); // 设置除法按钮和相应的处理函数关联

    m_dotButton.setFixedSize(60, 75); // 设置小数点按钮大小
    m_dotButton.setObjectName("dot_btn"); // 设置小数点按钮的类名
    connect(&m_dotButton, SIGNAL(clicked()), this, SLOT(dotButtonClicked())); // 设置小数点按钮和相应的槽函数关联

    m_equalButton.setFixedSize(60, 75); // 设置等号按钮大小
    m_equalButton.setObjectName("equal_btn"); // 设置等号按钮的类名
    connect(&m_equalButton, SIGNAL(clicked()), this, SLOT(equalButtonClicked())); // 设置等号按钮和相应的槽函数关联

}

/*******************************************************************************
 * 函数名称： layout
 * 功能描述： 布局
 * 参数列表： void
 *
 * 返回结果： void
 ******************************************************************************/
void Ui::layout()
{
    m_titleBar.move(0, 0); // 可拖拽窗体移动到相对于父窗体的(0, 0)位置
    m_display.move(0, 20); // 结果显示窗体移动到相对于父窗体的(0, 20)位置

    QGridLayout *pGridLayout = new QGridLayout; // 创建网格布局
    Q_CHECK_PTR(pGridLayout);

    pGridLayout->setContentsMargins(0, 0, 0, 0);

    // 第一行,把清除按按钮和删除按钮加到网格布局里
    pGridLayout->addWidget(&m_clearButton, 0, 0, 1, 3);
    pGridLayout->addWidget(&m_delButton,   0, 3, 1, 1);

    // 第二行,把数字按钮按按钮7,8,9 和除法按钮加到网格布局里
    pGridLayout->addWidget(&m_numButtons[7], 1, 0, 1, 1);
    pGridLayout->addWidget(&m_numButtons[8], 1, 1, 1, 1);
    pGridLayout->addWidget(&m_numButtons[9], 1, 2, 1, 1);
    pGridLayout->addWidget(&m_divButton,     1, 3, 1, 1);

    // 第三行,把数字按钮按按钮4,5,6和乘法按钮加到网格布局里
    pGridLayout->addWidget(&m_numButtons[4], 2, 0, 1, 1);
    pGridLayout->addWidget(&m_numButtons[5], 2, 1, 1, 1);
    pGridLayout->addWidget(&m_numButtons[6], 2, 2, 1, 1);
    pGridLayout->addWidget(&m_mulButton,     2, 3, 1, 1);

    // 第四行,把数字按钮按按钮1,2,3和减法按钮加到网格布局里
    pGridLayout->addWidget(&m_numButtons[1], 3, 0, 1, 1);
    pGridLayout->addWidget(&m_numButtons[2], 3, 1, 1, 1);
    pGridLayout->addWidget(&m_numButtons[3], 3, 2, 1, 1);
    pGridLayout->addWidget(&m_subButton,     3, 3, 1, 1);

    // 第五行,把清小数点按钮,数字0按钮和等号按钮加到网格布局里
    pGridLayout->addWidget(&m_dotButton,     4, 0, 1, 1);
    pGridLayout->addWidget(&m_numButtons[0], 4, 1, 1, 1);
    pGridLayout->addWidget(&m_equalButton,   4, 2, 1, 1);
    pGridLayout->addWidget(&m_addButton,     4, 3, 1, 1);

    pGridLayout->setContentsMargins(2,2,2,2);
    pGridLayout->setSpacing(2);

    QFrame *pGrindFrame = new QFrame(this); // 创建网格窗体(因为布局一定要放到widget才行)
    Q_CHECK_PTR(pGrindFrame);
    pGrindFrame->setFixedSize(324, 465); // 设置网格窗体的大小
    pGrindFrame->setObjectName("bg"); // 设置网格窗体的类名
    pGrindFrame->setLayout(pGridLayout); // 设置布局(把布局加到widget)
    pGrindFrame->move(0, 110);
}

void Ui::initWindowControls()
{
    auto *closeButton = new QPushButton(QStringLiteral("×"), &m_titleBar);
    closeButton->setObjectName("window_close");
    closeButton->setAccessibleName(tr("关闭窗口"));
    closeButton->setToolTip(closeButton->accessibleName());
    closeButton->setFixedSize(18, 18);
    closeButton->move(4, 1);
    connect(closeButton, &QPushButton::clicked, this, &QWidget::close);

    auto *minimizeButton = new QPushButton(QStringLiteral("−"), &m_titleBar);
    minimizeButton->setObjectName("window_minimize");
    minimizeButton->setAccessibleName(tr("最小化窗口"));
    minimizeButton->setToolTip(minimizeButton->accessibleName());
    minimizeButton->setFixedSize(18, 18);
    minimizeButton->move(26, 1);
    connect(minimizeButton, &QPushButton::clicked, this, &QWidget::showMinimized);
}

void Ui::initShortcuts()
{
    auto bind = [this](const QKeySequence &key, QAbstractButton *button) {
        auto *shortcut = new QShortcut(key, this);
        connect(shortcut, &QShortcut::activated, button, &QAbstractButton::click);
    };

    for (int digit = 0; digit < 10; ++digit)
        bind(QKeySequence(Qt::Key_0 + digit), &m_numButtons[digit]);

    bind(QKeySequence(Qt::Key_Plus), &m_addButton);
    bind(QKeySequence(Qt::Key_Minus), &m_subButton);
    bind(QKeySequence(Qt::Key_Asterisk), &m_mulButton);
    bind(QKeySequence(Qt::Key_Slash), &m_divButton);
    bind(QKeySequence(Qt::Key_Period), &m_dotButton);
    bind(QKeySequence(Qt::Key_Equal), &m_equalButton);
    bind(QKeySequence(Qt::Key_Return), &m_equalButton);
    bind(QKeySequence(Qt::Key_Enter), &m_equalButton);
    bind(QKeySequence(Qt::Key_Backspace), &m_delButton);
    bind(QKeySequence(Qt::Key_Escape), &m_clearButton);
}

void Ui::startEntry()
{
    if (m_state == State::Entering)
        return;

    m_display.setText(QStringLiteral("0"));
    m_state = State::Entering;
}

/*******************************************************************************
 * 函数名称： numButtonClicked
 * 功能描述： 数字按钮的槽函数
 * 参数列表： void
 *
 * 返回结果： void
 ******************************************************************************/
void Ui::numButtonClicked()
{
    // 把相应的数字按钮转化为响应的指针
    auto *button = qobject_cast<Button *>(sender());
    Q_ASSERT(button);
    startEntry();

    const QString text = m_display.text();
    if (text == QStringLiteral("0"))
        m_display.setText(button->getName());
    else if (text.size() < maxInputLength)
        m_display.setText(text + button->getName());
}

/*******************************************************************************
 * 函数名称： dotButtonClicked
 * 功能描述： 小数点按钮的槽函数
 * 参数列表： void
 *
 * 返回结果： void
 ******************************************************************************/
void Ui::dotButtonClicked()
{
    // 开始输入新的操作数时，显示从 0 起步
    startEntry();

    const QString text = m_display.text();
    // 如果还没有小数点，就添加小数点
    if (!text.contains(u'.') && text.size() < maxInputLength)
        m_display.setText(text + u'.');
}

/*******************************************************************************
 * 函数名称： delButtonClicked
 * 功能描述： 删除按钮的槽函数
 * 参数列表： void
 *
 * 返回结果： void
 ******************************************************************************/
void Ui::delButtonClicked()
{
    if (m_state == State::Error)
        clearButtonClicked();
    else if (m_state == State::Entering)
    {
        QString text = m_display.text();
        // 删除最后一个字符
        text.chop(1);
        // 如果只有一个数字,删除后重置为0
        m_display.setText(text.isEmpty() ? QStringLiteral("0") : text);
    }
}

/*******************************************************************************
 * 函数名称： clearButtonClicked
 * 功能描述： 清除按钮的槽函数
 * 参数列表： void
 *
 * 返回结果： void
 ******************************************************************************/
void Ui::clearButtonClicked()
{
    m_state = State::Entering; // 重置输入状态
    m_operation = EAlgorithmType_Unknown; // 重置算法类型
    m_accumulator = 0.0; // 重置左操作数
    m_display.setText(QStringLiteral("0")); // 重置文本内容
}

bool Ui::commitOperand()
{
    bool valid;
    const double operand = m_display.text().toDouble(&valid);
    if (valid)
    {
        // 有待执行的运算时，通过算法上下文类调用相应的算法
        // 记录结果,为下次计算
        m_accumulator = m_operation == EAlgorithmType_Unknown
            ? operand
            : AlgorithmContext(m_operation).Calculate(m_accumulator, operand);
        if (std::isfinite(m_accumulator))
        {
            if (m_operation != EAlgorithmType_Unknown)
                m_display.setText(QString::number(m_accumulator, 'g', displayPrecision));
            return true;
        }
    }

    m_state = State::Error;
    m_operation = EAlgorithmType_Unknown;
    m_accumulator = 0.0;
    m_display.setText(QStringLiteral("Error"));
    return false;
}

/*******************************************************************************
 * 函数名称： selectOperator
 * 功能描述： 响应加、减、乘、除按钮，选择下一步运算
 * 参数列表： operation - 运算类型
 *
 * 返回结果： void
 ******************************************************************************/
void Ui::selectOperator(EAlgorithmType operation)
{
    if (m_state == State::Error)
        return;
    // 等待右操作数时，只更新运算符
    if (m_state == State::Entering && !commitOperand())
        return;

    m_operation = operation;
    m_state = State::AwaitingOperand;
}

/*******************************************************************************
 * 函数名称： equalButtonClicked
 * 功能描述： 等号按钮的槽函数
 * 参数列表： void
 *
 * 返回结果： void
 ******************************************************************************/
void Ui::equalButtonClicked()
{
    if (m_state != State::Entering || m_operation == EAlgorithmType_Unknown)
        return;
    if (commitOperand())
    {
        m_operation = EAlgorithmType_Unknown; // 重置算法类型
        m_state = State::Result;
    }
}
