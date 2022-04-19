# Npool go service app template

[![Test](https://github.com/NpoolPlatform/sphinx-coininfo/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/NpoolPlatform/sphinx-coininfo/actions/workflows/main.yml)

[目录](#目录)
- [命令](#命令)
- [步骤](#步骤)
- [最佳实践](#最佳实践)
- [关于mysql](#关于mysql)
- [SoftDelete](#SoftDelete)
- [GRPC](#grpc)

-----------
### 功能

### 命令
* make init ```初始化仓库，创建go.mod```
* make verify ```验证开发环境与构建环境，检查code conduct```
* make verify-build ```编译目标```
* make test ```单元测试```
* make generate-docker-images ```生成docker镜像```
* make sphinx-coininfo ```单独编译服务```
* make sphinx-coininfo-image ```单独生成服务镜像```
* make deploy-to-k8s-cluster ```部署到k8s集群```

### 步骤
* 在github上将模板仓库https://github.com/NpoolPlatform/sphinx-coininfo.git import为https://github.com/NpoolPlatform/my-service-name.git
* git clone https://github.com/NpoolPlatform/my-service-name.git
* cd my-service-name
* mv cmd/sphinx-coininfo cmd/my-service
* 修改cmd/my-service/main.go中的serviceName为My Service
* mv cmd/my-service/ServiceSample.viper.yaml cmd/my-service/MyService.viper.yaml
* 将cmd/my-service/MyService.viper.yaml中的内容修改为当前服务对应内容
* 修改Dockerfile和k8s部署文档为当前服务对应内容
  * grep -rb "sphinx coininfo" ./*
  * grep -rb "sphinx-coininfo" ./*
  * grep -rb "sphinx-coininfo" ./*
  * grep -rb "service\.sample" ./*
  * grep -rb "service\*sample" ./*
  * grep -rb "ServiceSample" ./*
  * grep -rb "SphinxCoininfo" ./*
  * grep -rb "sphinx_coininfo" ./*
  * grep -rb "sphinx_coininfo" ./*
  * grep -rb "sample-service" ./*

### 最佳实践
* 每个服务只提供单一可执行文件，有利于docker镜像打包与k8s部署管理
* 每个服务提供http调试接口，通过curl获取调试信息
* 集群内服务间direct call调用通过服务发现获取目标地址进行调用
* 集群内服务间event call调用通过rabbitmq解耦

### 关于mysql
* 创建app后，从app.Mysql()获取本地mysql client
* [文档参考](https://entgo.io/docs/sql-integration)

### SoftDelete

考虑到安全因素,所有的数据库删除操作都是标志删除

代码集成流程

1. 在项目的 pkg/ent目录下新建文件**rule/rule.go**内容如下
```sql
package rule

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/privacy"
)

func FilterTimeRule() privacy.QueryMutationRule {
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		f.Where(entql.FieldEQ("deleted_at", 0))
		return privacy.Skip
	})
}

```
2. 数据库 **schema** 目录下新建文件 **mixin.go**内容如下
```sql
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/privacy"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/rule"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("created_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("updated_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("deleted_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

func (TimeMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTimeRule(),
		},
	}
}

```

3. 在业务 **schema** 内集成 **mixin.go** 的功能
```sql
func (CoinInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
```

4. make gen-ent


测试代码

```sql
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent"
	_ "github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/runtime"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	id := uuid.New()
	fmt.Println(client.CoinInfo.Create().SetID(id).SetName("name").SetUnit("unit").Exec(context.Background()))
	fmt.Println(client.CoinInfo.Query().All(context.Background()))
	fmt.Println(client.CoinInfo.UpdateOneID(id).SetDeletedAt(uint32(time.Now().Unix())).Exec(context.Background()))
	fmt.Println(client.CoinInfo.Query().All(context.Background()))
	fmt.Println(client.CoinInfo.UpdateOneID(id).SetDeletedAt(0).Exec(context.Background()))
	fmt.Println(client.CoinInfo.Query().All(context.Background()))
}

```


### GRPC
* [GRPC 环境搭建和简单学习](./grpc.md)
