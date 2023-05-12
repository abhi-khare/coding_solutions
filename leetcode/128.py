

def get_max_sequence_len(nums):

    hashmap = {}

    for ele in nums:

        if ele in hashmap:
            continue

        previous_count = 0
        if ele-1 in hashmap:
            previous_count = hashmap[ele-1]
        
        hashmap[ele] = previous_count + 1

        next = ele+1
        seq_length = previous_count + 1
        while next in hashmap:
            hashmap[next] = seq_length + 1
            seq_length = hashmap[next]
            next +=1
            
    
    ans = -1000000
    for key,value in hashmap.items():
        if value > ans:
            ans= value


    return ans

nums = [100,4,200,2,3,5,0,1,3,2]
max_seq_len = get_max_sequence_len(nums)

print(max_seq_len)