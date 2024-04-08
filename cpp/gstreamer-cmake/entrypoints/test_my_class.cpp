//
// Created by sunxy on 2024/3/23.
//

#include <gtest/gtest.h>
#include "test/2549_distinctIntegers.h"
#include "test/6_convert.h"

TEST(test_my_class, get_age)
{
  Solution solution;

  ASSERT_TRUE(solution.distinctIntegers(16) == 15) << "age is not 15";
}


TEST(_6_convert, tt1)
{
  _6_convert solution;

  ASSERT_TRUE(solution.convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR") << "convert is not PAHNAPLSIIGYIR";
}