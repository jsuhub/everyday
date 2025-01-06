use std::collections::HashSet;

struct Solution;

impl Solution {
    pub fn sum_counts(nums: Vec<i32>) -> i32 {
        let mut arr: Vec<i32> = Vec::new();
        for start in 0..nums.len() {
            for end in start..nums.len() {
                // arr.push(nums[start..=end].to_vec());
                let dist_nums = Self::count_distinct(nums[start..=end].to_vec());
                arr.push(dist_nums * dist_nums);
            }
        }
        arr.iter().sum()
    }
    pub fn count_distinct(nums: Vec<i32>) -> i32 {
        let unique_set: HashSet<i32> = nums.iter().cloned().collect();
        return unique_set.len() as i32;
    }
}

fn main() {
    let result = Solution::sum_counts(vec![1, 2, 3]);
    println!("{:?}", result);
}
