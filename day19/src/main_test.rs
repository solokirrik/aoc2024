#[cfg(test)]
mod tests {
    use crate::Solver;
    use std::fs;

    #[test]
    fn part1_example() {
        let content = fs::read_to_string("./ex");
        assert!(content.is_ok());

        let s: Solver = Solver::new().prep(content.unwrap());
        assert_eq!(6, s.part1());
    }

    #[test]
    fn part1_inp() {
        let content = fs::read_to_string("./inp");
        assert!(content.is_ok());

        let s: Solver = Solver::new().prep(content.unwrap());
        assert_eq!(336, s.part1());
    }

    #[test]
    fn part2_example() {
        let content = fs::read_to_string("./ex");
        assert!(content.is_ok());

        let s: Solver = Solver::new().prep(content.unwrap());
        assert_eq!(16, s.part2());
    }

    #[test]
    fn part2_inp() {
        let content = fs::read_to_string("./inp");
        assert!(content.is_ok());

        let s: Solver = Solver::new().prep(content.unwrap());
        assert_eq!(758890600222015, s.part2());
    }
}
