use std::result;

struct Solution;

impl Solution {
    pub fn largest_odd_number(num: String) -> String {
        let chars: Vec<char> = num.chars().collect();
        let n = chars.len();
        for i in (0..n).rev() {
            println!("{}", i);
            if chars[i].to_digit(10).unwrap() % 2 == 1 {
                return chars[..i].iter().collect();
            }
        }
        String::new()
    }
}

fn main() {
    let result = Solution::largest_odd_number(String::from("32782489638346578713315098393010310518347382"));
    println!("{}", result);
}
