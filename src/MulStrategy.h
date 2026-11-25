/*******************************************************************************
 * 类名: MulStrategy
 * 功能: 实现乘法
 ******************************************************************************/

#ifndef MULSTRATEGY_H
#define MULSTRATEGY_H

#include "AlgorithmStrategy.h"

class MulStrategy : public AlgorithmStrategy
{
public:
    double getResult(double leftOperand, double rightOperand) const override;
};

#endif // MULSTRATEGY_H
