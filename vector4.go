package vector

type Vector4[T Number] [4]T

func (vec *Vector4[T]) Add(other Vector4[T]) {
	vec[0] += other[0]
	vec[1] += other[1]
	vec[2] += other[2]
	vec[3] += other[3]
}

func (vec *Vector4[T]) Sub(other Vector4[T]) {
	vec[0] -= other[0]
	vec[1] -= other[1]
	vec[2] -= other[2]
	vec[3] -= other[3]
}