<!--  -->
<template>        
    <div class="login">
        <span class="name">openTool</span> 
        <a-form 
            layout="horizontal" 
            :model="modelRef" 
            @submit.prevent="handleSubmit"
          >
            <a-form-item
              name="ip"
              v-bind="validateInfos.ip"
            >
              <a-input v-model:value="modelRef.ip" size="large" placeholder="IP地址" allow-clear  autocomplete="new-ip" >
                
              </a-input>
            </a-form-item>

            <a-form-item
              name="username"
              v-bind="validateInfos.username"
            >
              <a-input v-model:value="modelRef.username" size="large" placeholder="用户" allow-clear  autocomplete="new-username" >
                
              </a-input>
            </a-form-item>
            <a-form-item name="password" v-bind="validateInfos.password">
              <a-input-password
                  autocomplete="new-password"
                  v-model:value="modelRef.password"
                  size="large"
                  placeholder="密码"
                  allow-clear
              >
                
              </a-input-password>
            </a-form-item>
            <a-form-item>
              <a-button type="primary" html-type="submit" size="large" :loading="modelRef.loading" block>
                登录
              </a-button>
            </a-form-item>
      </a-form>
    </div>
</template>

<script setup>
import { reactive, ref} from 'vue'
import { useRouter } from 'vue-router'
import axios from '@/api/axios'
import { Form } from 'ant-design-vue';
import { Base64 } from 'js-base64'
import {useStore} from 'vuex'


const store = useStore()

const useForm = Form.useForm; 

   const modelRef = reactive({
      username: '',
      password: '',
      loading: false,
      ip: import.meta.env.VITE_HOST
    })


    const validateIPAddress = (rule, value, callback)=>{

        let regexp = /(http|https):?/
        if (value == '') {
            return callback(new Error('请输入iP地址'));
        }else if(!regexp.test(value)){ 
            return callback(new Error('格式不正确'));
        }else{
          callback()
        }
    }  

  const rulesRef = reactive({
      ip: [
        {
          required: true,
          validator: validateIPAddress,
          trigger: 'blur'
        }
      ],
      username: [
        {
          required: true,
          message: '请输入用户名',
          trigger: 'blur'
        },
      ],
       password: [
        {
          required: true,
          message: '请输入密码',
          trigger: 'blur'
        },
      ]
  })    
  const { validate, validateInfos } = useForm(modelRef, rulesRef);


  const router = useRouter();

 

  const handleSubmit =  () => {
    validate()
        .then(() => {

          modelRef.loading = true
          submitForm()
        })
        .catch(err => {
          console.log('error', err);
        });
 
  }


const submitForm = async () => {


      modelRef.password = 'admin'
      var result = Base64.encode(modelRef.password)

      axios.post('/login', {
        host: modelRef.ip,
        user: modelRef.username,
        password: result

      }).then(res => {
        
        if(res.code == 200){
            sessionStorage.setItem('search-token', res.token)
            store.commit("updateToken", res.token)
            setTimeout(() => {
                console.log(store.state)
                router.push({
                  path: 'detail',
                  query: {
                      url:modelRef.ip
                      }
                  })
                modelRef.loading = false  
            }, 1000)
           
            
        }else{
             modelRef.loading = false
        }
      

      }).catch(()=>{
        modelRef.loading = false
      })
}


</script>
<style lang='less' scoped>
:deep(.ant-form) {
  width: 320px;
  .ant-form-item-label {
    padding-right: 6px;
  }
}
.login{

    display: flex;
    padding-top: 120px;
    flex-direction: column;
    align-items: center;  

}
.name{
  font-size: 32px;
  margin: 20px 0;
}
</style>