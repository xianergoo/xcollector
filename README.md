
# DataCollector

DataCollector 是一个用于数据采集和处理的系统。该系统通过 TCP 协议与硬件设备交互，收集数据并进行处理。系统能够支持类似地铁刷卡或门禁刷卡的操作，并能够根据采集的数据进行复杂的交互逻辑处理。

## 技术栈

- **后端**: Go 语言使用 Gin 框架
- **前端**: React 使用 Create React App 构建
- **网络通信**: RESTful API
- **数据采集**: TCP 协议

## 安装

### 后端安装

1. 安装 Go 语言环境:
   ```bash
   brew install go # 如果您使用 macOS/Linux
   choco install golang # 如果您使用 Windows
   ```

2. 克隆项目仓库:
   ```bash
   git clone https://github.com/yourusername/DataCollector.git
   cd DataCollector
   ```

3. 安装依赖:
   ```bash
   go mod init DataCollector
   go mod tidy
   ```

4. 运行后端服务:
   ```bash
   go run ~/main.go
   ```

### 前端安装

1. 安装 Node.js 和 npm:
   ```bash
   brew install node # 如果您使用 macOS/Linux
   choco install nodejs # 如果您使用 Windows
   ```

2. 安装前端依赖:
   ```bash
   cd frontend
   npm install
   ```

3. 运行前端开发服务器:
   ```bash
   npm start
   ```

## 运行

### 后端

1. 确保 TCP 服务配置正确。
2. 运行后端服务:
   ```bash
   go run ~/main.go
   ```

### 前端

1. 确保前端依赖已安装。
2. 运行前端开发服务器:
   ```bash
   npm start
   ```

## 功能

- **数据采集**: 通过 TCP 协议从硬件设备采集数据。
- **数据处理**: 对采集的数据进行处理，支持复杂的业务逻辑。
- **交互逻辑**: 根据数据处理结果与用户进行交互。
- **配置管理**: 提供基本的配置管理功能。

## 系统架构

### 后端架构

- **数据采集模块**: 负责与硬件设备建立 TCP 连接并读取数据。
- **数据处理模块**: 对采集的数据进行清洗、转换和存储。
- **交互逻辑模块**: 根据处理后的数据生成相应的交互逻辑。

### 前端架构

- **数据展示**: 展示采集到的数据和处理结果。
- **配置管理**: 允许用户配置数据采集和处理的参数。
- **交互界面**: 提供用户友好的交互界面。

## API 文档

### 数据采集

- **启动数据采集**:
  - **URL**: `/collect/start`
  - **Method**: `POST`
  - **Request Body**:
    ```json
    {
      "device_id": "device1",
      "interval": 5000
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Data collection started."
    }
    ```

- **停止数据采集**:
  - **URL**: `/collect/stop`
  - **Method**: `POST`
  - **Request Body**:
    ```json
    {
      "device_id": "device1"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Data collection stopped."
    }
    ```

### 数据处理

- **获取处理结果**:
  - **URL**: `/process/result`
  - **Method**: `GET`
  - **Query Parameters**:
    - `device_id`: 设备 ID
    - `timestamp`: 时间戳
  - **Response**:
    ```json
    {
      "device_id": "device1",
      "timestamp": "2024-08-01T00:00:00Z",
      "processed_data": {
        "status": "success",
        "details": "Card accepted."
      }
    }
    ```

## 贡献指南

1. 叉取 (fork) 该项目。
2. 创建一个新的分支:
   ```bash
   git checkout -b feature-name
   ```
3. 实现您的功能或修复 bug。
4. 提交更改:
   ```bash
   git commit -m "Add some feature"
   ```
5. 推送到您的分支:
   ```bash
   git push origin feature-name
   ```
6. 发起拉取请求 (pull request)。

## 许可证

本项目使用 MIT 许可证发布。更多信息参见 [LICENSE](LICENSE) 文件。
