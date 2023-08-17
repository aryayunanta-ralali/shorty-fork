// Package generator
package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

func questionType() string {
	var answer string
	scanner := bufio.NewScanner(os.Stdin)

	question := Question{
		Name:         "file-type",
		AskText:      "What type of the file is want to be generated ? [options: ucase | repository]: ",
		DefaultValue: "",
	}

	for {
		fmt.Print(question.AskText)
		scanner.Scan()
		answer = scanner.Text()
		count := len(strings.TrimSpace(answer))

		if count == 0 {
			continue
		}

		break
	}

	return answer
}

func BuildProject() {

	fileType := questionType()
	switch fileType {
	case "ucase":
		UcasePackageCreate()

	case "repository":
		RepositoryPackageCreate()

	default:
		fmt.Println("can't find template for provided package!")
	}

}

func RepositoryPackageCreate() {
	var (
		fileType       = "repository"
		fileName       string
		isCreateEntity string

		repoName = "shorty"
		qa       = mappedQuestion[fileType]
		stub     = mapStub[fileType]
	)

	scanner := bufio.NewScanner(os.Stdin)
	for _, v := range qa {
		for {
			fmt.Print(v.AskText)
			scanner.Scan()
			answer := scanner.Text()

			count := len(strings.TrimSpace(answer))
			if count == 0 {
				continue
			}

			if v.Name == "file-name" {
				fileName = answer
			}

			if v.Name == "create-entity" {
				isCreateEntity = answer
			}

			break
		}

	}

	dst := strings.Trim(stub.DestDir, "/")

	if dst != "" {
		DirectoryCreate(fmt.Sprintf("./%s", dst))
	}

	modelName := strings.ReplaceAll(strings.Title(strings.ReplaceAll(fileName, "_", " ")), " ", "")
	firstCharacter := string(modelName[0])
	lowerFirstCharacter := strings.ToLower(firstCharacter)
	functionName := lowerFirstCharacter + modelName[1:]

	if isCreateEntity == "yes" {
		stubEntity := mapStub["entity"]
		dstEntity := strings.Trim(stubEntity.DestDir, "/")
		for _, s := range stubEntity.SourceFiles {
			stubFile, e := ioutil.ReadFile(fmt.Sprintf("%s/%s", stubEntity.SourceDir, s))
			if e != nil {
				log.Fatal("error : ", e.Error())
			}

			v := util.Replacer(map[string]string{
				"{{tableName}}": functionName,
				"{{modelName}}": modelName,
				"{{fileName}}":  fileName,
			}, string(stubFile))

			fileName := fmt.Sprintf("./%s/%s", dstEntity, fileName+".go")
			ioutil.WriteFile(fileName, []byte(v), os.ModePerm)
			fmt.Println("Success create file: ", fileName)
		}
	}

	for _, s := range stub.SourceFiles {
		stubFile, e := ioutil.ReadFile(fmt.Sprintf("%s/%s", stub.SourceDir, s))
		if e != nil {
			log.Fatal("error : ", e.Error())
		}

		v := util.Replacer(map[string]string{
			"{{repoName}}":     repoName,
			"{{modelName}}":    modelName,
			"{{functionName}}": functionName,
			"{{fileName}}":     fileName,
		}, string(stubFile))

		fileName := fmt.Sprintf("./%s/%s", dst, fileName+".go")
		ioutil.WriteFile(fileName, []byte(v), os.ModePerm)
		fmt.Println("Success create file: ", fileName)
	}
}

func UcasePackageCreate() {
	var (
		fileType             = "ucase"
		packageName          string
		functionMethod       string
		functionName         string
		isCreatePresentation string

		repoName = "shorty"
		qa       = mappedQuestion[fileType]
		stub     = mapStub[fileType]
	)

	scanner := bufio.NewScanner(os.Stdin)
	for _, v := range qa {
		for {
			fmt.Print(v.AskText)
			scanner.Scan()
			answer := scanner.Text()

			count := len(strings.TrimSpace(answer))
			if count == 0 {
				continue
			}

			if v.Name == "package-name" {
				packageName = answer
			}

			if v.Name == "function-method" {
				if !util.InArray(answer, []string{"get", "post", "put"}) {
					fmt.Println("provided method is not supported. please fill with another method!")
					continue
				}

				functionMethod = answer
			}

			if v.Name == "function-name" {
				functionName = answer
			}

			if v.Name == "is-create-presentation" {
				isCreatePresentation = answer
			}

			break
		}

	}

	fullPath := stub.DestDir + "/" + packageName
	dst := strings.Trim(fullPath, "/")

	if dst != "" {
		DirectoryCreate(fmt.Sprintf("./%s", dst))
	}

	upperFunctionName := strings.ReplaceAll(strings.Title(strings.ReplaceAll(functionName, "_", " ")), " ", "")

	if isCreatePresentation == "yes" {
		stubPresentations := mapStub["presentations"]
		dstPresentations := strings.Trim(stubPresentations.DestDir, "/")
		for _, s := range stubPresentations.SourceFiles {
			if !strings.Contains(s, functionMethod) {
				continue
			}

			stubFile, e := ioutil.ReadFile(fmt.Sprintf("%s/%s", stubPresentations.SourceDir, s))
			if e != nil {
				log.Fatal("error : ", e.Error())
			}

			v := util.Replacer(map[string]string{
				"{{upperFunctionName}}": upperFunctionName,
			}, string(stubFile))

			fileName := fmt.Sprintf("./%s/presentation_%s", dstPresentations, functionName+".go")
			ioutil.WriteFile(fileName, []byte(v), os.ModePerm)
			fmt.Println("Success create file: ", fileName)
		}
	}

	for _, s := range stub.SourceFiles {
		if !strings.Contains(s, functionMethod) {
			continue
		}

		b, e := ioutil.ReadFile(fmt.Sprintf("%s/%s", stub.SourceDir, s))
		if e != nil {
			log.Fatal("error : ", e.Error())
		}

		lowerFunctionName := strings.ToLower(upperFunctionName[0:1]) + upperFunctionName[1:]
		v := util.Replacer(map[string]string{
			"{{repoName}}":          repoName,
			"{{packageNamespace}}":  packageName,
			"{{functionName}}":      functionName,
			"{{lowerFunctionName}}": lowerFunctionName,
			"{{upperFunctionName}}": upperFunctionName,
		}, string(b))

		if strings.Contains(s, "_test") {
			functionName += "_test"
		}

		fileName := fmt.Sprintf("./%s/%s", dst, functionName+".go")
		ioutil.WriteFile(fileName, []byte(v), os.ModePerm)
		fmt.Println("Success create file: ", fileName)
	}
}

func DirectoryCreate(paths ...string) {
	for _, p := range paths {
		fmt.Println("Create dir :", p)
		os.MkdirAll(p, os.ModePerm)
	}
}
