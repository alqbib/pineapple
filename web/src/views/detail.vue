<template>
  <div>
    <div class="cluster">
      <a-row >
        <a-col :span="17"><a-button type="primary" ghost>{{ info.cluster_name }}</a-button></a-col>
        <a-col :span="6">
          <a-button type="primary" shape="circle" size="small" style="margin-right:10px;" @click="refresh()">
              <template #icon><redo-outlined /></template>
          </a-button>
          <a-select
              ref="select"
              v-model:value="inter"
              @change="change"
            >
              <a-select-option value="15">15sec</a-select-option>
              <a-select-option value="30">30sec</a-select-option>
              <a-select-option value="60">1min</a-select-option>
              <a-select-option value="300">5min</a-select-option>
            </a-select>
            <span  style="padding-left:10px;font-size: 14px;">{{ url }}</span>
           
        </a-col>  
        <a-col :span="1">
          <a-button type="primary" @click="loginOut">退出</a-button>
        </a-col>
      </a-row>
      <a-divider :style="{'height': '3px', 'backgroundColor': health.status}" />

      <a-row :gutter="[16,16]">
        <a-col :span="4">
          <a-card>
              <a-statistic
                title="shards"
                :value="health.shards"
                :value-style="{ color: '#fff' }"
                style="margin-right: 50px"
              >
              </a-statistic>
            </a-card>
        </a-col>
        <a-col :span="4">
          <a-card>
              <a-statistic
                title="active_shards_percent"
                :value="health.active_shards_percent"
                :value-style="{ color: '#fff' }"
                style="margin-right: 50px"
              >
              </a-statistic>
            </a-card>
        </a-col>
        <a-col :span="4">
          <a-card>
              <a-statistic
                title="node.total"
                :value="health['node.total']"
                :value-style="{ color:'#fff' }"
                style="margin-right: 50px"
              >
              </a-statistic>
            </a-card>
        </a-col>
        <a-col :span="4">
          <a-card>
              <a-statistic
                title="node.data"
                :value="health['node.data']"
                :value-style="{ color: '#fff' }"
                style="margin-right: 50px"
              >
              </a-statistic>
            </a-card>
        </a-col>
        <a-col :span="4">
          <a-card>
              <a-statistic
                title="pri"
                :value="health.pri"
                :value-style="{ color: '#fff' }"
                style="margin-right: 50px"
              >
              </a-statistic>
            </a-card>
        </a-col>
        <a-col :span="4">
          <a-card>
              <a-statistic
                title="pending_tasks"
                :value="health.pending_tasks"
                :value-style="{ color: '#fff' }"
                style="margin-right: 50px"
              >
              </a-statistic>
            </a-card>
        </a-col>
      </a-row>
    </div>

    <a-table :columns="columns" :data-source="nodes" :pagination="pagination" bordered>
      <template #headerCell="{ column }">
        <template v-if="column.dataIndex != 'ip'">
          <span :style="{color: indices.filter(i=>{return i.index == column.dataIndex})[0].health}">
            {{ column.dataIndex }}
          </span>
          <br/>
          <span style="font-size: 10px;" v-html="getIndices(column.dataIndex)"></span>
        </template>
      </template>
      <template #bodyCell="{ column, text, record  }">
        <template v-if="column.dataIndex =='ip'">
            <p>{{ record.ip  }}</p>
            <a-row :gutter="[16, 16]">
              <a-col :span="12">
                <div :title="'ram ：' + record['ram.percent'] + '%'">
                  ram
                  <a-progress :percent="record['ram.percent']" size="small"   :showInfo="false" strokeWidth="3"/>
                </div>
                  
              </a-col>
              <a-col :span="12">
                <div :title="'heap：' + record['heap.percent'] + '%'">
                  heap
                  <a-progress :percent="record['heap.percent']" size="small"    :showInfo="false" strokeWidth="3"/>
                </div>
                 
              </a-col>
              
            </a-row>
        </template>
        <template v-else>
          <div class="flex">
            <template v-for="s of shards.filter(s=>{return s.index == column.dataIndex  && s.node == record.name})">
              <div class="square" :style="{border: s.prirep=='p'?'2px solid green':'2px dashed green'}">
                <div class="ele">
                    {{s.shard}}
                  </div>
              </div>
            </template> 
            </div>
        </template>
      </template>
    </a-table>

    
  </div>
</template>

<script setup>
import {ref,  onBeforeUnmount, reactive, onMounted} from 'vue'
import axios from '@/api/axios'
import { RedoOutlined } from '@ant-design/icons-vue';
import { useRouter } from 'vue-router'
import {useStore} from 'vuex'


