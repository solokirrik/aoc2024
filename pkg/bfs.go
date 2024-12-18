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
		curPos := curStep.dpos.GetPos()
		if _, wasVisited := visited[curPos.Hash()]; wasVisited {
			continue
		}

		visited[curPos.Hash()] = curStep.score

		if curPos.Eq(end) {
			return curStep.pathLen
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
					parent: NewDCoord(start, dir),
					dpos:   NewDCoord(start, dir),
				},
				mtx,
			)...,
		)
	}

	stE := Step{
		parent: NewDCoord(start, EAST),
		dpos:   NewDCoord(start, EAST),
	}
	stS := Step{
		parent: NewDCoord(start, SOUTH),
		dpos:   NewDCoord(start, SOUTH),
	}

	return append(
		options(map[CoordHash]int{}, stS, mtx),
		options(map[CoordHash]int{}, stE, mtx)...,
	)
}

func options(visited map[CoordHash]int, curStep Step, mtx [][]int) []Step {
	out := []Step{}
	for dir := range deltas {
		row, col := curStep.dpos.c.R+deltas[dir][0], curStep.dpos.c.C+deltas[dir][1]
		if dir == opposites[curStep.dpos.dir] ||
			row < 0 || col < 0 ||
			row >= len(mtx) || col >= len(mtx[0]) ||
			mtx[row][col] == 1 {
			continue
		}

		stepCoord := NewCoord(row, col)
		stepDirCoord := NewDCoord(stepCoord, dir)
		newScore := curStep.pathLen

		st := Step{
			parent:  curStep.dpos,
			dpos:    stepDirCoord,
			score:   newScore,
			pathLen: curStep.pathLen + 1,
		}

		if _, wasVisited := visited[stepCoord.Hash()]; wasVisited {
			continue
		}

		out = append(out, st)
	}

	return out
}
