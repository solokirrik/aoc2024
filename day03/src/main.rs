use regex::Regex;
use std::fs;

fn main() {
    let content = fs::read_to_string("./day03/inp");
    match content {
        Ok(content) => {
            let got1 = task1(&content);
            println!("Part 1: {:?} {:?}", got1, 192767529 == got1);

            let got2 = task2(&content);
            println!("Part 2: {:?} {:?}", got2, 104083373 == got2);
        }
        Err(e) => {
            println!("{:?}", e);
        }
    }
}

fn task1(code: &str) -> i64 {
    sum_mults(code)
}

fn sum_mults(code: &str) -> i64 {
    let re_mul = match Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)") {
        Ok(re) => re,
        Err(e) => {
            println!("{:?}", e);
            return -1;
        }
    };

    let mut sum = 0;

    for cap in re_mul.captures_iter(code) {
        if let (Some(x), Some(y)) = (cap.get(1), cap.get(2)) {
            let a: i64 = x.as_str().parse().unwrap_or(0);
            let b: i64 = y.as_str().parse().unwrap_or(0);

            sum += a * b;
        }
    }

    sum
}

fn task2(code: &str) -> i64 {
    let re_dos = match Regex::new(r"do\(\)|don't\(\)") {
        Ok(re) => re,
        Err(e) => {
            println!("{:?}", e);
            return -1;
        }
    };

    let base_starts = re_dos
        .find_iter(code)
        .map(|re_match| re_match.start())
        .collect::<Vec<_>>();

    let mut starts = if base_starts[0] != 0 {
        std::iter::once(0).chain(base_starts.into_iter()).collect::<Vec<_>>()
    } else {
        base_starts
    };

    starts.push(code.len());

    starts.windows(2)
        .map(|win| &code[win[0]..win[1]])
        .filter(|sub_srt| !sub_srt.starts_with("don't()"))
        .map(sum_mults)
        .sum()
}
