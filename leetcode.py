def main(nums1, m, nums2, n):
    if n == 0:
        return
    if m == 0:
        for num in nums2:
            nums1[m] == num
            m += 1
            return
    lastElementNums1 = m - 1
    lastElementNums2 = n - 1
    a = m + n - 1
    while lastElementNums2 >= 0 and lastElementNums1 >= 0:
        # nums1 bigger
        if nums1[lastElementNums1] > nums2[lastElementNums2]:
            nums1[a] = nums1[lastElementNums1]
            lastElementNums1 -= 1
        # nums2 bigger or equal
        else:
            nums1[a] = nums2[lastElementNums2]
            lastElementNums2 -= 1
        a -= 1
    # nums1 no numbers left
    while lastElementNums2 >= 0:
        nums1[a] = nums2[lastElementNums2]
    

nums1 = [0]
m = 0
nums2 = [1]
n = 1
print(nums1)
main(nums1, m, nums2, n)
print(nums1)