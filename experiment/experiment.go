package experiment

type StructureOfNumbers struct {
	numbers []struct {
		val int
	}
}

func (s *StructureOfNumbers) GetNumbers() []struct{ val int } {
	return s.numbers
}
