/*******************************************************************************
 * 类名: AlgorithmContext
 * 功能: 算法上下文类,用于调用相应的算法
 ******************************************************************************/

#ifndef CONTEXT_H
#define CONTEXT_H

#include <memory>
#include "AlgorithmStrategy.h"

class AlgorithmContext
{
public:
    explicit AlgorithmContext(EAlgorithmType operation);
    double Calculate(double leftOperand, double rightOperand) const;
private:
    std::unique_ptr<AlgorithmStrategy> m_strategy;
};

#endif // CONTEXT_H
