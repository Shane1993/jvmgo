package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	//先申请一个Entry数组
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {

	//遍历Entry数组让各自的Entry调用自己的readClass方法
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}

	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	//将Entry数组拼接成一个字符串
	strs := make([]string, len(self)) //申请一个大小为len(self)的字符串数组
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}