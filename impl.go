package go_hivereg

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/gabriel-samfira/go-hivex"
	"github.com/jc-lab/go-hivereg/model/regtype"
	"github.com/jc-lab/go-hivereg/pkg/hiveutil"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type storeImpl struct {
	Hive     *hivex.Hivex
	Writable bool
}

func NewStore(hive *hivex.Hivex, writable bool) (Store, error) {
	return &storeImpl{
		Hive:     hive,
		Writable: writable,
	}, nil
}

func OpenStore(store string, writable bool) (Store, error) {
	var flags = hivex.READ
	if writable {
		flags |= hivex.WRITE
	}
	h, err := hivex.NewHivex(store, flags)
	if err != nil {
		return nil, errors.Wrap(err, "opening hive file '"+store+"'")
	}

	return &storeImpl{
		Hive:     h,
		Writable: writable,
	}, nil
}

func (s *storeImpl) Close() error {
	var err error
	if s.Writable {
		_, err = s.Hive.Commit()
	}
	closeErr := s.Hive.Close()
	if err != nil {
		return err
	}
	return closeErr
}

func (s *storeImpl) AddKey(key string) error {
	// Check if store is writable
	if !s.Writable {
		return fmt.Errorf("store is not writable")
	}

	_, err := s.getKey(key, true)

	return err
}

func (s *storeImpl) DeleteKey(key string) error {
	// Check if store is writable
	if !s.Writable {
		return fmt.Errorf("store is not writable")
	}

	node, err := s.getKey(key, true)
	if err != nil {
		return err
	}

	_, err = s.Hive.NodeDeleteChild(node)
	if err != nil {
		return fmt.Errorf("failed to delete key: %v", err)
	}

	return nil
}

func (s *storeImpl) AddValue(key string, valueName string, dataType regtype.RegType, separator string, data string) error {
	// Convert RegType to hivex type
	hiveType := hiveutil.RegTypeToHive(dataType)

	// Check if store is writable
	if !s.Writable {
		return fmt.Errorf("store is not writable")
	}

	node, err := s.getKey(key, true)
	if err != nil {
		return err
	}

	// Convert data based on type
	var valueData []byte
	switch hiveType {
	case hiveutil.RegSz, hiveutil.RegExpandSz:
		valueData, err = hiveutil.StringToUtf16LE(data, true)
		if err != nil {
			return err
		}
	case hiveutil.RegMultiSz:
		items := strings.Split(data, separator)
		valueData, err = hiveutil.StringsToMultiUtf16LE(items)
		if err != nil {
			return err
		}
	case hiveutil.RegDword:
		valueNum, err := parseInt(32, data)
		if err != nil {
			return err
		}
		valueData = binary.LittleEndian.AppendUint32(nil, uint32(valueNum))
	case hiveutil.RegQword:
		valueNum, err := parseInt(64, data)
		if err != nil {
			return err
		}
		valueData = binary.LittleEndian.AppendUint64(nil, valueNum)
	case hiveutil.RegBinary:
		valueData, err = base64.StdEncoding.DecodeString(data)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported data type: %v", dataType)
	}

	// Add or modify the value
	_, err = s.Hive.NodeSetValue(node, hivex.HiveValue{
		Type:  int(hiveType),
		Key:   valueName,
		Value: valueData,
	})
	if err != nil {
		return fmt.Errorf("failed to set value: %v", err)
	}

	return nil
}

func (s *storeImpl) DeleteValue(key string, valueName string) error {
	// Check if store is writable
	if !s.Writable {
		return fmt.Errorf("store is not writable")
	}

	node, err := s.getKey(key, false)
	if err != nil {
		return err
	}

	valueNode, err := hiveutil.FindValue(s.Hive, node, valueName)
	if err != nil {
		return err
	}

	// Delete the value
	_, err = s.Hive.NodeDeleteChild(valueNode)
	if err != nil {
		return fmt.Errorf("failed to delete value: %v", err)
	}

	return nil
}

func (s *storeImpl) getKey(key string, mkall bool) (int64, error) {
	// Normalize the key path
	key = strings.ReplaceAll(key, "/", "\\")

	// Split the key into parts
	parts := strings.Split(key, "\\")

	// Start from the root node
	currentNode, err := s.Hive.Root()
	if err != nil {
		return 0, err
	}

	// Traverse or create intermediate keys
	for i := 0; i < len(parts); i++ {
		// Try to find existing subkey
		subkey, err := hiveutil.FindChild(s.Hive, currentNode, parts[i])
		if err != nil || subkey == 0 {
			if mkall {
				// If subkey doesn't exist, create it
				subkey, err = s.Hive.NodeAddChild(currentNode, parts[i])
				if err != nil {
					return 0, errors.Wrap(err, fmt.Sprintf("failed to add key %s", parts[i]))
				}
			} else {
				return 0, errors.Wrap(err, fmt.Sprintf("failed to get key %s", parts[i]))
			}
		}
		currentNode = subkey
	}

	return currentNode, nil
}

func parseInt(bitSize int, s string) (uint64, error) {
	base := 10
	if strings.HasPrefix(strings.ToLower(s), "0x") {
		s = s[2:]
		base = 16
	}
	return strconv.ParseUint(s, base, bitSize)
}
