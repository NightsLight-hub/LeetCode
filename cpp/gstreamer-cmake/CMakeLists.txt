cmake_minimum_required(VERSION 3.16)
SET(CMAKE_C_COMPILER /usr/bin/gcc)
SET(CMAKE_CXX_COMPILER /usr/bin/g++)
set(CMAKE_CXX_STANDARD 17)
project(gstreamer_cmake )


#set(CMAKE_BUILD_TYPE "Release")
set(CMAKE_BUILD_TYPE "Debug")
set(CMAKE_CXX_FLAGS_DEBUG "$ENV{CXXFLAGS} -O0 -Wall -g -ggdb")
set(CMAKE_CXX_FLAGS_RELEASE "$ENV{CXXFLAGS} -O3 -Wall")

# https://cmake.org/cmake/help/latest/module/FindPkgConfig.html
include(FindPkgConfig)
find_package(GTest REQUIRED)
# GTest ��ͷ�ļ�
pkg_check_modules(GST REQUIRED gstreamer-1.0)
pkg_check_modules(GSTAPP REQUIRED gstreamer-app-1.0)
# add all files in directory "package" to variable "DIR_SRCS"
file(GLOB_RECURSE  DIR_SRCS ${CMAKE_SOURCE_DIR}/package/*.cpp)
message(STATUS "src file list: ")
foreach (srcFile IN LISTS DIR_SRCS)
    message(STATUS ${srcFile} ",")
endforeach ()
message(STATUS "")

include_directories(${GST_INCLUDE_DIRS} ${GSTAPP_INCLUDE_DIRS} ${GTEST_INCLUDE_DIRS} ${CMAKE_SOURCE_DIR}/include)
link_directories( ${GST_LIBRARY_DIRS} ${GSTAPP_LIBRARY_DIRS})

set(RESTWEB_LIBRARIES "-lboost_system -lcrypto -lssl -lcpprest")

# <------------ add hiredis dependency --------------->
find_path(HIREDIS_HEADER hiredis PATHS /usr/local/include/hiredis)
find_library(HIREDIS_LIB hiredis)
# <------------ add redis-plus-plus dependency -------------->
find_path(REDIS_PLUS_PLUS_HEADER sw)
find_library(REDIS_PLUS_PLUS_LIB redis++)

# first executable, it must be defined before target link it!
add_executable(${PROJECT_NAME}  ./main.cpp ${DIR_SRCS})
target_include_directories(${PROJECT_NAME} PUBLIC "${PROJECT_BINARY_DIR}" "${HIREDIS_HEADER}" ${REDIS_PLUS_PLUS_HEADER})
target_link_libraries(${PROJECT_NAME} ${GST_LIBRARIES} ${GSTAPP_LIBRARIES} ${RESTWEB_LIBRARIES} ${HIREDIS_LIB} ${REDIS_PLUS_PLUS_LIB} -lpthread -lm)

# second executable
add_executable(entrypoint1 ./entrypoints/entrypoint1.cpp ${DIR_SRCS})
target_include_directories(entrypoint1 PUBLIC "${PROJECT_BINARY_DIR}" "${HIREDIS_HEADER}" ${REDIS_PLUS_PLUS_HEADER} )
target_link_libraries(entrypoint1 ${GST_LIBRARIES} ${GSTAPP_LIBRARIES} ${RESTWEB_LIBRARIES} -lpthread -lm ${HIREDIS_LIB} ${REDIS_PLUS_PLUS_LIB})

# third executable
add_executable(entrypoint2 ./entrypoints/entrypoint2.cpp ${DIR_SRCS})
target_include_directories(entrypoint2 PUBLIC "${PROJECT_BINARY_DIR}" "${HIREDIS_HEADER}" ${REDIS_PLUS_PLUS_HEADER} )
target_link_libraries(entrypoint2 ${GST_LIBRARIES} ${GSTAPP_LIBRARIES} ${RESTWEB_LIBRARIES} ${HIREDIS_LIB} ${REDIS_PLUS_PLUS_LIB} -lpthread -lm)

# https://blog.csdn.net/joelcat/article/details/90766192

message(STATUS "GTEST_MAIN_LIBRARIES is " ${GTEST_MAIN_LIBRARIES} " " ${GTEST_LIBRARIES} )
message(STATUS "gtest header is " ${GTEST_INCLUDE_DIRS})
enable_testing()

add_executable(test_my_class ./entrypoints/test_my_class.cpp)
target_include_directories(test_my_class PUBLIC "${PROJECT_BINARY_DIR}" "${HIREDIS_HEADER}" ${REDIS_PLUS_PLUS_HEADER} )
target_link_libraries(test_my_class ${GST_LIBRARIES} ${GSTAPP_LIBRARIES} ${RESTWEB_LIBRARIES} ${HIREDIS_LIB} ${REDIS_PLUS_PLUS_LIB} ${GTEST_BOTH_LIBRARIES} -lpthread -lm)

#ADD_SUBDIRECTORY(lib)
ADD_LIBRARY(demolib SHARED demolib.cpp ${DIR_SRCS})
target_include_directories(demolib PUBLIC "${PROJECT_BINARY_DIR}" "${HIREDIS_HEADER}" ${REDIS_PLUS_PLUS_HEADER} )
target_link_libraries(demolib ${GST_LIBRARIES} ${GSTAPP_LIBRARIES} ${RESTWEB_LIBRARIES} ${HIREDIS_LIB} ${REDIS_PLUS_PLUS_LIB} ${GTEST_BOTH_LIBRARIES} -lpthread -lm)

gtest_discover_tests(test_my_class)