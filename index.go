package main
import (
    "crypto/md5"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "regexp"
    "strings"
    "time"
)
type Image struct {
    Link     string `json:"link"`
    Filename string `json:"filename"`
    Time     int64  `json:"created_at"`
}
func saveImageHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        fmt.Fprintf(w, "use POST method for upload!")
        return
    }
    imageFile, imageHeader, err := r.FormFile("image")
    defer imageFile.Close()
    if err != nil {
        fmt.Fprintf(w, "error: %v", err)
        return
    }
    //读取前512个字节用于判断文件类型
    firstImageBytes := make([]byte, 512)
    _, err = imageFile.Read(firstImageBytes)
    if err != nil {
        fmt.Fprintf(w, "error: ", err)
        return
    }
    md5Checksum := md5.Sum(firstImageBytes)
    extensionMatcher := regexp.MustCompile("\\.\\w+$")
    imageName := extensionMatcher.ReplaceAllString(imageHeader.Filename, "")
    filetype := http.DetectContentType(firstImageBytes)
    var extension string
    switch filetype {
    case "image/jpeg", "image/jpg":
        extension = "jpg"
    case "image/gif":
        extension = "gif"
    case "image/png":
        extension = "png"
    case "application/pdf":
        extension = "pdf"
    default:
        fmt.Fprintf(w, "unknown filetype: %v", filetype)
        return
    }
    imageName = url.PathEscape(imageName)
    //重新格式化文件名
    uploadFilePath := fmt.Sprintf("./uploads/%x-%v.%v", md5Checksum, imageName, extension)
    uploadedFile, err := os.Create(uploadFilePath)
    if err != nil {
        fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
        return
    }
    defer uploadedFile.Close()
    //把文件游标定位到开始
    imageFile.Seek(0, 0)
    _, err = io.Copy(uploadedFile, imageFile)
    if err != nil {
        fmt.Fprintf(w, "filesave error: %v", err)
        return
    }
    url := "http://" + r.Host + "/image/"
    link := strings.Replace(uploadFilePath, "./uploads/", url, 1)
    timestamp := time.Now().UnixNano() / int64(time.Millisecond)
    image := Image{link, imageHeader.Filename, timestamp}
    jsonBytes, _ := json.Marshal(image)
    w.Write([]byte(`<!doctype html><meta charset="utf-8"><form>圖片儲存路徑:</form>`))
    w.Write(jsonBytes)
}
func showImageHandler(w http.ResponseWriter, r *http.Request) {
    filePath := strings.Replace(r.URL.Path, "/image/", "./uploads/", 1)
    w.Write([]byte("所有圖片連結:")
    http.ServeFile(w, r, filePath)
    w.Write([]byte(`<!doctype html><meta charset="utf-8"><input type="button" value="回到上傳首頁" onclick="location.href='http://192.168.1.106:8000/test/'">`))

}
func statusHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`<!doctype html><meta charset="utf-8"><form>伺服器正常</form>`))
    w.Write([]byte(`<!doctype html><meta charset="utf-8"><input type="button" value="回到上傳首頁" onclick="location.href='http://192.168.1.106:8000/test/'">`))
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`<!doctype html><meta charset="utf-8"><form action="/upload/"  enctype="multipart/form-data" method="post"><input type="file" name="image" /><input type="submit" value="upload" /></form>`))
    w.Write([]byte(`<!doctype html><meta charset="utf-8"><input type="button" value="伺服器健康" onclick="location.href='http://192.168.1.106:8000/ping/'">`))
    w.Write([]byte(`<!doctype html><meta charset="utf-8"><input type="button" value="圖片庫" onclick="location.href='http://192.168.1.106:8000/image/'">`))
}
func main() {
    var port = ":8000"
    http.HandleFunc("/upload/", saveImageHandler)
    http.HandleFunc("/image/", showImageHandler)
    http.HandleFunc("/ping/", statusHandler)
    http.HandleFunc("/test/", HomeHandler)
    fmt.Printf("Server running on port: %v...", port)
    http.ListenAndServe(port, nil)
}