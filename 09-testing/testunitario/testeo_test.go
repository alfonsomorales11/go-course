package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 10 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 10)

	}
}

func TestSuma2(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 51},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, item.c)
		}
	}
}

func TestMayorQue(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{10, 20, 20},
		{100, 150, 150},
		{200, 100, 200},
	}

	for _, item := range tabla {
		result := GetMax(item.a, item.b)

		if result != item.c {
			t.Errorf("Hay un error, se esperaba %d y el resultado fue %d", item.c, result)
		}
	}
}
