use std::fs;
use std::time::Instant;

fn main() {
    let content = fs::read_to_string("./inp");
    match content {
        Ok(content) => {
            let columns: [Vec<i64>; 2] = to_sorted_lines(&content);
            measure_time(|| task_1_plain(columns.clone()));
            measure_time(|| task_1_func(columns.clone()));
            measure_time(|| task_2_plain(columns.clone()));
            measure_time(|| task_2_func(columns.clone()));

            println!("{:?}", task_1_plain(columns.clone()));
            println!("{:?}", task_2_plain(columns.clone()));
        }
        Err(e) => {
            println!("{:?}", e);
        }
    }
}

fn to_sorted_lines(inp: &str) -> [Vec<i64>; 2] {
    let mut first_column: Vec<i64> = Vec::new();
    let mut second_column: Vec<i64> = Vec::new();

    for line in inp.lines() {
        let nums: Vec<i64> = line
            .split_whitespace()
            .filter_map(|word| word.parse::<i64>().ok())
            .collect();

        first_column.push(nums[0]);
        second_column.push(nums[1]);
    }

    first_column.sort_by(|a, b| a.cmp(b));
    second_column.sort_by(|a, b| a.cmp(b));

    [first_column, second_column]
}

fn task_1_plain(inp: [Vec<i64>; 2]) -> i64 {
    let mut diff: i64 = 0;

    for i in 0..inp[0].len() {
        diff += (inp[0][i] - inp[1][i]).abs()
    }

    diff
}


fn task_1_func(inp: [Vec<i64>; 2]) -> i64 {
    inp[0]
        .iter()
        .zip(inp[1].iter())
        .map(|(&a, &b)| (a - b).abs())
        .sum()
}

fn task_2_plain(inp: [Vec<i64>; 2]) -> i64 {
    let mut sim: i64 = 0;

    for i in 0..inp[0].len() {
        let count = count_eq_func(inp[0][i], &inp[1]);
        sim += inp[0][i] * count;
    }

    sim
}

fn task_2_func(inp: [Vec<i64>; 2]) -> i64 {
    inp[0]
        .iter()
        .map(|&x| x * count_eq_func(x, &inp[1]))
        .sum()
}

fn count_eq_func(target: i64, inp: &[i64]) -> i64 {
    inp.iter().
        take_while(|&&x| x <= target).
        filter(|&&x| x == target).
        count() as i64
}

fn count_eq(target: i64,inp: &Vec<i64>) -> i64 {
    let mut count: i64 = 0;

    for i in 0..inp.len() {
        if inp[i] > target{
            break;
        }
        if inp[i] == target {
            count += 1;
        }
    }

    count
}

fn measure_time<F, R>(f: F) -> R
where
    F: FnOnce() -> R,
{
    let start = Instant::now();
    let result = f();
    let duration = start.elapsed();

    println!("Execution time: {:?}", duration);

    result
}