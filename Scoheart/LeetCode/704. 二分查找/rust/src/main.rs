struct Solution;

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut left = 0;
        let mut right = nums.len();
        let mut result = -1;
        while left < right {
            let middle = left + (right - left) / 2;
            if target > nums[middle] {
                left = middle + 1;
            } else if target < nums[middle] {
                right = middle;
            } else {
                result = middle as i32;
                break;
            }
        }
        result
    }
}


fn main() {
    let re = Solution::search(vec![-1, 0, 3, 5, 9, 12], 9);
    println!("{}", re);
}
