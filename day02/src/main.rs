use std::fs;

fn main() {
    let content = fs::read_to_string("./day02/inp");
    match content {
        Ok(content) => {
            let reports = into_matrix(&content);
            println!("{:?}", task1(reports.clone()));
            println!("{:?}", task2(reports.clone()));
        }
        Err(e) => {
            println!("{:?}", e);
        }
    }
}

fn task1(reports: Vec<Vec<i64>>) -> i64 {
    let mut safe: i64 = 0;

    for report in reports {
        if is_safe(report){
            safe += 1;
        }
    }

    safe
}

fn task2(reports: Vec<Vec<i64>>) -> i64 {
    let mut safe: i64 = 0;

    for report in reports {
        let mut was_safe: bool = false;

        for mutation in get_mutations(report) {
            if is_safe(mutation){
                was_safe = true;
                break;
            }
        }

        if was_safe{
            safe += 1;
        }
    }

    safe
}

fn into_matrix(content: &str) -> Vec<Vec<i64>> {
    content.lines()
        .map(|line| {
            line.split_whitespace()
                .filter_map(|word| word.parse::<i64>().ok())
                .collect::<Vec<i64>>()
        })
        .collect()
}

fn is_safe(report: Vec<i64>) -> bool {
    let old_dir = to_dir(report[0], report[1]);

    if old_dir == ERR_UNSAFE {
        return false;
    }

    for i in 0..report.len() - 1 {
        if check(report[i], report[i + 1], old_dir) == ERR_UNSAFE {
            return false;
        }
    }

    true
}

fn check(a: i64, b: i64, old_dir: &str) -> &str {
    let dir = to_dir(a, b);

    if !(1..=3).contains(&(a-b).abs()) || dir == ERR_UNSAFE || old_dir != dir{
        return ERR_UNSAFE
    }

    return ""
}

const ERR_UNSAFE: &str = "ERR_UNSAFE";

fn to_dir(a: i64, b: i64) -> &'static str {
   match a.cmp(&b) {
        std::cmp::Ordering::Equal => ERR_UNSAFE,
        std::cmp::Ordering::Greater => "decr",
        std::cmp::Ordering::Less => "incr",
    }
}

fn get_mutations(report: Vec<i64>) -> Vec<Vec<i64>> {
    let mut out = Vec::new();

    out.push(report.clone());

    for i in 0..report.len() {
        let mut mutation = report.clone();
        mutation.remove(i);
        out.push(mutation);
    }

    out
}
