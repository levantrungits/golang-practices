package lodash

import "testing"

func TestMap(t *testing.T) {
	t.Run("With Case Param2 Is Function", func(t *testing.T) {
		arr := [3]int{4, 8, 10}
		actual := Map(arr, square)
		expect := [3]int{16, 64, 100}
		if actual != expect {
			t.Fail()
		}
	})

	t.Run("With Case Param2 Is Map Int", func(t *testing.T) {
		dicMap := map[string]int{
			"value1": 4,
			"value2": 8,
			"value3": 10,
		}
		actual := Map(dicMap, square)
		expect := [3]int{16, 64, 100}
		if actual != expect {
			t.Fail()
		}
	})

	t.Run("With Case Param2 Is Map String", func(t *testing.T) {
		dicMap := map[string]string{
			"user1": "ong A",
			"user2": "ong B",
			"user3": "ong C",
		}
		actual := Map(dicMap, "user")
		expect := [3]string{"ong A", "ong B", "ong C"}
		if actual != expect {
			t.Fail()
		}
	})
}

func square(n int) int {
	return n * n
}
