# SWapi

## 后端需求：

[参考](https://swapi.co/documentation#base)

### 一、基本操作

1. base url（所以其他请求的url必须满足此前缀）

   `HOST/api`like`localhost:8080/api/`（服务端默认在本地8080端口监听，或者可以传入端口参数，详见boltdb/main.go）

2. 请求频率限制： 10,000 API request per day

3. 验证方式：open API，GET请求

4. 支持Schema，请求`/api/<resource>/schema`会返回data包括的所有字段

5. query查询：`api/people/?search=r2`，使用case-insensitive partial matches返回所有符合条件的

6. 使用JSON或者Wookiee格式输出（JSON为默认，Wookiee需要加query`/api/planets/1/?format=wookiee`）



### 二、资源访问

首先在`SwapiService/boltdb`目录下使用`go run main.go`指令启动客户端

随后进入浏览器可按顺序访问下列指令：

#### 2.1 访问 root

指令：

`http://localhost:8080/api`

response:

```go
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

#### 2.2 访问 people 

- 指令 ：

`http://localhost:8080/api/people/1`

- 页面响应 :

```
{"name":"Luke Skywalker","height":"172","mass":"77","hair_color":"blond","skin_color":"fair","eye_color":"blue","birth_year":"19BBY","gender":"male","homeworld":"https://swapi.co/api/planets/1/","films":["https://swapi.co/api/films/2/","https://swapi.co/api/films/6/","https://swapi.co/api/films/3/","https://swapi.co/api/films/1/","https://swapi.co/api/films/7/"],"species":["https://swapi.co/api/species/1/"],"vehicles":["https://swapi.co/api/vehicles/14/","https://swapi.co/api/vehicles/30/"],"starships":["https://swapi.co/api/starships/12/","https://swapi.co/api/starships/22/"],"created":"2014-12-09T13:50:51.644000Z","edited":"2014-12-20T21:17:56.891000Z","url":"https://swapi.co/api/people/1/"}
```

#### 2.3 访问 planet

- 指令：

`http://localhost:8080/api/planets/1`

- 页面响应 :

```
{"name":"Tatooine","rotation_period":"23","orbital_period":"304","diameter":"10465","climate":"arid","gravity":"1 standard","terrain":"desert","surface_water":"1","population":"200000","residents":["https://swapi.co/api/people/1/","https://swapi.co/api/people/2/","https://swapi.co/api/people/4/","https://swapi.co/api/people/6/","https://swapi.co/api/people/7/","https://swapi.co/api/people/8/","https://swapi.co/api/people/9/","https://swapi.co/api/people/11/","https://swapi.co/api/people/43/","https://swapi.co/api/people/62/"],"films":["https://swapi.co/api/films/5/","https://swapi.co/api/films/4/","https://swapi.co/api/films/6/","https://swapi.co/api/films/3/","https://swapi.co/api/films/1/"],"created":"2014-12-09T13:50:49.641000Z","edited":"2014-12-21T20:48:04.175778Z","url":"https://swapi.co/api/planets/1/"}
```

#### 2.4 访问 film

- 指令

`http://localhost:8080/api/films/1`

- 页面响应

```
{"title":"A New Hope","episode_id":4,"opening_crawl":"It is a period of civil war.\r\nRebel spaceships, striking\r\nfrom a hidden base, have won\r\ntheir first victory against\r\nthe evil Galactic Empire.\r\n\r\nDuring the battle, Rebel\r\nspies managed to steal secret\r\nplans to the Empire's\r\nultimate weapon, the DEATH\r\nSTAR, an armored space\r\nstation with enough power\r\nto destroy an entire planet.\r\n\r\nPursued by the Empire's\r\nsinister agents, Princess\r\nLeia races home aboard her\r\nstarship, custodian of the\r\nstolen plans that can save her\r\npeople and restore\r\nfreedom to the galaxy....","director":"George Lucas","producer":"Gary Kurtz, Rick McCallum","characters":["https://swapi.co/api/people/1/","https://swapi.co/api/people/2/","https://swapi.co/api/people/3/","https://swapi.co/api/people/4/","https://swapi.co/api/people/5/","https://swapi.co/api/people/6/","https://swapi.co/api/people/7/","https://swapi.co/api/people/8/","https://swapi.co/api/people/9/","https://swapi.co/api/people/10/","https://swapi.co/api/people/12/","https://swapi.co/api/people/13/","https://swapi.co/api/people/14/","https://swapi.co/api/people/15/","https://swapi.co/api/people/16/","https://swapi.co/api/people/18/","https://swapi.co/api/people/19/","https://swapi.co/api/people/81/"],"planets":["https://swapi.co/api/planets/2/","https://swapi.co/api/planets/3/","https://swapi.co/api/planets/1/"],"starships":["https://swapi.co/api/starships/2/","https://swapi.co/api/starships/3/","https://swapi.co/api/starships/5/","https://swapi.co/api/starships/9/","https://swapi.co/api/starships/10/","https://swapi.co/api/starships/11/","https://swapi.co/api/starships/12/","https://swapi.co/api/starships/13/"],"vehicles":["https://swapi.co/api/vehicles/4/","https://swapi.co/api/vehicles/6/","https://swapi.co/api/vehicles/7/","https://swapi.co/api/vehicles/8/"],"species":["https://swapi.co/api/species/5/","https://swapi.co/api/species/3/","https://swapi.co/api/species/2/","https://swapi.co/api/species/1/","https://swapi.co/api/species/4/"],"created":"2014-12-10T14:23:31.880000Z","edited":"2015-04-11T09:46:52.774897Z","url":"https://swapi.co/api/films/1/"}
```

#### 2.5 访问 vehicle

- 指令

`http://localhost:8080/api/vehicles/4`

- 页面响应

```
{"name":"Sand Crawler","model":"Digger Crawler","manufacturer":"Corellia Mining Corporation","cost_in_credits":"150000","length":"36.8","max_atmosphering_speed":"30","crew":"46","passengers":"30","cargo_capacity":"50000","consumables":"2 months","vehicle_class":"wheeled","pilots":[],"films":["https://swapi.co/api/films/5/","https://swapi.co/api/films/1/"],"created":"2014-12-10T15:36:25.724000Z","edited":"2014-12-22T18:21:15.523587Z","url":"https://swapi.co/api/vehicles/4/"}
```

#### 2.6 访问 species

- 指令

`http://localhost:8080/api/species/1`

- 页面响应

```
{"name":"Human","classification":"mammal","designation":"sentient","average_height":"180","skin_colors":"caucasian, black, asian, hispanic","hair_colors":"blonde, brown, black, red","eye_colors":"brown, blue, green, hazel, grey, amber","average_lifespan":"120","homeworld":"https://swapi.co/api/planets/9/","language":"Galactic Basic","people":["https://swapi.co/api/people/1/","https://swapi.co/api/people/4/","https://swapi.co/api/people/5/","https://swapi.co/api/people/6/","https://swapi.co/api/people/7/","https://swapi.co/api/people/9/","https://swapi.co/api/people/10/","https://swapi.co/api/people/11/","https://swapi.co/api/people/12/","https://swapi.co/api/people/14/","https://swapi.co/api/people/18/","https://swapi.co/api/people/19/","https://swapi.co/api/people/21/","https://swapi.co/api/people/22/","https://swapi.co/api/people/25/","https://swapi.co/api/people/26/","https://swapi.co/api/people/28/","https://swapi.co/api/people/29/","https://swapi.co/api/people/32/","https://swapi.co/api/people/34/","https://swapi.co/api/people/43/","https://swapi.co/api/people/51/","https://swapi.co/api/people/60/","https://swapi.co/api/people/61/","https://swapi.co/api/people/62/","https://swapi.co/api/people/66/","https://swapi.co/api/people/67/","https://swapi.co/api/people/68/","https://swapi.co/api/people/69/","https://swapi.co/api/people/74/","https://swapi.co/api/people/81/","https://swapi.co/api/people/84/","https://swapi.co/api/people/85/","https://swapi.co/api/people/86/","https://swapi.co/api/people/35/"],"films":["https://swapi.co/api/films/2/","https://swapi.co/api/films/7/","https://swapi.co/api/films/5/","https://swapi.co/api/films/4/","https://swapi.co/api/films/6/","https://swapi.co/api/films/3/","https://swapi.co/api/films/1/"],"created":"2014-12-10T13:52:11.567000Z","edited":"2015-04-17T06:59:55.850671Z","url":"https://swapi.co/api/species/1/"}
```

#### 2.7 访问 starship

- 指令 :

````
http://localhost:8080/api/starships/3
````

- 页面响应 :

```
{"name":"Star Destroyer","model":"Imperial I-class Star Destroyer","manufacturer":"Kuat Drive Yards","cost_in_credits":"150000000","length":"1,600","max_atmosphering_speed":"975","crew":"47060","passengers":"0","cargo_capacity":"36000000","consumables":"2 years","hyperdrive_rating":"2.0","MGLT":"60","starship_class":"Star Destroyer","pilots":[],"films":["https://swapi.co/api/films/2/","https://swapi.co/api/films/3/","https://swapi.co/api/films/1/"],"created":"2014-12-10T15:08:19.848000Z","edited":"2014-12-22T17:35:44.410941Z","url":"https://swapi.co/api/starships/3/"}
```

## 具体实现

1. 注册路由
2. 数据库访问
3. JSON格式和状态码输出

...