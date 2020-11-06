package pattern

import (
	"fmt"
	"strings"
)

type Git interface {
	Clone(url string) bool
}

type AWSCodeCommit struct {}

func (a AWSCodeCommit)Clone(url string) bool {
	if strings.HasPrefix(url, "https") {
		fmt.Println("clone from CodeCommit: " + url)
		return true
	}

	fmt.Println("failed to clone from " + url)
	return false
}
type GitHub struct {}

func (g GitHub)Clone(url string) bool {
	if strings.HasPrefix(url, "https") {
		fmt.Println("clone from GitHub: " + url)
		return true
	}

	fmt.Println("failed to clone from " + url)
	return false
}
//代理类
type GitBash struct {
	bash Git
}

func (g GitBash)Clone(url string)bool  {
	return g.bash.Clone(url)
}

type Engineer struct {}

func (e Engineer) GetCode(url string,git int)  {
	gitBash := GetGit(git)
	if gitBash.Clone(url) {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}
func GetGit(t int) Git {
	if t == 1 {
		return GitBash{bash: GitHub{}}
	}else if t ==2{
		return GitBash{bash:AWSCodeCommit{}}
	}

	return nil // 可能还有其他的git源
}
func TestProxy()  {
	me := new(Engineer)
	me.GetCode("https://aaaaaa",1)
	me.GetCode("https://aaaaaa",2)

}