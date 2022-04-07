# Authentication

该 Demo 用于展示业界内常用的认证方式：

- Basic
- Digest
- OAuth
- Bearer

# 前言

由于认证和授权是很容易混为一谈的两个概念，故而此处特此先声明两个概念的区别：

- 认证(Authentication, 缩写 authn): 认证用于验证用户是否具有系统的访问权限，如果通过，该用户就可以访问、操作系统上的支持的资源；
- 授权(Authorization, 缩写 authz): 用于验证用户是否具有访问某个资源的权限，如果通过，就可以进行该资源相应的操作。

如果把项目比做一个工厂，则认证相当于厂门的守卫，而授权则相当于厂内的各个门禁；**认证证明了你是谁，授权决定了你能干什么**

# 四种认证方式

常见的认证方式有四种，分别是 Basic、Digest、OAuth 和 Bearer。

# Basic

Basic 认证（基础认证），是最简单的认证方式。通过将 用户名:密码 进行 base64 编码后，放到 HTTP Authorization Header
中传递至后端，后端通过解析获取用户名和密码，并将用户名和密码与数据库进行比对，如果匹配则通过。例如:

```shell
$ basic=`echo -n 'admin:Admin@1234'|base64`
$ curl -XPOST -H"Authorization: Basic ${basic}" http://127.0.0.1:8080/login
```

虽然通过 base64 进行编码，可以将密码以非明文的方式进行传递，相对于明文而言增加了一定的安全性；但，base64
不是加密技术，攻击者仍然可以截获编码后的字符串，并反编译获取用户名、密码；甚至攻击者可以直接使用编码后的字符串进行重放攻击。

故而，Basic 认证虽然简单，但极不安全。使用 Basic 认证唯一的场景就是将它和 SSL 配合使用，来确保整个过程安全。

# Digest

Digest 认证（摘要认证）是另一种 HTTP 认证协议，与基本认证兼容，但修复了基本认证的缺陷，其特点如下：

- 不实用明文方式在传输中发送密码；
- 可以有效防止重放攻击；
- 可以有选择地防止对报文内容的篡改。

Digest 认证流程如下：

1. 客户端请求服务端资源；
2. 在客户端可以证明它知道密码而从确认身份之前，服务端一律按认证失败处理，返回 401 Unauthorized 状态，并返回WWW-Authenticate头，里面包含认证需要的信息；
3. 客户端根据 WWW-Authenticate 头中信息，选择加密算法，并使用密码随机数 nonce，计算出密码摘要 response，并再次请求服务端；
4. 服务端将客户提供的密码摘要与服务端内部计算出的摘要进行对比，如果匹配，则认证通过，并返回一些与授权相关的附加信息，放入 Authorization-Info 中。

WWW-Authenticate 头中包含的头部信息如下：

|         字段          |                   说明                   |
|:-------------------:|:--------------------------------------:|
|      username       |                  用户名                   |
|        realm        |                   域名                   |
|        nonce        |            服务器发送给客户端的随机字符串             |
|   nc(nonceCount)    |          请求的次数，用于标记、技术、防止重放攻击          |
| cnonce(clientNonce) |      客户端发送给服务端的随机字符串，用于客户端对服务器的认证      |
|         qop         | 保护质量参数，一般是 auth 或 auth-int，该字段会影响摘要的算法 |
|         uri         |                请求的 uri                 |
|      response       |            客户端根据算法算出的密码摘要值             |

