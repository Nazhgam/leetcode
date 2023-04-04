func nextPermutation(nums []int) {
	// var mustChanged int
	fmt.Println(nums)
	var helper []int
	var changeIndx int
	length := len(nums)
	if length == 1 {
		return
	}
	if length == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}
	for i := length - 1; i >= 1; i-- {
		helper = append(helper, nums[i])
		if nums[i] > nums[i-1] {
			changeIndx = i - 1
			break
		}
		if i == 1 && nums[i] < nums[i-1] {
			sortArr(nums)
			return
		}
	}
	// fmt.Printf("helper[]: %v, changeIndx: %d\n", helper, changeIndx)
	for i := range helper {
		if nums[changeIndx] < helper[i] {
			nums[changeIndx], helper[i] = helper[i], nums[changeIndx]
			break
		}
	}
	// fmt.Printf("step2 --- helper[]: %v, changeIndx: %d\n", helper, changeIndx)
	sortArr(helper)
	var j = 0
	for i := changeIndx + 1; i < length; i++ {
		nums[i] = helper[j]
		j++
	}
}

func sortArr(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return
}