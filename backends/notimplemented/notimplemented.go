// Package notimplemented implements a backends.Builder interface that throws a "Not implemented"
// exception to all operations.
//
// This can help bootstrap any backend implementation.
package notimplemented

import (
	"fmt"
	"github.com/gomlx/gomlx/backends"
	"github.com/gomlx/gomlx/types/shapes"
	"github.com/pkg/errors"
)

// NotImplementedError is thrown by every method.
var NotImplementedError = fmt.Errorf("not implemented")

// Builder implements backends.Builder by throwing "Not implemented" exception for every operation.
type Builder struct{}

var _ backends.Builder = Builder{}

//go:generate go run ../../internal/cmd/notimplemented_generator

func (b Builder) Name() string {
	return "Dummy \"not implemented\" backend, please override this method"
}

func (b Builder) Compile(outputs ...backends.Op) backends.Executable {
	panic(errors.Wrapf(NotImplementedError, "Compile()"))
	return nil
}

func (b Builder) OpShape(op backends.Op) shapes.Shape {
	panic(errors.Wrapf(NotImplementedError, "OpShape()"))
}

func (b Builder) Parameter(name string, shape shapes.Shape) backends.Op {
	panic(errors.Wrapf(NotImplementedError, "Parameter()"))
}

func (b Builder) Constant(flat any, dims ...int) backends.Op {
	panic(errors.Wrapf(NotImplementedError, "Constant()"))
}

func (b Builder) Identity(x backends.Op) backends.Op {
	panic(errors.Wrapf(NotImplementedError, "Identity()"))
}

func (b Builder) ReduceWindow(x backends.Op, reductionType backends.ReduceOpType, windowDimensions, strides, baseDilations, windowDilations []int, paddings [][2]int) backends.Op {
	panic(errors.Wrapf(NotImplementedError, "ReduceWindow()"))
}

func (b Builder) RngBitGenerator(state backends.Op, shape shapes.Shape) (newState, values backends.Op) {
	panic(errors.Wrapf(NotImplementedError, "RngBitGenerator()"))
}

func (b Builder) BatchNormForInference(operand, scale, offset, mean, variance backends.Op, epsilon float32, axis int) backends.Op {
	panic(errors.Wrapf(NotImplementedError, "BatchNormForInference()"))
}

func (b Builder) BatchNormForTraining(operand, scale, offset backends.Op, epsilon float32, axis int) (normalized, batchMean, batchVariance backends.Op) {
	panic(errors.Wrapf(NotImplementedError, "BatchNormForTraining()"))
}

func (b Builder) BatchNormGradient(operand, scale, mean, variance, gradOutput backends.Op, epsilon float32, axis int) (gradOperand, gradScale, gradOffset backends.Op) {
	panic(errors.Wrapf(NotImplementedError, "BatchNormGradient()"))
}

func (b Builder) BitCount(operand backends.Op) backends.Op {
	panic(errors.Wrapf(NotImplementedError, "BitCount()"))
}
