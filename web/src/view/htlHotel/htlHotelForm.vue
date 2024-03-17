<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="酒店地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="酒店英文地址:" prop="addressEn">
          <el-input v-model="formData.addressEn" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="百度维度:" prop="baiduLat">
          <el-input v-model="formData.baiduLat" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="百度经度:" prop="baiduLon">
          <el-input v-model="formData.baiduLon" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="取消政策:" prop="canclePolicy">
          <el-input v-model="formData.canclePolicy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="城市名称:" prop="cityName">
          <el-input v-model="formData.cityName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建人:" prop="createBy">
          <el-input v-model="formData.createBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createTime">
          <el-date-picker v-model="formData.createTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="支持的信用卡:" prop="creditCards">
          <el-input v-model="formData.creditCards" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="装修时间:" prop="decorateTime">
          <el-date-picker v-model="formData.decorateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="目的地标签:" prop="destinationLabel">
          <el-input v-model.number="formData.destinationLabel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="酒店设施列表:" prop="facilities">
          <el-input v-model="formData.facilities" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="高德经度:" prop="gaodeLat">
          <el-input v-model="formData.gaodeLat" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="高德维度:" prop="gaodeLon">
          <el-input v-model="formData.gaodeLon" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="谷歌维度:" prop="googleLat">
          <el-input v-model="formData.googleLat" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="谷歌经度:" prop="googleLon">
          <el-input v-model="formData.googleLon" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="酒店编号:" prop="hotelId">
          <el-input v-model.number="formData.hotelId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="酒店名称:" prop="hotelName">
          <el-input v-model="formData.hotelName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="英文名:" prop="hotelNameEn">
          <el-input v-model="formData.hotelNameEn" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电话:" prop="hotelPhone">
          <el-input v-model="formData.hotelPhone" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="酒店介绍:" prop="introduction">
          <el-input v-model="formData.introduction" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开业时间:" prop="openingTime">
          <el-date-picker v-model="formData.openingTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="星级(1/2/3/4/5):" prop="star">
          <el-input v-model.number="formData.star" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="修改人:" prop="updateBy">
          <el-input v-model="formData.updateBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="修改时间:" prop="updateTime">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createHtlHotel,
  updateHtlHotel,
  findHtlHotel
} from '@/api/htlHotel'

defineOptions({
    name: 'HtlHotelForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHtlHotel({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehtlHotel
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
