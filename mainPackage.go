package gengo

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const mainPackageReplacement = "gengomain"
const mainPackageReplacementPath = gengoDir + "/" + mainPackageReplacement

func (s *Service) setupTempMainPackage() error {
	var pkgName string

	newMainPkgPath := path.Join(s.getGenGoDirPath(), mainPackageReplacement)
	if _, err := os.Stat(newMainPkgPath); err == nil {
		err := os.RemoveAll(newMainPkgPath)
		if err != nil {
			return err
		}
	}
	err := os.Mkdir(newMainPkgPath, 0755)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(s.WorkDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {

			fileName := path.Join(s.WorkDir, file.Name())
			fileData, err := os.ReadFile(fileName)
			if err != nil {
				return err
			}

			fileData = []byte(strings.ReplaceAll(string(fileData), "package main", fmt.Sprintf("package %s", mainPackageReplacement)))

			newFilePath := path.Join(newMainPkgPath, file.Name())
			err = os.WriteFile(newFilePath, fileData, 0644)
			if err != nil {
				return err
			}
		} else if strings.HasSuffix(file.Name(), ".mod") {
			fileName := path.Join(s.WorkDir, file.Name())
			fileData, err := os.ReadFile(fileName)
			if err != nil {
				return err
			}

			pkgName = strings.Replace(strings.Split(string(fileData), "\n")[0], "module ", "", 1)
		}
	}

	newMainPkg := path.Join(pkgName, mainPackageReplacementPath)

	s.tempMainPackageName = newMainPkg

	return nil
}

func (s *Service) cleanupTempMainPackage() error {
	newMainPkgPath := path.Join(s.getGenGoDirPath(), mainPackageReplacement)
	if _, err := os.Stat(newMainPkgPath); err == nil {
		err := os.RemoveAll(newMainPkgPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) GetMainPackageName() (string, error) {
	if s.tempMainPackageName == "" {
		err := s.setupTempMainPackage()
		if err != nil {
			return "", err
		}
	}
	return s.tempMainPackageName, nil
}
