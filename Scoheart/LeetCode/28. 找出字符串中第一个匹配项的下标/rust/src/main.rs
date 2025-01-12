struct Solution;

impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        let mut index = 0;
        for begin in 0..haystack.len() {
            for end in begin..haystack.len() {
                let sub = &haystack[begin..=end];
                if (sub == needle) {
                    index = begin;
                    return index as i32;
                }
            }
        }
        return -1;
    }
}

fn main() {
    let result = Solution::str_str(String::from("demo"), String::from("e"));
    // assert!(result == 0);
    println!("{}", result);
}
