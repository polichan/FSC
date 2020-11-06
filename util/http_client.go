package util

import (
    "crypto/tls"
    "encoding/json"
    "errors"
    "fmt"
    "fsc/global"
    "fsc/global/response"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "sync"
)
 
var (
    GET_METHOD    = "GET"
    POST_METHOD   = "POST"
    SENDTYPE_FROM = "from"
    SENDTYPE_JSON = "json"
)
 
type FSCHttpClient struct {
    Link     string
    SendType string
    Header   map[string]string
    Body     map[string]string
    sync.RWMutex
}
 
func NewFSCHttpClientSend(link string) *FSCHttpClient {
    return &FSCHttpClient{
        Link:     link,
        SendType: SENDTYPE_FROM,
    }
}
 
func (h *FSCHttpClient) SetBody(body map[string]string) {
    h.Lock()
    defer h.Unlock()
    h.Body = body
}
 
func (h *FSCHttpClient) SetHeader(header map[string]string) {
    h.Lock()
    defer h.Unlock()
    h.Header = header
}
 
func (h *FSCHttpClient) SetSendType(send_type string) {
    h.Lock()
    defer h.Unlock()
    h.SendType = send_type
}
 
func (h *FSCHttpClient) Get() ([]byte, error) {
    return h.send(GET_METHOD)
}
 
func (h *FSCHttpClient) Post() ([]byte, error) {
    return h.send(POST_METHOD)
}
 
func GetUrlBuild(link string, data map[string]string) string {
    u, _ := url.Parse(link)
    q := u.Query()
    for k, v := range data {
        q.Set(k, v)
    }
    u.RawQuery = q.Encode()
    return u.String()
}

func (h *FSCHttpClient) send(method string) ([]byte, error) {
    var (
        req       *http.Request
        resp      *http.Response
        client    http.Client
        send_data string
        err       error
    )
 
    if len(h.Body) > 0 {
        if strings.ToLower(h.SendType) == SENDTYPE_JSON {
            send_body, json_err := json.Marshal(h.Body)
            if json_err != nil {
                return nil, json_err
            }
            send_data = string(send_body)
        } else {
            send_body := http.Request{}
            send_body.ParseForm()
            for k, v := range h.Body {
                send_body.Form.Add(k, v)
            }
            send_data = send_body.Form.Encode()
        }
    }
 
    //忽略https的证书
    client.Transport = &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
 
    req, err = http.NewRequest(method, h.Link, strings.NewReader(send_data))
    if err != nil {
        return nil, err
    }
    defer req.Body.Close()
 
    //设置默认header
    if len(h.Header) == 0 {
        //json
        if strings.ToLower(h.SendType) == SENDTYPE_JSON {
            h.Header = map[string]string{
                "Content-Type": "application/json; charset=utf-8",
                "User-Agent": "CollegeSports2/2.2.14 (iPhone; iOS 11.0; Scale/3.00)",
                "versionCode": "309",
                "versionName": "2.2.15",
                "xxversionxx": "20180601",
                "uuid": "uuid",
                "platform": "iOS",
            }
        } else { //form
            h.Header = map[string]string{
                "Content-Type": "application/x-www-form-urlencoded",
                "User-Agent": "CollegeSports2/2.2.14 (iPhone; iOS 11.0; Scale/3.00)",
                "versionCode": "309",
                "versionName": "2.2.15",
                "xxversionxx": "20180601",
                "uuid": "uuid",
                "platform": "iOS",
            }
        }
    }
 
    for k, v := range h.Header {
        if strings.ToLower(k) == "host" {
            req.Host = v
        } else {
            req.Header.Add(k, v)
        }
    }

     if global.FSC_USER != nil {
         // 如果以及登录过了，就添加 UToken 进去
        req.Header.Add("utoken", global.FSC_USER.UToken)
     }


    resp, err = client.Do(req)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New(fmt.Sprintf("error http code :%d", resp.StatusCode))
    }

    defer req.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    scRep := response.SCResponse{}
    err = json.Unmarshal(body, &scRep)
    if scRep.Code != "200" {
        return nil, errors.New(scRep.Msg)
    }else{
        return body, err
    }
}