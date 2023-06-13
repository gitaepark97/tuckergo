package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct {	// 찾은 결과 정보
	lineNo 	int
	line 		string
}

// 파일 내 라인 정보
type FindInfo struct {
	filename 	string
	lines 		[]LineInfo
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{ filename, []LineInfo{} }
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)

		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)		// 스캐너를 만듭니다.
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {	// 한 줄씩 읽으며 단어 포함 여부 검색
			findInfo.lines = append(findInfo.lines, LineInfo{ lineNo, line })
		}
		lineNo++
	}

	return findInfo
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllfiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)

		return findInfos
	}
	for _, filename := range filelist {	// 각 파일별로 검색
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}

	return findInfos
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.3 word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}
	for _, path := range files {
		// 파일 찾기
		findInfos = append(findInfos, FindWordInAllfiles(word, path)...)
	}
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------")
		fmt.Println()
	}
}