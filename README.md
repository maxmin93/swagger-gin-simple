# Simple Gin + GORM & SQLite + Swagger Docs

> 간단한 Swagger 튜토리얼 따라하기

## [깃허브/swagger-gin-simple](/)

출처

- [Tutorial: Generate Swagger Specification and SwaggerUI for Gin Go Web Framework](https://levelup.gitconnected.com/tutorial-generate-swagger-specification-and-swaggerui-for-gin-go-web-framework-9f0c038483b5)
- [My Favourite Setup for REST Microservices in Go](https://betterprogramming.pub/my-favourite-setup-for-rest-microservices-in-go-770ca18615ba)

### 1) 내용

소스 코드로부터 API 와 설명문을 추출하여 Swagger 문서 생성

- 추출된 주석 [json](http://localhost:3000/swagger/doc.json)
- Gin 웹프레임워크에 Swagger 페이지 연결

```go
module swagger-gin-simple

func main() {
  r := gin.New()

  // The url pointing to API definition
  url := ginSwagger.URL("/swagger/doc.json")
  r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

```

| ![png](https://miro.medium.com/max/1400/1*cxgihTKKB3J1R34pcL7Mzg.png){: width="600"} |
| :----------------------------------------------------------------------------------: |
|                               OpenAPI Specification UI                               |

### 2) Command-lines

```bash
# project root path
$ mkdir swagger-test && cd swagger-test
# directories for gin root & swagger docs
$ mkdir -p docs/ginsimple

# create go.mod
$ go mod init swagger-gin-simple

# install dependencies for gin
$ go get -u github.com/gin-gonic/gin

# write code `main.go`
# ... 문서화 시킬 주석 포함해 코드 작성

# run gin-server
$ go run main.go

# install dependencies for swagger
$ go get -v github.com/swaggo/swag/cmd/swag
$ go get -v github.com/swaggo/gin-swagger
$ go get -v github.com/swaggo/files

# generate swagger pages
$ swag init --parseDependency --parseInternal --parseDepth 1 -g main.go -o docs/ginsimple
2022/11/04 20:22:51 Generate swagger docs....
2022/11/04 20:22:51 Generate general API Info, search dir:./
2022/11/04 20:22:51 create docs.go at  docs/ginsimple/docs.go
2022/11/04 20:22:51 create swagger.json at  docs/ginsimple/swagger.json
2022/11/04 20:22:51 create swagger.yaml at  docs/ginsimple/swagger.yaml
# ==> /swagger/doc.json

# Update main.go to add SwaggerUI.
# ... swagger 페이지 route 생성

# re-run gin-server
$ go run main.go
```
