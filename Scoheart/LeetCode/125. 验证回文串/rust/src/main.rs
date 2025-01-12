use std::result;

struct Solution;
impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let mut str2 = String::from(s);
        let filter_str: String = str2.chars().filter(|ch| ch.is_ascii_alphanumeric()).collect();
        let reversed_str: String = filter_str.chars().rev().collect();
        filter_str.to_lowercase() == reversed_str.to_lowercase()
    }
}

fn main() {
    let str = String::from("A man, a plan, a canal: Panama");
    let result = Solution::is_palindrome(str);
    print!("{}", result);
}
