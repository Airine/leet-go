# 滑动窗口类型

滑动窗口类型的题目经常是用来执行数组或是链表上某个区间（窗口）上的操作。比如找最长的全为1的子数组长度。滑动窗口一般从第一个元素开始，一直往右边一个一个元素挪动。当然了，根据题目要求，我们可能有固定窗口大小的情况，也有窗口的大小变化的情况。

![](https://img-upic.oss-accelerate.aliyuncs.com/uPic/2020/05/ZcTYjw.jpg)

该图中，我们的窗子不断往右一格一个移动

## 下面是一些我们用来判断我们可能需要上滑动窗口策略的方法：

- 这个问题的输入是一些线性结构：比如链表呀，数组啊，字符串啊之类的
- 让你去求最长/最短子字符串或是某些特定的长度要求

## 题目

| No.  | Name                                                  | Hardness | Status             |
| ---- | ----------------------------------------------------- | -------- | ------------------ |
|      | Maximum Sum Subarray of Size K                        | easy     | :heavy_check_mark: |
|      | Smallest Subarray with a given sum                    | easy     |                    |
|      | Longest Substring with K Distinct Characters          | medium   |                    |
|      | Fruits into Baskets                                   | medium   |                    |
|      | No-repeat Substring                                   | hard     |                    |
|      | Longest Substring with Same Letters after Replacement | hard     |                    |
|      | Longest Subarray with Ones after Replacement          | hard     |                    |
