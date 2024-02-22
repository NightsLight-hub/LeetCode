cmake_minimum_required(VERSION 2.8)
project(gstreamer_cmake )

set(CMAKE_CXX_STANDARD 14)

#set(CMAKE_BUILD_TYPE "Release")
set(CMAKE_BUILD_TYPE "Debug")
set(CMAKE_CXX_FLAGS_DEBUG "$ENV{CXXFLAGS} -O0 -Wall -g -ggdb")
set(CMAKE_CXX_FLAGS_RELEASE "$ENV{CXXFLAGS} -O3 -Wall")

# https://cmake.org/cmake/help/latest/module/FindPkgConfig.html
include(FindPkgConfig)
pkg_check_modules(GST REQUIRED gstreamer-1.0)
pkg_check_modules(GSTAPP REQUIRED gstreamer-app-1.0)

# add all files in directory "package" to variable "DIR_SRCS"
aux_source_directory(package DIR_SRCS)
message(STATUS ${DIR_SRCS})
include_directories(${GST_INCLUDE_DIRS} ${GSTAPP_INCLUDE_DIRS} "./include")
link_directories( ${GST_LIBRARY_DIRS} ${GSTAPP_LIBRARY_DIRS})

set(RESTWEB_LIBRARIES "-lboost_system -lcrypto -lssl -lcpprest")

# first executable, it must be defined before target link it!
add_executable(${PROJECT_NAME}  ./main.cpp ${DIR_SRCS})
target_include_directories(${PROJECT_NAME} PUBLIC "${PROJECT_BINARY_DIR}")
target_link_libraries(${PROJECT_NAME} ${GST_LIBRARIES} ${GSTAPP_LIBRARIES} ${RESTWEB_LIBRARIES} -lpthread -lm)


# second executable
add_executable(entrypoint1 ./entrypoints/entrypoint1.cpp ${DIR_SRCS})
target_include_directories(entrypoint1 PUBLIC "${PROJECT_BINARY_DIR}")
target_link_libraries(entrypoint1 ${GST_LIBRARIES} ${GSTAPP_LIBRARIES} ${RESTWEB_LIBRARIES} -lpthread -lm)

