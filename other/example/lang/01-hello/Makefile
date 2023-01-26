
# 没有命令项，默认执行
.DEFAULT_GOAL := build

# 区别，同名目录 e.g. fmt/
.PHONY: fmt vet build

test:
	@printf '\e[32m测试\e[0m\n'


# 格式化
fmt:
	go fmt ./...

# 检查隐藏的语法错误
vet: fmt
	go vet ./...

# 编译
build: vet
	go build -o main main.go


clean: build
	rm -rf ./main


# front: end  target 命令项以及依赖关系

