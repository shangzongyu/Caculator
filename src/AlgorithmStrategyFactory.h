/*******************************************************************************
 * 类名: AlgorithmStrategyFactory
 * 功能: 算法产生工厂(用于产生相应的的算法)
 ******************************************************************************/

#ifndef ALGORITHMSTRATEGYFACTORY_H
#define ALGORITHMSTRATEGYFACTORY_H

#include <memory>
#include "AlgorithmStrategy.h"

class AlgorithmStrategyFactory
{
public:
    static std::unique_ptr<AlgorithmStrategy> createStrategy(EAlgorithmType operation);
};

#endif // ALGORITHMSTRATEGYFACTORY_H
