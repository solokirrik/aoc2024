mod main_test;

use std::fs;
use std::ops::Neg;

fn main() {
    let content = fs::read_to_string("./day04/inp");
    match content {
        Ok(content) => {
            let s: Solver = Solver::new().prep(&content);
            println!("{:?}", s.part1());
            println!("{:?}", s.part2());
        }
        Err(e) => {
            println!("{:?}", e);
        }
    }
}

struct Solver {
    mtx: Vec<Vec<char>>,
}

impl Solver {
    fn new() -> Solver {
        Solver { mtx: vec![] }
    }

    fn prep(mut self, inp: &str) -> Self {
        self.mtx = inp
            .lines()
            .map(|line| line.chars().collect::<Vec<char>>())
            .collect();
        self
    }

    fn part1(&self) -> i64 {
        self.mtx
            .iter()
            .enumerate()
            .flat_map(|(i, row)| {
                row.iter().enumerate().filter_map(move |(j, &cell)| {
                    if cell == 'X' {
                        Some(self.count_xmas(i, j))
                    } else {
                        None
                    }
                })
            })
            .sum()
    }

    fn part2(&self) -> i64 {
        0
    }

    fn count_xmas(&self, row: usize, col: usize) -> i64 {
        let mut found: i64 = 0;

        for dir in XMAS_POINTS {
            for i in 0..XMAS.len() {
                let letter = &XMAS[i];
                let step_row: usize;
                let step_col: usize;

                if let Some([new_row, new_col]) = self.new_row_col(dir, i, row, col) {
                    step_row = new_row;
                    step_col = new_col;
                } else {
                    break;
                }

                if !self.is_in_bound(step_row, step_col)
                    || self.mtx[step_row][step_col].cmp(letter).is_ne()
                {
                    break;
                }

                if *letter == 'S'
                    && self.mtx[step_row][step_col]
                        .cmp(XMAS.last().unwrap())
                        .is_eq()
                {
                    found += 1;
                }
            }
        }

        found
    }

    fn new_row_col(&self, dir: [i64; 2], i: usize, row: usize, col: usize) -> Option<[usize; 2]> {
        let row_offset = dir[0];
        let col_offset = dir[1];

        let test_cur_row: Option<usize> = if row_offset < 0 {
            row.checked_sub(i * row_offset.neg() as usize)
        } else {
            row.checked_add(i * row_offset as usize)
        };

        let test_cur_col: Option<usize> = if col_offset < 0 {
            col.checked_sub(i * col_offset.neg() as usize)
        } else {
            col.checked_add(i * col_offset as usize)
        };

        if test_cur_row.is_none() || test_cur_col.is_none() {
            return None;
        }

        Some(<[usize; 2]>::try_from(vec![test_cur_row.unwrap(), test_cur_col.unwrap()]).unwrap())
    }

    fn is_in_bound(&self, row: usize, col: usize) -> bool {
        row < self.mtx.len() && col < self.mtx[row].len()
    }
}

const XMAS: [char; 4] = ['X', 'M', 'A', 'S'];
const XMAS_POINTS: [[i64; 2]; 8] = [
    [1, 0],   // down
    [0, 1],   // right
    [-1, 0],  // up
    [0, -1],  // left
    [1, 1],   // down-right
    [-1, 1],  // up-right
    [-1, -1], // up-left
    [1, -1],  // down-left[
];

#[allow(dead_code)]
const MAS_BOUNDS: [[i64; 2]; 4] = [[-1, -1], [-1, 1], [1, 1], [1, -1]];