const store = useStore()



let token = (sessionStorage.getItem("search-token")??'') != ''?sessionStorage.getItem("search-token"):store.state.token

const info = ref({
  cluster_name:''
})

const inter = ref('15')

var health = ref({
  status: '',
  shards: '',
  active_shards_percent:'',
  'node.data':'',
  'node.total': '',
  pending_tasks:''
})
const nodes = ref([])
const columns = ref([
  {
    title: '',
    dataIndex: 'ip',
    width: 250,
  }
])

var loading = ref(true)
var timer = ref(0)

const indices = ref([])
const shards = ref([])

const { currentRoute } = useRouter()
const router = useRouter();
const url = ref(currentRoute.value.query.url)

const pagination = reactive({
    total: 0,
    current: 1,
    pageSize: 6,
    showQuickJumper: false,
    showSizeChanger: false,
    showTotal: (total, range) => `第${range.join("-")}条/总共${total}条`,
    onChange: (page, pageSize) => changePage(page, pageSize),
    onShowSizeChange: (current, size) => showSizeChange(current, size),
})


const changePage = (page, pageSize)=>{

  columns.value.splice(1, pageSize)

  indices.value.slice((page - 1) * pageSize, page* pageSize).forEach(i=>{
            columns.value.push(
              {

                title: i.index,
                dataIndex: i.index,
                width: 300
              })
      })
  pagination.current = page
}

const showSizeChange = (current, size)=>{
  console.log(current)
  console.log(size)
}


const getIndex = ()=>{
  axios.get('/index',{
            headers: {
                'Authorization': token,
            }
    }).then(res => {

    if(res.code == 200){
      info.value = res.data
    }

  })
}

const getIndices = (dataIndex)=>{
  const data = indices.value.filter(i=>{return i.index == dataIndex})[0]
  const shard = shards.value.filter(s=>{return s.index == dataIndex && s.state == 'UNASSIGNED'})
  const tmp = shard.length > 0?`<br/><span style=\"color: yellow\">UNASSIGNED：${shard.length}</span>`:''
  return `docs：${data['docs.count']}|size：${data['store.size']}|status：${data['status']}${tmp}`
}


const refresh = ()=>{


  axios.get('/refresh',{
            headers: {
                'Authorization': token,
            }
    }).then(res => {

    if(res.code == 200){

      nodes.value = []
      health.value = []
      indices.value = []
      columns.value = [{
        title: '',
        dataIndex: 'ip',
        width: 220,
      }]

      loading.value = false
      health.value = res.data.health[0]
      health.value.nodeData = res.data.health[0]['node.data']
      health.value.nodeTotal = res.data.health[0]['node.total']
    
      res.data.indices.slice((pagination.current - 1) * pagination.pageSize, pagination.current * pagination.pageSize).forEach(i=>{
            columns.value.push(
              {
                title: i.index,
                dataIndex: i.index,
                width: 300
              })
      })                
      indices.value = res.data.indices
      shards.value = res.data.shards
      nodes.value = res.data.nodes
      pagination.total = indices.value.length

    }

  }).catch(()=>{
        loading.value = false
  })
}


const loginOut = ()=>{
  sessionStorage.removeItem('search-token')
  router.push({ path: '/' })
}

const change = ()=>{
  if (timer.value) {

        clearInterval(timer.value)
        timer.value = setInterval(() => {

           refresh()

        }, inter.value * 1000) 
  }
}

timer.value = setInterval(() => {

    refresh()

}, inter.value * 1000) 




onMounted(() => {

  getIndex()
  refresh()

})



onBeforeUnmount(()=>{
    if (timer.value) {
        clearInterval(timer.value); 
    }
})
</script>
<style lang='less' scoped>
.cluster{
  padding: 10px;
}
.ant-divider-horizontal{
  margin: 12px 0;
}
:deep(.ant-card){
  background-color: rgb(105,156,228,0.8);
  
}
:deep(.ant-card .ant-card-body){
  padding: 5px;
}
:deep(.ant-statistic .ant-statistic-title){
  color: #fff;
}
:deep(.ant-statistic-content){
  float: right;
}

.ele{
  position: absolute;
  margin-left: 7px;
  margin-bottom: 4px;
  font-size: 12px;
}
.square {
  width: 25px;
  height: 25px;
  margin-right: 5px;
}
.flex {
  display: flex;
  flex-wrap: wrap
}

</style>