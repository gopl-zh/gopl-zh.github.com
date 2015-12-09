**練習 1.7:** 函數調用io.Copy(dst, src)會從src中讀取內容，併將讀到的結果寫入到dst中，使用這個函數替代掉例子中的ioutil.ReadAll來拷貝響應結構體到os.Stdout，避免申請一個緩衝區(例子中的b)來存儲。記得處理io.Copy返迴結果中的錯誤。

