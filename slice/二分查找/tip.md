### 二分查找

## 使用的前提条件：
1. 有序数组
2. 元素唯一

## 注意
1. 左右区间规定好：left，right
2. 中间值位置确定好：middle
3. 结束循环的条件
4. 查找判断的条件

## 写法
1. `[left,right]` , 在左闭右闭的区间中
   - `for left <= right` ,因为此时left == right有意义
   - `if nums[middle]>target {right = middle-1}` ,因为此时middle位置的值一定不是target，右边界可以进一步缩小
2. `[left,right)` ,在左闭右开的区间中
   - `for left < right` ,因为此时left == right无意义
   - `if nums[middle]>target {right = middle}` ,因为此时right作为区间开边界，下次判断不会找到right