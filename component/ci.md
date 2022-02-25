### Pipeline
- commit/MR会触发pipeline执行

### Stages

- Pipeline包含多个Stages
- 多个Stages依次执行
- 包含多个流程，如：安装依赖、运行测试、编译、部署测试服务器、部署生产服务器

### Jobs

- 一个Stages中包含多个Jods
- 多个Jods并行执行

### Runner

- 用来执行构建任务
- 一般把runner安装到其他机器

### 一个例子

```yaml
stages:
  - install_deps
  - test
  - build
  - deploy_test
  - deploy_production

cache:
  key: ${CI_BUILD_REF_NAME}
  paths:
    - node_modules/
    - dist/


# 安装依赖
install_deps:
  stage: install_deps
  only:
    - develop
    - master
  script:
    - npm install


# 运行测试用例
test:
  stage: test
  only:
    - develop
    - master
  script:
    - npm run test


# 编译
build:
  stage: build
  only:
    - develop
    - master
  script:
    - npm run clean
    - npm run build:client
    - npm run build:server


# 部署测试服务器
deploy_test:
  stage: deploy_test
  only:
    - develop
  script:
    - pm2 delete app || true
    - pm2 start app.js --name app


# 部署生产服务器
deploy_production:
  stage: deploy_production
  only:
    - master
  script:
    - bash scripts/deploy/deploy.sh

```







