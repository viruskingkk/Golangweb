<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>

<body>
<p><strong>URL </strong><strong>規劃</strong><strong> </strong></p>
<p>對&nbsp;url&nbsp;的規劃如下: </p>
<table border="0" cellspacing="0" cellpadding="0" width="412">
  <tr>
   <td><pre>
      /upload/     用來接收檔上傳  只支援 POST 方法 <br />
      /images/     用來展示上傳檔 <br />
      /test/       用來顯示上傳表單 <br />
      /ping/       用作微服務健康狀態檢查 </pre></td>
  </tr>
</table>
<div>
  <h3>引入相關類庫 </h3>
</div>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>package main 
      import ( 
          &quot;crypto/md5&quot;    // 用於生成 md5 值 <br />
          &quot;encoding/json&quot; // 用來序列化返回值，生成 json 字串 <br />
          &quot;fmt&quot;           // 用來格式化輸出，主要用在錯誤輸出上 <br />
          &quot;io&quot;            // 用在保存檔上 <br />
          &quot;net/http&quot;      // 生成 http 服務 <br />
          &quot;net/url&quot;       // 用來 url encode <br />
          &quot;os&quot;            // 用在保存檔上 <br />
          &quot;regexp&quot;        // 用在格式化檔案名上 <br />
          &quot;strings&quot;       // 字串處理 <br />
          &quot;time           // 用來生成時間戳記 <br />
      ) </pre></td>
  </tr> 
</table>

<div>
  <h3>定義微服務處理 url 處理函數 </h3>
</div>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>func main() { 
          var port = &quot;:8000&quot; <br />
          http.HandleFunc(&quot;/upload/&quot;, saveImageHandler) <br />
          http.HandleFunc(&quot;/image/&quot;, showImageHandler) <br />
          http.HandleFunc(&quot;/ping/&quot;, statusHandler) <br />
          http.HandleFunc(&quot;/test/&quot;, HomeHandler) <br />
          fmt.Printf(&quot;Server running on port: %v...&quot;, port) <br />
          http.ListenAndServe(port, nil) <br />
      } </pre></td>
  </tr>
</table>

<p>監聽&nbsp;8000&nbsp;埠 </p>
<div>
  <h3>編寫 ping 服務 </h3>
</div>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>func statusHandler(w http.ResponseWriter, r *http.Request) { <br />
          w.Write([]byte(`&lt;!doctype html&gt;&lt;meta charset=&quot;utf-8&quot;&gt;&lt;form&gt;伺服器正常&lt;/form&gt;`))<br />    
          w.Write([]byte(`&lt;!doctype html&gt;&lt;meta charset=&quot;utf-8&quot;&gt;&lt;input type=&quot;button&quot; value=&quot;回到上傳首頁&quot; onclick=&quot;location.href='http://192.168.1.106:8000/test/'&quot;&gt;`))<br /> 
      } </pre></td>
  </tr>
</table>

<p>這裡沒有用到&nbsp;fmt&nbsp;系列函數，我覺得能少一些函式呼叫最好不過的<br />
<div>
  <h3>編寫上傳表單頁面&nbsp;/test </h3>
</div>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>func HomeHandler(w http.ResponseWriter, r *http.Request) { <br /> 
    w.Write([]byte(`&lt;!doctype html&gt;&lt;meta charset=&quot;utf-8&quot;&gt;&lt;form action=&quot;/upload/&quot;  enctype=&quot;multipart/form-data&quot; method=&quot;post&quot;&gt;&lt;input type=&quot;file&quot; name=&quot;image&quot; /&gt;&lt;input type=&quot;submit&quot; value=&quot;upload&quot; /&gt;&lt;/form&gt;`))<br /> 
    w.Write([]byte(`&lt;!doctype html&gt;&lt;meta charset=&quot;utf-8&quot;&gt;&lt;input type=&quot;button&quot; value=&quot;伺服器健康&quot; onclick=&quot;location.href='http://192.168.1.106:8000/ping/'&quot;&gt;`))<br />    
    w.Write([]byte(`&lt;!doctype html&gt;&lt;meta charset=&quot;utf-8&quot;&gt;&lt;input type=&quot;button&quot; value=&quot;圖片庫&quot; onclick=&quot;location.href='http://192.168.1.106:8000/image/'&quot;&gt;`)) <br /> 
      } </pre></td>
  </tr>
</table>

<p>這裡仍然只用了&nbsp;w.Write&nbsp;至於後面的&nbsp;<em>html</em>&nbsp;一大堆的，是從其他地方規規矩矩寫好的，然後壓縮而來的，然後，還省去了好多的部分，比如&nbsp;head&nbsp;title&nbsp;body,&nbsp;html&nbsp;之類的，這些在現代流覽器中都是非必需品，能省一個位元組是一個位元組 </p>
<div>
  <h3>展示上傳的檔 </h3>
