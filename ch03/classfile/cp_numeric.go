package classfile

import (
	"math"
)

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit32()
	this.val = int32(bytes)
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit32()
	this.val = math.Float32frombits(bytes)
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = int64(bytes)
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint65()
	this.val = math.Float64frombits(bytes)
}
