use pkg::measure_time;
use std::collections::HashMap;
use std::fs;
mod main_test;

fn main() {
    let content = fs::read_to_string("./day19/inp");
    match content {
        Ok(content) => {
            let s: Solver = Solver::new().prep(content);
            measure_time(|| s.part1());
            measure_time(|| s.part2());
        }
        Err(e) => {
            println!("{:?}", e);
        }
    }
}

struct Solver {
    patterns: Vec<String>,
    designs: Vec<String>,
}

impl Solver {
    fn new() -> Solver {
        Solver {
            patterns: vec![],
            designs: vec![],
        }
    }

    fn prep(mut self, inp: String) -> Solver {
        let lines: Vec<&str> = inp.split("\n\n").collect();

        self.patterns = lines
            .first()
            .unwrap()
            .split(", ")
            .map(|s| s.to_string())
            .collect();
        self.designs = lines
            .last()
            .unwrap()
            .lines()
            .map(|s| s.to_string())
            .collect();

        self
    }

    fn part1(&self) -> i64 {
        let mut cache: HashMap<String, i64> = HashMap::new();

        let mut res: i64 = 0;

        for i in 0..self.designs.len() {
            if count_occurances(
                &mut cache,
                &self.designs[i],
                &filter_patterns(&self.designs[i], &self.patterns),
            ) > 0
            {
                res += 1;
            }
        }

        res
    }

    fn part2(&self) -> i64 {
        let mut cache: HashMap<String, i64> = HashMap::new();

        let mut res: i64 = 0;

        for i in 0..self.designs.len() {
            res += count_occurances(
                &mut cache,
                &self.designs[i],
                &filter_patterns(&self.designs[i], &self.patterns),
            );
        }

        res
    }
}

fn filter_patterns(design: &String, patterns: &[String]) -> Vec<String> {
    let mut out: Vec<String> = Vec::new();

    for p in patterns {
        if design.contains(p) {
            out.push(p.to_string());
        }
    }

    out
}

fn count_occurances(cache: &mut HashMap<String, i64>, design: &String, patterns: &[String]) -> i64 {
    if design.is_empty() {
        return 1;
    }

    if let Some(val) = cache.get(design) {
        return *val;
    }

    let mut res: i64 = 0;
    for pt in patterns.iter() {
        if design.starts_with(pt) {
            res += count_occurances(cache, &design[pt.len()..].to_string(), patterns);
        }
    }

    cache.insert(design.clone(), res);

    return res;
}
