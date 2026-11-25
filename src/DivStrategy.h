/*******************************************************************************
 * 类名: DivStrategy
 * 功能: 实现除法
 ******************************************************************************/

#ifndef DIVSTRATEGY_H
#define DIVSTRATEGY_H

#include "AlgorithmStrategy.h"

class DivStrategy : public AlgorithmStrategy
{
public:
    double getResult(double leftOperand, double rightOperand) const override;
};

#endif // DIVSTRATEGY_H
