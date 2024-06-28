package vector_test

import (
	"log"
	"testing"

	"github.com/baldurstod/go-vector"
)

func TestVector3(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	v1 := vector.Vector3[float32]{1, 2, 3}
	v2 := vector.Vector3[float32]{10, 20, 30}

	v1.Add(&v2)
	log.Println(v1, v2)
}
