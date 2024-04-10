//
// Created by sunxy on 2024/3/23.
//

#include "test/2549_distinctIntegers.h"
#include "test/6_convert.h"
#include <gtest/gtest.h>

TEST(test_my_class, get_age) {
  double a = 0.1151;
  printf("%.2f%%", a*100);
}

TEST(_6_convert, tt1) {
  _6_convert solution;

  ASSERT_TRUE(solution.convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
      << "convert is not PAHNAPLSIIGYIR";
}