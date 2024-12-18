use std::time::Instant;

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

pub fn measure_time<F, R>(f: F) -> R
where
    F: FnOnce() -> R,
{
    let start = Instant::now();
    let result = f();
    let duration = start.elapsed();

    println!("Execution time: {:?}", duration);

    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
