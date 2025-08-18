# GINP-API 命令行工具

## 简介

GINP-API 是一个基于 Gin 框架的 API 开发工具，提供了代码生成、实体管理等功能。该工具可以快速生成实体、控制器、服务、模型等代码，大大提高开发效率。

## 安装

### 方法一：从源码编译

1. 克隆仓库
```bash
git clone https://github.com/yourusername/ginp-api.git
cd ginp-api
```

2. 编译命令行工具
```bash
# 使用项目自带的编译脚本
./cmd/gencode/build_gapi.sh

# 或者手动编译
go build -o ./build/gapi ./cmd/gencode/main.go
chmod +x ./build/gapi
```

3. 添加到 PATH（推荐）
```bash
# 临时添加到当前会话
export PATH="$(pwd)/build:$PATH"

# 永久添加到 PATH（根据您使用的 shell 选择）
# Bash
echo 'export PATH="/完整路径/ginp-api/build:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Zsh
echo 'export PATH="/完整路径/ginp-api/build:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

4. 创建系统级符号链接（可选，需要管理员权限）
```bash
sudo ln -sf "$(pwd)/build/gapi" /usr/local/bin/gapi
```

## 使用方法

### 查看帮助和版本
```bash
# 查看帮助信息
gapi --help

# 查看版本
gapi -v
# 或
gapi --version
```

### 创建实体并生成 CRUD 代码
```bash
# 交互式方式（会提示输入实体名称）
gapi gen entity

# 直接指定实体名称
gapi gen entity -c UserGroup
# 或
gapi gen entity --create UserGroup
```

### 批量生成多个实体的 CRUD 代码
```bash
# 交互式方式（会提示输入实体名称列表）
gapi gen crud

# 直接指定实体名称列表，多个实体用逗号分隔
gapi gen crud -e "UserGroup,UserRole,UserPermission"
# 或
gapi gen crud --entities "UserGroup,UserRole,UserPermission"
```

### 生成实体字段常量
```bash
gapi gen const
```

### 新增 API 接口
```bash
# 交互式方式（会提示输入 API 名称和目录）
gapi gen api

# 直接指定 API 名称和目录
gapi gen api -a GetUserInfo -d user/cuser
# 或
gapi gen api --add GetUserInfo --dir user/cuser
```

## 命令说明

### 根命令
- `gapi`: 显示帮助信息
- `gapi -v` 或 `gapi --version`: 显示版本信息
- `gapi --help`: 显示帮助信息

### 生成代码命令
- `gapi gen`: 代码生成命令组
  - `gapi gen entity`: 创建实体并生成 CRUD 代码
    - `-c, --create`: 指定实体名称（大驼峰命名，如 UserGroup）
  - `gapi gen crud`: 批量生成多个实体的 CRUD 代码
    - `-e, --entities`: 指定实体名称列表，多个实体用逗号分隔（如 "UserGroup,UserRole"）
  - `gapi gen const`: 生成实体字段常量
  - `gapi gen api`: 新增 API 接口控制器
    - `-a, --add`: 指定 API 名称（大驼峰命名，如 GetUserInfo）
    - `-d, --dir`: 指定目录路径（如 user/cuser）

### 生成的文件说明

当使用 `gapi gen entity` 或 `gapi gen crud` 命令生成 CRUD 代码时，会为每个实体生成以下文件：

1. **实体文件** (entity): `internal/app/gapi/entity/{entity_name}.e.go`
2. **路由文件** (router): `internal/app/gapi/controller/c{entity_name}/{entity_name}.r.go`
3. **控制器文件** (controller): `internal/app/gapi/controller/c{entity_name}/{entity_name}.c.go`
4. **服务文件** (service): `internal/app/gapi/service/s{entity_name}/{entity_name}.s.go`
5. **模型文件** (model): `internal/app/gapi/model/m{entity_name}/{entity_name}.m.go`
6. **字段常量文件** (fields): `internal/app/gapi/model/m{entity_name}/{entity_name}.f.go`

## 示例

### 创建单个实体并生成 CRUD 代码
```bash
# 创建 UserGroup 实体
gapi gen entity -c UserGroup

# 查看生成的文件
ls -la internal/app/gapi/entity/user_group.e.go
ls -la internal/app/gapi/controller/cusergroup/user_group.c.go
ls -la internal/app/gapi/controller/cusergroup/user_group.r.go
ls -la internal/app/gapi/service/susergroup/user_group.s.go
ls -la internal/app/gapi/model/musergroup/user_group.m.go
ls -la internal/app/gapi/model/musergroup/user_group.f.go
```

### 批量生成多个实体的 CRUD 代码
```bash
# 批量生成 UserGroup、UserRole 和 UserPermission 实体的 CRUD 代码
gapi gen crud -e "UserGroup,UserRole,UserPermission" -p user

# 或者使用交互式方式
gapi gen crud
# 然后输入: UserGroup,UserRole,UserPermission
```

### 在 controller/user/cuser 目录下新增 GetUserInfo API
```bash
gapi gen api -a GetUserInfo -d user/cuser
```

### 生成实体字段常量
```bash
# 生成所有实体的字段常量
gapi gen const
```

## 开发流程示例

以下是使用 GINP-API 开发一个新功能的典型流程：

1. **创建实体**
   ```bash
   gapi gen entity -c Product
   ```

2. **修改实体定义**
   编辑 `internal/app/gapi/entity/product.e.go` 文件，添加所需字段

3. **生成字段常量**
   ```bash
   gapi gen const
   ```

4. **添加自定义 API**
   ```bash
   gapi gen api -a GetProductsByCategory -d product/cproduct
   ```

5. **实现业务逻辑**
   编辑生成的控制器、服务和模型文件，实现具体业务逻辑

## 常见问题

### Q: 如何在现有实体中添加新字段？
A: 直接编辑实体文件（.e.go），然后运行 `gapi gen const` 生成更新后的字段常量。

### Q: 如何批量生成多个已有实体的 CRUD 代码？
A: 使用 `gapi gen crud -e "Entity1,Entity2,Entity3"` 命令。

### Q: 生成的代码在哪里？
A: 所有生成的代码都在 `internal/app/gapi/` 目录下，按照实体、控制器、服务和模型分类存放。

## 贡献

欢迎提交 Pull Request 或提出 Issue。如果您有任何改进建议或发现了 bug，请随时联系我们。

## 许可证

[MIT](LICENSE)