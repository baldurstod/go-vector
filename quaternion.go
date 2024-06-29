package vector

import "fmt"

type Quaternion[T Number] [4]T

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

func (quat *Quaternion[T]) String() string {
	return fmt.Sprintf("[%f, %f, %f, %f]", float32(quat[0]), float32(quat[1]), float32(quat[2]), float32(quat[3]))
}
