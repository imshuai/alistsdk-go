package alistsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

const (
	VERSION = "1.0.0"
)

type Client struct {
	base     string              //Alist API base url
	token    string              //Alist API access token
	username string              //Alist username
	password string              //Alist password
	header   map[string][]string //http request header
}

// NewClient creates a new instance of the Client struct.
//
// Parameters:
// - endpoint: the base URL of the client.
// - username: the username for authentication.
// - password: the password for authentication.
//
// Returns:
// - *Client: a pointer to the newly created Client instance.
func NewClient(endpoint, username, password string) *Client {
	return &Client{
		base:     endpoint,
		username: username,
		password: password,
		header: map[string][]string{
			"Content-Type": {"application/json; charset=utf-8"},
		},
	}
}

// NewClientWithToken creates a new client with the given token.
//
// Parameters:
//   - token: the token to be used for authentication.
//
// Returns:
//   - *Client: a pointer to the newly created client.
func NewClientWithToken(token string) *Client {
	return &Client{
		token: token,
		header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}
}

// Login performs the login operation for the Client.
//
// It sends a POST request to the "/api/auth/login" endpoint with the provided
// username and password in the request body as JSON. If successful, it
// extracts the JWT token from the response and sets it as the authorization
// header for future requests. Then, it sends a GET request to the "/api/me"
// endpoint to retrieve the user information. If successful, it returns the
// user information as a *User object.
//
// Returns:
// - *User: The user information if the login is successful.
// - error: An error if the login or user information retrieval fails.
func (c *Client) Login() (*User, error) {
	body := `{
        "username": "` + c.username + `",
        "password": "` + c.password + `"
    }`

	respByts, err := do("POST", c.base+"/api/auth/login", c.header, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	loginResp := &LoginResp{}
	err = json.Unmarshal(respByts, loginResp)
	if err != nil {
		return nil, err
	}
	if loginResp.Code != 200 {
		return nil, errors.New(loginResp.Message)
	}
	c.token = loginResp.Data.Token
	c.header["Authorization"] = []string{c.token}

	respByts = respByts[0:0]

	respByts, err = do("GET", c.base+"/api/me", c.header, nil)
	if err != nil {
		return nil, err
	}
	userResp := &UserInfoResp{}
	err = json.Unmarshal(respByts, userResp)
	if err != nil {
		return nil, err
	}
	if userResp.Code != 200 {
		return nil, errors.New(userResp.Message)
	}
	u := &User{}
	u = userResp.Data
	return u, nil
}

func (c *Client) isLogin() bool {
	return c.token != ""
}

// MkDir creates a directory at the specified path.
//
// path: the path of the directory to be created.
// error: returns an error if the directory creation fails.
func (c *Client) MkDir(path string) error {
	if !c.isLogin() {
		return errors.New("not login yet")
	}
	body := `{
        "path": "` + path + `"
    }`

	respByts, err := do("POST", c.base+"/api/fs/mkdir", c.header, bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	comResp := &CommonResp{}
	err = json.Unmarshal(respByts, comResp)
	if err != nil {
		return err
	}
	if comResp.Code != 200 {
		return errors.New(comResp.Message)
	}
	return nil
}

// Rename renames a file or directory to a new name.
//
// Parameters:
// - newName: the new name to rename to (string)
// - path: the path of the file or directory to rename (string)
//
// Returns:
// - error: an error if the renaming process fails (error)
func (c *Client) Rename(newName, path string) error {
	if !c.isLogin() {
		return errors.New("not login yet")
	}
	body := `{
        "name": "` + newName + `",
        "path": "` + path + `"
    }`
	respByts, err := do("POST", c.base+"/api/fs/rename", c.header, bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	comResp := &CommonResp{}
	err = json.Unmarshal(respByts, comResp)
	if err != nil {
		return err
	}
	if comResp.Code != 200 {
		return errors.New(comResp.Message)
	}
	return nil
}

// Remove removes the specified files or directories from the client's filesystem.
//
// Parameters:
//   - dir (string): The directory where the files or directories are located.
//   - names ([]string): The names of the files or directories to be removed.
//
// Returns:
//   - error: An error if the removal operation fails, nil otherwise.
func (c *Client) Remove(dir string, names []string) error {
	if !c.isLogin() {
		return errors.New("not login yet")
	}
	body := `{
        "names": ["` + strings.Join(names, `","`) + `"],
        "dir": "` + dir + `"
    }`

	respByts, err := do("POST", c.base+"/api/fs/remove", c.header, bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	comResp := &CommonResp{}
	err = json.Unmarshal(respByts, comResp)
	if err != nil {
		return err
	}
	if comResp.Code != 200 {
		return errors.New(comResp.Message)
	}
	return nil
}

// RemoveEmptyDir removes an empty directory.
//
// The `dir` parameter is a string that represents the directory path to be removed.
// It returns an error if there was a problem removing the directory.
func (c *Client) RemoveEmptyDir(dir string) error {
	if !c.isLogin() {
		return errors.New("not login yet")
	}
	body := `{
        "src_dir": "` + dir + `"
    }`

	respByts, err := do("POST", c.base+"/api/fs/remove_empty_directory", c.header, bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	comResp := &CommonResp{}
	err = json.Unmarshal(respByts, comResp)
	if err != nil {
		return err
	}
	if comResp.Code != 200 {
		return errors.New(comResp.Message)
	}
	return nil
}

// Copy copies files from source directory to destination directory.
//
// srcDir: the source directory from which files will be copied.
// destDir: the destination directory to which files will be copied.
// names: a slice of string containing the names of the files to be copied.
// error: an error indicating any failure during the copy operation.
func (c *Client) Copy(srcDir, destDir string, names []string) error {
	if !c.isLogin() {
		return errors.New("not login yet")
	}
	body := `{
        "src_dir": "` + srcDir + `",
        "dst_dir": "` + destDir + `",
        "names": ["` + strings.Join(names, `","`) + `"]
    }`

	respByts, err := do("POST", c.base+"/api/fs/copy", c.header, bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	comResp := &CommonResp{}
	err = json.Unmarshal(respByts, comResp)
	if err != nil {
		return err
	}
	if comResp.Code != 200 {
		return errors.New(comResp.Message)
	}
	return nil
}

// RecursiveMove 递归移动
// srcDir: 源目录
// destDir: 目标目录
func (c *Client) RecursiveMove(srcDir, destDir string) error {
	if !c.isLogin() {
		return errors.New("not login yet")
	}
	body := `{
        "src_dir": "` + srcDir + `",
        "dst_dir": "` + destDir + `"
    }`

	respByts, err := do("POST", c.base+"/api/fs/recursive_move", c.header, bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	comResp := &CommonResp{}
	err = json.Unmarshal(respByts, comResp)
	if err != nil {
		return err
	}
	if comResp.Code != 200 {
		return errors.New(comResp.Message)
	}
	return nil
}

// Move 移动文件
// srcDir: 源目录
// destDir: 目标目录
// names: 文件名
func (c *Client) Move(srcDir, destDir string, names []string) error {
	if !c.isLogin() {
		return errors.New("not login yet")
	}
	body := `{
        "src_dir": "` + srcDir + `",
        "dst_dir": "` + destDir + `",
        "names": ["` + strings.Join(names, `","`) + `"]
    }`

	respByts, err := do("POST", c.base+"/api/fs/move", c.header, bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	comResp := &CommonResp{}
	err = json.Unmarshal(respByts, comResp)
	if err != nil {
		return err
	}
	if comResp.Code != 200 {
		return errors.New(comResp.Message)
	}
	return nil
}

func (c *Client) batchRename(endpoint, srcDir string, nameKV map[string]string) error {
	if !c.isLogin() {
		return errors.New("not login yet")
	}
	kvs := make([]string, len(nameKV))
	for k, v := range nameKV {
		kvs = append(kvs, `"`+k+`":"`+v+`"`)
	}

	body := `{
        "src_dir": "` + srcDir + `",
        "rename_objects": [{` + strings.Join(kvs, `,`) + `}]
    }`

	respByts, err := do("POST", c.base+endpoint, c.header, bytes.NewBufferString(body))
	if err != nil {
		return err
	}
	comResp := &CommonResp{}
	err = json.Unmarshal(respByts, comResp)
	if err != nil {
		return err
	}
	if comResp.Code != 200 {
		return errors.New(comResp.Message)
	}
	return nil
}

// RegexRename 正则重命名
// srcDir: 源目录
// srcName: 源文件名正则匹配表达式
// newName: 新文件名正则引用表达式
func (c *Client) RegexRename(srcDir string, regexKV map[string]string) error {
	return c.batchRename("/api/fs/regex_rename", srcDir, regexKV)
}

// BatchRename 批量重命名
// BatchRename renames multiple files in a given source directory.
//
// Parameters:
// - srcDir: the source directory where the files are located.
// - batchKV: a map containing the source file names as keys and the new file names as values.
//
// Returns:
// - error: an error if the batch rename operation fails.
func (c *Client) BatchRename(srcDir string, batchKV map[string]string) error {
	return c.batchRename("/api/fs/batch_rename", srcDir, batchKV)
}

// Dirs 获取目录
// Dirs retrieves a list of directories from the server.
//
// It takes the following parameters:
// - path: the path of the directory to retrieve.
// - dirPassword: the password for the directory.
// - forceRoot: a flag indicating whether to the root directory.
//
// It returns a slice of Dir structs and an error.
func (c *Client) Dirs(path, dirPassword string, forceRoot bool) ([]Dir, error) {
	if !c.isLogin() {
		return nil, errors.New("not login yet")
	}
	body := `{
        "path": "` + path + `",
        "password": "` + dirPassword + `",
        "force_root": ` + strconv.FormatBool(forceRoot) + `
    }`

	respByts, err := do("POST", c.base+"/api/fs/dirs", c.header, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	dirsResp := &DirsResp{}
	err = json.Unmarshal(respByts, dirsResp)
	if err != nil {
		return nil, err
	}
	if dirsResp.Code != 200 {
		return nil, errors.New(dirsResp.Message)
	}
	return dirsResp.Data, nil
}

// List 列出文件目录
// List returns a list of files in the specified directory.
//
// Parameters:
// - path: the path of the directory.
// - dirPassword: the password of the directory (if applicable).
// - pageNum: the page number for pagination.
// - pageSize: the number of files per page.
// - refresh: indicates whether to refresh the file list.
//
// Returns:
// - a slice of File structs representing the files in the directory.
// - an error if any error occurs.
func (c *Client) List(path, dirPassword string, pageNum, pageSize int, refresh bool) ([]File, error) {
	if !c.isLogin() {
		return nil, errors.New("not login yet")
	}
	body := `{
        "path": "` + path + `",
        "password": "` + dirPassword + `",
        "page_num": ` + strconv.Itoa(pageNum) + `,
        "per_page": ` + strconv.Itoa(pageSize) + `,
        "refresh": ` + strconv.FormatBool(refresh) + `
    }`

	respByts, err := do("POST", c.base+"/api/fs/list", c.header, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	listResp := &ListResp{}
	err = json.Unmarshal(respByts, listResp)
	if err != nil {
		return nil, err
	}
	if listResp.Code != 200 {
		return nil, errors.New(listResp.Message)
	}
	return listResp.Data.Content, nil
}

// Get 获取某个文件/目录信息
// Get retrieves a file or directory from the server.
//
// Parameters:
// - path: the path of the file or directory.
// - dirPassword: the password of the directory (if applicable).
//
// Returns:
// - a File struct representing the file or directory.
// - an error if any error occurs.
func (c *Client) Get(path, dirPassword string) (*File, error) {
	if !c.isLogin() {
		return nil, errors.New("not login yet")
	}
	body := `{
        "path": "` + path + `",
        "password": "` + dirPassword + `"
    }`

	respByts, err := do("POST", c.base+"/api/fs/get", c.header, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	getResp := &GetResp{}
	err = json.Unmarshal(respByts, getResp)
	if err != nil {
		return nil, err
	}
	if getResp.Code != 200 {
		return nil, errors.New(getResp.Message)
	}
	return getResp.Data, nil
}

// GetSettings 获取设置
// GetSettings retrieves the settings from the client.
//
// This function does not take any parameters.
// It returns a pointer to Settings and an error.
func (c *Client) GetSettings() (*Settings, error) {
	if !c.isLogin() {
		return nil, errors.New("not login yet")
	}
	respByts, err := do("GET", c.base+"/api/settings", c.header, nil)
	if err != nil {
		return nil, err
	}
	settingsResp := &SettingsResp{}
	err = json.Unmarshal(respByts, settingsResp)
	if err != nil {
		return nil, err
	}
	if settingsResp.Code != 200 {
		return nil, errors.New(settingsResp.Message)
	}
	return settingsResp.Data, nil
}
