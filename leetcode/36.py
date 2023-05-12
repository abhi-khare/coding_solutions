
def check_validity(board):

    row_maps = [{} for i in range(9)]
    col_maps = [{} for i in range(9)]
    group_maps = [{} for i in range(9)]


    for jdx, row in enumerate(board):
        for idx, _ in enumerate(row):

            ele = board[jdx][idx]

            if ele == '.':
                continue
            if ele in row_maps[idx]:
                return False
            else:
                row_maps[idx][ele] = 1
            
            if ele in col_maps[jdx]:
                return False
            else:
                col_maps[jdx][ele] = 1
            
            group_idx = int(idx/3) + 3*int(jdx/3) 
            if ele in group_maps[group_idx]:
                return False
            else:
                group_maps[group_idx][ele] = 1
    
    return True


board = [["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]


status = check_validity(board)

print(status)