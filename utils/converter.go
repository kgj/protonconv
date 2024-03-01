package utils

import (
	"errors"
	"flag"
	"os"
)

type Params struct {
	JsonFilename string
	DBFileName   string
	Password     string
}

func (p Params) Parse() Params {
	p.JsonFilename = *flag.String("in", "", "JSON file to convert")
	p.DBFileName = *flag.String("db", "", "Name for the new KeepassDB")
	p.Password = *flag.String("pass", "", "Password for the new KeepassDB")
	flag.Parse()
	return p
}

func (p Params) Validate() error {
	if len(p.JsonFilename) == 0 {
		return errors.New("Please select a valid JSON file")
	}

	if len(p.DBFileName) == 0 {
		return errors.New("Please Enter a valid file name for the new keepass db")
	}

	if len(p.Password) == 0 {
		return errors.New("Please enter a password for the new keepass db")
	}

	return nil
}

func Convert(params Params) error {

	err := params.Validate()
	if err != nil {
		return err
	}

	passData, err := Parse(params.JsonFilename)
	if err != nil {
		return err
	}

	rootGroup := FillDB(passData)

	newKeepassFile, err := OpenDB(params.DBFileName)
	if err != nil {
		return err
	}

	err = CloseDB(newKeepassFile, params.Password, rootGroup)
	if err != nil {
		return err
	}

	defer func(newKeepassFile *os.File) {
		err = newKeepassFile.Close()
	}(newKeepassFile)

	return err
}
