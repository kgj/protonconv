package utils

import (
	"encoding/json"
	"os"
)

type Content struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Urls     []string `json:"urls"`
	TOTPURI  string   `json:"totpUri"`
	PassKeys []string `json:"passkeys"`
}

type Metadata struct {
	Name     string `json:"name"`
	Note     string `json:"note"`
	ItemUUID string `json:"itemuuid"`
}

type Data struct {
	Metadata    Metadata `json:"metadata"`
	ExtraFields []string `json:"extrafields"`
	Type        string   `json:"type"`
	Content     Content  `json:"content"`
}

type Item struct {
	ItemId               string `json:"itemId"`
	ShareId              string `json:"shareId"`
	Data                 Data   `json:"data"`
	State                int    `json:"state"`
	AliasEmail           string `json:"aliasEmail"`
	ContentFormatVersion int    `json:"contentFormatVersion"`
	CreateTime           int    `json:"createTime"`
	ModifyTime           int    `json:"modifyTime"`
	Pinned               bool   `json:"pinned"`
}

type Display struct {
	Color int `json:"color"`
	Icon  int `json:"icon"`
}

type Vault struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Display     Display `json:"display"`
	Items       []Item  `json:"items"`
}

type PassData struct {
	Encrypted bool             `json:"encrypted"`
	UserID    string           `json:"userId"`
	Vaults    map[string]Vault `json:"vaults"`
	Version   string           `json:"version"`
}

func Parse(jsonFilename string) (PassData, error) {
	jsonFile, err := os.ReadFile(jsonFilename)
	if err != nil {
		return PassData{}, err
	}

	passData := PassData{}

	err = json.Unmarshal(jsonFile, &passData)
	if err != nil {
		return PassData{}, err
	}

	return passData, nil
}
