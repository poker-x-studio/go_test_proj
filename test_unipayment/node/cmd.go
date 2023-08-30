package node

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// 调用js的签名函数
func Call_javascript_sign(client_id string, client_secret string, uri string, requestMethod string, body string) (string, error) {
	exe_path, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("获取可执行文件路径失败:%s", err.Error())
	}
	exe_dir := filepath.Dir(exe_path)

	filepath := path.Join(exe_dir, "../javascript/index.js")
	cmd := exec.Command("node", filepath, client_id, client_secret, uri, requestMethod, body)

	var std_out bytes.Buffer
	cmd.Stdout = &std_out

	//处理错误日志
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	go handler_std_err(errReader)

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	//等待结束
	cmd.Wait()

	cmd_result := std_out.String()
	fmt.Println("cmd_result: " + cmd_result)

	sign := "sign:"
	if strings.Contains(cmd_result, sign) {
		index := strings.Index(cmd_result, sign)
		extract_sign := cmd_result[index+len(sign) : len(cmd_result)-1]
		fmt.Println(index)
		extract_sign = strings.TrimSpace(extract_sign)
		fmt.Println(extract_sign)
		return extract_sign, nil
	}

	return "", errors.New("错误")
}

// 处理错误信息
func handler_std_err(errReader io.ReadCloser) {
	//logrus_entry := log.New_logrus_entry("handler_std_err()")
	in := bufio.NewScanner(errReader)
	for in.Scan() {
		str := string(in.Bytes())
		//logrus_entry.Errorf("JS,%s", str)
		fmt.Println(str)
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
