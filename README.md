# alistsdk-go
Alist SDK for golang
  
## Usage
```go
import (
      "github.com/aliyun/alistsdk-go"
)

func main() {
    client := alistsdk.NewClient("endpoint", "username", "password")
    user,err:= client.Login()
    if err != nil {
        panic(err)
    }
    fmt.Println(user)

    //do something with client
}
```  
## API Endpoints  
 - auth 验证类
   - [x] POST /api/auth/login  Token获取
   - [ ] POST /api/auth/login/hash  Token获取hash
   - [ ] POST /api/auth/2fa/generate 生成2FA密钥
   - [ ] POST /api/auth/2fa/verify 验证2FA Code
 - fs 文件类 
   - [x] GET /api/me 获取用户信息
   - [x] POST /api/fs/mkdir 创建文件夹
   - [x] POST /api/fs/rename 重命名文件
   - [ ] PUT /api/fs/form 表单上传文件
   - [x] POST /api/fs/list 列出文件目录
   - [x] POST /api/fs/get 获取某个文件/目录信息
   - [ ] POST /api/fs/search 搜索文件或文件夹
   - [x] POST /api/fs/dirs 获取目录
   - [x] POST /api/fs/batch_rename 批量重命名
   - [x] POST /api/fs/regex_rename 正则重命名
   - [x] POST /api/fs/move 移动文件
   - [x] POST /api/fs/recursive_move 聚合移动
   - [x] POST /api/fs/copy 复制文件
   - [x] POST /api/fs/remove 删除文件或文件夹
   - [x] POST /api/fs/remove_empty_directory 删除空文件夹
   - [ ] PUT /api/fs/put 流式上传文件
   - [ ] POST /api/fs/add_aria2 添加aria2下载
   - [ ] POST /api/fs/add_qbit 添加qBittorrent下载
 - public 公共类
   - [ ] GET /ping ping检测
   - [x] GET /api/public/settings 获取站点设置