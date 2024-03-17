<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
          <el-table-column align="left" label="酒店名称" prop="hotelName" :show-overflow-tooltip="true" width="250" />
          <el-table-column align="left" label="英文名" prop="hotelNameEn" :show-overflow-tooltip="true" width="250" />
          <el-table-column align="left" label="酒店地址" prop="address":show-overflow-tooltip="true"  width="250" />
          <el-table-column align="left" label="酒店英文地址" prop="addressEn" width="250" />
          <el-table-column align="left" label="城市名称" prop="cityName" :show-overflow-tooltip="true" width="250" />
          <el-table-column align="left" label="支持的信用卡" prop="creditCards" width="120" />
        <!--<el-table-column align="left" label="百度维度" prop="baiduLat" width="120" />-->
        <!--<el-table-column align="left" label="百度经度" prop="baiduLon" width="120" />-->
        <!--<el-table-column align="left" label="取消政策" prop="canclePolicy" width="120" />-->
        <!--<el-table-column align="left" label="创建人" prop="createBy" width="120" />-->
        <!-- <el-table-column align="left" label="创建时间" width="180">-->
        <!--    <template #default="scope">{{ formatDate(scope.row.createTime) }}</template>-->
        <!-- </el-table-column>-->
        <!-- <el-table-column align="left" label="装修时间" width="180">-->
        <!--    <template #default="scope">{{ formatDate(scope.row.decorateTime) }}</template>-->
        <!-- </el-table-column>-->
        <!--<el-table-column align="left" label="目的地标签" prop="destinationLabel" width="120" />-->
        <!--<el-table-column align="left" label="酒店设施列表" prop="facilities" width="120" />-->
        <!--<el-table-column align="left" label="高德经度" prop="gaodeLat" width="120" />-->
        <!--<el-table-column align="left" label="高德维度" prop="gaodeLon" width="120" />-->
        <!--<el-table-column align="left" label="谷歌维度" prop="googleLat" width="120" />-->
        <!--<el-table-column align="left" label="谷歌经度" prop="googleLon" width="120" />-->
        <!--<el-table-column align="left" label="酒店编号" prop="hotelId" width="120" />-->
        <!--<el-table-column align="left" label="电话" prop="hotelPhone" width="120" />-->
        <!--<el-table-column align="left" label="酒店介绍" prop="introduction" width="120" />-->
        <!-- <el-table-column align="left" label="开业时间" width="180">-->
        <!--    <template #default="scope">{{ formatDate(scope.row.openingTime) }}</template>-->
        <!-- </el-table-column>-->
        <!--<el-table-column align="left" label="星级(1/2/3/4/5)" prop="star" width="120" />-->
        <!--<el-table-column align="left" label="修改人" prop="updateBy" width="120" />-->
        <!-- <el-table-column align="left" label="修改时间" width="180">-->
        <!--    <template #default="scope">{{ formatDate(scope.row.updateTime) }}</template>-->
        <!-- </el-table-column>&ndash;&gt;-->
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateHtlHotelFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="酒店地址:"  prop="address" >
              <el-input v-model="formData.address" :clearable="true"  placeholder="请输入酒店地址" />
            </el-form-item>
            <el-form-item label="酒店英文地址:"  prop="addressEn" >
              <el-input v-model="formData.addressEn" :clearable="true"  placeholder="请输入酒店英文地址" />
            </el-form-item>
            <el-form-item label="百度维度:"  prop="baiduLat" >
              <el-input v-model="formData.baiduLat" :clearable="true"  placeholder="请输入百度维度" />
            </el-form-item>
            <el-form-item label="百度经度:"  prop="baiduLon" >
              <el-input v-model="formData.baiduLon" :clearable="true"  placeholder="请输入百度经度" />
            </el-form-item>
            <el-form-item label="取消政策:"  prop="canclePolicy" >
              <el-input v-model="formData.canclePolicy" :clearable="true"  placeholder="请输入取消政策" />
            </el-form-item>
            <el-form-item label="城市名称:"  prop="cityName" >
              <el-input v-model="formData.cityName" :clearable="true"  placeholder="请输入城市名称" />
            </el-form-item>
            <el-form-item label="创建人:"  prop="createBy" >
              <el-input v-model="formData.createBy" :clearable="true"  placeholder="请输入创建人" />
            </el-form-item>
            <el-form-item label="创建时间:"  prop="createTime" >
              <el-date-picker v-model="formData.createTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="支持的信用卡:"  prop="creditCards" >
              <el-input v-model="formData.creditCards" :clearable="true"  placeholder="请输入支持的信用卡" />
            </el-form-item>
            <el-form-item label="装修时间:"  prop="decorateTime" >
              <el-date-picker v-model="formData.decorateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="目的地标签:"  prop="destinationLabel" >
              <el-input v-model.number="formData.destinationLabel" :clearable="true" placeholder="请输入目的地标签" />
            </el-form-item>
            <el-form-item label="酒店设施列表:"  prop="facilities" >
              <el-input v-model="formData.facilities" :clearable="true"  placeholder="请输入酒店设施列表" />
            </el-form-item>
            <el-form-item label="高德经度:"  prop="gaodeLat" >
              <el-input v-model="formData.gaodeLat" :clearable="true"  placeholder="请输入高德经度" />
            </el-form-item>
            <el-form-item label="高德维度:"  prop="gaodeLon" >
              <el-input v-model="formData.gaodeLon" :clearable="true"  placeholder="请输入高德维度" />
            </el-form-item>
            <el-form-item label="谷歌维度:"  prop="googleLat" >
              <el-input v-model="formData.googleLat" :clearable="true"  placeholder="请输入谷歌维度" />
            </el-form-item>
            <el-form-item label="谷歌经度:"  prop="googleLon" >
              <el-input v-model="formData.googleLon" :clearable="true"  placeholder="请输入谷歌经度" />
            </el-form-item>
            <el-form-item label="酒店编号:"  prop="hotelId" >
              <el-input v-model.number="formData.hotelId" :clearable="true" placeholder="请输入酒店编号" />
            </el-form-item>
            <el-form-item label="酒店名称:"  prop="hotelName" >
              <el-input v-model="formData.hotelName" :clearable="true"  placeholder="请输入酒店名称" />
            </el-form-item>
            <el-form-item label="英文名:"  prop="hotelNameEn" >
              <el-input v-model="formData.hotelNameEn" :clearable="true"  placeholder="请输入英文名" />
            </el-form-item>
            <el-form-item label="电话:"  prop="hotelPhone" >
              <el-input v-model="formData.hotelPhone" :clearable="true"  placeholder="请输入电话" />
            </el-form-item>
            <el-form-item label="酒店介绍:"  prop="introduction" >
              <el-input v-model="formData.introduction" :clearable="true"  placeholder="请输入酒店介绍" />
            </el-form-item>
            <el-form-item label="开业时间:"  prop="openingTime" >
              <el-date-picker v-model="formData.openingTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="星级(1/2/3/4/5):"  prop="star" >
              <el-input v-model.number="formData.star" :clearable="true" placeholder="请输入星级(1/2/3/4/5)" />
            </el-form-item>
            <el-form-item label="修改人:"  prop="updateBy" >
              <el-input v-model="formData.updateBy" :clearable="true"  placeholder="请输入修改人" />
            </el-form-item>
            <el-form-item label="修改时间:"  prop="updateTime" >
              <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="酒店地址">
                        {{ formData.address }}
                </el-descriptions-item>
                <el-descriptions-item label="酒店英文地址">
                        {{ formData.addressEn }}
                </el-descriptions-item>
                <el-descriptions-item label="百度维度">
                        {{ formData.baiduLat }}
                </el-descriptions-item>
                <el-descriptions-item label="百度经度">
                        {{ formData.baiduLon }}
                </el-descriptions-item>
                <el-descriptions-item label="取消政策">
                        {{ formData.canclePolicy }}
                </el-descriptions-item>
                <el-descriptions-item label="城市名称">
                        {{ formData.cityName }}
                </el-descriptions-item>
                <el-descriptions-item label="创建人">
                        {{ formData.createBy }}
                </el-descriptions-item>
                <el-descriptions-item label="创建时间">
                      {{ formatDate(formData.createTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="支持的信用卡">
                        {{ formData.creditCards }}
                </el-descriptions-item>
                <el-descriptions-item label="装修时间">
                      {{ formatDate(formData.decorateTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="目的地标签">
                        {{ formData.destinationLabel }}
                </el-descriptions-item>
                <el-descriptions-item label="酒店设施列表">
                        {{ formData.facilities }}
                </el-descriptions-item>
                <el-descriptions-item label="高德经度">
                        {{ formData.gaodeLat }}
                </el-descriptions-item>
                <el-descriptions-item label="高德维度">
                        {{ formData.gaodeLon }}
                </el-descriptions-item>
                <el-descriptions-item label="谷歌维度">
                        {{ formData.googleLat }}
                </el-descriptions-item>
                <el-descriptions-item label="谷歌经度">
                        {{ formData.googleLon }}
                </el-descriptions-item>
                <el-descriptions-item label="酒店编号">
                        {{ formData.hotelId }}
                </el-descriptions-item>
                <el-descriptions-item label="酒店名称">
                        {{ formData.hotelName }}
                </el-descriptions-item>
                <el-descriptions-item label="英文名">
                        {{ formData.hotelNameEn }}
                </el-descriptions-item>
                <el-descriptions-item label="电话">
                        {{ formData.hotelPhone }}
                </el-descriptions-item>
                <el-descriptions-item label="酒店介绍">
                        {{ formData.introduction }}
                </el-descriptions-item>
                <el-descriptions-item label="开业时间">
                      {{ formatDate(formData.openingTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="星级(1/2/3/4/5)">
                        {{ formData.star }}
                </el-descriptions-item>
                <el-descriptions-item label="修改人">
                        {{ formData.updateBy }}
                </el-descriptions-item>
                <el-descriptions-item label="修改时间">
                      {{ formatDate(formData.updateTime) }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createHtlHotel,
  deleteHtlHotel,
  deleteHtlHotelByIds,
  updateHtlHotel,
  findHtlHotel,
  getHtlHotelList
} from '@/api/htlHotel'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'HtlHotel'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        address: '',
        addressEn: '',
        baiduLat: '',
        baiduLon: '',
        canclePolicy: '',
        cityName: '',
        createBy: '',
        createTime: new Date(),
        creditCards: '',
        decorateTime: new Date(),
        destinationLabel: 0,
        facilities: '',
        gaodeLat: '',
        gaodeLon: '',
        googleLat: '',
        googleLon: '',
        hotelId: 0,
        hotelName: '',
        hotelNameEn: '',
        hotelPhone: '',
        introduction: '',
        openingTime: new Date(),
        star: 0,
        updateBy: '',
        updateTime: new Date(),
        })


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getHtlHotelList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteHtlHotelFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteHtlHotelByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateHtlHotelFunc = async(row) => {
    const res = await findHtlHotel({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rehtlHotel
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHtlHotelFunc = async (row) => {
    const res = await deleteHtlHotel({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findHtlHotel({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.rehtlHotel
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          address: '',
          addressEn: '',
          baiduLat: '',
          baiduLon: '',
          canclePolicy: '',
          cityName: '',
          createBy: '',
          createTime: new Date(),
          creditCards: '',
          decorateTime: new Date(),
          destinationLabel: 0,
          facilities: '',
          gaodeLat: '',
          gaodeLon: '',
          googleLat: '',
          googleLon: '',
          hotelId: 0,
          hotelName: '',
          hotelNameEn: '',
          hotelPhone: '',
          introduction: '',
          openingTime: new Date(),
          star: 0,
          updateBy: '',
          updateTime: new Date(),
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        address: '',
        addressEn: '',
        baiduLat: '',
        baiduLon: '',
        canclePolicy: '',
        cityName: '',
        createBy: '',
        createTime: new Date(),
        creditCards: '',
        decorateTime: new Date(),
        destinationLabel: 0,
        facilities: '',
        gaodeLat: '',
        gaodeLon: '',
        googleLat: '',
        googleLon: '',
        hotelId: 0,
        hotelName: '',
        hotelNameEn: '',
        hotelPhone: '',
        introduction: '',
        openingTime: new Date(),
        star: 0,
        updateBy: '',
        updateTime: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createHtlHotel(formData.value)
                  break
                case 'update':
                  res = await updateHtlHotel(formData.value)
                  break
                default:
                  res = await createHtlHotel(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
