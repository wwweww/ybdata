# ybdata
易班数据获取
---
## API
| route | params | paramsType | return | returnType | remark |
|-|-|-|-|-|-|
| /search/class | College | string | Class | []map[string]string | |
| /search/student | Class | string | students  | []string | |

## ex
```http
POST http://localhost:8080/search/student HTTP/1.1
Content-Type: application/json


{
    "class": "营销2301"
}

###

POST http://yb.zjhzcc.edu.cn/ybdata/search/class HTTP/1.1
Content-Type: application/json

{
    "College": "管理学院"
}
```
