//
// Created by sunxy on 2024/3/23.
//

#ifndef CLION_DEMO_6_CONVERT_H
#define CLION_DEMO_6_CONVERT_H

#include "cstdlib"
#include "cstring"
#include <iostream>
#include "vector"
using namespace std;

/**
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

    P   A   H   N
    A P L S I I G
    Y   I   R
  之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

    请你实现这个将字符串进行指定行数变换的函数：

    string convert(string s, int numRows);
 */

class _6_convert {
public:
  string convert(string s, int numRows) {
        if (numRows == 1) return s;
        vector<string> rows(min(numRows, int(s.size())));
        int curRow = 0;
        bool goingDown = false;
        for (char c : s) {
          rows[curRow] += c;
          if (curRow == 0 || curRow == numRows - 1) goingDown = !goingDown;
          curRow += goingDown ? 1 : -1;
        }
        string ret;
        for (string row : rows) {
          ret += row;
        }
        return ret;
  }
};

#endif // CLION_DEMO_6_CONVERT_H
