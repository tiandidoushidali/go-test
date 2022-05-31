<!-- package=med.xim -->
- [GET /med.xim.Group/CreateOrRelate](#get-medximgroupcreateorrelate)  创建或者关联一个群
- [GET /med.xim.Group/InfoBatch](#get-medximgroupinfobatch)  批量获取群信息
- [GET /med.xim.Group/JoinOrCreate](#get-medximgroupjoinorcreate)  加群或创建群组
- [GET /med.xim.Group/History](#get-medximgrouphistory)  获取群组历史消息
- [GET /med.xim.Group/Quit](#get-medximgroupquit)  退群

## GET /med.xim.Group/CreateOrRelate
### 创建或者关联一个群

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|userId|否|多个integer||
|groupName|否|string||
|groupId|否|integer||
|businessType|否|integer||
|businessId|否|integer||
|avatar|否|string||

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
        "groupId": 0
    }
}
```


## GET /med.xim.Group/InfoBatch
### 批量获取群信息

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|uid|否|integer||
|groupId|否|多个integer||

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
        "info": [
            {
                "id": 0,
                "name": "",
                "owner": 0,
                "businessType": 0,
                "businessId": 0,
                "isPublic": 0,
                "isForbidden": 0,
                "insertTime": 0,
                "avatarList": [
                    ""
                ],
                "content": "",
                "number": 0,
                "broadcast": "",
                "qrCodeUrl": "",
                "shareH5Url": "",
                "joinStatus": 0,
                "isNewOrigin": true,
                "ownerUser": {
                    "userId": 0,
                    "name": "",
                    "sex": 0,
                    "avatar": "",
                    "type": 0,
                    "hospital": "",
                    "hospitalId": 0,
                    "sectionId": 0,
                    "sectionName": "",
                    "titleId": 0,
                    "titleName": "",
                    "isSurgeon": 0,
                    "userPlatform": "",
                    "jumpUrl": ""
                }
            }
        ]
    }
}
```


## GET /med.xim.Group/JoinOrCreate
### 加群或创建群组

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|groupId|否|integer||
|userId|否|integer||
|groupName|否|string||
|businessType|否|integer||
|businessId|否|integer||
|avatar|否|string||

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
        "groupId": 0
    }
}
```


## GET /med.xim.Group/History
### 获取群组历史消息

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|groupId|否|integer||
|start|否|integer||
|limit|否|integer||
|sort|否|bool||
|minId|否|integer||

#### 响应

```javascript
{
    "code": 0,
    "errcode": 0,
    "message": "ok",
    "errmsg": "ok",
    "data": {
        "next": 0,
        "more": true,
        "list": [
            {
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
            }
        ]
    }
}
```


## GET /med.xim.Group/Quit
### 退群

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|groupId|否|integer||
|userId|否|integer||
|businessType|否|integer||
|businessId|否|integer||

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

