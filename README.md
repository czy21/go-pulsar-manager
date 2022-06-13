## go-pulsar-manager
pulsar 多环境管理平台，基于apache/pulsar-manager的基础功能进行拓展
## Development

### Requirements
go 1.18.3+

node 16.14.2+
```bash
cd go-pulsar-manager
sed -i 's|http://gitea.cluster.com/czyhome|https://github.com/czy21|g' .gitmodules
git submodule update --recursive --init
```
### Quickstart
init mysql-schema.sql
```bash
bash build.sh
cp api/app.yaml api/build/
# modify db.url in api/build/app.yaml
cd api/build/
chmod +x main
GIN_MODE=release ./main
```
open browser: http://localhost:3001
#### Frontend
```bash
yarn watch:share
yarn start:web
```
