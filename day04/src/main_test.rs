use std::fs;
use crate::Solver;

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_example() {
        let content = fs::read_to_string("./ex");
        assert!(content.is_ok());

        let s: Solver = Solver::new().prep(content.unwrap().as_str());
        assert_eq!(18, s.part1());
    }

    #[test]
    fn part1() {
        let content = fs::read_to_string("./inp");
        assert!(content.is_ok());

        let s: Solver = Solver::new().prep(content.unwrap().as_str());
        assert_eq!(2547, s.part1());
    }

    #[test]
    fn part2_example() {
        let content = fs::read_to_string("./ex");
        assert!(content.is_ok());

        let s: Solver = Solver::new().prep(content.unwrap().as_str());
        assert_eq!(9, s.part2());
    }

    #[test]
    fn part2() {
        let content = fs::read_to_string("./inp");
        assert!(content.is_ok());

        let s: Solver = Solver::new().prep(content.unwrap().as_str());
        assert_eq!(1939, s.part2());
    }
}
