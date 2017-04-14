//
//date:2017-4-15
//describe:按行遍历整个数组，找到值为1的位置，检查这个值上下左右是否为1.
//			
//一点点技巧：记录下一行为1的位置,即nextRowStart。最终遍历的次数就是n.
//
int islandPerimeter(int** grid, int gridRowSize, int gridColSize) {
    if(gridRowSize <= 0 || gridColSize <= 0 || grid == NULL)
        return 0;
    int nextRowStart = 0;
    int sum = 0;
    for(int i = 0;i < gridRowSize;i++)
    {
        int start = nextRowStart;
        for(int j=start - 1;j>=0;j--)
        {
            if(grid[i][j] == 1)
            {
                if( j-1 < 0 || grid[i][j-1] != 1)
                    sum++;
                if( j+1 >= gridColSize || grid[i][j+1] != 1)
                    sum++;
                if( i-1 < 0 || grid[i-1][j] != 1)
                    sum++;
                if( i+1 >= gridRowSize || grid[i+1][j] !=1 )
                    sum++;
                else
                    nextRowStart = j;
            }
        }
        
        for(int j=start;j < gridColSize;j++)
        {
            if(grid[i][j] == 1)
            {
                
                if( j-1 < 0 || grid[i][j-1] != 1)
                    sum++;
                if( j+1 >= gridColSize || grid[i][j+1] != 1)
                    sum++;
                if( i-1 < 0 || grid[i-1][j] != 1)
                    sum++;
                if( i+1 >= gridRowSize || grid[i+1][j] !=1 )
                    sum++;
                else
                    nextRowStart = j;
            }
        }
    }
    return sum;
}
