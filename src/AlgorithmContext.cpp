#include "AlgorithmContext.h"
#include "AlgorithmStrategyFactory.h"

#include <limits>

/*******************************************************************************
 * 函数名称： AlgorithmContext
 * 功能描述： 根据操作类型创建相应的算法策略
 * 参数列表： operation - 操作类型
 *
 * 返回结果： 无
 ******************************************************************************/
AlgorithmContext::AlgorithmContext(EAlgorithmType operation)
    : m_strategy(AlgorithmStrategyFactory::createStrategy(operation))
{
}

/*******************************************************************************
 * 函数名称： Calculate
 * 功能描述： 计算
 * 参数列表： leftOperand  - 左操作数
 *           rightOperand - 右操作数
 *
 * 返回结果： 结果，操作类型未知时返回 NaN
 ******************************************************************************/
double AlgorithmContext::Calculate(double leftOperand, double rightOperand) const
{
    if (!m_strategy)
        return std::numeric_limits<double>::quiet_NaN();

    return m_strategy->getResult(leftOperand, rightOperand);
}
