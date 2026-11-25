#include "AlgorithmStrategyFactory.h"
#include "AddStrategy.h"
#include "DivStrategy.h"
#include "MulStrategy.h"
#include "SubStrategy.h"

/*******************************************************************************
 * 函数名称： createStrategy
 * 功能描述： 创建相应的算法策略
 * 参数列表： operation - 算法类型
 *
 * 返回结果： 相应的算法策略，类型未知时返回空指针
 ******************************************************************************/
std::unique_ptr<AlgorithmStrategy> AlgorithmStrategyFactory::createStrategy(EAlgorithmType operation)
{
    switch (operation)
    {
        case EAlgorithmType_Add:
            return std::make_unique<AddStrategy>();
        case EAlgorithmType_Sub:
            return std::make_unique<SubStrategy>();
        case EAlgorithmType_Mul:
            return std::make_unique<MulStrategy>();
        case EAlgorithmType_Div:
            return std::make_unique<DivStrategy>();
        case EAlgorithmType_Unknown:
            return nullptr;
    }
    return nullptr;
}
