#include "DivStrategy.h"

#include <limits>

/*******************************************************************************
 * 函数名称： getResult
 * 功能描述： 获得结果
 * 参数列表： leftOperand  - 左操作数
 *          rightOperand - 右操作数
 *
 * 返回结果： 结果，除数为 0 时返回 NaN
 ******************************************************************************/
double DivStrategy::getResult(double leftOperand, double rightOperand) const
{
    if (rightOperand == 0.0)
        return std::numeric_limits<double>::quiet_NaN();

    return leftOperand / rightOperand;
}
