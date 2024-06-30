package vector

import (
	"fmt"
	"math"
)

type Vector3[T Number] [3]T

func (vec *Vector3[T]) Clone() *Vector3[T] {
	return &Vector3[T]{vec[0], vec[1], vec[2]}
}

func (vec *Vector3[T]) Copy(other *Vector3[T]) {
	vec[0] = other[0]
	vec[1] = other[1]
	vec[2] = other[2]
}

func (vec *Vector3[T]) Add(other *Vector3[T]) {
	vec[0] += other[0]
	vec[1] += other[1]
	vec[2] += other[2]
}

func (vec *Vector3[T]) Sub(other *Vector3[T]) {
	vec[0] -= other[0]
	vec[1] -= other[1]
	vec[2] -= other[2]
}

func (vec *Vector3[T]) Len() float64 {
	return float64(math.Sqrt(float64(vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2])))
}

func (vec *Vector3[T]) Identity() {
	vec[0] = 0
	vec[1] = 0
	vec[2] = 0
}

func (vec *Vector3[T]) String() string {
	return fmt.Sprintf("[%f, %f, %f]", float32(vec[0]), float32(vec[1]), float32(vec[2]))
}