</div>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>func showImageHandler(w http.ResponseWriter, r *http.Request) { <br /> 
          filePath := strings.Replace(r.URL.Path, &quot;/image/&quot;, &quot;./uploads/&quot;, 1) <br /> 
          http.ServeFile(w, r, filePath)<br /> 
          w.Write([]byte(`&lt;!doctype html&gt;&lt;meta charset=&quot;utf-8&quot;&gt;&lt;input type=&quot;button&quot; value=&quot;回到上傳首頁&quot; onclick=&quot;location.href='http://192.168.1.106:8000/test/'&quot;&gt;`))<br /> 
      } </pre></td>
  </tr>
</table>

<p>利用&nbsp;image&nbsp;來隱藏真實的上傳檔位址&nbsp;uploads&nbsp; </p>
<div>
  <h3>編寫接收上傳微服務 </h3>
</div>
<p>終於到了最關鍵的部分了，這裡先展示所有的代碼，然後再一一解釋 </p>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>type Image struct { <br /> 
          Link     string `json:&quot;link&quot;` <br /> 
         Filename string `json:&quot;filename&quot;` <br /> 
          Time     int64  `json:&quot;created_at&quot;` <br /> 
     } <br /> 
      func saveImageHandler(w http.ResponseWriter, r *http.Request) { <br /> 
          if r.Method != &quot;POST&quot; { <br /> 
              fmt.Fprintf(w, &quot;use POST method for upload!&quot;) <br /> 
              return <br /> 
          } <br /> 
          imageFile, imageHeader, err := r.FormFile(&quot;image&quot;) <br /> 
         defer imageFile.Close() <br /> 
          if err != nil { <br /> 
              fmt.Fprintf(w, &quot;error: %v&quot;, err) <br /> 
              return <br /> 
          } <br /> 
          //讀取前512個位元組用於判斷檔案類型 <br /> 
          firstImageBytes := make([]byte, 512) <br /> 
          _, err = imageFile.Read(firstImageBytes) <br /> 
         if err != nil { <br /> 
             fmt.Fprintf(w, &quot;error: &quot;, err) <br /> 
              return <br /> 
          } <br /> 
          md5Checksum := md5.Sum(firstImageBytes) <br /> 
          extensionMatcher := regexp.MustCompile(&quot;\\.\\w+$&quot;) <br /> 
          imageName := extensionMatcher.ReplaceAllString(imageHeader.Filename, &quot;&quot;) <br /> 
          filetype := http.DetectContentType(firstImageBytes) <br /> 
          var extension string <br /> 
          switch filetype { <br /> 
          case &quot;image/jpeg&quot;, &quot;image/jpg&quot;: <br /> 
              extension = &quot;jpg&quot; <br /> 
          case &quot;image/gif&quot;: <br /> 
              extension = &quot;gif&quot; <br /> 
          case &quot;image/png&quot;: <br /> 
              extension = &quot;png&quot; <br /> 
          case &quot;application/pdf&quot;: <br /> 
              extension = &quot;pdf&quot; <br /> 
          default: <br /> 
              fmt.Fprintf(w, &quot;unknown filetype: %v&quot;, filetype) <br /> 
              return <br /> 
          } <br /> 
          imageName = url.PathEscape(imageName) <br /> 
          //重新格式化檔案名 <br /> 
          uploadFilePath := fmt.Sprintf(&quot;./uploads/%x-%v.%v&quot;, md5Checksum, imageName, extension) <br /> 
          uploadedFile, err := os.Create(uploadFilePath) <br /> 
          if err != nil { <br /> 
              fmt.Fprintf(w, &quot;Unable to create the file for writing. Check your write access privilege&quot;) <br /> 
              return <br /> 
          } <br /> 
          defer uploadedFile.Close() <br /> 
          //把檔游標定位到開始 <br /> 
          imageFile.Seek(0, 0) <br /> 
          _, err = io.Copy(uploadedFile, imageFile) <br /> 
          if err != nil { <br /> 
              fmt.Fprintf(w, &quot;filesave error: %v&quot;, err) <br /> 
              return <br /> 
          } <br /> 
          url := &quot;http://&quot; + r.Host + &quot;/image/&quot; <br /> 
          link := strings.Replace(uploadFilePath, &quot;./uploads/&quot;, url, 1) <br /> 
          timestamp := time.Now().UnixNano() / int64(time.Millisecond) <br /> 
          image := Image{link, imageHeader.Filename, timestamp} <br /> 
          jsonBytes, _ := json.Marshal(image) <br /> 
          w.Write(jsonBytes) <br /> 
      } </pre></td>
  </tr>
