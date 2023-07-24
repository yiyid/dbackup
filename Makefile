# 设置变量
BINARY_NAME = dbackup

# 默认目标
all: build

# 编译目标
build:
	go build -ldflags "-w -s" -o $(BINARY_NAME) .

# 清理目标
clean:
	go clean
	rm -f $(BINARY_NAME)

# 运行目标
run:
	go run .

# 帮助目标
help:
	@echo "使用说明："
	@echo "make          编译项目"
	@echo "make clean    清理项目"
	@echo "make run      运行项目"
	@echo "make help     显示帮助信息"
