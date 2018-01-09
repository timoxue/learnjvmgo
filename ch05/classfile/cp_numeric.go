package classfile

import (
	"math"
)

//ConstantIntegerInfo #
type ConstantIntegerInfo struct {
	val int32
}

//ConstantFloatInfo #
type ConstantFloatInfo struct {
	val float32
}

//ConstantLongInfo #
type ConstantLongInfo struct {
	val int64
}

//ConstantDoubleInfo #
type ConstantDoubleInfo struct {
	val float64
}

func (cii *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit32()
	cii.val = int32(bytes)
}

func (cii *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit32()
	cii.val = math.Float32frombits(bytes)
}

func (cii *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cii.val = int64(bytes)
}

func (cii *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cii.val = math.Float64frombits(bytes)
}
