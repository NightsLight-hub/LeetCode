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
 * ��һ�������ַ��� s ���ݸ��������� numRows ���Դ������¡������ҽ��� Z �������С�

���������ַ���Ϊ "PAYPALISHIRING" ����Ϊ 3 ʱ���������£�

    P   A   H   N
    A P L S I I G
    Y   I   R
  ֮����������Ҫ�����������ж�ȡ��������һ���µ��ַ��������磺"PAHNAPLSIIGYIR"��

    ����ʵ��������ַ�������ָ�������任�ĺ�����

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
