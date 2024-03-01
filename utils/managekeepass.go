package utils

import (
	"github.com/tobischo/gokeepasslib"
	"os"
	"strconv"
)

func fillEntry(item Item) gokeepasslib.Entry {
	entry := gokeepasslib.NewEntry()
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "Title", Value: gokeepasslib.V{Content: item.Data.Metadata.Name}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "UserName", Value: gokeepasslib.V{Content: item.Data.Content.Username}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "Notes", Value: gokeepasslib.V{Content: item.Data.Metadata.Note}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "Password", Value: gokeepasslib.V{Content: item.Data.Content.Password, Protected: true}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "AliasEmail", Value: gokeepasslib.V{Content: item.AliasEmail}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "ItemID", Value: gokeepasslib.V{Content: item.ItemId}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "ShareID", Value: gokeepasslib.V{Content: item.ShareId}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "ItemUUID", Value: gokeepasslib.V{Content: item.Data.Metadata.ItemUUID}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "CreateTime", Value: gokeepasslib.V{Content: strconv.Itoa(item.CreateTime)}})
	entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "ModifyTime", Value: gokeepasslib.V{Content: strconv.Itoa(item.ModifyTime)}})

	if len(item.Data.Content.Urls) > 0 {
		entry.Values = append(entry.Values, gokeepasslib.ValueData{Key: "URL", Value: gokeepasslib.V{Content: item.Data.Content.Urls[0]}})
	}

	return entry
}

func fillGroup(vault Vault) gokeepasslib.Group {
	vaultGroup := gokeepasslib.NewGroup()
	vaultGroup.Name = vault.Name
	vaultGroup.Notes = vault.Description

	for _, item := range vault.Items {
		entry := fillEntry(item)
		vaultGroup.Entries = append(vaultGroup.Entries, entry)
	}

	return vaultGroup
}

func OpenDB(dbFileName string) (*os.File, error) {
	newKeepassFile, err := os.Create(dbFileName)
	if err != nil {
		return nil, err
	}

	return newKeepassFile, nil
}

func CloseDB(newKeepassFile *os.File, password string, rootGroup gokeepasslib.Group) error {
	db := &gokeepasslib.Database{
		Signature:   &gokeepasslib.DefaultSig,
		Headers:     gokeepasslib.NewFileHeaders(),
		Credentials: gokeepasslib.NewPasswordCredentials(password),
		Content: &gokeepasslib.DBContent{
			Meta: gokeepasslib.NewMetaData(),
			Root: &gokeepasslib.RootData{
				Groups: []gokeepasslib.Group{rootGroup},
			},
		},
	}

	err := db.LockProtectedEntries()
	if err != nil {
		return err
	}

	keepassEncoder := gokeepasslib.NewEncoder(newKeepassFile)
	if err := keepassEncoder.Encode(db); err != nil {
		return err
	}

	return nil
}

func FillDB(data PassData) gokeepasslib.Group {

	rootGroup := gokeepasslib.NewGroup()
	rootGroup.Name = "root group"

	for _, vault := range data.Vaults {
		vaultGroup := fillGroup(vault)
		rootGroup.Groups = append(rootGroup.Groups, vaultGroup)
	}

	if len(rootGroup.Groups) == 1 {
		return rootGroup.Groups[0]
	}

	return rootGroup
}
