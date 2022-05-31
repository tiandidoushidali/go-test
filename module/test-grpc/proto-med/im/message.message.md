<!-- package=med.xim -->
- [GET /med.xim.Message/Handle](#get-medximmessagehandle)  发消息
- [GET /med.xim.Message/SendSystem](#get-medximmessagesendsystem)  发送系统消息
- [GET /med.xim.Message/DelChat](#get-medximmessagedelchat)  删除指定用户的会话入口(不会删除群组的其他用户会话入口)
- [GET /med.xim.Message/AddToWhiteList](#get-medximmessageaddtowhitelist)  添加白名单

## GET /med.xim.Message/Handle
### 发消息

#### 方法：GET

#### 请求参数

```javascript
{
    "data": {
        "id": 0,
        "idStr": "",
        "from": {
            "id": 0,
            "reference": 0,
            "name": "",
            "avatar": "",
            "type": 0,
            "hospital": "",
            "title": "",
            "section": ""
        },
        "user": {
            "id": 0,
            "reference": 0,
            "name": "",
            "avatar": "",
            "type": 0,
            "hospital": "",
            "title": "",
            "section": ""
        },
        "group": {
            "id": 0,
            "name": "",
            "avatar": [
                ""
            ],
            "amount": 0,
            "owner": 0,
            "businessType": 0,
            "businessId": 0,
            "avatarUrl": "",
            "isPublic": 0,
            "isForbidden": 0,
            "isWechat": 0
        },
        "text": {
            "content": "",
            "anchors": [
                {
                    "href": "",
                    "start": 0,
                    "len": 0,
                    //  链接位置，居左传1，居中传2，居右传3，不传默认是0
                    "style": 0,
                    //  区分不同端，key代表不同端标识；如doctor-app、doctor-h5、panel、patient-h5等
                    //  详见 https://wiki.medlinker.com/pages/viewpage.action?pageId=54669120
                    "anchorTag": {
                        "mapKey": {
                            //  不同端的跳转链接
                            "tagHref": "",
                            //  不同端是否展示此链接文本
                            "tagIsShow": true,
                            //  链接颜色
                            "color": "",
                            //  扩展字段
                            "ext": ""
                        }
                    },
                    //  链接颜色
                    "color": "",
                    //  扩展字段
                    "ext": ""
                }
            ],
            "atUsers": [
                {
                    "id": 0,
                    "reference": 0,
                    "name": "",
                    "avatar": "",
                    "type": 0,
                    "hospital": "",
                    "title": "",
                    "section": ""
                }
            ]
        },
        "voice": {
            "url": "",
            "duration": 0
        },
        "card": {
            "type": 0,
            "label": "",
            "title": "",
            "subtitle": "",
            "summary": "",
            "image": "",
            "url": "",
            "display": 0,
            "extra": ""
        },
        "image": {
            "url": "",
            "size": 0,
            "preview": ""
        },
        "location": {
            "longitude": 0.1,
            "latitude": 0.1
        },
        "notice": {
            "content": "",
            "anchors": [
                {
                    "href": "",
                    "start": 0,
                    "len": 0,
                    //  链接位置，居左传1，居中传2，居右传3，不传默认是0
                    "style": 0,
                    //  区分不同端，key代表不同端标识；如doctor-app、doctor-h5、panel、patient-h5等
                    //  详见 https://wiki.medlinker.com/pages/viewpage.action?pageId=54669120
                    "anchorTag": {
                        "mapKey": {
                            //  不同端的跳转链接
                            "tagHref": "",
                            //  不同端是否展示此链接文本
                            "tagIsShow": true,
                            //  链接颜色
                            "color": "",
                            //  扩展字段
                            "ext": ""
                        }
                    },
                    //  链接颜色
                    "color": "",
                    //  扩展字段
                    "ext": ""
                }
            ]
        },
        "businessCard": {
            "user": {
                "id": 0,
                "reference": 0,
                "name": "",
                "avatar": "",
                "type": 0,
                "hospital": "",
                "title": "",
                "section": ""
            }
        },
        "json": {
            "content": "",
            "type": 0
        },
        "notify": {
            "type": 0
        },
        "login": {
            "device": 0,
            "platform": 0,
            "token": "",
            "session_key": "",
            "version": ""
        },
        "withdraw": {
            "id": 0,
            "idStr": ""
        },
        "groupOperation": {
            "operation": 0,
            "group": 0,
            "members": [
                {
                    "id": 0,
                    "reference": 0,
                    "name": "",
                    "avatar": "",
                    "type": 0,
                    "hospital": "",
                    "title": "",
                    "section": ""
                }
            ],
            "content": ""
        },
        "friendOperation": {
            "operation": 0,
            "opposite": 0,
            "content": "",
            "reference": 0
        },
        "video": {
            "url": "",
            "duration": 0,
            "cover": "",
            "title": "",
            "id": 0,
            "width": 0,
            "height": 0
        },
        "ext": "",
        "isShow": {
            "mapKey": true
        }
    },
    "fromProxy": true
}
```

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
        "msgId": 0,
        "info": ""
    }
}
```


## GET /med.xim.Message/SendSystem
### 发送系统消息

#### 方法：GET

#### 请求参数

```javascript
{
    "userIds": [
        0
    ],
    "type": "",
    "senderId": 0,
    "text": {
        "content": "",
        "anchors": [
            {
                "href": "",
                "start": 0,
                "len": 0,
                //  链接位置，居左传1，居中传2，居右传3，不传默认是0
                "style": 0,
                //  区分不同端，key代表不同端标识；如doctor-app、doctor-h5、panel、patient-h5等
                //  详见 https://wiki.medlinker.com/pages/viewpage.action?pageId=54669120
                "anchorTag": {
                    "mapKey": {
                        //  不同端的跳转链接
                        "tagHref": "",
                        //  不同端是否展示此链接文本
                        "tagIsShow": true,
                        //  链接颜色
                        "color": "",
                        //  扩展字段
                        "ext": ""
                    }
                },
                //  链接颜色
                "color": "",
                //  扩展字段
                "ext": ""
            }
        ],
        "atUsers": [
            {
                "id": 0,
                "reference": 0,
                "name": "",
                "avatar": "",
                "type": 0,
                "hospital": "",
                "title": "",
                "section": ""
            }
        ]
    },
    "card": {
        "type": 0,
        "label": "",
        "title": "",
        "subtitle": "",
        "summary": "",
        "image": "",
        "url": "",
        "display": 0,
        "extra": ""
    },
    "notify": {
        "type": 0
    }
}
```

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
    }
}
```


## GET /med.xim.Message/DelChat
### 删除指定用户的会话入口(不会删除群组的其他用户会话入口)

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|userId|否|integer||
|mapId|否|string||

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
    }
}
```


## GET /med.xim.Message/AddToWhiteList
### 添加白名单

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|userIdList|否|多个integer||

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
    }
}
```

