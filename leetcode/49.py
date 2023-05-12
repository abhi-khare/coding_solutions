
maps = {
	'a':1,'b':2,'c':3,'d':4,'e':5,'f':6,'g':7,'h':8,'i':9,'j':10,'k':11,'l':12,'m':13,'n':14,'o':15,'p':16,'q':17,'r':18,'s':19,'t':20,'u':21,'v':22,'w':23,'x':24,'y':25,'z':26
}

def get_hash_string(word: str) -> str:

	hash_map = [0]*26
	for char in word:
		hash_map[maps[char]-1] +=1

	hash_str = '#'.join( [str(count) for count in hash_map])

	return hash_str

def get_anagram_groups(strs: list) -> list:

	ans = []
	hash_Strings = [((get_hash_string(word)), word) for word in strs]

	hash_Strings.sort()

	

	current_hash, current_word = hash_Strings[0][0], hash_Strings[0][1]
	temp = [current_word]
	for hash, word in hash_Strings[1:]:
		if hash == current_hash:
			temp.append(word)
		else:
			ans.append(temp)
			temp = [word]
		
		current_hash, current_word = hash, word
	
	ans.append(temp)
	return ans



strs = ["bdddddddddd","bbbbbbbbbbc"]

group_anagrams = get_anagram_groups(strs)

print(group_anagrams)
