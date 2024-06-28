package vector

type Quaternion[T Number] [4]T

func (vec *Quaternion[T]) Add(other Quaternion[T]) {
	vec[0] += other[0]
	vec[1] += other[1]
	vec[2] += other[2]
	vec[3] += other[3]
}
