/*
@Time : 2021/7/16 16:00
@Author : sunxy
@File : main
@description:
*/
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal("command can not run!")
	}
}

// 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func exec_shell(s string) (string, error) {
	// 函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	// 读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)
	return out.String(), err
}

func checkKvmName(cmd string) *[]string {
	outPut, error := exec_shell(cmd)
	hostSlice := make([]string, 0)
	checkErr(error)
	fileSlice := strings.Split(outPut, "\n")
	backName := os.Args[1:]
	for _, name := range backName {
		for _, kvmName := range fileSlice {
			kvmRealName := strings.Split(kvmName, ".")
			if name == kvmRealName[0] {
				hostSlice = append(hostSlice, kvmName)
			}
		}
	}
	return &hostSlice
}

// 清除字符串空格的函数
func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	return strings.Join(strings.Fields(input), "")
}

// 查找kvm虚拟机所属磁盘路径
func diskPath() *[]string {
	diskSlice := make([]string, 0)
	cmd1 := "ls -l /etc/libvirt/qemu |grep -v 'total' |awk '{print $NF}'"
	kvmHostName := checkKvmName(cmd1)
	for _, kvmName := range *kvmHostName {

		cmd2 := "cat /etc/libvirt/qemu/" + kvmName + "| grep 'source file'"
		outPut, _ := exec_shell(cmd2)
		aa := StringStrip(outPut)
		//		diskNameSlice := strings.Split(stdout, "\n")
		name := strings.Split(aa, "'")
		for _, i := range name {
			r, _ := regexp.Compile("^/data.*")
			diskNameStr := r.FindStringSubmatch(i)
			for _, name := range diskNameStr {
				diskSlice = append(diskSlice, name)

			}

		}

	}
	return &diskSlice
}

// 物理机hostName
func getHostName() string {
	outPut, error := exec_shell("echo $HOSTNAME")
	checkErr(error)
	return outPut
}

func main() {
	value := diskPath()
	baseHostName := strings.TrimSpace(getHostName())
	for _, i := range *value {
		cmd := fmt.Sprintf("restic -r sftp:root@172.18.60.17:/data/kvm_back/%v  backup %v ", baseHostName, i)
		fmt.Print(cmd)
		outPut, error := exec_shell(cmd)
		checkErr(error)
		fmt.Print(outPut)

	}

}
