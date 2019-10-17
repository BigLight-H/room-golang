package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str) )
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

//发送邮件
//调用方法:
//mailTo := []string{
//"****@qq.com",
//"****@qq.com",
//}
//subject := "Hello"
//body := "Good"
//_ = SendEmail(mailTo, subject, body)
func SendEmail(mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": beego.AppConfig.String("set_email"),
		"pass": beego.AppConfig.String("set_pass"),
		"host": beego.AppConfig.String("set_stmp"),
		"port": beego.AppConfig.String("set_port"),
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "AXICOO.COM"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                                  //发送给多个用户
	m.SetHeader("Subject", subject)                               //设置邮件主题
	m.SetBody("text/html", body)                                  //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

func TimeSet() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// 自定义http get请求函数
// 入参：请求url
// 返回值：rlt，天气数据。err，错误信息
// 网络请求
func DoHttpGetRequest(url string) (rlt string, err error) {

	// http.Get在net/http中，所以要import "net/http"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	} else {
		// 使用efer resp.Body.Close()。当doHttpGetRequest成功return之后，执行此行语句。多用于句柄关闭
		defer resp.Body.Close()

		// io流数据读取。需要引用io/ioutil
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		} else {
			return string(body), err
		}

	}

}
//int64转string
func ChangeInt(number int64) (str string) {
	string := strconv.FormatInt(number,10)
	return string
}

//时间格式转换
func DateChange(str string) time.Time  {
	//转化所需模板

	timeLayout := "2006年01月02日 15:04:05"

	//重要：获取时区

	loc, _ := time.LoadLocation("Asia/Shanghai")

	//使用模板在对应时区转化为time.time类型

	theTime, _ := time.ParseInLocation(timeLayout, str, loc)

	return theTime
}