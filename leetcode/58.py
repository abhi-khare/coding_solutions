class Solution:
    def lengthOfLastWord(self, s: str) -> int:

        word_len = 0
        flag = 0
        for idx in range(len(s)-1,-1,-1):
            if s[idx] != ' ':
                flag = 1
                word_len +=1
            if flag and s[idx] == ' ':
                break
        
        return word_len


s = " "

sol = Solution()
word_len = sol.lengthOfLastWord(s)
print(word_len)