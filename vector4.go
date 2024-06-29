package vector

import "fmt"

type Vector4[T Number] [4]T

func (vec *Vector4[T]) Clone() *Vector4[T] {
	return &Vector4[T]{vec[0], vec[1], vec[2], vec[3]}
}

func (vec *Vector4[T]) Copy(other *Vector4[T]) {
	vec[0] = other[0]
	vec[1] = other[1]
	vec[2] = other[2]
	vec[3] = other[3]
}

func (vec *Vector4[T]) Add(other *Vector4[T]) {
	vec[0] += other[0]
	vec[1] += other[1]
	vec[2] += other[2]
	vec[3] += other[3]
}

func (vec *Vector4[T]) Sub(other *Vector4[T]) {
	vec[0] -= other[0]
	vec[1] -= other[1]
	vec[2] -= other[2]
	vec[3] -= other[3]
}

func (vec *Vector4[T]) String() string {
	return fmt.Sprintf("[%f, %f, %f, %f]", float32(vec[0]), float32(vec[1]), float32(vec[2]), float32(vec[3]))
}
