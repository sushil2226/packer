package arm

import (
	"fmt"

	"github.com/hashicorp/packer/builder/azure/common"
)

const (
	TempNameAlphabet     = "0123456789bcdfghjklmnpqrstvwxyz"
	TempPasswordAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type TempName struct {
	AdminPassword       string
	CertificatePassword string
	ComputeName         string
	DeploymentName      string
	KeyVaultName        string
	ResourceGroupName   string
	OSDiskName          string
	NicName             string
	SubnetName          string
	PublicIPAddressName string
	VirtualNetworkName  string
}

func NewTempName() *TempName {
	tempName := &TempName{}

	suffix := common.RandomString(TempNameAlphabet, 10)
	tempName.ComputeName = fmt.Sprintf("pkrvm%s", suffix)
	tempName.DeploymentName = fmt.Sprintf("pkrdp%s", suffix)
	tempName.KeyVaultName = fmt.Sprintf("pkrkv%s", suffix)
	tempName.OSDiskName = fmt.Sprintf("pkros%s", suffix)
	tempName.NicName = fmt.Sprintf("pkrni%s", suffix)
	tempName.PublicIPAddressName = fmt.Sprintf("pkrip%s", suffix)
	tempName.SubnetName = fmt.Sprintf("pkrsn%s", suffix)
	tempName.VirtualNetworkName = fmt.Sprintf("pkrvn%s", suffix)
	tempName.ResourceGroupName = fmt.Sprintf("packer-Resource-Group-%s", suffix)

	tempName.AdminPassword = common.RandomString(TempPasswordAlphabet, 32)
	tempName.CertificatePassword = common.RandomString(TempPasswordAlphabet, 32)

	return tempName
}
