package singlenum

func singleNumber(nums []int) int {
    set := map[int]bool{}

    for _, n := range nums {
        if set[n] {
            delete(set, n)
            continue
        }

        set[n] = true
    }

    for n := range set {
        return n
    }

    return -1
}
