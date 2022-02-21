package main

import (
	"example.com/ex00/imgconv"
)

func main() {
	imgconv.JpgToPng()
	// flag.Parse()
	// args := flag.Args()
	// if len(args) == 0 {
	// 	fmt.Fprintln(os.Stderr, "error: invalid argument")
	// 	return
	// }
	// for _, dir := range args {
	// 	filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "error: %s: no such file or directory\n", path)
	// 			return err
	// 		}
	// 		if info.IsDir() || strings.HasSuffix(path, ".png") {
	// 			return nil
	// 		}
	// 		if !strings.HasSuffix(path, ".jpg") {
	// 			fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
	// 			return nil
	// 		}
	// 		in_path := path
	// 		out_path := replaceSuffix(path, ".jpg", ".png")
	// 		if err := convertImage(in_path, out_path); err != nil {
	// 			fmt.Fprintf(os.Stderr, "error: %s %v\n", path, err)
	// 			return nil
	// 		}
	// 		return nil
	// 	})
	// }
}
