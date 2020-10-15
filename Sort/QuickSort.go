package Sort

func QuickSort(nums []int,start,end int){
	if start >= end {
		return
	}

	mid := partition(nums,start,end)
	QuickSort(nums,start,mid)
	QuickSort(nums,mid+1,end)
}
func partition(nums []int,start,end int)  int{
	temp:= nums[start]
	low := start
	height := end

	for low < height{
		for low < height && temp < nums[height] {
			height--
		}
		if low < height {
			nums[low] = nums[height]
		}

		for low < height && temp > nums[low] {
			low++
		}

		if low < height {
			nums[height] = nums[low]
		}
	}

	nums[low] = temp

	return low
}
