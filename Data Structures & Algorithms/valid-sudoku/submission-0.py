class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = defaultdict(set)
        cols = defaultdict(set)
        sq = defaultdict(set)
        for i in range(9):
            for j in range(9):
                val = board[i][j]
                if val == '.':
                    continue
                if val in rows[i]:
                    return False
                rows[i].add(val)
                if val in cols[j]:
                    return False
                cols[j].add(val)
                if val in sq[(i // 3, j // 3)]:
                    return False
                sq[(i // 3, j // 3)].add(val)
        return True