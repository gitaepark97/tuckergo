package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		fileList, err := GetFileList(path)	// 파일 목록 가져오기
		if err != nil {
			fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
			return
		}

		fmt.Println("찾으려는 파일 리스트")
		for _, name := range fileList {
			fmt.Println(name)
		}
	}
}

func main() {
	if len(os.Args) < 3 {	// 실행 인수 개수 확인
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	word := os.Args[1]	// 실행 인수 가져오기
	files := os.Args[2:]
	fmt.Println("찾으려는 단어:", word)
	PrintAllFiles(files)
}