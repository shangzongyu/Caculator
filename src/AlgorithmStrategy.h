/*******************************************************************************
 * 类名: AlgorithmStrategy
 * 功能: 纯虚算法策略类型
 ******************************************************************************/

#ifndef ALGORITHMSTRATEGY_H
#define ALGORITHMSTRATEGY_H

typedef enum AlgorithmType
{
    EAlgorithmType_Add    = 0,
    EAlgorithmType_Sub    = 1,
    EAlgorithmType_Mul    = 2,
    EAlgorithmType_Div    = 3,
    EAlgorithmType_Unknown = 100
}EAlgorithmType;

class AlgorithmStrategy
{
public:
    virtual ~AlgorithmStrategy() = default;
    virtual double getResult(double leftOperand, double rightOperand) const = 0;
};

#endif // ALGORITHMSTRATEGY_H
