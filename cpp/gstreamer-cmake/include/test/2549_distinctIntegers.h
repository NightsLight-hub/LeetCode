//
// Created by sunxy on 2024/3/23.
// https://leetcode.cn/problems/count-distinct-numbers-on-board/description/?envType=daily-question&envId=2024-03-23
//

using namespace std;

class Solution {
public:
  int distinctIntegers(int n) {
    return n == 1 ? 1 : n - 1;
  }
};