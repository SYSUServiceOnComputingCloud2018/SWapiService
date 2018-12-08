# SWapi

## 后端需求：

[参考](https://swapi.co/documentation#base)

can

### 一、基本操作

1. base url（所以其他请求的url必须满足此前缀）

   `HOST/api`like`localhost:8080/api/`（服务端默认在本地8080端口监听，或者可以传入端口参数，详见boltdb/main.go）

2. 请求频率限制： 10,000 API request per day

3. 验证方式：open API，GET请求

4. 支持Schema，请求`/api/<resource>/schema`会返回data包括的所有字段

5. query查询：`api/people/?search=r2`，使用case-insensitive partial matches返回所有符合条件的

6. 使用JSON或者Wookiee格式输出（JSON为默认，Wookiee需要加query`/api/planets/1/?format=wookiee`）



### 二、资源访问

[详见](https://swapi.co/documentation#base)

root的reponse

```
HTTP/1.0 200 OK
Content-Type: application/json
{
    "films": "https://swapi.co/api/films/",
    "people": "https://swapi.co/api/people/",
    "planets": "https://swapi.co/api/planets/",
    "species": "https://swapi.co/api/species/",
    "starships": "https://swapi.co/api/starships/",
    "vehicles": "https://swapi.co/api/vehicles/"
}
```

...



## 具体实现

1. 注册路由
2. 数据库访问
3. JSON格式和状态码输出

...