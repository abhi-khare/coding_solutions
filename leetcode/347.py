def get_k_frequent_element(nums, k):
	
    count_map = {}

    for ele in nums:
        if ele not in count_map:
            count_map[ele] = 1
        else:
            count_map[ele] += 1
    
    count_map = [ (key,value) for key,value in count_map.items()]

    count_map.sort(key=lambda x: x[1], reverse=True)

    ans = [ele[0] for ele in count_map[:k]]
    return ans

nums = [1,1,1,2,2,3]
k = 2 

k_frequent_ele = get_k_frequent_element(nums, k)

print(k_frequent_ele)
