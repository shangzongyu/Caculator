/*******************************************************************************
 * 类名: SubStrategy
 * 功能: 实现减法
 ******************************************************************************/

#ifndef SUBSTRATEGY_H
#define SUBSTRATEGY_H

#include "AlgorithmStrategy.h"

class SubStrategy : public AlgorithmStrategy
{
public:
    double getResult(double leftOperand, double rightOperand) const override;
};

#endif // SUBSTRATEGY_H
