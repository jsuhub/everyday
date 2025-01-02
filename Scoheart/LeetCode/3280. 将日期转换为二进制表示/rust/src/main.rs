struct Solution;

impl Solution {
    pub fn convert_date_to_binary(date: String) -> String {
        let slice = &date[0..4];
        let slice2 = &date[5..7];
        let slice3 = &date[8..10];
        let result = octal_to_binary(slice.to_string());
        let result2 = octal_to_binary(slice2.to_string());
        let result3 = octal_to_binary(slice3.to_string());
        let result4 = format!("{}-{}-{}", result, result2, result3);
        return result4;
    }
}

fn octal_to_binary(num: String) -> String {
    let mut bits = 0;
    let mut temp = num.parse::<i32>().unwrap();
    while temp > 0 {
        temp = temp / 2;
        bits += 1;
    }
    let mut vec = Vec::with_capacity(bits);
    let mut temp2 = num.parse::<i32>().unwrap();
    while temp2 > 0 {
        vec.push(temp2 % 2);
        temp2 = temp2 / 2;
    }

    vec.reverse();
    let result = vec.iter().map(|x| x.to_string()).collect::<String>();
    return result;
}

fn main() {
    let date = String::from("2080-02-29");
    let result = Solution::convert_date_to_binary(date);
    assert_eq!(result, "100000100000-10-11101");
}