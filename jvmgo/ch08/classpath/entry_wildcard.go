package classpath

import "os"
import "path/filepath"
import "strings"

//由于通配符路径其实可以转化成混合路径，因此可以用CompositeEntry而不用创建WildcardEntry
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //remove *
	//申请Entry数组作为CompositeEntry
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		//由于通配符不匹配子目录，所以要跳过子目录
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		//遇到jar后缀的才进去查找类文件
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			//创建一个ZipEntry对象
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}

		return nil

	}

	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}