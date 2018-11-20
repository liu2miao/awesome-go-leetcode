package problem0016

import "sort"
import "fmt"

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    result := nums[0] + nums[1] + nums[2]
    for i := range nums{
        if nums[i] >= target { break }
        for j := i + 1; j < len(nums)-1; j++{
            l := j + 1
            r := len(nums) - 1
            if l > r { break }
            tmpTarget := target - nums[i] - nums[j]
            tmp := nums[l]
            tmp2 := nums[r]
            for (l < r){
                m := (l + r ) / 2
                if tmpTarget == nums[m] {
                    return target
                } else if tmpTarget < nums[m] {
                    r = m - 1
                } else {
                    l = m + 1
                }
                tmp = nums[m]
            }
            if delta(result, target) > delta( nums[i] + nums[j] + tmp, target) {
                result = nums[i]+ nums[j] + tmp
                fmt.Println(result, nums[i], nums[j], tmp)
            }
            if delta(result, target) > delta( nums[i] + nums[j] + tmp2, target) {
                result = nums[i]+ nums[j] + tmp2
                fmt.Println(result, nums[i], nums[j], tmp2)
            }
            fmt.Println(result, nums[i], nums[j], tmp, "-------------")
            for j < len(nums) && nums[j-1] == nums[j] { j++ }

        }
        for i > 0 && i < len(nums) && nums[i-1] == nums[i] { i++ }
    }
    return result
}

func delta(a int, b int) int{
    if a > b {
        return a - b
    } else {
        return b - a
    }
}