</table>

<p>先定義上傳後的返回&nbsp;json&nbsp;串格式 </p>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>type Image struct { <br /> 
          Link     string `json:&quot;link&quot;`       //上傳檔訪問url <br /> 
          Filename string `json:&quot;filename&quot;`   //上傳檔的檔案名 <br /> 
          Time     int64  `json:&quot;created_at&quot;` //上傳時間 <br /> 
      } </pre></td>
  </tr>
</table>

<p>定義都很字面量吧，如果不懂，看後面的注釋，這裡缺了一個 欄位就是&nbsp;<em>MIME Type</em> <br />
  然後是判斷請求是否為&nbsp;POST&nbsp;方法，判斷&nbsp;image&nbsp;表單欄位是否設置 </p>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>if r.Method != &quot;POST&quot; { <br /> 
          fmt.Fprintf(w, &quot;use POST method for upload!&quot;) <br /> 
          return <br /> 
      } <br /> 
      imageFile, imageHeader, err := r.FormFile(&quot;image&quot;) <br /> 
      defer imageFile.Close() <br /> 
      if err != nil { <br /> 
          fmt.Fprintf(w, &quot;error: %v&quot;, err) <br /> 
          return <br /> 
      } </pre></td>
  </tr>
</table>

<p>重點中的重點來了，讀取上傳文件前 512  個位元組，然後根據內容判斷檔案類型&nbsp;MIME Type </p>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>//讀取前512個位元組用於判斷檔案類型 <br />
          firstImageBytes := make([]byte, 512) <br />
          _, err = imageFile.Read(firstImageBytes) <br />
          if err != nil { <br />
              fmt.Fprintf(w, &quot;error: &quot;, err) <br />
              return <br />
          } <br />
          md5Checksum := md5.Sum(firstImageBytes) <br />
          extensionMatcher := regexp.MustCompile(&quot;\\.\\w+$&quot;) <br />
          imageName := extensionMatcher.ReplaceAllString(imageHeader.Filename, &quot;&quot;) <br />
          filetype := http.DetectContentType(firstImageBytes) <br />
          var extension string <br />
          switch filetype { <br />
          case &quot;image/jpeg&quot;, &quot;image/jpg&quot;: <br />
              extension = &quot;jpg&quot; <br />
          case &quot;image/gif&quot;: <br />
              extension = &quot;gif&quot; <br />
          case &quot;image/png&quot;: <br />
              extension = &quot;png&quot; <br />
          case &quot;application/pdf&quot;: <br />
              extension = &quot;pdf&quot; <br />
          default: <br />
              fmt.Fprintf(w, &quot;unknown filetype: %v&quot;, filetype) <br />
              return <br />
          } </pre></td>
  </tr>
</table>
<p> 最後就是格式化檔案名和保存檔了 </p>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>imageName = url.PathEscape(imageName) <br />
      //重新格式化檔案名 <br />
      uploadFilePath := fmt.Sprintf(&quot;./uploads/%x-%v.%v&quot;, md5Checksum, imageName, extension) <br />
      uploadedFile, err := os.Create(uploadFilePath) <br />
      if err != nil { <br />
          fmt.Fprintf(w, &quot;Unable to create the file for writing. Check your write access privilege&quot;) <br />
          return <br />
      } <br />
      defer uploadedFile.Close() <br />
      //把檔游標定位到開始 <br />
      imageFile.Seek(0, 0) <br />
      _, err = io.Copy(uploadedFile, imageFile) <br />
      if err != nil { <br />
          fmt.Fprintf(w, &quot;filesave error: %v&quot;, err) <br />
          return <br />
      } <br />
      url := &quot;http://&quot; + r.Host + &quot;/image/&quot; <br />
      link := strings.Replace(uploadFilePath, &quot;./uploads/&quot;, url, 1) <br />
      timestamp := time.Now().UnixNano() / int64(time.Millisecond) <br />
      image := Image{link, imageHeader.Filename, timestamp} <br />
      jsonBytes, _ := json.Marshal(image) <br />
      w.Write(jsonBytes) </pre></td>
  </tr>
</table>
<p>這裡有一點需要注意的是,因為我們前面讀取了 512 個位元組，現在檔游標還停留在 512 這個位置 </p>
<div>
  <h3 id="终章">終章 </h3>
</div>
<p>直接&nbsp;go run index.go&nbsp;運行,然後訪問http://localhost:8000/test/&nbsp;一展成果 </p>
</body>
</html>
