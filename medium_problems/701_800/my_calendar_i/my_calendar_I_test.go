package my_calendar_I

import (
	"testing"
)

func TestBookTestCase1(t *testing.T) {

	t.Run("Testing Book", func(test *testing.T) {
		obj := Constructor()
		got := obj.Book(10, 20)

		if got == false {
			test.Errorf("expected true, got false available")
		}

		got = obj.Book(15, 25)

		if got == true {
			test.Errorf("expected false, got true available")
		}

		got = obj.Book(20, 30)

		if got == false {
			test.Errorf("expected true, got false available")
		}

	})

}

func TestBookTestCase2(t *testing.T) {

	t.Run("Testing Book", func(test *testing.T) {
		obj := Constructor()
		obj.Book(47, 50)
		obj.Book(33, 41)
		obj.Book(25, 32)
		got := obj.Book(19, 25)
		if got == false {
			test.Errorf("expected true, got false available")
		}
	})

}
