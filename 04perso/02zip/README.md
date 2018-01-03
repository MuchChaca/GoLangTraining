# archive/zip package - official library - go
## TODOs
[ ] Check if it handles password protected zip archives

## Sources
* [Some examples / explanations](http://blog.ralch.com/tutorial/golang-working-with-zip/)


## Peek
### Extracting
The content of zip package can be read by using zip reader.
To extract the package, we need to recreate all directories and files:
```go
// unzip extract all directories an files from a zip archive
// by recreating all of its files and directories
func unzip(archive, target string) error {
	// openning a zip reader
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	// Recreating all directories
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	// Recreating all files
	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
```

#### Notes
I guess the zip file appears like a regular directory, we can't remove stuff from it but copy them somewhere else.

### Compressing
Compressing is more complicated than extracting.
It is possible to **compress a single file** or a **hierarchy of directories**. In both cases it is required to change the file header name depending on its type.  
If the content copied is directory, the file header name should be changed to ``<directory_name>/``.  
For regular files, its header name is a relative path ``<directory_name>/<file_name>``.
Like so:  
```go
// zipit compress the source to the target
func zipit(source, target string) error {
	// create the zip file that will receive the compressed source
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	// create a writer on the zip file
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// Get the file info + structure of the file
	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	//* if the file the compress is a file, baseDir will be set at 0 value
	// otherwise it will be take the last element of the path, without seperators.
	// if the path is empty it will return "."
	// if the path consists only of seperators, it will return "/"
	// check the comments to see where/what
	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}
	// https://godoc.org/path/filepath#Walk
	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// https://godoc.org/archive/zip#FileHeader
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		// if it's not a dir then it's a file
		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}
		// but if it's a dir
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
```

## Next
**archive/tar:** [link](https://godoc.org/archive/tar)
