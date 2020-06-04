// +----------------------------------------------------------------------
// | ssh方法
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年04月02日
// +----------------------------------------------------------------------

package gohelp

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"os/exec"
	"strconv"
	"time"
)

type cli struct {
	IP       string //IP地址
	Username string //用户名
	Password string //密码
	Port     int    //端口号
}

func (c cli) ssh() (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(c.Password))
	clientConfig = &ssh.ClientConfig{
		User:    c.Username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr = fmt.Sprintf("%s:%d", c.IP, c.Port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

// 执行远程SSH命令行
// info 服务器信息
// cmd 执行的命令行
func ExecSSH(info map[string]string, cmd string) ([]byte, error) {
	port, _ := strconv.Atoi(info["port"])
	cli := cli{IP: info["ip"], Username: info["username"], Password: info["password"], Port: port}
	session, err := cli.ssh()
	defer session.Close()
	if err != nil {
		return nil, err
	}
	buf, e := session.Output(cmd)
	return buf, e
}

// 执行本地命令行
// cmd 执行的命令行
func ExecCommand(c string) (string, error) {
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("sh")
	cmd.Stdin = in
	in.WriteString(c + "\n")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	}
	return out.String(), nil
}
