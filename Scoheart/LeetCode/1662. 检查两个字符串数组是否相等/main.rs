impl Solution {
    pub fn array_strings_are_equal(word1: Vec<String>, word2: Vec<String>) -> bool {
        let str1 = word1.join("");
        let str2 = word2.join("");
        return str1 == str2;
    }
}