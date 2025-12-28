fn main() {
    println!("Hello, world!");
}

pub fn two_sum_brute(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut res = Vec::new();
    for i in 0..nums.len() {
        for j in (i + 1)..nums.len() {
            if nums[i] + nums[j] == target {
                res = vec![i as i32, j as i32];
            }
        }
    }
    res
}

pub fn two_sum_hashmap(nums: Vec<i32>, target: i32) -> Vec<i32> {
    use std::collections::HashMap;

    let mut res = Vec::new();
    let mut map: HashMap<i32, i32> = HashMap::new();

    nums.iter().enumerate().for_each(|(k, v)| {
        map.insert(*v, k as i32);
    });

    for (i, v) in nums.iter().enumerate() {
        let maybe_pair = target - *v;
        if let Some(&j) = map.get(&maybe_pair)
            && j != i as i32 {
                res.push(i as i32);
                res.push(j);
                break;
            }
    }
    res
}

pub fn two_sum_one_pass(nums: Vec<i32>, target: i32) -> Vec<i32> {
    use std::collections::HashMap;

    let mut map: HashMap<i32, i32> = HashMap::with_capacity(nums.len());

    for (i, v) in nums.iter().enumerate() {
        let int_i = i as i32;
        let maybe_pair = target - *v;
        if let Some(&j) = map.get(&maybe_pair)
            && j != int_i {
                return vec![j, int_i];
            }
        map.insert(*v, int_i);
    }
    vec![]
}

pub fn max_profit(prices: Vec<i32>) -> i32 {
    let mut min_val = prices[0];
    let mut max_profit = 0;

    for &price in &prices[1..] {
        if price < min_val {
            min_val = price;
            continue;
        }

        max_profit = i32::max(max_profit, price - min_val);
    }

    max_profit
}
