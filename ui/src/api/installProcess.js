import request from '@/plugin/axios'

//  获取全局状态
export function getState () {
  return request({
    url: `/cluster/status`,
    method: 'get'
  })
}
//  获取检测状态
export function getDetectionState () {
  return request({
    url: `/cluster/precheck`,
    method: 'get'
  })
}

export function getClusterInitConfig () {
  return request({
    url: `/cluster/status-info`,
    method: 'get'
  })
}

//  获取全局状态
export function putInit () {
  return request({
    url: `/cluster/init`,
    method: 'post'
  })
}

//  保存安装记录
export function putRecord (data) {
  return request({
    url: 'https://log.rainbond.com/log/install',
    method: 'post',
    data
  })
}
//  重置处理镜像
export function putrestartpackage () {
  return request({
    url: '/cluster/install/restartpackage',
    method: 'post'
  })
}

//  获取集群配置信息
export function getClusterInfo () {
  return request({
    url: `/cluster/configs`,
    method: 'get'
  })
}

//  查询安装检测结果
export function detectionCluster () {
  return request({
    url: `/cluster/install/status`,
    method: 'get'
  })
}

//  修改集群配置信息
export function putClusterConfig (data) {
  return request({
    url: `/cluster/configs`,
    method: 'PUT',
    data
  })
}
// 集群安装
export function installCluster (data) {
  return request({
    url: `/cluster/install`,
    method: 'post',
    data
  })
}
//  安装集群配置结果
export function getClusterInstallResults () {
  return request({
    url: `/cluster/install/status`,
    method: 'get'
  })
}
//  安装集群配置结果
export function getClusterInstallResultsState (params) {
  return request({
    url: `/cluster/components`,
    method: 'get',
    params: {
      isInit: params ? params.isInit : false
    }
  })
}
//  访问地址
export function getAccessAddress () {
  return request({
    url: `/cluster/address`,
    method: 'get'
  })
}
//  获取可升级版本
export function getUpVersions () {
  return request({
    url: `/upgrade/versions`,
    method: 'get'
  })
}
//  升级版本
export function postUpVersions (data) {
  return request({
    url: `/upgrade/versions`,
    method: 'post',
    data
  })
}

//  平台安装包卸载
export function deleteUnloadingPlatform () {
  return request({
    url: `/cluster/uninstall`,
    method: 'DELETE'
  })
}

export function queryNode (params) {
  return request({
    url: params.mock ? 'http://doc.goodrain.org/mock/48/cluster/nodes' : `/cluster/nodes`,
    method: 'GET',
    params: params
  })
}