下图取自极客时间的[Go 语言开发实战](https://time.geekbang.org/column/intro/100079601?tab=catalog)
![](./imgs/digest.webp)

# OAuth

OAuth（开放授权）是一个开放的授权标准，允许用户让第三方应用访问该用户在某一 Web 服务上存储的私密资源（例如照片、视频、音频等），而无需将用户名和密码提供给第三方应用。OAuth 目前的版本是 2.0 版。

OAuth2.0 共有4种授权方式，分别为密码式、隐藏式、拼接式、授权码模式。

## 密码式

密码式指用户直接将账号、密码告诉第三方，由第三方使用账号、密码向应用换取令牌；所以，使用此种授权方式的前提是，无法采用其他授权方式，并且用户高度信任某第三方应用。

具体流程如下：

1. 网站 A 向用户发出获取用户名和密码的请求；
2. 用户同意后，网站 A 凭借用户名和密码向网站 B 换取令牌；
3. 网站 B 验证身份后，向网站 A 出具令牌，网站 A 凭借令牌可以访问网站 B 对应权限的资源。

## 隐藏式

这种方式适用于前端应用，流程如下：

1. A 网站提供一个跳转到 B 网站的链接，用户点击跳转 B 网站后，并向用户请求授权；
2. 用户登陆 B 网站后，同意授权后，跳转回 A 网站指定的重定向 redirect_url 地址，并携带 B 网站返回的令牌，用户在 B 网站的数据给 A 网站使用。
   这种方式存在着 "中间人攻击"的风险，故而只能用于一些安全性要求不高的场景，并且令牌的有效期要非常短。

## 凭借式

这种方式是在命令行中请求授权，适用于没有前端的命令行应用。认证流程如下：

1. 应用 A 在命令行向应用 B 请求授权，此时应用 A 需要携带应用 B 提前颁发的 secretID 和 secretKey，其中 secretKey 出于安全性考虑，需要后端发送；
2. 应用 B 接收到 secretID 和 secretKey，并进行身份验证，验证通过后返回给应用 A 令牌。

## 授权码模式

种方式就是第三方应用先提前申请一个授权码，然后再使用授权码来获取令牌。相对来说，这种方式安全性更高，前端传送授权码，后端存储令牌，与资源的通信都是在后端，可以避免令牌的泄露导致的安全问题。认证流程如下：

1. 网站 A 提供一个跳转到网站 B 的连接 + redirect_url, 用户点击后跳转至 B 网站；
2. 用户携带向 B 网站提前申请的 client_id，向 B 网站发起身份验证请求；
3. 用户登陆 B 网站，通过验证，授予 A 网站权限，此时网站跳回 redirect_url，其中会有 B 网站通过验证后的授权码附在该 url 后。
4. 网站 A 携带授权码向网站 B 请求令牌，网站 B 验证授权码后，返回令牌即 access_token。

# Bearer

Bearer 认证，也称令牌认证，是一种 HTTP 身份验证方法。Bearer 认证的核心是 bearer token。bearer token 是一个加密字符串，通常由服务端根据密钥生成。客户端请求服务端时，必须在请求头包含
Authorization: Bearer <token>。服务端收到请求后，解析出 <token>, 并校验 <token> 的合法性，如果校验通过，则认证通过。跟基本认证一样， Bearer 认证需要配合 HTTPS
一起使用，来确保认证安全性。

目前最流行的 token 编码方式是 JSON Web Token（JWT）。

在典型的业务场景中，为了区分用户和保障安全，会对 API 请求进行鉴权，但是又不能要求每一个请求都进行登陆操作，合理的做法是，第一次登陆后，服务器端产生一个具有一定时效性的 token，并将其存储在缓存或数据库中；客户端将它存储在浏览器的
Cookie 或 LocalStorage 中，之后的每次请求都携带这个 token，请求到达服务器后，服务器端用这个 token 进行请求鉴权；

或者也可以使用更简单的方式，直接用密钥来签发 token；这样，就可以省下额外的存储，也可以减少每一次请求时对数据库的压力；这种方式在业界内已经有一种标准的实现方式，就是 JWT。

## JWT

JWT 是 Bearer Token 的一个具体实现，由 JSON 数据格式组成，通过 HASH 散列算法生成一个字符串；该字符串可以用来进行授权或信息交换。

使用 JWT Token 进行认证有很多优点，例如，无需在服务端存储该用户数据，可以减轻服务器压力；而且采用 JSON 数据格式，比较易读；除此之外，使用 JWT Token 还有跨语言、轻量级等优点。

### JWT 认证流程

下图取自极客时间的[Go 语言开发实战](https://time.geekbang.org/column/intro/100079601?tab=catalog)
![](./imgs/jwt.jpg)

1. 客户端使用账号、密码请求登陆；
2. 服务端收到登陆请求，进行验证，如果通过，则签发一个 Token 返回给客户端；反之，通知客户端，验证失败；
3. 客户端收到 Token 后，将其进行缓存，然后每次请求携带该 Token；
4. 服务端收到请求后，对请求中的 Token 进行验证，如果验证通过，则进行后续业务。

### JWT 格式

JWT 由三部分组成，分别是：Header、Payload 和 Signature，它们之间使用圆点 . 连接，例如：

`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NDI4NTY2MzcsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2MzUwODA2MzcsInN1YiI6ImFkbWluIn0.Shw27RKENE_2MVBq7-c8OmgYdF92UmdwS8xE-Fts2FM`

JWT 中，每部分包含信息如下：

下图取自极客时间的[Go 语言开发实战](https://time.geekbang.org/column/intro/100079601?tab=catalog)
![](./imgs/jwt_data.jpg)

三部分的信息如下：

1. Header

JWT Token 的 Header 中，包含两部分信息，一是 Token 的类型，二是 Token 所使用的加密算法，例如：

```
{
  "typ": "JWT",
  "alg": "HS256"
}
```

参数说明：

- typ: 说明 Token 的类型是 JWT；
- alg: 说明 Token 的加密算法，此处是 HS256（alg 算法可以有多种）。


2. Payload

Payload 中携带 Token 的具体内容由三部分组成：JWT 标准中注册的声明（可选）、公共的声明、私有的声明。

标准声明如下：

| 字段                   | 描述                                                                       |
|----------------------|--------------------------------------------------------------------------|
| iss(Issuer)          | JWT Token 签发者，该值因该是大小写敏感的字符串或 uri                                        |
| sub(Subject)         | 主题，sub 可以用来鉴别一个用户                                                        |
| exp(Expiration time) | JWT Token 过期时间                                                           |
| aud(Audience)        | 接收 Token 的一方，必须是大小写敏感的字符串或 uri，一般是 App 或模块或服务端的安全策略在签发和验证时，必须确保 aud 的一致。 |
| iat(Issued At)       | JWT Token 签发时间                                                           |
| nbf(Not Before)      | JWT Token 生效时间                                                           |
| jti(JWT ID)          | JWT Token ID,令牌的唯一标识符，通常用于一次性消费的令牌                                       |

Payload 例子：

```
{ 
   "aud": "jwt-client.com", 
   "exp": 1604158987, 
   "iat": 1604151787, 
   "iss": "jwt-demo", 
   "nbf": 1604151787
}
```

除此之外，还有公共的声明和私有的声明。公共的声明可以添加任何的需要的信息，一般添加用户的相关信息或其他业务需要的信息，注意不要添加敏感信息；私有声明是客户端和服务端所共同定义的声明，因为 base64
是对称解密的，所以一般不建议存放敏感信息。

3. Signature（签名）

Signature 是 Token 的签名部分，通过如下方式生成：将 Header 和 Payload 分别 base64 编码后，用 . 连接。然后再使用 Header 中声明的加密方式，利用 secretKey
对连接后的字符串进行加密，加密后的字符串即为最终的 Signature。

签名后服务端会返回生成的 Token，客户端下次请求会携带该 Token。服务端收到 Token 后会解析出 header.payload，然后用相同的加密算法和密钥对 header.payload 再进行一次加密，得到
Signature。并且，对比加密后的 Signature 和收到的 Signature 是否相同，如果相同则验证通过，不相同则返回 HTTP 401 Unauthorized 的错误。

JWT 使用建议：

- 不要存放敏感信息在 Token 里；
- Payload 中的 exp 值不要设置得太大，一般开发版本 7 天，线上版本 2 小时。当然，也可以根据需要自行设置。 