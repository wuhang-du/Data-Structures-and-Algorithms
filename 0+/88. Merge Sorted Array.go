//date: 2017-8-8
// 1,3,4
// 2,5
//
//

func merge(nums1 []int, m int, nums2 []int, n int)  {
    if len(nums2) == 0 || n == 0 {
        return
    }
    end := m + n - 1  
    for n > 0 {
        if m == 0 {
            copy(nums1,nums2[:n])
            return
        }
        if nums1[m - 1] < nums2[n - 1] {
            nums1[end] = nums2[n - 1]
            n--
        } else {
            nums1[end] = nums1[m - 1]
            m--
        }
        end--
    }
    return
}

