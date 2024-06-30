package vector

import (
	"fmt"
	"math"
)

type Vector2[T Number] [2]T

func (vec *Vector2[T]) Clone() *Vector2[T] {
	return &Vector2[T]{vec[0], vec[1]}
}

func (vec *Vector2[T]) Copy(other *Vector2[T]) {
	vec[0] = other[0]
	vec[1] = other[1]
}

func (vec *Vector2[T]) Add(other *Vector2[T]) {
	vec[0] += other[0]
	vec[1] += other[1]
}

func (vec *Vector2[T]) Sub(other *Vector2[T]) {
	vec[0] -= other[0]
	vec[1] -= other[1]
}

func (vec *Vector2[T]) Len() float64 {
	return float64(math.Sqrt(float64(vec[0]*vec[0] + vec[1]*vec[1])))
}

func (vec *Vector2[T]) Identity() {
	vec[0] = 0
	vec[1] = 0
}

func (vec *Vector2[T]) String() string {
	return fmt.Sprintf("[%f, %f]", float32(vec[0]), float32(vec[1]))
}
