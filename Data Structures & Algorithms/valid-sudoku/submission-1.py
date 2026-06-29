class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = [0] * 9
        cols = [0] * 9
        sq = [0] * 9 
        for i in range(9):
            for j in range(9):
                val = board[i][j]
                if val == ".":
                    continue
                val = int(val)    
                idx = 3 * (i // 3) + (j // 3)
                mask = 1 << val
                if rows[i] & mask:
                    return False
                if cols[j] & mask:
                    return False
                if sq[idx] & mask:
                    return False       
                rows[i] |= mask
                cols[j] |= mask
                sq[idx] |= mask
        return True