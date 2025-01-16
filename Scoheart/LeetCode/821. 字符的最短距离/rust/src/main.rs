impl Solution {
    pub fn shortest_to_char(s: String, c: char) -> Vec<i32> {
        let s = s.as_bytes(); // 将字符串转换为字节数组
        let n = s.len(); // 字符串的长度
        let mut answer = vec![0; n]; // 初始化结果数组
        let mut prev = i32::MIN / 2; // 初始化 prev 为一个很小的值

        // 从左到右遍历，计算每个字符到左边最近的字符 c 的距离
        for i in 0..n {
            if s[i] == c as u8 {
                prev = i as i32; // 更新 prev 为当前索引
            }
            answer[i] = (i as i32 - prev).abs(); // 计算距离
        }

        // 从右到左遍历，计算每个字符到右边最近的字符 c 的距离，并取最小值
        prev = i32::MAX / 2; // 重置 prev 为一个很大的值
        for i in (0..n).rev() {
            if s[i] == c as u8 {
                prev = i as i32; // 更新 prev 为当前索引
            }
            answer[i] = answer[i].min((i as i32 - prev).abs()); // 取最小值
        }

        answer // 返回结果数组
    }
}