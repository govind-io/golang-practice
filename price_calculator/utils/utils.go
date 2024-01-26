package utils

type mapList func(float64, int)

func Map(list *[]float64, iterator mapList) {

	for index, value := range *list {
		iterator(value, index)
	}
}
