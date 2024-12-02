// Copyright 2024 JC-Lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package regtype

import (
	"fmt"
	"strings"
)

// RegType represents the registry value types
type RegType string

const (
	REG_NONE                       RegType = "REG_NONE"
	REG_SZ                         RegType = "REG_SZ"
	REG_EXPAND_SZ                  RegType = "REG_EXPAND_SZ"
	REG_BINARY                     RegType = "REG_BINARY"
	REG_DWORD                      RegType = "REG_DWORD"
	REG_DWORD_BIG_ENDIAN           RegType = "REG_DWORD_BIG_ENDIAN"
	REG_LINK                       RegType = "REG_LINK"
	REG_MULTI_SZ                   RegType = "REG_MULTI_SZ"
	REG_RESOURCE_LIST              RegType = "REG_RESOURCE_LIST"
	REG_FULL_RESOURCE_DESC         RegType = "REG_FULL_RESOURCE_DESC"
	REG_RESOURCE_REQUIREMENTS_LIST RegType = "REG_RESOURCE_REQUIREMENTS_LIST"
	REG_QWORD                      RegType = "REG_QWORD"
)

// ValidateRegType checks if the provided type is a valid registry type
func ValidateRegType(t string) (RegType, error) {
	// Convert input to uppercase to make comparison case-insensitive
	t = strings.ToUpper(t)

	// Try to match the input with known types
	switch t {
	case "NONE", "REG_NONE":
		return REG_NONE, nil
	case "SZ", "REG_SZ":
		return REG_SZ, nil
	case "EXPAND_SZ", "REG_EXPAND_SZ":
		return REG_EXPAND_SZ, nil
	case "BINARY", "REG_BINARY":
		return REG_BINARY, nil
	case "DWORD", "REG_DWORD":
		return REG_DWORD, nil
	case "DWORD_BIG_ENDIAN", "REG_DWORD_BIG_ENDIAN":
		return REG_DWORD_BIG_ENDIAN, nil
	case "LINK", "REG_LINK":
		return REG_LINK, nil
	case "MULTI_SZ", "REG_MULTI_SZ":
		return REG_MULTI_SZ, nil
	case "RESOURCE_LIST", "REG_RESOURCE_LIST":
		return REG_RESOURCE_LIST, nil
	case "FULL_RESOURCE_DESC", "REG_FULL_RESOURCE_DESC":
		return REG_FULL_RESOURCE_DESC, nil
	case "RESOURCE_REQUIREMENTS_LIST", "REG_RESOURCE_REQUIREMENTS_LIST":
		return REG_RESOURCE_REQUIREMENTS_LIST, nil
	case "QWORD", "REG_QWORD":
		return REG_QWORD, nil
	default:
		return "", fmt.Errorf("invalid registry type: %s", t)
	}
}

// String returns the string representation of the RegType
func (r RegType) String() string {
	return string(r)
}

// GetSupportedTypes returns a list of all supported registry types
func GetSupportedTypes() []RegType {
	return []RegType{
		REG_NONE,
		REG_SZ,
		REG_EXPAND_SZ,
		REG_BINARY,
		REG_DWORD,
		REG_DWORD_BIG_ENDIAN,
		REG_LINK,
		REG_MULTI_SZ,
		REG_RESOURCE_LIST,
		REG_FULL_RESOURCE_DESC,
		REG_RESOURCE_REQUIREMENTS_LIST,
		REG_QWORD,
	}
}
