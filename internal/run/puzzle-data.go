package run

var listOfPuzzles = []PuzzleData{
}

type puzzleCollection = map[int]map[int]PuzzleData

func getPuzzleCollection() puzzleCollection {
	collection := make(puzzleCollection)
	for _, puzzle := range listOfPuzzles {
		day, dayExists := collection[puzzle.Day]
		if !dayExists {
			collection[puzzle.Day] = make(map[int]PuzzleData)
			day, _ = collection[puzzle.Day]
		}

		day[puzzle.Part] = puzzle
	}
	return collection
}
