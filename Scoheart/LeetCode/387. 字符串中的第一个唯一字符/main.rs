struct Solution;
use std::collections::HashMap;
impl Solution {
    pub fn first_uniq_char(s: String) -> i32 {
        let mut char_count:HashMap<char, i32> = HashMap::new();
        for ch in s.chars() {
            char_count.insert(ch, char_count.get(&ch).unwrap_or(&0) + 1);
        }
        let mut index: i32 = -1;
        for (i, ch) in s.chars().enumerate() {
            println!("{} {} {}", i, ch, char_count.get(&ch).unwrap());
            if char_count.get(&ch).unwrap() == &1 {
                index = i as i32;
                break;
            }
        }
        return index;
    }
}
fn main() {
    Solution::first_uniq_char(String::from("leetcode"));
    println!("{}", &1 == &1);
    println!("Hello, world!");
}
