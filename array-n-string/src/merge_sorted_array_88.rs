use std::error::Error;

/// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
/// and two integers m and n, representing the number of elements in nums1 and nums2
/// respectively.
///
/// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
///
/// The final sorted array should not be returned by the function, but instead be stored
/// inside the array nums1.To accommodate this, nums1 has a length of m + n, where the
/// first m elements denote the elements that should be merged, and the last n elements
/// are set to 0 and should be ignored. nums2 has a length of n.
///
/// ## Example 1:
/// - Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
/// - Output: [1,2,2,3,5,6]
/// - Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
/// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
///
/// ## Example 2:
/// - Input: nums1 = [1], m = 1, nums2 = [], n = 0
/// - Output: [1]
/// - Explanation: The arrays we are merging are [1] and [].
/// The result of the merge is [1].
///
/// ## Example 3:
/// - Input: nums1 = [0], m = 0, nums2 = [1], n = 1
/// - Output: [1]
/// - Explanation: The arrays we are merging are [] and [1].
/// The result of the merge is [1].
/// Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the
/// merge result can fit in nums1.
///
/// ## Constraints:
///
/// - nums1.length == m + n
/// - nums2.length == n
/// - 0 <= m, n <= 200
/// - 1 <= m + n <= 200
/// - -109 <= nums1[i], nums2[j] <= 109
///
/// ### Follow up
/// Can you come up with an algorithm that runs in O(m + n) time?
pub fn test() -> Result<(), Box<dyn Error>> {
    let mut case1_nums1 = vec![1, 2, 3, 0, 0, 0];
    merge_and_sort(case1_nums1.as_mut(), 3, vec![2, 5, 6], 3)?;
    println!("Case 1: {:?}", case1_nums1);

    let mut case2_nums1 = vec![1];
    merge_and_sort(case2_nums1.as_mut(), 1, vec![], 0)?;
    println!("Case 2: {:?}", case2_nums1);

    let mut case3_nums1 = vec![0];
    merge_and_sort(case3_nums1.as_mut(), 0, vec![1], 1)?;
    println!("Case 3: {:?}", case3_nums1);

    Ok(())
}

fn merge_and_sort(
    nums1: &mut Vec<i32>,
    m: i32,
    nums2: Vec<i32>,
    n: i32,
) -> Result<(), Box<dyn Error>> {
    let mut combined = nums1
        .iter()
        .take(m as usize)
        .cloned()
        .chain(nums2.iter().cloned())
        .collect::<Vec<_>>();

    combined.sort();
    nums1.clear();
    nums1.extend(combined.into_iter().take((m + n) as usize));

    Ok(())
}
