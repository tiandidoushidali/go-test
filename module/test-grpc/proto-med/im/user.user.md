<!-- package=med.xim -->
- [GET /med.xim.User/CheckOnLine](#get-medximusercheckonline)  检查用户是否在线
- [GET /med.xim.User/CheckOnLineBatch](#get-medximusercheckonlinebatch)  批量检查用户是否在线

## GET /med.xim.User/CheckOnLine
### 检查用户是否在线

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|userId|否|integer||
|reference|否|integer||

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
        "isOnline": true
    }
}
```


## GET /med.xim.User/CheckOnLineBatch
### 批量检查用户是否在线

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|userIdList|否|多个integer||
|reference|否|integer||

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
        "list": [
            {
                "userId": 0,
                "isOnline": true
            }
        ]
    }
}
```

