import json
import uuid

import requests

from mysports.original_json import headers, host
from mysports.sports import get_md5_code


def login(mobile, psd, type):
    # 生成 uuid
    headers['uuid'] = uuid.uuid4().hex.upper()

    # 启动 session
    s = requests.Session()

    # 设置请求头
    s.headers = headers
    print('<LoginModule>：header 请求头为：', s.headers)

    # POST 所需要的 data 数据
    login_data = json.dumps(
        {"info": headers['uuid'], "mobile": mobile, "password": psd, "type": type})
    print('<LoginModule>：请求数据:', login_data)
    print('<LoginModule>：登陆接口处md5加密:', get_md5_code(login_data))
    print('<LoginModule>：准备发起请求')
    params = {'sign': get_md5_code(login_data), 'data': login_data}
    print(params)
    login_res = s.get(host + '/api/reg/login', params={'sign': get_md5_code(login_data), 'data': login_data})
    print('<LoginModule>：发起请求成功')
    print('<LoginModule>：正在处理数据')
    # convert to JSON
    login_rd = login_res.json()
    # catch login failed
    print('<LoginModule>：接口返回的信息:', login_rd)

    try:
        userid = login_rd['data']['userid']
        utoken = login_rd['data']['utoken']
        school = login_rd['data']['school']
    except:
        print(login_rd)
        raise Exception

    s.headers.update({'utoken': utoken})
    return userid, s, school
