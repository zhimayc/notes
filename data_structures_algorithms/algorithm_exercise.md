算法题练习
===================================================================

1. 给定一个数组A[0,...n-1]，求A的连续子数组，使得该子数组的和最大。
    >例如：
    >
    >数组：1,-2,3,10,-4,7,2,-5
    >
    >最大子数组：3,10,-4,7,2


    >>解法：
    >>
    >> 1.暴力法

    ```
        int MaxSubArray(int * arr,int len,int * start,int * end) {
            int i, j, k;
            int maxSum = arr[0];
            int currSum;
            for (i = 0; i < len;i++) {
                for (j = i; j < len;j++) {
                    currSum = 0;
                    for (k = i; k <= j;k++) {
                        currSum += arr[k];
                    }

                    if (currSum > maxSum) {
                        maxSum = currSum;
                        *start = i;
                        *end = j;
                    }
                }
            }

            return maxSum;
        }

        int main()
        {
            int arr[] = {1,-2,3,10,-4,7,2,-5};
            int start = 0;
            int end = 0;
            int i;
            int len = sizeof(arr) / sizeof(int);
            int maxSum;
            maxSum = MaxSubArray(arr, len, &start, &end);
            
            for (i = start; i <= end;i++) {
                printf("%d ",arr[i]);
            }
            printf("\n");

            printf("maxSum = %d \n",maxSum);
            system("pause");
            return 0;
        }

    ```

    >> 2.分治法

    ```
        int max(int a,int b) {
            int max;
            return max = a > b ? a : b;
        }

        int max(int a,int b,int c) {
            int max;
            max = a > b ? a : b;
            max = c > max ? c : max;
            return max;
        }


        int MaxAddSub(int * arr,int from,int to) {
            if (to == from) {
                return arr[from];
            }

            int middle = (from+to)/2;
            int m1 = MaxAddSub(arr,from,middle);
            int m2 = MaxAddSub(arr,middle+1,to);
            int i, left = arr[middle], now = arr[middle];
            
            for (i = middle - 1; i >= from;--i) {
                now += arr[i];
                left = max(now,left);
            }

            int right = arr[middle+1];
            now = arr[middle+1];
            for (i = middle + 2; i <= to;++i) {
                now += arr[i];
                right = max(now,right);
            }

            int m3 = left + right;

            return max(m1,m2,m3);
        }

    ```

    >> 3. 动态规划法

    ```
        int MaxSubArray(int * arr,int len) {
            int result;
            int i,sum;
            result = arr[0];
            sum = arr[0];
            for (i = 1; i < len;i++) {
                if (sum > 0) {
                    sum += arr[i];
                }
                else {
                    sum = arr[i];
                }
                if (sum > result) {
                    result = sum;
                }
            }

            return result;
        }

    ```

