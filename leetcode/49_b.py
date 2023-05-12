def get_anagram_groups(strs: list) -> list:

    hash_map = {}

    strs = [(word, ''.join(sorted(word))) for word in strs]

    strs.sort(key=lambda x: x[1])

    for ele in strs:
        if ele[1] not in hash_map:
            hash_map[ele[1]] = [ele[0]]
        else:
            hash_map[ele[1]].append(ele[0])

    return list(hash_map.values())


strs = ["eat","tea","tan","ate","nat","bat"]

group_anagrams = get_anagram_groups(strs)

print(group_anagrams)
