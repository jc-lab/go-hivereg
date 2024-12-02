package hiveutil

import (
	"github.com/gabriel-samfira/go-hivex"
	"github.com/jc-lab/go-hivereg/model/regtype"
)

type ValueType int

const (
	RegNone                     ValueType = hivex.RegNone
	RegSz                       ValueType = hivex.RegSz
	RegExpandSz                 ValueType = hivex.RegExpandSz
	RegBinary                   ValueType = hivex.RegBinary
	RegDword                    ValueType = hivex.RegDword
	RegDwordBigEndian           ValueType = hivex.RegDwordBigEndian
	RegLink                     ValueType = hivex.RegLink
	RegMultiSz                  ValueType = hivex.RegMultiSz
	RegResourceList             ValueType = hivex.RegResourceList
	RegFullResourceDescriptor   ValueType = hivex.RegFullResourceDescriptor
	RegResourceRequirementsList ValueType = hivex.RegResourceRequirementsList
	RegQword                    ValueType = hivex.RegQword
)

func RegTypeToHive(input regtype.RegType) ValueType {
	switch input {
	case regtype.REG_NONE:
		return RegNone
	case regtype.REG_SZ:
		return RegSz
	case regtype.REG_EXPAND_SZ:
		return RegExpandSz
	case regtype.REG_BINARY:
		return RegBinary
	case regtype.REG_DWORD:
		return RegDword
	case regtype.REG_DWORD_BIG_ENDIAN:
		return RegDwordBigEndian
	case regtype.REG_LINK:
		return RegLink
	case regtype.REG_MULTI_SZ:
		return RegMultiSz
	case regtype.REG_RESOURCE_LIST:
		return RegResourceList
	case regtype.REG_FULL_RESOURCE_DESC:
		return RegFullResourceDescriptor
	case regtype.REG_RESOURCE_REQUIREMENTS_LIST:
		return RegResourceRequirementsList
	case regtype.REG_QWORD:
		return RegQword
	}
	return RegNone
}

func RegTypeFromHive(input ValueType) regtype.RegType {
	switch input {
	case RegNone:
		return regtype.REG_NONE
	case RegSz:
		return regtype.REG_SZ
	case RegExpandSz:
		return regtype.REG_EXPAND_SZ
	case RegBinary:
		return regtype.REG_BINARY
	case RegDword:
		return regtype.REG_DWORD
	case RegDwordBigEndian:
		return regtype.REG_DWORD_BIG_ENDIAN
	case RegLink:
		return regtype.REG_LINK
	case RegMultiSz:
		return regtype.REG_MULTI_SZ
	case RegResourceList:
		return regtype.REG_RESOURCE_LIST
	case RegFullResourceDescriptor:
		return regtype.REG_FULL_RESOURCE_DESC
	case RegResourceRequirementsList:
		return regtype.REG_RESOURCE_REQUIREMENTS_LIST
	case RegQword:
		return regtype.REG_QWORD
	}
	return regtype.REG_NONE
}
