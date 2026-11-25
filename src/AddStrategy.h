/*******************************************************************************
 * 类名: AddStrategy
 * 功能: 实现加法
 ******************************************************************************/

#ifndef ADDSTRATEGY_H
#define ADDSTRATEGY_H

#include "AlgorithmStrategy.h"

class AddStrategy : public AlgorithmStrategy
{
public:
    double getResult(double leftOperand, double rightOperand) const override;
};

#endif // ADDSTRATEGY_H
