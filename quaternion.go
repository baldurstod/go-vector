package vector

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type Quaternion[T constraints.Float] [4]T

func (vec *Quaternion[T]) Set(x T, y T, z T, w T) {
	vec[0] = x
	vec[1] = y
	vec[2] = z
	vec[3] = w
}

func (quat *Quaternion[T]) Clone() *Quaternion[T] {
	return &Quaternion[T]{quat[0], quat[1], quat[2], quat[3]}
}

func (quat *Quaternion[T]) Copy(other *Quaternion[T]) {
	quat[0] = other[0]
	quat[1] = other[1]
	quat[2] = other[2]
	quat[3] = other[3]
}

func (quat *Quaternion[T]) Add(other *Quaternion[T]) {
	quat[0] += other[0]
	quat[1] += other[1]
	quat[2] += other[2]
	quat[3] += other[3]
}

func (quat *Quaternion[T]) Len() float64 {
	return float64(math.Sqrt(float64(quat[0]*quat[0] + quat[1]*quat[1] + quat[2]*quat[2] + quat[3]*quat[3])))
}

func (quat *Quaternion[T]) Identity() {
	quat[0] = 0
	quat[1] = 0
	quat[2] = 0
	quat[3] = 1
}

func (quat *Quaternion[T]) Invert() {
	a0 := quat[0]
	a1 := quat[1]
	a2 := quat[2]
	a3 := quat[3]

	dot := a0*a0 + a1*a1 + a2*a2 + a3*a3

	var invDot T
	if dot == 0 {
		quat[0] = 0
		quat[1] = 0
		quat[2] = 0
		quat[3] = 0
	} else {
		invDot = 1.0 / dot

		quat[0] = -a0 * invDot
		quat[1] = -a1 * invDot
		quat[2] = -a2 * invDot
		quat[3] = a3 * invDot
	}
}

func (quat *Quaternion[T]) RotateX(angle float64) {
	angle *= 0.5

	ax := float64(quat[0])
	ay := float64(quat[1])
	az := float64(quat[2])
	aw := float64(quat[3])
	bx := math.Sin(angle)
	bw := math.Cos(angle)

	quat[0] = T(ax*bw + aw*bx)
	quat[1] = T(ay*bw + az*bx)
	quat[2] = T(az*bw - ay*bx)
	quat[3] = T(aw*bw - ax*bx)
}

func (quat *Quaternion[T]) RotateY(angle float64) {
	angle *= 0.5

	ax := float64(quat[0])
	ay := float64(quat[1])
	az := float64(quat[2])
	aw := float64(quat[3])
	by := math.Sin(angle)
	bw := math.Cos(angle)

	quat[0] = T(ax*bw - az*by)
	quat[1] = T(ay*bw + aw*by)
	quat[2] = T(az*bw + ax*by)
	quat[3] = T(aw*bw - ay*by)
}

func (quat *Quaternion[T]) RotateZ(angle float64) {
	angle *= 0.5

	ax := float64(quat[0])
	ay := float64(quat[1])
	az := float64(quat[2])
	aw := float64(quat[3])
	bz := math.Sin(angle)
	bw := math.Cos(angle)

	quat[0] = T(ax*bw + ay*bz)
	quat[1] = T(ay*bw - ax*bz)
	quat[2] = T(az*bw + aw*bz)
	quat[3] = T(aw*bw - az*bz)
}

func (quat *Quaternion[T]) FromEuler(x T, y T, z T) {
	x *= T(0.5)
	y *= T(0.5)
	z *= T(0.5)

	sx := math.Sin(float64(x))
	sy := math.Sin(float64(y))
	sz := math.Sin(float64(z))

	cx := math.Cos(float64(x))
	cy := math.Cos(float64(y))
	cz := math.Cos(float64(z))

	quat[0] = T(sx*cy*cz - cx*sy*sz)
	quat[1] = T(cx*sy*cz + sx*cy*sz)
	quat[2] = T(cx*cy*sz - sx*sy*cz)
	quat[3] = T(cx*cy*cz + sx*sy*sz)
}

func (quat *Quaternion[T]) String() string {
	return fmt.Sprintf("[%f, %f, %f, %f]", float32(quat[0]), float32(quat[1]), float32(quat[2]), float32(quat[3]))
}
