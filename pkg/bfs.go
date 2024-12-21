package pkg

const NONE = -1

func BFS(mtx [][]int, start, end Coord, startDirs []Direction) int {
	q := NewStepQueue()
	visited := make(map[CoordHash]int, len(mtx[0])*len(mtx))

	opts := startOptions(mtx, start, startDirs)
	for _, o := range opts {
		q.Push(o)
	}

	for q.Len() > 0 {
		curStep := q.Get()
		curPos := curStep.DPos.GetPos()
		if _, wasVisited := visited[curPos.Hash()]; wasVisited {
			continue
		}

		visited[curPos.Hash()] = curStep.Score

		if curPos.Eq(end) {
			return curStep.PathLen
		}

		opts := options(visited, curStep, mtx)
		for o := range opts {
			q.Push(opts[o])
		}
	}

	return NONE
}

func startOptions(mtx [][]int, start Coord, startDirs []Direction) []Step {
	out := make([]Step, 0, len(startDirs)*3)
	for _, dir := range startDirs {
		out = append(out,
			options(
				map[CoordHash]int{},
				Step{
					Parent: NewDCoord(start, dir),
					DPos:   NewDCoord(start, dir),
				},
				mtx,
			)...,
		)
	}

	stE := Step{
		Parent: NewDCoord(start, EAST),
		DPos:   NewDCoord(start, EAST),
	}
	stS := Step{
		Parent: NewDCoord(start, SOUTH),
		DPos:   NewDCoord(start, SOUTH),
	}

	return append(
		options(map[CoordHash]int{}, stS, mtx),
		options(map[CoordHash]int{}, stE, mtx)...,
	)
}

func options(visited map[CoordHash]int, curStep Step, mtx [][]int) []Step {
	out := []Step{}
	for dir := range Deltas {
		row, col := curStep.DPos.C.R+Deltas[dir][0], curStep.DPos.C.C+Deltas[dir][1]
		if dir == Opposites[curStep.DPos.Dir] ||
			row < 0 || col < 0 ||
			row >= len(mtx) || col >= len(mtx[0]) ||
			mtx[row][col] == 1 {
			continue
		}

		stepCoord := NewCoord(row, col)
		stepDirCoord := NewDCoord(stepCoord, dir)
		newScore := curStep.PathLen

		st := Step{
			Parent:  curStep.DPos,
			DPos:    stepDirCoord,
			Score:   newScore,
			PathLen: curStep.PathLen + 1,
		}

		if _, wasVisited := visited[stepCoord.Hash()]; wasVisited {
			continue
		}

		out = append(out, st)
	}

	return out
}
