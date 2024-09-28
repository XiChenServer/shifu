# Shifu 部署与运行指南

## 具体任务

1. **部署并运行 Shifu**  
   请参考以下指南：[Shifu 安装教程](https://shifu.dev/docs/tutorials/demo-install/)

2. **运行酶标仪的数字孪生**  
   请参考以下链接：[酶标仪数字孪生交互](https://shifu.dev/docs/tutorials/demo-try/#3-interact-with-the-microplate-reader)

3. **编写 Go 应用**  
   - 定期轮询获取酶标仪的 `/get_measurement` 接口。
   - 将返回值平均后打印出来，轮询时间可自定义。
   - Go 应用需要容器化并运行在 Shifu 的 Kubernetes 集群中。
   - 最终通过 `kubectl logs` 命令查看打印的值。

## 环境配置

以下是我的电脑及所使用软件的配置：

- **操作系统**: 
  - **Distributor ID**: Ubuntu
  - **Description**: Ubuntu 22.04.5 LTS
  - **Release**: 22.04
  - **Codename**: jammy

- **Docker**:
  ```bash
  $ docker --version
  Docker version 27.2.1, build 9e34c9b
- **Go**:
  ```bash
  $ go --version
  go version go1.21.4 linux/amd64
- **kind**:
  ```bash
  $ kind --version
  kind v0.20.0 go1.20.4 linux/amd64
## 主要命令
```bash
docker build --tag measurement:v0.0.1 .
sudo kind load docker-image measurement:v0.0.1
sudo kubectl run measurement --image=measurement:v0.0.1
sudo kubectl logs measurement -f
sudo kubectl delete pod measurement
```
# 启动测量应用
$ kubectl run measurement --image=measurement:v0.0.1

# 查看日志
$ kubectl logs measurement -f

![截图 2024-09-28 17-15-58](https://github.com/user-attachments/assets/93bb9ba2-bb3d-4189-91ce-e2dcad7eef58)
![截图 2024-09-28 17-15-47](https://github.com/user-attachments/assets/3cb0e820-2886-491c-9101-151817a8152b)
![截图 2024-09-28 10-48-39](https://github.com/user-attachments/assets/9e75b9c7-fa1b-417c-9b97-753226c2f97c)
![截图 2024-09-28 10-48-14](https://github.com/user-attachments/assets/70f76abd-0f48-47ea-9588-0ed0997b7640)
![截图 2024-09-28 10-47-23](https://github.com/user-attachments/assets/51d85a4f-f59a-4d7a-8182-1238dac9086a)
![截图 2024-09-27 17-21-57](https://github.com/user-attachments/assets/d4800b40-3c71-4a05-884e-8621de494422)
![截图 2024-09-27 17-20-24](https://github.com/user-attachments/assets/dc2e42bb-6d42-4887-8002-e2e1beea6a58)
