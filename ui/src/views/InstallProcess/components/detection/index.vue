<template>
  <div class="text item d2-p-20">
    <el-row :gutter="12" class="d2-mb">
      <el-col :span="19" class="d2-f-16">名称</el-col>
      <el-col :span="5" class="d2-f-16 d2-text-cen">状态</el-col>
    </el-row>
    <div class="componentTitle">
      <el-row
        :gutter="12"
        v-for="(item, index) in list"
        :key="index"
        style="border-bottom: 1px solid #ebeef5 !important;"
      >
        <div class="componentBox">
          <el-col :span="19" class="d2-f-16">
            {{ textMap[item.type] || item.type }}
          </el-col>
          <el-col :span="5" class="d2-f-16 d2-text-cen">
            <i
              v-if="item.status == 'True'"
              class="el-icon-circle-check d2-f-20"
              style="color:#52c41a"
            ></i>
            <i
              v-else
              style="color:#606266"
              class="el-icon-circle-check  d2-f-20"
            ></i>
          </el-col>
        </div>

        <div v-if="item.reason || item.message">
          <div class="componentBox">
            <el-row>
              <el-col :span="4" class="d2-f-16  minComponentColor">原因</el-col>
              <el-col :span="17" class="d2-f-16  descColor">{{
                item.reason
              }}</el-col>
            </el-row>
          </div>

          <div class="componentBox">
            <el-row>
              <el-col :span="4" class="d2-f-16  minComponentColor">消息</el-col>
              <el-col :span="17" class="d2-f-16  descColor">{{
                item.message
              }}</el-col>
            </el-row>
          </div>
        </div>
      </el-row>
    </div>
    <div class="d2-text-cen" style="margin:1rem 0">
      <el-button
        type="primary"
        style="margin-right:50px"
        :loading="loading"
        @click="handleUpstep"
      >
        {{ $t("page.install.config.upstep") }}
      </el-button>
      <el-button
        v-if="pass"
        type="primary"
        :loading="loading"
        @click="handleInstall"
      >
        {{ $t("page.install.config.startInstall") }}
      </el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'detection',
  data () {
    return {
      list: [],
      loading: false,
      pass: false,
      textMap: {
        ClusterInitialized: '集群初始化',
        DatabaseRegion: 'Region 数据库',
        DatabaseConsole: 'Console 数据库',
        ImageRepository: '镜像仓库',
        KubernetesVersion: 'Kubernetes 版本',
        KubernetesStatus: 'Kubernetes 集群',
        ContainerNetwork: '容器网络',
        Storage: '共享存储'
      }
    }
  },
  created () {
    document.documentElement.scrollTop = 0
    this.handleDetection()
  },
  beforeDestroy () {
    this.timer && clearInterval(this.timer)
  },
  methods: {
    fetchErrMessage (err) {
      return err && typeof err === 'object' ? JSON.stringify(err) : '/cluster/install'
    },
    handleDetection () {
      this.timer && clearInterval(this.timer)
      this.$store.dispatch('fetchDetectionState').then(res => {
        if (res && res.code === 200) {
          this.list = res.data.conditions
          this.pass = res.data.pass
        }
      })
      this.timer = setTimeout(() => {
        this.handleDetection()
      }, 1000)
    },
    handleInstall () {
      this.loading = true
      this.$store
        .dispatch('installCluster', {})
        .then(en => {
          if (en && en.code === 200) {
            this.$emit('onhandleStartRecord')
            this.$emit('onResults')
          } else {
            this.handleCancelLoading(en)
          }
        })
        .catch(err => {
          this.handleCancelLoading(err)
        })
    },
    handleCancelLoading (msg) {
      const message = this.fetchErrMessage(msg)
      this.$emit('onhandleErrorRecord', 'failure', `${message}`)
      this.loading = false
    },
    handleUpstep () {
      if (this.list.length > 0 && !this.pass) {
        let mag = ''
        this.list.map(item => {
          if (item.reason) {
            mag += `原因:${item.reason};消息:'${item.message};`
          }
        })
        if (mag) {
          this.$emit('onhandleErrorRecord', 'failure', `${mag}`)
        }
      }
      this.$emit('onUpstep')
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.d2-p-20 {
  padding: 20px;
}
.descColor {
  color: #606266;
}
.minComponentColor {
  color: #99a9bf !important;
}
.d2-f-16 {
  font-size: 16px;
}
.d2-text-cen {
  text-align: center;
}
.d2-f-14 {
  font-size: 14px;
}
.componentBox {
  min-height: 39px;
  line-height: 39px;
}
.d2-f-20 {
  font-size: 20px;
}
</style>

<style lang="scss">
.componentTitle {
  .el-collapse-item__header {
    border-bottom: 1px solid #ebeef5 !important;
    width: 100% !important;
  }
  .el-collapse-item__header:hover {
    background: #f5f7fa;
  }
}
</style>
