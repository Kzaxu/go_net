package tensor

import (
	"MyNet/common/utils"
	"MyNet/core/types"
	"fmt"
)

type Tensor[dType types.Numeric] struct {
	data_  []dType
	diff_  []dType
	shape_ []uint
	count_ uint
}

func NewTensor[dType types.Numeric](num ...uint) *Tensor[dType] {

	d := Tensor[dType]{
		shape_: num,
	}
	return &d
}

func NewTensorS[dType types.Numeric](size []uint) *Tensor[dType] {
	d := Tensor[dType]{}
	return &d
}

func (t *Tensor[dType]) Reshape(num ...uint) {

}

func (t *Tensor[dType]) ReshapeS(size []uint) {

}

func (t *Tensor[dType]) Size() []uint {
	return t.shape_
}

func (t *Tensor[dType]) AxisSize(d int) uint {
	if d < 0 {
		d += len(t.shape_)
	}
	if d >= len(t.shape_) {
		panic(fmt.Sprintf("tensor shape index %v illegal", d))
	}
	return t.shape_[d]
}

func (t *Tensor[dType]) Count(axis ...uint) uint {

	if len(axis) == 0 {
		return t.count_
	}

	start := axis[0]
	end := uint(len(t.shape_))
	if len(axis) == 2 {
		end = axis[1]
	}

	cnt := uint(1)
	for i := start; i < end; i++ {
		cnt *= t.shape_[i]
	}
	return cnt
}

func (t *Tensor[dType]) Offset(nums ...uint) uint {
	return t.OffsetS(nums)
}

func (t *Tensor[dType]) OffsetS(nums []uint) uint {
	ans := uint(0)
	utils.Check(len(nums) > len(t.shape_), "shape mismatch")
	nums = append(nums, make([]uint, len(t.shape_)-len(nums))...)

	for i := 0; i < len(t.shape_); i++ {
		ans *= t.shape_[i]
		ans += nums[i]
	}
	return ans
}

func (t *Tensor[dType]) CopyFrom(
	src *Tensor[dType], copyDiff bool, reshape bool) {

}

func (t *Tensor[dType]) DataAt(nums ...uint) dType {
	return t.DataAtS(nums)
}

func (t *Tensor[dType]) DataAtS(nums []uint) dType {
	return t.data_[t.OffsetS(nums)]
}

func (t *Tensor[dType]) DiffAt(nums ...uint) dType {
	return t.DiffAtS(nums)
}

func (t *Tensor[dType]) DiffAtS(nums []uint) dType {
	return t.diff_[t.OffsetS(nums)]
}
